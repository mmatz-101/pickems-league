package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mmatz101/go-odds/cmd/api"
	databases "github.com/mmatz101/go-odds/databases/sqlc"
	"github.com/valyala/fasthttp"
)

const ENDPOINT = "https://gql.vegasinsider.com/graphql"

type Data struct {
	Data Games `json:"data"`
}

type Games struct {
	Games []Game `json:"getTNEvents"`
}

type Game struct {
	ID        string
	HomeTeam  Team         `json:"home"`
	AwayTeam  Team         `json:"away"`
	Channel   string       `json:"channel"`
	Date      string       `json:"date"`
	Status    string       `json:"status"`
	Week      int32        `json:"week"`
	WeekName  string       `json:"weekName"`
	Consensus ConsensusBet `json:"consensus"`
	HomeScore int32        `json:"homescore"`
	AwayScore int32        `json:"awayscore"`
}

type Team struct {
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	LogoUrl   string `json:"LogoUrl"`
}

type ConsensusBet struct {
	Spread ConsensusSpread `json:"spread"`
}

type ConsensusSpread struct {
	AwayOrOverLine  float64 `json:"awayoroverline"`
	AwayOrOverOdd   int32   `json:"awayoroverodd"`
	HomeOrUnderLine float64 `json:"homeorunderline"`
	HomeOrUnderOdd  int32   `json:"homeorunderodd"`
	SportsbookId    int32   `json:"sportsbookid"`
	CreatedAt       string  `json:"created"`
}

func vegasInsiderCall(ctx *fiber.Ctx) error {
	// Parse the URL for league, week, year, season
	year, err := ctx.ParamsInt("year")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to parse year.")
	}
	week, err := ctx.ParamsInt("week")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to parse week.")
	}
	seasonType := ctx.Params("season_type")
	if seasonType != "nfl" && seasonType != "ncaaf" {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to use this season type.")
	}
	league := ctx.Params("league")

	date := fmt.Sprintf("%d-%s-%d", year, seasonType, week)

	jsonQuery := fmt.Sprintf(`{
		"query": "query GetTNEvents($sport: String!, $searchBy: TNEventsSearchBy) { getTNEvents(sport: $sport, searchBy: $searchBy) { home { fullName shortName logoUrl } away { fullName shortName logoUrl } channel date status week weekName consensus { spread { awayOrOverLine awayOrOverOdd homeOrUnderLine homeOrUnderOdd sportsbookId created } } homeScore awayScore }}",
		"variables": {"sport": "%s", "searchBy": {"dates": ["%s"]}}
	}`, league, date)

	// TRYING TO MAKE THE POST REQUEST
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(ENDPOINT)
	req.Header.Set("Content-Type", "application/json")
	req.SetBody([]byte(jsonQuery))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = fasthttp.Do(req, resp)
	if err != nil {
		fmt.Printf("Client get failed: %s\n", err)
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	var jsonData Data
	err = json.Unmarshal(resp.Body(), &jsonData)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to unmarshal data.")
	}

	// Store the parsed out jsonDATA into a database
	var DBGames []databases.CreateGameParams
	for _, game := range jsonData.Data.Games {
		// Create index for game
		game_idx := game.HomeTeam.ShortName + "." + game.AwayTeam.ShortName + "." + game.Date
		if game.HomeTeam.ShortName == "" || game.AwayTeam.ShortName == "" {
			return ctx.SendString("No team short name. Need a short team name for game id.")
		}
		// Date parser
		date_parsed, err := time.Parse("2006-01-02T15:04:05", game.Date)
		if err != nil {
			return ctx.SendString("Unable to parse data.")
		}
		// Determine the spread results if avaliable
		var matchResults databases.SpreadWinner
		if game.Status == "Final" {
			if float64(game.HomeScore)+game.Consensus.Spread.HomeOrUnderLine > float64(game.AwayScore) {
				matchResults = databases.SpreadWinnerHOME
			} else if float64(game.HomeScore)+game.Consensus.Spread.HomeOrUnderLine < float64(game.AwayScore) {
				matchResults = databases.SpreadWinnerAWAY
			} else {
				matchResults = databases.SpreadWinnerPUSH
			}
		} else {
			matchResults = databases.SpreadWinnerUNDETERMINED
		}
		DBGames = append(DBGames, databases.CreateGameParams{
			ID:                game_idx,
			HometeamFullname:  game.HomeTeam.FullName,
			HometeamShortname: game.HomeTeam.ShortName,
			HometeamLogourl:   game.HomeTeam.LogoUrl,
			AwayteamFullname:  game.AwayTeam.FullName,
			AwayteamShortname: game.AwayTeam.ShortName,
			AwayteamLogourl:   game.AwayTeam.LogoUrl,
			Channel:           game.Channel,
			Date:              date_parsed,
			Status:            game.Status,
			Year:              int32(year),
			Week:              game.Week,
			Weekname:          game.WeekName,
			Homeorunderline:   game.Consensus.Spread.HomeOrUnderLine,
			Homeorunderodd:    game.Consensus.Spread.HomeOrUnderOdd,
			Awayoroverline:    game.Consensus.Spread.AwayOrOverLine,
			Awayoroverodd:     game.Consensus.Spread.AwayOrOverOdd,
			CreatedAtVegas:    game.Consensus.Spread.CreatedAt,
			Homescore:         game.HomeScore,
			Awayscore:         game.AwayScore,
			GameSpreadWinner:  matchResults,
			SeasonType:        databases.SeasonType(strings.ToUpper(seasonType)),
			League:            databases.League(strings.ToUpper(league)),
		})
	}

	for _, game := range DBGames {
		_, err := api.Store.CreateGame(context.Background(), game)
		if err != nil {
			log.Println("Unable to create game duplicate key found. Attempting to update game.")
			// converting CreateGame struct to UpdatedGame struct
			updateGame := databases.UpdateGameAlreadyFoundParams(game)
			_, err = api.Store.UpdateGameAlreadyFound(context.Background(), updateGame)
			if err != nil {
				log.Fatalln("Unable to create game or update game", err)
			}
		}
	}

	return ctx.Status(fiber.StatusCreated).SendString("DB Updated.")
}

func retrieveGames(ctx *fiber.Ctx) error {
	// Parse the URL for league, week, year, season
	year, err := ctx.ParamsInt("year")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to parse year.")
	}
	week, err := ctx.ParamsInt("week")
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to parse week.")
	}
	seasonType := ctx.Params("season_type")
	if seasonType != "nfl" && seasonType != "ncaaf" {
		ctx.Status(fiber.StatusBadRequest).SendString("Unable to use this season type.")
	}
	league := ctx.Params("league")

	arg := databases.ListGamesByWeekParams{
		Week:       int32(week),
		Year:       int32(year),
		SeasonType: databases.SeasonType(strings.ToUpper(seasonType)),
		League:     databases.League(strings.ToUpper(league)),
	}

	games, err := api.Store.ListGamesByWeek(context.Background(), arg)
	if err != nil {
		fmt.Println("Unable to find list of games.")
		return ctx.Status(fiber.StatusBadRequest).SendString(":(")
	}

	return ctx.Status(fiber.StatusAccepted).JSON(games)
}

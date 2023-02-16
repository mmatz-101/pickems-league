package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	databases "github.com/mmatz101/go-odds/databases/sqlc"
)

// CreatePick ###########################################################
type createPickRequest struct {
	GameID       string                `json:"game_id"`
	Year         int32                 `json:"year"`
	Week         int32                 `json:"week"`
	UserPick     databases.SpreadPicks `json:"user_pick"`
	UserPickType databases.SpreadType  `json:"user_pick_type"`
}

func CreatePick(ctx *fiber.Ctx) error {
	// parse json
	var req createPickRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, err)
	}

	// validate Request
	err := validate.Struct(req)
	if err != nil {
		return errorResponse(ctx, err)
	}

	// get username
	userJWT := ctx.Locals("user").(*jwt.Token)
	claims := userJWT.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	// get game
	game, err := Store.GetGameByID(ctx.Context(), req.GameID)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	// create pick to database
	arg := databases.CreatePickParams{
		Username:         username,
		GameID:           game.ID,
		Year:             req.Year,
		Week:             req.Week,
		League:           game.League,
		UserPick:         req.UserPick,
		UserPickType:     req.UserPickType,
		GameSpreadWinner: game.GameSpreadWinner,
	}

	pick, err := Store.CreatePick(ctx.Context(), arg)
	if err != nil {
		log.Println(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{
		"succes": true,
		"data":   pick,
	})
}

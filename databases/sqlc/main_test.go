package databases_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	databases "github.com/mmatz101/go-odds/databases/sqlc"
	"github.com/mmatz101/go-odds/utils"
	"github.com/stretchr/testify/require"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	dbDriver = "pgx"
	dbURL    = "postgresql://postgres:secret@localhost:5432/league_db?sslmode=disable"
)

var testQueries *databases.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbURL)
	if err != nil {
		log.Fatalln("Cannot connect to database: ", err)
	}

	testQueries = databases.New(conn)

	os.Exit(m.Run())
}

// Create random user for our unit test.
func createRandomUser(t *testing.T) databases.User {
	hashedPassword, err := utils.HashPassword("password")
	require.NoError(t, err)

	arg := databases.CreateUserParams{
		Username:     fmt.Sprint(utils.RandomStringGenerator(4), " ", utils.RandomStringGenerator(5)),
		FullName:     fmt.Sprint(utils.RandomStringGenerator(4), " ", utils.RandomStringGenerator(5)),
		Email:        fmt.Sprint(utils.RandomStringGenerator(4), "@", utils.RandomStringGenerator(5)),
		HashPassword: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	// checking the values we add
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)

	// checking the values generated by postgres
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	databases "github.com/mmatz101/go-odds/databases/sqlc"
	"github.com/mmatz101/go-odds/utils"
	"gopkg.in/go-playground/validator.v9"
)

var validate = validator.New()

// CREATE USER ####################################################

type createUserRequest struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type createUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(ctx *fiber.Ctx) error {
	// parse json
	var req createUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, err)
	}

	// validate Request
	err := validate.Struct(req)
	if err != nil {
		return errorResponse(ctx, err)
	}

	// hash Password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errorResponse(ctx, err)
	}

	// create argument and send to datbase
	arg := databases.CreateUserParams{
		Username:     req.Username,
		FullName:     req.FullName,
		Email:        req.Email,
		HashPassword: hashedPassword,
	}

	// write to database
	user, err := Store.CreateUser(ctx.Context(), arg)
	if err != nil {
		return errorResponse(ctx, err)
	}

	// create Response
	resp := createUserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": true,
		"data":    resp,
	})
}

// Login USER ###############################################################
// loginUserRequest struct for the request to login a user.
type loginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginUser from the json request containing username and password. Search for the
// username in the datebase the check the password against the hash_password stored
// in the database. Create a JWT if success as well as return user information.
func LoginUser(ctx *fiber.Ctx) error {
	// parse login request
	var req loginUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return errorResponse(ctx, err)
	}

	// query database for user
	user, err := Store.GetUser(ctx.Context(), req.Username)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "either username or password incorrect.",
		})
	}

	// test if the password matches the password of user
	if err = utils.CheckPassword(req.Password, user.HashPassword); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "either username or password incorrect.",
		})
	}

	// response and token
	resp := loginUserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// create the claim
	claims := jwt.MapClaims{
		"username": resp.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	// todo: dont return hash password.
	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    resp,
		"token":   t,
	})

}

// Welcome Exlusive User ##################################################
func SecureArea(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)
	return ctx.SendString("Welcome + " + name)
}

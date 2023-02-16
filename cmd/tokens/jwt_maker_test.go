package tokens

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mmatz101/go-odds/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMarker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomStringGenerator(32))
	require.NoError(t, err)

	username := utils.RandomStringGenerator(6)
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTMarker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomStringGenerator(32))
	require.NoError(t, err)

	username := utils.RandomStringGenerator(6)

	token, err := maker.CreateToken(username, -time.Second)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgoNone(t *testing.T) {
	payload, err := NewPayload(utils.RandomStringGenerator(32), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.RandomStringGenerator(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}

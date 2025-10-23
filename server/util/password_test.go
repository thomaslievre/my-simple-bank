package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashedPassword(t *testing.T) {
	password := RandomString(6)

	hashPassword1, err := HashPassword(password) // ignore error for the sake of simplicity
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword1)

	match := CheckPasswordHash(password, hashPassword1)

	require.NoError(t, match)

	wrongPassword := RandomString(6)
	match = CheckPasswordHash(wrongPassword, hashPassword1)
	require.EqualError(t, match, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashPassword2, err := HashPassword(password) // ignore error for the sake of simplicity
	require.NoError(t, err)
	require.NotEmpty(t, hashPassword2)

	require.NotEqual(t, hashPassword1, hashPassword2)

}

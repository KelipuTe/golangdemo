package crypto

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {
	pswd := "123456"
	encrypt, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(encrypt))
	err = bcrypt.CompareHashAndPassword(encrypt, []byte(pswd))
	require.NoError(t, err)
}

package login

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	logtest "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"os"
	"testing"
	"time"
)

var hook *logtest.Hook
var log *logrus.Entry

func TestMain(m *testing.M) {
	var logger *logrus.Logger
	logger, hook = logtest.NewNullLogger()
	log = initLogger(logger)

	code := m.Run()
	os.Exit(code)
}

func TestNewLogin(t *testing.T) {
	login := NewLogin("url.123")

	assert.Equal(t, login.callbackURI, "url.123")
}

func TestLogin_SaveToken(t *testing.T) {
	login := NewLogin("url.123")
	tokenName, err := uuid.NewV4()
	assert.NoError(t, err)

	t.Run("ValidToken", func(t *testing.T) {
		tim := time.Now()

		err = login.SaveToken(tokenName.String(), &oauth2.Token{
			AccessToken:  "aaa",
			TokenType:    "ttt",
			RefreshToken: "rrr",
			Expiry:       tim,
		})
		assert.NoError(t, err)

		err = os.Remove(tokenName.String())
		assert.NoError(t, err)
	})
}

func TestCreateCodeVerifier(t *testing.T) {
	code := createCodeVerifier(10)
	assert.Equal(t, 14, len(code))
}

func TestCreateVerifierChallenge(t *testing.T) {
	code := createVerifierChallenge("12345abcde")
	fmt.Println(code)
	assert.Equal(t, "PDc_SVO4XN6liOBDbBNMgZ9XC3LB23QOs1z8lCuqK84", code)
}

func TestBase64Escape(t *testing.T) {
	var escape = []byte("any + old & data")

	escaped := base64Escape(escape)
	assert.Equal(t, "YW55ICsgb2xkICYgZGF0YQ", escaped)
}

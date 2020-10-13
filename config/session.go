package config

import (
	"io/ioutil"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func init() {
	authKeyOne, err1 := ioutil.ReadFile("./config/auth")
	encryptionKeyOne, err2 := ioutil.ReadFile("./config/encr")

	setCookieStore(authKeyOne, encryptionKeyOne, err1, err2)
}

func setCookieStore(auth []byte, encr []byte, err1 error, err2 error) {
	if err1 == nil && err2 == nil {

		Store = sessions.NewCookieStore(auth, encr)
		Store.Options = &sessions.Options{
			MaxAge:   60 * 60 * 24,
			HttpOnly: true,
		}
	} else {
		newAuth := securecookie.GenerateRandomKey(64)
		newEncr := securecookie.GenerateRandomKey(32)

		errWAuth := ioutil.WriteFile("./config/auth", newAuth, 0644)
		errWEncr := ioutil.WriteFile("./config/encr", newEncr, 0644)

		setCookieStore(newAuth, newEncr, errWAuth, errWEncr)
	}
}

package icecream_middleware

import (
	"encoding/base64"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	APP_KEY = "zalora.com"
)

func isBasicAuthentication(r *http.Request) bool {
	return len(r.Header["Authorization"]) > 0 && strings.HasPrefix(string(r.Header["Authorization"][0]), "Basic")
}

func TokenHandler(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		r = setTokenOnHeader(w, r)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func setTokenOnHeader(w http.ResponseWriter, r *http.Request) *http.Request {

	var username, password string

	if isBasicAuthentication(r) {
		log.Print(r.Context(), "Basic authentication")
		decodeData, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(string(r.Header["Authorization"][0]), "Basic "))
		if err != nil {
			log.Fatal(r.Context(), "failed to decode data, error: ", err.Error())
			return r
		}
		credentials := strings.Split(string(decodeData), ":")
		username, password = credentials[0], credentials[1]
	}

	if username != "zalora" || password != "zalora" {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"error":"invalid_credentials"}`)
		return r
	}

	// We are happy with the credentials, so build a token. We've given it
	// an expiry of 1 hour.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": username,
		"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(APP_KEY))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error":"token_generation_failed"}`)
		return r
	}

	var bearer = "Bearer " + tokenString
	r.Header.Del("Authorization")
	r.Header.Add("Authorization", bearer)

	return r

}

// AuthMiddleware is our middleware to check our token is valid. Returning
// a 401 status to the client if it is not valid.
func AuthMiddleware(next http.Handler) http.Handler {

	if len(APP_KEY) == 0 {
		log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
	}
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(APP_KEY), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return jwtMiddleware.Handler(next)
}

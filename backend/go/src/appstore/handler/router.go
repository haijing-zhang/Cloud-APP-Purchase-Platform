package handler

// Use Mux library to handle HTTP routing (Unlike Spring, Go cannot support annotation-based HTTP routing)

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

var mySigningKey []byte

func InitRouter(config *util.TokenInfo) http.Handler {
    mySigningKey = []byte(config.Secret)

    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })

    router := mux.NewRouter()

    router.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST")
  router.Handle("/checkout", jwtMiddleware.Handler(http.HandlerFunc(checkoutHandler))).Methods("POST")    router.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET")

    router.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST")
    router.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST")

    return router
}


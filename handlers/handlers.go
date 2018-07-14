package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

// User ...
// Custom object which can be stored in the claims
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthToken ...
// This is what is retured to the user
type AuthToken struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthTokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	User
}

type Claim struct {
	Exp float64
	Iat float64
	Sub string
}

// ErrorMsg ...
// Custom error object
type ErrorMsg struct {
	Message string `json:"message"`
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func Proxyhandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.ServeHTTP(w, r)
	}
}

type Authenticater interface {
	Authenticater() bool
}

type CreateToken interface {
	CreateToken() AuthToken
}

func GetAuthentication(auth Authenticater) bool {
	return auth.Authenticater()
}

func GetJWTToken(tok CreateToken) AuthToken {
	return tok.CreateToken()
}

func (u *User) Authenticater() bool {
	log.Println("User: authenticating", u.Username)

	authenticated := true // <--- replace this with real authentication logic
	return authenticated
}

func (u *User) CreateToken() AuthToken {
	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": expiresAt,
		"iat": time.Now().Unix(),
		"sub": u.Username,
	}

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	authToken := AuthToken{
		TokenType: "Bearer",
		Token:     tokenString,
		ExpiresIn: expiresAt,
	}

	return authToken

}

func AuthenticateHandler(w http.ResponseWriter, req *http.Request) {

	user := &User{
		Username: req.Header.Get("username"),
		Password: req.Header.Get("password"),
	}

	userAuthenticated := GetAuthentication(user)

	if userAuthenticated == true {
		authToken := GetJWTToken(user) // if user successfully authenticates then create JWT

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&authToken)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorMsg{Message: "An authentication error occurred"})
	}

}

func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(ErrorMsg{Message: error.Error()})
					return
				}
				if token.Valid {

					var user Claim

					err := mapstructure.Decode(token.Claims, &user)

					if err != nil {
						panic(err)
					}

					vars := mux.Vars(req)
					name := vars["id"]

					if name != user.Sub {
						json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
						return
					}

					context.Set(req, "decoded", token.Claims)
					w.Header().Set("Content-Type", "application/json")
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(ErrorMsg{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(ErrorMsg{Message: "An authorization header is required"})
		}
	})
}

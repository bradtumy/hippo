package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func createJWT(user string) AuthToken {
	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"exp": expiresAt,
		"iat": time.Now().Unix(),
		"sub": user,
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

func checkCredentials(user string, pass string) bool {
	var strAuthenticationStatus = true
	// Read properties file to get the IdRepo Configuration
	// Create connection to IdRepo
	// Authenticate the user using credentials provided by the user
	// send username and password to identity repo
	return strAuthenticationStatus
}

func AuthenticateHandler(w http.ResponseWriter, req *http.Request) {
	var authenticated = false
	var user User
	_ = json.NewDecoder(req.Body).Decode(&user)

	user.Username = req.Header.Get("username")
	user.Password = req.Header.Get("password")

	// now actually authenticate the user
	authenticated = checkCredentials(user.Username, user.Password)

	if authenticated == true {
		authToken := createJWT(user.Username)
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

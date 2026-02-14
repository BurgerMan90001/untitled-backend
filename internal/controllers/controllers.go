package controllers

import (
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/auth"
	"github.com/BurgerMan90001/untitled-backend/internal/model/responses"
	"github.com/BurgerMan90001/untitled-backend/internal/util"
	"golang.org/x/oauth2"
)

var oauthConfig = oauth2.Config{
	ClientID:     "your_client_id",
	ClientSecret: "your_client_secret",
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"profile", "email"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://oauth_provider.com/oauth2/authorize",
		TokenURL: "https://oauth_provider.com/oauth2/token",
	},
}

func Root(w http.ResponseWriter, r *http.Request) {

	
	util.WriteJSON(w, responses.Response{Message: "root"})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Still alive"})
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Root"})
}

func Secret(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}
	auth.VerifyJWTToken(cookie.Value)
	//util.WriteJSON()
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	auth.CreateJWTToken(r.FormValue("userId"))
}

func OathLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
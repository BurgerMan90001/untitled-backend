package controllers

import (
	"net/http"

	"github.com/BurgerMan90001/untitled-backend/internal/model/responses"
	"github.com/BurgerMan90001/untitled-backend/internal/util"
)

oauthConfig := oauth2.Config{
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

	url := oauthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	//util.WriteJSON(w, responses.Response{Message: "root"})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Still alive"})
}

func TestRoute(w http.ResponseWriter, r *http.Request) {
	util.WriteJSON(w, responses.Response{Message: "Root"})
}

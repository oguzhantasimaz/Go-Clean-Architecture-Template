package controller

import (
	"net/http"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleController struct {
	GoogleUseCase domain.GoogleUseCase
	Env           *bootstrap.Env
}

var googleOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:8080/api/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: google.Endpoint,
}

func (gc *GoogleController) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := gc.GoogleUseCase.GenerateStateOauthCookie(w)
	googleOauthConfig.ClientSecret = gc.Env.GoogleClientSecret
	googleOauthConfig.ClientID = gc.Env.GoogleClientID
	u := googleOauthConfig.AuthCodeURL(oauthState)

	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (gc *GoogleController) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	googleOauthConfig.ClientSecret = gc.Env.GoogleClientSecret
	googleOauthConfig.ClientID = gc.Env.GoogleClientID
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Error("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := gc.GoogleUseCase.GetUserDataFromGoogle(googleOauthConfig, r.FormValue("code"), oauthGoogleUrlAPI)
	if err != nil {
		log.Error(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	accessToken, refreshToken, err := gc.GoogleUseCase.GoogleLogin(ctx, data, gc.Env)
	if err != nil {
		log.Error(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// write access token and refresh token to cookie
	utils.SetCookie(w, "access_token", accessToken)
	utils.SetCookie(w, "refresh_token", refreshToken)

	// redirect to home page
	http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
}

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/utils"

	log "github.com/sirupsen/logrus"
)

type SignupController struct {
	SignupUseCase domain.SignupUseCase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request domain.SignupRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := sc.SignupUseCase.SignUp(ctx, request, sc.Env)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, signupResponse)
	return
}

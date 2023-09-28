package controller

import (
	"encoding/json"
	"net/http"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/utils"

	log "github.com/sirupsen/logrus"
)

type LoginController struct {
	LoginUseCase domain.LoginUseCase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	var request domain.LoginRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := lc.LoginUseCase.Login(ctx, request, lc.Env)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	response := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, response)
	return
}

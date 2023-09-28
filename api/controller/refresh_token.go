package controller

import (
	"encoding/json"
	"net/http"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/utils"

	log "github.com/sirupsen/logrus"
)

type RefreshTokenController struct {
	RefreshTokenUseCase domain.RefreshTokenUseCase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request domain.RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, refreshToken, err := rtc.RefreshTokenUseCase.RefreshToken(ctx, request, rtc.Env)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	utils.JSON(w, http.StatusOK, refreshTokenResponse)
	return
}

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/domain"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/utils"

	log "github.com/sirupsen/logrus"
)

type UserController struct {
	UserUseCase domain.UserUseCase
	Env         *bootstrap.Env
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	users, err := uc.UserUseCase.GetUsers(ctx)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, users)
	return
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	id := fmt.Sprintf("%v", ctx.Value("user_id"))

	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := uc.UserUseCase.GetUserById(ctx, intId)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, user)
	return
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	var user *domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id := fmt.Sprintf("%v", ctx.Value("user_id"))

	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user.Id = userId

	err = uc.UserUseCase.UpdateUser(ctx, user)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, "Success")
	return
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "user_id", r.Context().Value("user_id"))

	id, err := strconv.Atoi(fmt.Sprintf("%v", ctx.Value("user_id")))
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = uc.UserUseCase.DeleteUser(ctx, id)
	if err != nil {
		log.Error(err)
		utils.JSON(w, http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	utils.JSON(w, http.StatusOK, "Success")
	return
}

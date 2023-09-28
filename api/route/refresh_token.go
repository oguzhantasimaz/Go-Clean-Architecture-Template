package route

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/api/controller"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/bootstrap"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/repository"
	"github.com/oguzhantasimaz/Go-Clean-Architecture-Template/usecase"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	ur := repository.NewUserRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUseCase: usecase.NewRefreshTokenUseCase(ur, timeout),
		Env:                 env,
	}

	r.HandleFunc("/refresh_token", rtc.RefreshToken).Methods("POST")
}

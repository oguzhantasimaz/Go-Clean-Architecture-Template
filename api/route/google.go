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

func NewGoogleRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	ur := repository.NewUserRepository(db)
	gc := &controller.GoogleController{
		GoogleUseCase: usecase.NewGoogleUseCase(ur, timeout),
		Env:           env,
	}

	r.HandleFunc("/google/login", gc.HandleGoogleLogin).Methods("GET")
	r.HandleFunc("/google/callback", gc.HandleGoogleCallback).Methods("GET")
}

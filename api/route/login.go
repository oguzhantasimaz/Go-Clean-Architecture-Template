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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUseCase: usecase.NewLoginUseCase(ur, timeout),
		Env:          env,
	}

	r.HandleFunc("/login", lc.Login).Methods("POST")
}

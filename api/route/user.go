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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *sqlx.DB, r *mux.Router) {
	ur := repository.NewUserRepository(db)
	uc := &controller.UserController{
		UserUseCase: usecase.NewUserUseCase(ur, timeout),
		Env:         env,
	}

	// USER ROUTES
	group := r.PathPrefix("/user").Subrouter()
	group.HandleFunc("/all", uc.GetUsers).Methods("GET")
	group.HandleFunc("", uc.GetUserById).Methods("GET")
	group.HandleFunc("", uc.UpdateUser).Methods("PUT")
	group.HandleFunc("", uc.DeleteUser).Methods("DELETE")
}

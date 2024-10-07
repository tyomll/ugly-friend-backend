package routes

import (
	"net/http"
	"ugly-friend/handler"
	"ugly-friend/models"
	"ugly-friend/utils"

	"github.com/go-chi/chi/v5"
)

const (
	LoginRoute        = "/login"
	RegisterRoute     = "/register"
	SearchFriendRoute = "/search-friend"
	AddFriendRoute    = "/friend"
	RemoveFriendRoute = "/friend"
	MyProfileRoute    = "/me"
	LeaderboardRoute  = "/leaderboard"
	CreateDebtRoute   = "/debt"
	GetDebtsRoute     = "/debts"
)

func SetupRoutes(router *chi.Mux, handler *handler.Handler) {
	router.Route("/api", func(r chi.Router) {
		// r.Post(LoginRoute, Login)
		r.Post(RegisterRoute, func(w http.ResponseWriter, r *http.Request) {
			var createUserReq models.CreateUserReq
			utils.HandleRouterBodyRequest(w, r, &createUserReq, func(requestBody interface{}) (interface{}, error) {
				return handler.Register(&createUserReq)
			})
		})
		// r.Get(SearchFriendRoute, SearchFriend)
		// r.Post(AddFriendRoute, AddFriend)
		// r.Delete(RemoveFriendRoute, RemoveFriend)
		// r.Get(MyProfileRoute, MyProfile)
		// r.Get(LeaderboardRoute, Leaderboard)
		// r.Post(CreateDebtRoute, CreateDebt)
		// r.Get(GetDebtsRoute, GetDebts)
	})
}

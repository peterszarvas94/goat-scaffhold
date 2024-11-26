package handlers

import (
	"net/http"

	l "github.com/peterszarvas94/goat/logger"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	// db, err := database.Get()
	// if err != nil {
	// 	server.TemplShow(pages.ServerError(), w, r, http.StatusInternalServerError)
	// 	return
	// }

	// queries := models.New(db)
	// user, err := queries.Login(context.Background(), models.LoginParams{
	// 	Email:    email,
	// 	Password: password,
	// })

	// if err != nil {
	// 	l.Logger.Error(err.Error())
	// 	return
	// }

	// server.TemplShow(pages.Index(pages.IndexProps{
	// 	Partial: "login",
	// }), w, r, http.StatusOK)

	l.Logger.Debug("Logged out")

	LoginWidget(w, r)
}

package main

import (
	"awesomeProject/db"
	"awesomeProject/internal/handler"
	"awesomeProject/internal/router"
	"awesomeProject/internal/store"
	"log"
)

func main() {

	app := router.New()
	//app.Get("/swagger/*", swagger.HandlerDefault)

	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)
	as := store.NewArticleStore(d)

	h := handler.NewHandler(us, as)
	h.RegisterRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

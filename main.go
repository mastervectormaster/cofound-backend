package main

import (
	"fmt"

	"github.com/gofiber/swagger"
	"github.com/mastervectormaster/cofound-backend/db"
	"github.com/mastervectormaster/cofound-backend/handler"
	"github.com/mastervectormaster/cofound-backend/router"
	"github.com/mastervectormaster/cofound-backend/store"
)

func main() {
	r := router.New()
	r.Get("/swagger/*", swagger.HandlerDefault)
	d := db.New()
	db.AutoMigrate(d)

	us := store.NewUserStore(d)

	h := handler.NewHandler(us)
	h.Register(r)
	err := r.Listen(":8585")
	if err != nil {
		fmt.Printf("%v", err)
	}
}

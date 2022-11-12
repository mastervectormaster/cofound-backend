package main

import (
	"fmt"

	"github.com/gofiber/swagger"
	"github.com/mastervectormaster/fiber-realworld/router"
)

func main() {
	r := router.New()
	r.Get("/swagger/*", swagger.HandlerDefault)
	// d := db.New()
	// db.AutoMigrate(d)

	// us := store.NewUserStore(d)
	// as := store.NewArticleStore(d)

	// h := handler.NewHandler(us, as)
	// h.Register(r)
	err := r.Listen(":8585")
	if err != nil {
		fmt.Printf("%v", err)
	}
}

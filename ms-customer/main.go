package main

import (
	"ms-customer/db"
	"ms-customer/service"

	"github.com/gin-gonic/gin"
)

func main() {

	secrets := db.GetSecretValue()
	db := db.GetDB(secrets)
	app := service.GetApp(db)
	r:= gin.Default()

	r.POST("/customers",app.PostHandler)
	r.GET("/customers/:customerId", app.PutHandler)
	r.PUT("/customers/:customerId", app.PutHandler)
	r.DELETE("/customers/:customerId", app.DeleteHandler)

	r.Run("localhost:8080")
}
package service

import (
	"ms-customer/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	db *db.PostgresDB	
}

func GetApp(db *db.PostgresDB) *App {
	return &App{
		db: db,
	}
}

func (app *App) PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
func (app *App) GetHandler(c *gin.Context) {}
func (app *App) PutHandler(c *gin.Context) {}
func (app *App) DeleteHandler(c *gin.Context) {}
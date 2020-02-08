package router

import (
	"net/http"

	h "github.com/ShiinaOrez/MarxProjectBackend/handler/history"
	"github.com/ShiinaOrez/MarxProjectBackend/handler/sd"
	s "github.com/ShiinaOrez/MarxProjectBackend/handler/study"
	"github.com/ShiinaOrez/MarxProjectBackend/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The history handlers, requiring authentication
	history := g.Group("/api/history")
	{
		history.GET("/news", middleware.Page, h.HistoryNews)
		history.GET("/new/:id", middleware.Id, h.HistoryNew)
	}

	// The study handlers, requiring authentication
	study := g.Group("/api/study")
	{
		study.GET("/news", middleware.Page, s.StudyNews)
		study.GET("/new/:id", middleware.Id, s.StudyNew)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

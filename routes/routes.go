package routes

import "github.com/gin-gonic/gin"

func RouteInit(r *gin.RouterGroup) {
	AuthRoutes(r)
	UserRoutes(r)
	CountryRoutes(r)
	TripRoute(r)
	TransacRoute(r)
}

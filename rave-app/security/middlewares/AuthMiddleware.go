package middlewares

import (
	control "github.com/djfemz/rave/rave-app/controllers"
	"github.com/djfemz/rave/rave-app/security"
	"github.com/djfemz/rave/rave-app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Routers(router *gin.Engine) {
	organizerController := control.NewOrganizerController()
	eventController := control.NewEventController()
	protected := router.Group("/protected", AuthMiddleware())
	{
		protected.POST("/event", organizerController.CreateEvent)
		protected.PUT("/event", eventController.EditEvent)
		protected.POST("/event-staff", organizerController.AddEventStaff)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(utils.AUTHORIZATION)
		token := extractTokenFrom(authHeader)
		org, err := security.ExtractUserFrom(token)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusForbidden, err)
			return
		}
		if org != nil {
			ctx.Next()
		}
	}
}

func extractTokenFrom(authHeader string) string {
	authValue := strings.Split(authHeader, " ")
	token := authValue[len(authValue)-1]
	return token
}

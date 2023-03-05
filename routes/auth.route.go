package routes

import ( 
	"github.com/gin-gonic/gin"
	"github.com/Kikozai/Moon/controllers"
	db "github.com/Kikozai/Moon/db/sqlc"
	"github.com/Kikozai/Moon/middleware"
)

type AuthRoutes struct { 
	authController controllers.AuthController
	db *db.Queries
}

func NewAuthRoutes(authController controllers.AuthController, db *db.Queries) AuthRoutes { 
	return AuthRoutes{authController,db}
}

func(rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(rc.db), rc.authController.LogoutUser)
}
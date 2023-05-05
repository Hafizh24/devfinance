package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hafizh24/devfinance/internal/app/controller"
	"github.com/hafizh24/devfinance/internal/app/repository"
	"github.com/hafizh24/devfinance/internal/app/service"
	"github.com/hafizh24/devfinance/internal/pkg/config"
	"github.com/hafizh24/devfinance/internal/pkg/db"
	"github.com/hafizh24/devfinance/internal/pkg/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	cfg      config.Config
	DBConn   *sqlx.DB
	enforcer *casbin.Enforcer
)

func init() {

	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic("db not established")
	}
	DBConn = db

	// Setup logrus
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // appyly log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json

	// setup casbin
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		panic("cannot load app casbin enforcer")
	}
	enforcer = e

}

func main() {

	r := gin.New()

	// implement middleware
	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
		cors.Default(),
	)

	// ---------------------------------------------------------------------------------------
	tokenMaker := service.NewTokenMaker(
		cfg.AccessTokenKey,
		cfg.RefreshTokenKey,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
	)

	categoryRepository := repository.NewCategoryRepository(DBConn)
	registrationRepository := repository.NewUserRepository(DBConn)
	userRepository := repository.NewUserRepository(DBConn)
	authRepository := repository.NewAuthRepository(DBConn)
	currencyRepository := repository.NewCurrencyRepository(DBConn)

	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(registrationRepository)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)
	currencyService := service.NewCurrencyService(currencyRepository)

	categoryController := controller.NewCategoryController(categoryService)
	registrationController := controller.NewRegistrationController(registrationService)
	sessionController := controller.NewSessionController(sessionService, tokenMaker)
	currencyController := controller.NewCurrencyController(currencyService)

	// Entrypoint

	route := r.Group("/api/auth")
	{
		route.POST("/signup", registrationController.Register)
		route.POST("/signin", sessionController.Login)
		route.GET("/refresh", sessionController.Refresh)
	}

	secured := r.Group("/api").Use(middleware.AuthMiddleware(tokenMaker))
	{
		secured.GET("/auth/signout", sessionController.Logout)
		secured.GET("/auth/showprofile", sessionController.ShowProfile)

		secured.GET("/categories", categoryController.BrowseCategory)
		secured.GET("/categories/:id", categoryController.DetailCategory)
		secured.POST("/categories", categoryController.CreateCategory)
		secured.PATCH("/categories/:id", categoryController.UpdateCategory)
		secured.DELETE("/categories/:id", categoryController.DeleteCategory)

		secured.GET("/currencies", currencyController.BrowseCurrency)
		secured.GET("/currencies/:id", currencyController.DetailCurrency)
		secured.POST("/currencies", currencyController.CreateCurrency)
		secured.PATCH("/currencies/:id", currencyController.UpdateCurrency)
		secured.DELETE("/currencies/:id", currencyController.DeleteCurrency)
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	// nolint:errcheck
	r.Run(appPort)
}

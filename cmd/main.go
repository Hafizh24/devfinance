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
	transactionRepository := repository.NewTransactionRepository(DBConn)

	categoryService := service.NewCategoryService(categoryRepository)
	registrationService := service.NewRegistrationService(registrationRepository)
	sessionService := service.NewSessionService(userRepository, authRepository, tokenMaker)
	currencyService := service.NewCurrencyService(currencyRepository)
	transactionService := service.NewTransactionService(transactionRepository, authRepository)

	categoryController := controller.NewCategoryController(categoryService)
	registrationController := controller.NewRegistrationController(registrationService)
	sessionController := controller.NewSessionController(sessionService, tokenMaker)
	currencyController := controller.NewCurrencyController(currencyService)
	transactionController := controller.NewTransactionController(transactionService)

	// Entrypoint

	route := r.Group("/api/auth")
	{
		route.POST("/signup", registrationController.Register)
		route.POST("/signin", sessionController.Login)
		route.GET("/refresh", sessionController.Refresh)
	}

	secured := r.Group("/api").Use(middleware.AuthMiddleware(tokenMaker))
	{
		secured.GET("/auth/signout", middleware.AuthorizationMiddleware("auth", "read", enforcer), sessionController.Logout)
		secured.GET("/auth/showprofile", middleware.AuthorizationMiddleware("auth", "read", enforcer), sessionController.ShowProfile)
		secured.DELETE("/auth/delete/:id", middleware.AuthorizationMiddleware("auth", "write", enforcer), registrationController.DeleteUser)

		secured.GET("/categories", middleware.AuthorizationMiddleware("categories", "read", enforcer), categoryController.BrowseCategory)
		secured.GET("/categories/:id", middleware.AuthorizationMiddleware("categories", "read", enforcer), categoryController.DetailCategory)
		secured.POST("/categories", middleware.AuthorizationMiddleware("categories", "write", enforcer), categoryController.CreateCategory)
		secured.PATCH("/categories/:id", middleware.AuthorizationMiddleware("categories", "write", enforcer), categoryController.UpdateCategory)
		secured.DELETE("/categories/:id", middleware.AuthorizationMiddleware("categories", "write", enforcer), categoryController.DeleteCategory)

		secured.GET("/currencies", middleware.AuthorizationMiddleware("currencies", "read", enforcer), currencyController.BrowseCurrency)
		secured.GET("/currencies/:id", middleware.AuthorizationMiddleware("currencies", "read", enforcer), currencyController.DetailCurrency)
		secured.POST("/currencies", middleware.AuthorizationMiddleware("currencies", "write", enforcer), currencyController.CreateCurrency)
		secured.PATCH("/currencies/:id", middleware.AuthorizationMiddleware("currencies", "write", enforcer), currencyController.UpdateCurrency)
		secured.DELETE("/currencies/:id", middleware.AuthorizationMiddleware("currencies", "write", enforcer), currencyController.DeleteCurrency)

		secured.GET("/transactions/show", middleware.AuthorizationMiddleware("transactions", "read", enforcer), transactionController.ShowRecaps)
		secured.GET("/transactions/show/:type", middleware.AuthorizationMiddleware("transactions", "read", enforcer), transactionController.ShowByType)
		secured.POST("/transactions", middleware.AuthorizationMiddleware("transactions", "write", enforcer), transactionController.CreateTransaction)
		secured.DELETE("/transactions/:id", middleware.AuthorizationMiddleware("transactions", "write", enforcer), transactionController.DeleteTransaction)
		secured.PATCH("/transactions/:id", middleware.AuthorizationMiddleware("transactions", "write", enforcer), transactionController.UpdateTransaction)
	}

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	// nolint:errcheck
	r.Run(appPort)
}

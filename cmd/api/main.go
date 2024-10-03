package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sakaguchi-0725/echo-onion-arch/application/usecase"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/db"
	"github.com/sakaguchi-0725/echo-onion-arch/infra/persistence"
	"github.com/sakaguchi-0725/echo-onion-arch/pkg/config"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/handler"
	"github.com/sakaguchi-0725/echo-onion-arch/presentation/api/router"
)

func main() {
	e := echo.New()

	cfg := config.NewConfig()
	db, err := db.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

	userRepo := persistence.NewUserRepository(db)
	profileRepo := persistence.NewProfileRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, profileRepo)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)

	authHandler := handler.NewAuthHandler(authUsecase, cfg.App)
	profileHandler := handler.NewProfileHandler(profileUsecase)

	deps := &router.HandlerDependencies{
		AuthHandler:    authHandler,
		ProfileHandler: profileHandler,
	}

	router.NewRouter(e, deps, *cfg.App)
	e.Start(":8080")
}

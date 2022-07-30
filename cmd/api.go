package cmd

import (
	"context"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/ybalcin/user-management/docs"
	"github.com/ybalcin/user-management/internal/api"
	"github.com/ybalcin/user-management/internal/application"
	"github.com/ybalcin/user-management/internal/config"
	"github.com/ybalcin/user-management/internal/infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func RunApi() {
	cfg, e := config.Read()
	if e != nil {
		panic(e)
	}

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(cfg.DatabaseSettings.DatabaseURI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	if e = cli.Ping(ctx, nil); e != nil {
		panic(e)
	}

	db := cli.Database(cfg.DatabaseSettings.DatabaseName)

	userRepository := infrastructure.NewUserMongoRepository(db)
	userService := application.NewUserManagementService(userRepository)

	a := api.New(userService)

	a.App().Get("/swagger/*", fiberSwagger.WrapHandler)

	if e := a.Listen(); e != nil {
		panic(e)
	}
}

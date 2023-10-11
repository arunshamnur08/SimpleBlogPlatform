package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Server struct {
	Ctx         context.Context
	Log         log.Logger
	Mongoclient *mongo.Client
	GinEngine   *gin.Engine
}

func NewServer(ctx context.Context, logger log.Logger, mongoclient *mongo.Client, engine *gin.Engine) *Server {
	return &Server{
		Ctx:         ctx,
		Log:         logger,
		Mongoclient: mongoclient,
		GinEngine:   engine,
	}
}

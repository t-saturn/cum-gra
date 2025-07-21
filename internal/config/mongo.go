package config

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

// ConnectMongo establece la conexión con MongoDB y asigna los valores globales.
func ConnectMongo() {
	cfg := GetConfig() // Carga la configuración global
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.Mongo.URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		logger.Log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verifica la conexión con un ping
	if err := client.Ping(ctx, nil); err != nil {
		logger.Log.Fatalf("MongoDB no responde al ping: %v", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(cfg.Mongo.DBName)

	logger.Log.Infof("Conectado a MongoDB en %s", cfg.Mongo.URI)
}

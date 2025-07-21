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

// ConnectMongo establece la conexión con MongoDB y asigna los valores a MongoClient y MongoDatabase.
func ConnectMongo() {
	cfg := GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MONGO_URI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		logger.Log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verificar la conexión
	if err := client.Ping(ctx, nil); err != nil {
		logger.Log.Fatalf("MongoDB no responde al ping: %v", err)
	}

	MongoClient = client
	MongoDatabase = client.Database(cfg.MONGO_DB_NAME)

	logger.Log.Infof("Conectado a MongoDB exitosamente en %s", cfg.MONGO_URI)
}

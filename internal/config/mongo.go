package config

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

// ConnectMongo establece la conexión a la base de datos MongoDB utilizando los datos de configuración.
func ConnectMongo() {
	cfg := GetConfig()

	// contexto con timeout solo para conectar/ping
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Usa exactamente lo que viene en MONGO_URI (NO SetHosts manual)
	clientOptions := options.Client().ApplyURI(cfg.Mongo.URI).
		SetMaxPoolSize(10).
		SetMinPoolSize(0).
		SetMaxConnIdleTime(30 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Log.Fatalf("Error al conectar con MongoDB: %v", err)
	}

	// ping
	if err := client.Ping(ctx, nil); err != nil {
		logger.Log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	MongoDB = client.Database(cfg.Mongo.DBName)
	logger.Log.Infof("Conexión exitosa establecida a MongoDB en la base de datos: %s", cfg.Mongo.DBName)
}

// DisconnectMongo cierra la conexión a MongoDB de forma segura
func DisconnectMongo() {
	if MongoDB != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := MongoDB.Client().Disconnect(ctx); err != nil {
			logger.Log.Errorf("Error al desconectar MongoDB: %v", err)
		} else {
			logger.Log.Info("Desconexión exitosa de MongoDB")
		}
	}
}

// GetMongoDB retorna la instancia de la base de datos MongoDB
func GetMongoDB() *mongo.Database {
	if MongoDB == nil {
		logger.Log.Fatal("MongoDB no está conectado")
	}
	return MongoDB
}

// GetMongoCollection retorna una colección específica de MongoDB
func GetMongoCollection(collectionName string) *mongo.Collection {
	if MongoDB == nil {
		logger.Log.Fatal("MongoDB no está conectado")
	}
	return MongoDB.Collection(collectionName)
}

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
	cfg := GetConfig() // Obtiene la configuración global

	// Crear contexto con timeout para la conexión
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configurar opciones del cliente
	clientOptions := options.Client().ApplyURI(cfg.Mongo.URI)

	// Forzar conexión IPv4 para evitar problemas con localhost
	clientOptions.SetHosts([]string{"127.0.0.1:27017"})

	// Configurar el pool de conexiones (opcional)
	clientOptions.SetMaxPoolSize(10)
	clientOptions.SetMinPoolSize(5)
	clientOptions.SetMaxConnIdleTime(30 * time.Second)

	// Conectar al cliente de MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Log.Fatalf("Error al conectar con MongoDB: %v", err)
	}

	// Verificar la conexión con ping
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Log.Fatalf("Error al hacer ping a MongoDB: %v", err)
	}

	// Asignar la base de datos
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

// GetMongoCollection retorna una colección específica de MongoDB
func GetMongoCollection(collectionName string) *mongo.Collection {
	if MongoDB == nil {
		logger.Log.Fatal("MongoDB no está conectado")
	}
	return MongoDB.Collection(collectionName)
}

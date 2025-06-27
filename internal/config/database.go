package config

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

// DatabaseConfig representa las conexiones compartidas con la base de datos
type DatabaseConfig struct {
	Mongo *mongo.Database
}

func InitMongoDB() {
	uri := GetEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := GetEnv("MONGO_DB_NAME", "mongo_db")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		Logger.WithError(err).Fatal("Error connecting to MongoDB")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		Logger.WithError(err).Fatal("MongoDB not responding to ping")
	}

	Logger.WithFields(logrus.Fields{
		"uri":      uri,
		"database": dbName,
	}).Info("Successful connection to MongoDB")

	MongoClient = client
	MongoDatabase = client.Database(dbName)
}

func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Mongo: MongoDatabase,
	}
}

package repository

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Nombre de la colección en Mongo
const userCredentialsCollection = "user_credentials"

// Obtener colección
func getUserCredentialsCollection() *mongo.Collection {
	return config.MongoDatabase.Collection(userCredentialsCollection)
}

// Insertar una nueva credencial
func CreateUserCredential(ctx context.Context, credential *models.UserCredential) error {
	collection := getUserCredentialsCollection()

	_, err := collection.InsertOne(ctx, credential)
	if err != nil {
		config.Logger.WithError(err).Error("Error al insertar user_credential")
		return err
	}

	config.Logger.WithField("email", credential.Email).Info("✅ Credencial insertada")
	return nil
}

// Obtener todas las credenciales
func GetAllUserCredentials(ctx context.Context) ([]*models.UserCredential, error) {
	collection := getUserCredentialsCollection()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		config.Logger.WithError(err).Error("Error al obtener credenciales")
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*models.UserCredential
	for cursor.Next(ctx) {
		var credential models.UserCredential
		if err := cursor.Decode(&credential); err != nil {
			config.Logger.WithError(err).Warn("Error decodificando documento")
			continue
		}
		results = append(results, &credential)
	}

	return results, nil
}

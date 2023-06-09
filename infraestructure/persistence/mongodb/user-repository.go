package mongodb

import (
	"context"
	"time"

	"github.com/fepc18/go-users-cleanarchitecture/core/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) models.IUserRepository {

	return &userRepository{
		collection: db.Collection("users"),
	}
}

func (r *userRepository) Create(user *models.User) (*models.User, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancela el contexto cuando la función se termine de ejecutar

	//u.Password, _ = EncryptPassword(u.Password)

	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return user, false, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, true, nil

}

func (r *userRepository) CheckUserExist(email string) (models.User, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancela el contexto cuando la función se termine de ejecutar

	condition := bson.M{"email": email}
	var resultado models.User
	err := r.collection.FindOne(ctx, condition).Decode(&resultado) // decode json the result in the variable resultado

	if err != nil {
		return resultado, false, err
	}

	//ID := resultado.ID.Hex()
	return resultado, true, nil

}

package contact

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"srv_order/main/api/presenter"
	"srv_order/main/pkg/entities"
)

type Repository interface {
	CreateOrder(book *entities.Order) (*entities.Order, error)
	ReadOrder() (*[]presenter.Order, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateOrder(contact *entities.Order) (*entities.Order, error) {
	contact.ID = primitive.NewObjectID()
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), contact)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (r *repository) ReadOrder() (*[]presenter.Order, error) {
	var contacts []presenter.Order
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var contact presenter.Order
		_ = cursor.Decode(&contact)
		contacts = append(contacts, contact)
	}
	return &contacts, nil
}

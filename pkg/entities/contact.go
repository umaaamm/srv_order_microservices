package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Nama      string             `json:"nama" bson:"nama"`
	NoHp      string             `json:"nohp" bson:"nohp,omitempty"`
	Order     string             `json:"order" bson:"order"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

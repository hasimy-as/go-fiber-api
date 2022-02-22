package models

import (
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UserId    uuid.UUID          `json:"userId,omitempty" bson:"userId"`
	Name      string             `json:"name,omitempty" validate:"required" bson:"name"`
	Address   string             `json:"address,omitempty" validate:"required" bson:"address"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manga struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required"`
	Character string             `json:"character,omitempty" validate:"required"`
	Type      string             `json:"type,omitempty" validate:"required"`
}

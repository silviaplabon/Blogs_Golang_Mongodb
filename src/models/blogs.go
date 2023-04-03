package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	Email     string             `json:"email"`
	Title     string             `json:"title"`
	Subtitle  string             `json:"subtitle"`
	Details   string             `json:"details"`
	MinAge    int                `json:"minAge"`
	MaxAge    int                `json:"maxAge"`
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Interests []string           `json:"interests"`
}


package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DisplayName        string             `bson:"display_name,omitempty" json:"display_name,omitempty"`
	Password           string             `bson:"password,omitempty" json:"password,omitempty"`
	Country            string             `bson:"country,omitempty" json:"country,omitempty"`
	Points             int                `bson:"points,omitempty" json:"points,omitempty"`
	Rank               int                `bson:"rank,omitempty" json:"rank"`
	LastScoreTimestamp string             `bson:"last_score_timestamp" json:"last_score_timestamp"`
}

type UserScore struct {
	ID                 primitive.ObjectID `json:"user_id" validate:"required"`
	Score              int                `json:"score_worth" validate:"required"`
	LastScoreTimestamp string             `json:"timestamp,omitempty"`
}

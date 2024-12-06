package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
    ID          primitive.ObjectID `json:"id" bson:"_id"`
    Name        string             `json:"name" bson:"name"`
    Ingredients []Ingredient       `json:"ingredients" bson:"ingredients"`
    CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

type Ingredient struct {
    Name  string             `json:"name" bson:"name"`
    Type  string             `json:"type" bson:"type"`
}
package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection
var MongoClient *mongo.Client
var Ctx = context.TODO()
var BlogsCollection *mongo.Collection


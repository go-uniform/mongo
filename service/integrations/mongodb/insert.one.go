package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (s *mongodb) InsertOne(database, collection string, document bson.M) (*mongo.InsertOneResult, error) {
	c, cancel := s.connect(database, collection)
	defer cancel()

	delete(document, "_id")
	document["createdAt"] = time.Now()
	document["modifiedAt"] = time.Now()

	return c.InsertOne(context.Background(), document)
}

package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (s *mongodb) UpdateOne(database, collection string, query bson.D, document bson.M) (*mongo.UpdateResult, error) {
	c, cancel := s.connect(database, collection)
	defer cancel()

	document["modifiedAt"] = time.Now()

	return c.UpdateOne(context.Background(), query, document)
}

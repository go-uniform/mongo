package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *mongodb) DeleteOne(database, collection string, query bson.D) (*mongo.DeleteResult, error) {
	c, cancel := s.connect(database, collection)
	defer cancel()

	return c.DeleteOne(context.Background(), query)
}

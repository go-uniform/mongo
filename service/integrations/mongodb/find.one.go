package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *mongodb) FindOne(database, collection string, options *options.FindOneOptions, query bson.D) *mongo.SingleResult {
	c, cancel := s.connect(database, collection)
	defer cancel()

	return c.FindOne(context.Background(), query, options)
}

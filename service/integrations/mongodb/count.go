package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *mongodb) Count(database, collection string, query bson.D) (int64, error) {
	c, cancel := s.connect(database, collection)
	defer cancel()

	return c.CountDocuments(context.Background(), query)
}

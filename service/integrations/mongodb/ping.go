package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (s *mongodb) ping() (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	if s.Client == nil {
		return false
	}

	if err := s.Client.Ping(context.Background(), readpref.Primary()); err != nil {
		return false
	}
	return true
}
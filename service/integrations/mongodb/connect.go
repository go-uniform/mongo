package mongodb

import (
	"context"
	"crypto/tls"
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
)

// make the connect routine thread safe
var connectLock = sync.Mutex{}

// connect to the given database/collection (pooling connection)
func (s *mongodb) connect(database, collection string) (c *mongo.Collection, close context.CancelFunc) {
	if err := s.Page.Scope("connect", func(p diary.IPage) {
		connectLock.Lock()
		defer connectLock.Unlock()

		if s.Client == nil || s.ClientExpiresAt.Before(time.Now()) {
			if !s.ping() {
				s.Page.Notice("reconnect", diary.M{
					"uri": s.Uri,
					"expires-at": s.ClientExpiresAt,
				})

				clientOptions := options.Client().ApplyURI(s.Uri)

				/* Security Credentials */
				if s.Username != "" {
					clientOptions.SetAuth(options.Credential{
						AuthSource: s.AuthSource,
						Username:   s.Username,
						Password:   s.Password,
					})
				}

				/* Communication Encryption */
				if s.CertFile != "" || s.KeyFile != "" {
					cert, err := tls.LoadX509KeyPair(s.CertFile, s.KeyFile)
					if err != nil {
						panic(err)
					}
					clientOptions.SetTLSConfig(&tls.Config{
						InsecureSkipVerify: s.Verify,
						Certificates: 	[]tls.Certificate{cert},
						MinVersion: 	tls.VersionTLS12,
					})
				}

				// create new client connection
				client, err := mongo.NewClient(clientOptions)
				if err != nil {
					panic(err)
				}
				ctx, cancel := context.WithTimeout(context.Background(), s.ConnectionTimeout)

				// connect the new client
				if err := client.Connect(ctx); err != nil {
					cancel()
					panic(err)
				}

				// test the new client
				if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
					cancel()
					panic(err)
				}

				// disconnect old client connection
				if s.Client != nil {
					func() {
						defer func() {
							// ignore disconnect issues
							_ = recover()
						}()
						_ = s.Client.Disconnect(context.Background())
					}()
				}

				// replace existing client with new client
				s.Client = client
			}
		}

		// open interface to given database/collection
		c = s.Client.Database(database).Collection(collection)
		close = func() {
			// keep connection open since we are connection pooling
		}

		s.Page.Info("collection", diary.M{
			"database": database,
			"collection": collection,
		})

		// add client expire interval to client expiry time since connection has just had some activity
		s.ClientExpiresAt = time.Now().Add(s.ClientExpireInterval)
	}); err != nil {
		panic(err)
	}

	return
}

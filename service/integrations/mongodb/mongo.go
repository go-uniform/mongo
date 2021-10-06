package mongodb

import (
	"github.com/go-diary/diary"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/* MongoDB

Reference:
- https://docs.mongodb.com/manual/introduction/
- https://docs.mongodb.com/drivers/go/current/#installation
- https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.7.2/mongo
*/

type mongodb struct {
	Client *mongo.Client
	ClientExpiresAt time.Time
	Page diary.IPage
	ClientExpireInterval time.Duration
	ConnectionTimeout time.Duration

	/* Target */
	Uri string

	/* Security Credentials */
	AuthSource string
	Username string
	Password string

	/* Communication Encryption */
	CertFile string
	KeyFile string
	Verify bool
}

type IMongoDb interface {
	Count(database, collection string, query bson.D) (int64, error)
	FindOne(database, collection string, options *options.FindOneOptions, query bson.D) *mongo.SingleResult
	InsertOne(database, collection string, document bson.M) (*mongo.InsertOneResult, error)
	UpdateOne(database, collection string, query bson.D, document bson.M) (*mongo.UpdateResult, error)
	DeleteOne(database, collection string, query bson.D) (*mongo.DeleteResult, error)
}

func NewMongoConnector(page diary.IPage, uri, authSource, username, password, certFile, keyFile string, verify bool) IMongoDb {
	var instance IMongoDb
	page.Scope("mongo", func(p diary.IPage) {
		instance = &mongodb{
			Page: p,
			ClientExpireInterval: time.Second * 30,
			ConnectionTimeout: time.Minute,

			Uri: uri,
			AuthSource: authSource,
			Username: username,
			Password: password,
			CertFile: certFile,
			KeyFile: keyFile,
			Verify: verify,
		}
	})
	return instance
}
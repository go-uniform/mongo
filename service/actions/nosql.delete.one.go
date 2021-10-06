package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/nosql"
	"go.mongodb.org/mongo-driver/bson"
	"service/service/_base"
	"service/service/info"
	"time"
)

func init() {
	_base.Subscribe(_base.TargetAction("nosql", "delete.one"), func(r uniform.IRequest, p diary.IPage) {
		var model nosql.DeleteOneRequest
		r.Read(&model)

		var identifier interface{}
		if model.SoftDelete {
			result, err := info.Mongo.UpdateOne(model.Database, model.Collection, model.Query, bson.M{
				"deletedAt": time.Now(),
			})
			if err != nil {
				panic(err)
			}
			identifier = result.UpsertedID
		} else {
			result := info.Mongo.FindOne(model.Database, model.Collection, nil, model.Query)
			var document bson.M
			if err := result.Decode(&document); err != nil {
				panic(err)
			}
			identifier = document["_id"]

			_, err := info.Mongo.DeleteOne(model.Database, model.Collection, bson.D{
				{"_id", identifier},
			})
			if err != nil {
				panic(err)
			}
		}

		document := findOne(r, p, nosql.FindOneRequest{
			Database: model.Database,
			Collection: model.Collection,
			Query: bson.D{
				{ "_id", identifier },
			},
		})

		if r.CanReply() {
			if err := r.Reply(uniform.Request{
				Model: document,
			}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
	})
}
package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/nosql"
	"go.mongodb.org/mongo-driver/bson"
	"service/service/_base"
	"service/service/events"
	"service/service/info"
)

func init() {
	_base.Subscribe(_base.TargetAction("nosql", "insert.one"), func(r uniform.IRequest, p diary.IPage) {
		var model struct {
			Database string
			Collection string
			Document bson.M
		}
		r.Read(&model)

		model.Document = events.EntityValidate(r, p, model.Document, model.Database, model.Collection)
		events.EntityConstraints(r, p, "", model.Document, model.Database, model.Collection)

		result, err := info.Mongo.InsertOne(model.Database, model.Collection, model.Document)

		if err != nil {
			panic(err)
		}

		document := findOne(r, p, nosql.FindOneRequest{
			Database: model.Database,
			Collection: model.Collection,
			Query: bson.D{
				{ "_id", result.InsertedID },
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
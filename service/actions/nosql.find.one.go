package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/common/nosql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"service/service/_base"
	"service/service/events"
	"service/service/info"
)

var findOne = func(r uniform.IRequest, p diary.IPage, model nosql.FindOneRequest) bson.M {
	options := createFindOneOptions(model.Sort, model.Skip)
	if model.Query == nil {
		model.Query = bson.D{}
	}

	result := info.Mongo.FindOne(model.Database, model.Collection, options, model.Query)

	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(nosql.ErrNoResults)
		}
		panic(result.Err())
	}

	var document bson.M
	if err := result.Decode(&document); err != nil {
		panic(err)
	}

	// request logic service to decrypt the encrypted fields of the document
	document = events.EntityDecrypt(r, p, document, model.Database, model.Collection)

	return document
}

func init() {
	_base.Subscribe(_base.TargetAction("nosql", "find.one"), func(r uniform.IRequest, p diary.IPage) {
		var model nosql.FindOneRequest
		r.Read(&model)

		document := findOne(r, p, model)

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

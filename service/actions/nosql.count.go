package actions

import (
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"github.com/go-uniform/uniform/nosql"
	"go.mongodb.org/mongo-driver/bson"
	"service/service/_base"
	"service/service/info"
)

func init() {
	_base.Subscribe(_base.TargetAction("nosql", "count"), func(r uniform.IRequest, p diary.IPage) {
		var model nosql.FindOneRequest
		r.Read(&model)

		if model.Query == nil {
			model.Query = bson.D{}
		}
		count, err := info.Mongo.Count(model.Database, model.Collection, model.Query)
		if err != nil {
			panic(err)
		}

		if r.CanReply() {
			if err := r.Reply(uniform.Request{
				Model: count,
			}); err != nil {
				p.Error("reply", err.Error(), diary.M{
					"err": err,
				})
			}
		}
	})
}
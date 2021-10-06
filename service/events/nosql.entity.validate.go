package events

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"service/service/_base"
)

func EntityValidate(r uniform.IRequest, p diary.IPage, document interface{}, database, collection string) (model bson.M) {
	if err := p.Scope("validate", func(p diary.IPage) {
		if err := r.Conn().Request(p, _base.TargetEvent("entity", fmt.Sprintf("%s.validate", database)), r.Remainder(), uniform.Request{
			Parameters: uniform.P{
				"source": "mongodb",
				"database": database,
				"collection": collection,
			},
			Model: document,
		}, func(r uniform.IRequest, p diary.IPage) {
			if r.HasError() {
				panic(r.Error())
			}
			r.Read(&model)
		}); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
	return
}
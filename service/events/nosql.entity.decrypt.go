package events

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"go.mongodb.org/mongo-driver/bson"
	"service/service/_base"
)

func EntityDecrypt(r uniform.IRequest, p diary.IPage, document bson.M, database, collection string) (model bson.M) {
	if err := p.Scope("decrypt", func(p diary.IPage) {
		if err := r.Conn().Request(p, _base.TargetEvent("entity", fmt.Sprintf("%s.decrypt", database)), r.Remainder(), uniform.Request{
			Parameters: uniform.P{
				"source": "mongodb",
				"database": database,
				"collection": collection,
			},
			Model: document,
		}, func(r uniform.IRequest, p diary.IPage) {
			r.Read(&model)
		}); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
	return
}
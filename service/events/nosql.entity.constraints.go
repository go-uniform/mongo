package events

import (
	"fmt"
	"github.com/go-diary/diary"
	"github.com/go-uniform/uniform"
	"service/service/_base"
)

func EntityConstraints(r uniform.IRequest, p diary.IPage, identifier string, document interface{}, database, collection string) {
	if err := p.Scope("constraints", func(p diary.IPage) {
		if err := r.Conn().Request(p, _base.TargetEvent("entity", fmt.Sprintf("%s.constraints", database)), r.Remainder(), uniform.Request{
			Parameters: uniform.P{
				"source": "mongodb",
				"database": database,
				"collection": collection,
				"identifier": identifier,
			},
			Model: document,
		}, func(r uniform.IRequest, p diary.IPage) {
			if r.HasError() {
				panic(r.Error())
			}
		}); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
	return
}
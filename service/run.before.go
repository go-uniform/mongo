package service

import (
	"github.com/go-diary/diary"
	"service/service/info"
	"service/service/integrations/mongodb"
	"sync"
)

func RunBefore(shutdown chan bool, group *sync.WaitGroup, p diary.IPage) {
	uri, ok := info.Args["uri"].(string)
	if !ok {
		panic("mongo uri must be a string")
	}
	authSource, ok := info.Args["authSource"].(string)
	if !ok {
		panic("mongo authSource must be a string")
	}
	username, ok := info.Args["username"].(string)
	if !ok {
		panic("mongo username must be a string")
	}
	password, ok := info.Args["password"].(string)
	if !ok {
		panic("mongo password must be a string")
	}
	certFile, ok := info.Args["certFile"].(string)
	if !ok {
		panic("mongo certFile must be a string")
	}
	keyFile, ok := info.Args["keyFile"].(string)
	if !ok {
		panic("mongo keyFile must be a string")
	}
	verify, ok := info.Args["verify"].(bool)
	if !ok {
		panic("mongo verify must be a bool")
	}

	info.Uri = uri
	info.AuthSource = authSource
	info.Username = username
	info.Password = password
	info.CertFile = certFile
	info.KeyFile = keyFile
	info.Verify = verify

	info.Mongo = mongodb.NewMongoConnector(p, uri, authSource, username, password, certFile, keyFile, verify)
}
package info

import "service/service/integrations/mongodb"

var Uri string
var AuthSource string
var Username string
var Password string
var CertFile string
var KeyFile string
var Verify bool

var Mongo mongodb.IMongoDb
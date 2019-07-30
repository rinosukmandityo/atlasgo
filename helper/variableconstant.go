package helper

import (
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var (
	PathSeparator string = string(os.PathSeparator)
	WD, _                = os.Getwd()
)

type AppConfig struct {
	URI             string `bson:"uri" json:"uri"`
	PerformanceTest bool   `bson:"performancetest" json:"performancetest"`
	ShowResult      bool   `bson:"showresult" json:"showresult"`
	CollectionTest  string `bson:"collectiontest" json:"collectiontest"`
}

type Connection struct {
	ShowResult     bool
	Database       string
	CollectionTest *mgo.Collection
	session        *mgo.Session
}

type Person struct {
	Name      string
	Phone     string
	Timestamp time.Time
}

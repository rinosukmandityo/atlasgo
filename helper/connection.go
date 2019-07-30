package helper

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewConnection(uri string) (c *Connection, e error) {
	c = new(Connection)
	if e = c.Connect(uri); e != nil {
		log.Println(e.Error())
	}
	return
}

func (c *Connection) Connect(uri string) (e error) {
	dialInfo, e := mgo.ParseURL(uri)
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, e := mgo.DialWithInfo(dialInfo)
	c.session = session
	c.Database = dialInfo.Database

	return
}

func (c *Connection) GetCollections() (res []string) {
	if c.session == nil {
		return []string{}
	}

	mgoDb := c.session.DB(c.Database)

	res = []string{}

	cols, err := mgoDb.CollectionNames()
	if err != nil {
		return []string{}
	}

	for _, col := range cols {
		res = append(res, col)
	}

	return res
}

func (c *Connection) PerformanceTest(colname string) {
	t := time.Now()
	cols := c.GetCollections()
	log.Println("get collections duration", time.Since(t))
	if c.ShowResult {
		log.Println(cols)
	}

	c.CollectionTest = c.session.DB(c.Database).C(colname)
	c.DropCollection()
	t = time.Now()
	c.InsertData()
	log.Println("insert data duration", time.Since(t))
	t = time.Now()
	c.QueryOne()
	log.Println("query one data duration", time.Since(t))
	t = time.Now()
	c.QueryAll()
	log.Println("query all data duration", time.Since(t))
	t = time.Now()
	c.UpdateData()
	log.Println("update data duration", time.Since(t))
	t = time.Now()
	c.DeleteData()
	log.Println("delete data duration", time.Since(t))
	t = time.Now()
	c.QueryAll()
	log.Println("query all data duration", time.Since(t))
}

func (c *Connection) InsertData() (e error) {
	e = c.CollectionTest.Insert(&Person{Name: "Person 1", Phone: "+62 44 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Person 2", Phone: "+62 33 1234 5678", Timestamp: time.Now()})

	return
}

func (c *Connection) QueryOne() (e error) {
	result := Person{}
	e = c.CollectionTest.Find(bson.M{"name": "Person 1"}).One(&result)
	if e != nil {
		return
	}
	if c.ShowResult {
		log.Printf("Query One Result == %+v\n", result)
	}
	return
}
func (c *Connection) QueryAll() (e error) {
	var results []Person
	e = c.CollectionTest.Find(nil).Sort("-timestamp").All(&results)

	if e != nil {
		return
	}
	if c.ShowResult {
		log.Println("Results All: ", results)
	}
	return
}

func (c *Connection) UpdateData() (e error) {
	q := bson.M{"name": "Person 1"}
	set := bson.M{"$set": bson.M{"phone": "+62 77 8888 9999", "timestamp": time.Now()}}
	e = c.CollectionTest.Update(q, set)
	if e != nil {
		return
	}
	return
}

func (c *Connection) DeleteData() (e error) {
	selector := bson.M{"name": "Person 2"}
	e = c.CollectionTest.Remove(selector)
	if e != nil {
		return
	}
	return
}

func (c *Connection) DropCollection() (e error) {
	e = c.CollectionTest.DropCollection()
	if e != nil {
		return
	}
	return
}

func (c *Connection) Close() {
	if c.session != nil {
		c.session.Close()
	}
}

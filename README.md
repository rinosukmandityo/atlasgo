# Connect to MongoDB Atlas with Golang
To connect our apps into MongoDB Atlas with simple golang apps using mgo library and adding CRUD feature as a simple performance test

How to use
---
1. Change config file inside `configs/configs.json` with appropriate information about your MongoDB Atlas URI and other info
2. `uri` is connection string used to connect into your MongoDB Atlas database
3.	`performancetest` will do simple CRUD to your mongoDB Atlas database if you set this value as `true`
4.	`showresult` will show a result for each CRUD function if you set this value as `true`
5.	`collectiontest` is your collection name to test CRUD feature
6. To test it just run `go run main.go` or with file config specified `go run main.go -config yourconfigfilelocation`.

![Demo](static/demo.png)
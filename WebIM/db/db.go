// This program provides a sample application for using MongoDB with
// the mgo driver.
package db

import (
  "gopkg.in/mgo.v2"
  //"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	MongoDBHosts = "localhost:27017"
	AuthDatabase = "test"
	AuthUserName = ""
	AuthPassword = ""
	TestDatabase = "Test"
)

type (
	// BuoyCondition contains information for an individual station.
	Product struct {
		Name string
		Price float64
	}

  ProductSet []Product
)

// main is the entry point for the application.
func GetResult() int {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving forward, the data read will be consistent
	// across sequential queries in the same session, and modifications made
	// within the session will be observed in following queries (read-your-writes).
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode
	mongoSession.SetMode(mgo.Monotonic, true)

	// Create a wait group to manage the goroutines.
	//var waitGroup sync.WaitGroup

	// Perform 10 concurrent queries against the database.
	//waitGroup.Add(10)
  res := make(chan ProductSet)
	go RunQuery(mongoSession, res)

	// Wait for all the queries to complete.
	//waitGroup.Wait()

  result := <- res

  log.Println("Result")
  log.Println(result)

	log.Println("All Queries Completed")

  return len(result)
}

// RunQuery is a function that is launched as a goroutine to perform
// the MongoDB work.
func RunQuery(mongoSession *mgo.Session, res chan ProductSet) {
	// Decrement the wait group count so the program knows this
	// has been completed once the goroutine exits.
	//defer waitGroup.Done()

	// Request a socket connection from the session to process our query.
	// Close the session when the goroutine exits and put the connection back
	// into the pool.
	sessionCopy := mongoSession.Copy()
	defer sessionCopy.Close()

	// Get a collection to execute the query against.
	collection := sessionCopy.DB(TestDatabase).C("Products")

  /*erre := collection.Insert(&Product{"Test Prod2", 13.23},
	                          &Product{"Test Prod3", 12.12})
        if erre != nil {
                log.Fatal(erre)
        }*/

	log.Printf("RunQuery : Executing\n")

	// Retrieve the list of stations.
	var products []Product
	err := collection.Find(nil).All(&products)
	if err != nil {
		log.Printf("RunQuery : ERROR : %s\n", err)
		return
	}

	log.Printf("RunQuery:  Count[%d]\n",  len(products))

  res <- products;
}

package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"log"
	"math/rand"
	"time"
)

type EventModel struct {
	EventName		string
	Date			string
	Location		string
	Description		string
	CreatedBy		string
	EmailId			string
	ContactNo		string
	EmpId			string
	Time			string
	EventStatus		bool
	Comment			string
	EventType		string
	EventId			string
	EventComment	string
}
type EventCounter struct {

	Count		string
	EventId  string
}

func (event EventModel)InsertEvent()bool{

	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		fmt.Println("something happend")
		os.Exit(1)
	}
	defer sess.Close()
	fmt.Println("kkkkkk")
	sess.SetSafe(&mgo.Safe{})
	var r *rand.Rand

	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "012345abcdefghijklmnopqrstuvwxyz6789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}

	fmt.Println("gfghggghg", string(result))
	event.EventId="Ev"+ string(result)
	con := sess.DB("mirafrautilityapp").C("event")
	err = con.Insert(event)
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return false
	}
	//fmt.Println("haiiii")

	return true
}
func DisplayEventDetails()(bool,[] EventModel){
	var Event []EventModel
	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("mirafrautilityapp").C("event")


	err = collection.Find(bson.M{"eventstatus":false}).All(&Event)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Event

	}
	fmt.Println("event details struct",Event)
	return true,Event

}

func (event EventModel)EventUpdate()bool{

	uri := os.Getenv("MONGOLAB_URL")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	sess, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("mirafrautilityapp").C("event")


	//updating the event

	err = collection.Update(bson.M{"EventId":event.EventId }, bson.M{"$set": bson.M{"EventStatus": event.EventStatus,"EventComment":event.EventComment}})
	if err !=nil{
		return false
	}
	return true
}
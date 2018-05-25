package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"os"
	"log"
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

	sess.SetSafe(&mgo.Safe{})
	eventCounter :=EventCounter{}
	eventFind  :=EventCounter{}


	collection := sess.DB("mirafrautilityapp").C("eventCounter")
	err =collection.Find(nil).One(eventFind)
	fmt.Println("hhhhhh",eventFind.EventId)
	if len(eventFind.EventId) ==0{
		eventCounter.Count="0"
		eventCounter.EventId ="EventId"
		fmt.Println("event",eventCounter)
		err = collection.Insert(eventCounter)
		if err != nil {
			log.Fatal("Problem inserting data: ", err)
			return false
		}
	}


	//updating the sequence
	err =collection.Find(nil).One(eventCounter)
	i, err := strconv.Atoi(eventCounter.Count)
	counter :=i+1

	selector := bson.M{"eventId": "EventId"}
	updator := bson.M{"$set": bson.M{"count": counter}}
	if err := collection.Update(selector, updator); err != nil {
		fmt.Println("hhhhhgddgfdfdf",err)
	}





	event.EventId="event00"+ strconv.Itoa(counter)
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


	err = collection.Find(bson.M{"eventstatus": true}).All(&Event)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Event

	}
	fmt.Println("event details struct",Event)
	return true,Event

}

func (event EventModel)EventUpdate()bool{

	uri := os.Getenv("mongodb://aswathyashok:aswathy@ds133550.mlab.com:33550/mirafrautilityapp")
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
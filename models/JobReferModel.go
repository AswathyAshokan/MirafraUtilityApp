package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
)

type JobRefer struct {
	EmpId					string
	JobTitle				string
	CandidateName			string
	CandidatePhoneNumber	string
	Location				string
	Resume 					string
	Experience 				string
	JobId					string
	ReferStatus				bool

}

func(jobRefer JobRefer)InsertJobRrfer()bool{
 fmt.Println("ggggfffffffff")
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
	collection := sess.DB("mirafrautilityapp").C("jobRefer")

	if err := collection.Insert(jobRefer); err != nil {
		return false
	}
	return true
}
func DisplayJobReferDetails()(bool,[]JobRefer){
	var JobRefer []JobRefer
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
	collection := sess.DB("mirafrautilityapp").C("jobRefer")
	err = collection.Find(bson.M{"referstatus":false}).All(&JobRefer)
	if err != nil {
		fmt.Println("error2",err)
		return  false,JobRefer

	}
	fmt.Println("some place ")
	fmt.Println("jobrefer details struct",JobRefer)
	return true,JobRefer

}

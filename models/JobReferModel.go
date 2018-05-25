package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
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

	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("jobRefer")

	if err := c.Insert(jobRefer); err != nil {
		return false
	}
	return true
}
func DisplayJobReferDetails()(bool,[]JobRefer){
	var JobRefer []JobRefer
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Println("error1",err)
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("jobRefer")
	err = c.Find(bson.M{"ReferStatus": false}).All(&JobRefer)
	if err != nil {
		fmt.Println("error2",err)
		return  false,JobRefer

	}
	fmt.Println("jobrefer details struct",JobRefer)
	return true,JobRefer

}

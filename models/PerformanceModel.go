package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

type PerformanceModel struct {

	EmpId		string
	Name		string
	Comment		string
	Photo		string
}
func (performace PerformanceModel)InsertAward()bool{

	//UserData :=User{}
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
	collection := sess.DB("mirafrautilityapp").C("performance")
	//err = c.Find(bson.M{"EmpId": performace.EmpId}).One(&UserData)

	//performace.Photo =UserData.UploadPhoto
	// Insert
	if err := collection.Insert(performace); err != nil {
		return false
	}
	return true
}
func GetPerformance()(bool,[]PerformanceModel){
	var Performance []PerformanceModel
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
	collection := sess.DB("mirafrautilityapp").C("performance")
	err = collection.Find(nil).All(&Performance)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Performance

	}
	fmt.Println("performance details struct",Performance)
	return true,Performance
}
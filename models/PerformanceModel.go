package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

type PerformanceModel struct {

	EmpId		string
	Name		string
	Comment		string
	Photo		string
}
func (performace PerformanceModel)InsertAward()bool{

	//UserData :=User{}
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("performance")
	//err = c.Find(bson.M{"EmpId": performace.EmpId}).One(&UserData)

	//performace.Photo =UserData.UploadPhoto
	// Insert
	if err := c.Insert(performace); err != nil {
		return false
	}
	return true
}
func GetPerformance()(bool,[]PerformanceModel){
	var Performance []PerformanceModel
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Println("error1",err)
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("performance")
	err = c.Find(nil).All(&Performance)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Performance

	}
	fmt.Println("performance details struct",Performance)
	return true,Performance
}
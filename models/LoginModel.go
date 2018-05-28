package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"

	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)





type Refferals struct{
	reffreralId				string
	jobType					string
	userId					string
	candidateName			string
	candidatePhoneNumber	int64
	location				string
	experience				string
	resumePath				string
}

type Events struct{
	eventId			string
	userId			string
	eventTitle		string
	createdBy		string
	date			string
	time			string
	venue			string
	eventName		string
	description  	string
	eventStatus		bool
	comment			string
}

type EventRegister struct{
	eventid		string
	userid		string
}

type Login struct{
	Email		string
	Password	[]byte
}


func(loginDetails *Login)CheckLogin( )(bool,string){


	//loginDetails :=Login{}
	//loginDetails.UserId =userId
	//loginDetails.Password =password

	Result :=User{}
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
	fmt.Println("kkkkkk")
	defer sess.Close()
	collection := sess.DB("mirafrautilityapp").C("user")
	//err = collection.Find(bson.M{"DateOfJoin": todayDate}).All(&NewJoiners)
	//collection := sess.DB("mirafrautilityapp").C("user")
	fmt.Println("connection",collection)
	fmt.Println("kkkkk",loginDetails.Email)
	err = collection.Find(bson.M{"email": loginDetails.Email}).One(&Result)
	if err !=nil{
		log.Println("mmmmmm",err)
		return false,Result.PositionType

	}
	err = bcrypt.CompareHashAndPassword(Result.Password, loginDetails.Password)
    fmt.Println("result",Result.PositionType)
	if err != nil {
		fmt.Println("connection error ",err)
		return  false,Result.PositionType

	}else{
		return true,Result.PositionType

	}
	return  true,Result.PositionType

}

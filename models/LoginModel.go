package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"

	"golang.org/x/crypto/bcrypt"
	"log"
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
	Host := []string{
		"127.0.0.1:27017",

	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: Host,

	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	dbConnect := session.DB("MirafraUtility").C("user")
	fmt.Println("connection",dbConnect)
	err = dbConnect.Find(bson.M{"Email": loginDetails.Email}).One(&Result)
	if err !=nil{
		log.Println(err)
		return false,Result.PositionType

	}
	err = bcrypt.CompareHashAndPassword(Result.Password, loginDetails.Password)

	if err != nil {
		fmt.Println("connection error ",err)
		return  false,Result.PositionType

	}else{
		return true,Result.PositionType

	}
	return  true,Result.PositionType

}

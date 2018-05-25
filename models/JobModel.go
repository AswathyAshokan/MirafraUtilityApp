package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type JobRequirement struct{

	JobId				string
	EmpId				string
	Designation			string
	JobLocation			string
	KeySkills			string
	JobDescription		string
	KeyResponsibility	string
	MinExperience		string
	MaxExperience		string



}

type JobCounter struct {
	JobId  string
	seq		int64
}




func (jobDetails *JobRequirement)InserInJobRequirement( job JobRequirement) (bool) {

	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	jobCounter :=JobCounter{}
	jobCounter.JobId ="JobId"
	jobCounter.seq= 0
	db :=session.DB("MirafraUtility").C("jobCounter")
	if db !=nil{

		fmt.Println("error",db)
	}
	err = db.Insert(jobCounter)
	if err !=nil {
		fmt.Println("error login ",err)
	}

	//updating the sequence
	err =db.Find(nil).All(jobCounter)
	counter :=jobCounter.seq+1
	err = db.Update(bson.M{"JobId":jobCounter.JobId }, bson.M{"$set": bson.M{"seq": counter}})

	db =session.DB("MirafraUtility").C("jobRequirements")
	if db !=nil{

		fmt.Println("error",db)
	}
	job.JobId="job00"+ strconv.FormatInt(counter, 10)
	err = db.Insert(job)
	if err !=nil {
		fmt.Println("error login ",err)
	}

return true
}
func DisplayJobDetails()(bool,[]JobRequirement){
	var Job []JobRequirement
	session,err:=mgo.Dial("127.0.0.1")

	if err != nil {
		fmt.Println("error1",err)
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("MirafraUtility").C("jobRequirements")
	err = c.Find(nil).All(&Job)
	if err != nil {
		fmt.Println("error2",err)
		return  false,Job

	}
	fmt.Println("job details struct",Job)
	return true,Job

}

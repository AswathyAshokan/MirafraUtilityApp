

package controllers

type HelpController struct {
	BaseController

 }



//func getStudentSequence(sequenceName string) {
//	session,err:=mgo.Dial("127.0.0.1")
//
//	if err != nil {
//		panic(err)
//	}
//
//	defer session.Close()
//	c := session.DB("MirafraUtility")
//	var result = c.jobCounter.findAndModify({query: { _id: sequenceName },update: { $inc: { seq: 1 } }, new: true})
//	return result.seq;
//}
package controllers
//
//import (
//"github.com/gorilla/securecookie"
//"net/http"
//"log"
//
//)
//
//type SessionController struct{
//	BaseController
//}
//
//type SessionValues struct{
//	UserId		string
//	Password    string
//	AccountType string
//
//
//}
//
//var cookieToken = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
//
//
////function for set session values
//func SetSession(w http.ResponseWriter, sessionValues SessionValues){
//
//	value := make(map[string]string)
//	value["UserId"] = sessionValues.UserId
//	value["Password"] = sessionValues.Password
//	value["AccountType"] = sessionValues.AccountType
//
//
//	if encoded, err := cookieToken.Encode("session",value);err == nil{
//		cookie := &http.Cookie{
//			Name:  "session",
//			Value: encoded,
//			Path:  "/",
//		}
//		http.SetCookie(w,cookie)
//		log.Println("Session is Set!")
//	}
//}
//
////function to read the session values
//func ReadSession (w http.ResponseWriter, r *http.Request, companyTeamName string) (SessionValues) {
//	sessionValues := SessionValues{}
//	if cookie, err := r.Cookie("session"); err == nil {
//		value := make(map[string]string)
//		if err = cookieToken.Decode("session", cookie.Value, &value); err == nil {
//
//
//				sessionValues.UserId = value["UserId"]
//				sessionValues.Password = value["Password"]
//				sessionValues.AccountType = value["AccountType"]
//
//		} else {
//			http.Redirect(w, r, "/login", 302)
//			log.Println("Access Denied! You are not logged in!")
//		}
//	} else {
//		log.Println(err)
//		http.Redirect(w, r, "/login", 302)
//		log.Println("Access Denied! You are not logged in!")
//	}
//	return sessionValues
//}
//
//
////function to clear all session values
//func ClearSession(w http.ResponseWriter) {
//	cookie := &http.Cookie{
//		Name:   "session",
//		Value:  "",
//		Path:   "/",
//		MaxAge: -1,
//	}
//	http.SetCookie(w, cookie)
//	log.Println("Logged out Successfully!")
//	log.Println("The value in session after Logout:", cookie.Value)
//
//}
//
//
//
//
//
//
//
//
//

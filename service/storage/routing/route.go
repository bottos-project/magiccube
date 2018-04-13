package routing

//import (
//	"fmt"
//	"log"
//	"net/http"

//	"github.com/julienschmidt/httprouter"

//)

//// finds all possible routes that satisfy a given specification
//type StorageRoute interface {
//	//FetchRoutesForSpecification(rs storage.RouteSpecification) []storage.nodeinfo
//	FetchRoutesForSpecification(rs string) []string
//}
//func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
//	fmt.Fprint(w, "Welcome!\n")
//}

//func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//}

//func getuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	uid := ps.ByName("uid")
//	fmt.Fprintf(w, "you are get user %s", uid)
//}

//func modifyuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	uid := ps.ByName("uid")
//	fmt.Fprintf(w, "you are modify user %s", uid)
//}

//func deleteuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	uid := ps.ByName("uid")
//	fmt.Fprintf(w, "you are delete user %s", uid)
//}

//func adduser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	// uid := r.FormValue("uid")
//	uid := ps.ByName("uid")
//	fmt.Fprintf(w, "you are add user %s", uid)
//}

//func UrlRouting(){
//	router := httprouter.New()
//	router.GET("/", Index)
//	router.GET("/hello/:name", Hello)

//	router.GET("/user/:uid", getuser)
//	router.GET("/adduser/:uid", adduser)
//	router.GET("/deluser/:uid", deleteuser)
//	log.Fatal(http.ListenAndServe(":8080", router))
//}

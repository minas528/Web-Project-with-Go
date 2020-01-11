package main

import (
	"html/template"
	"net/http"

	eventRepo "../../../github.com/minas528/Online-voting-System/Event/repository"
	eventServ "../../../github.com/minas528/Online-voting-System/Event/service"
	"../../../github.com/minas528/Online-voting-System/delivery/http/handler"
	voteRepo "../../../github.com/minas528/Online-voting-System/votes/repository"
	voteServ "../../../github.com/minas528/Online-voting-System/votes/service"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var temp = template.Must(template.ParseGlob("ui/templates/*"))

func login(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "admin.voters", nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "signup", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "", nil)
}
func newEvnet(w http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(w, "new.event", nil)
}

func RoutesForAdmin() {

}
func main() {

	dbconn, err := gorm.Open("postgres", "postgres://postgres:default@localhost/votedb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	// errs := dbconn.CreateTable(&entities.Events{}, &entities.Post{}).GetErrors()
	// if 0 < len(errs) {
	// 	panic(errs)
	// }

	// errs := dbconn.CreateTable(&entities.RegParties{}, &entities.RegVoters{}).GetErrors()
	// if 0 < len(errs) {
	//	panic(errs)
	// }

	/*postRepo := postRepo.NewPostGormRepo(dbconn)
	postserv := postServ.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(temp, postserv)*/

	voteRepo := voteRepo.NewVoteGormRepo(dbconn)
	voteserv := voteServ.NewVoteService(voteRepo)
	voteHandler := handler.NewVotesHandler(temp, voteserv)

	eventRep := eventRepo.NewEventRepository(dbconn)
	eventserv := eventServ.NewEventService(eventRep)
	eventHandle := handler.NewEventHandler(temp, eventserv)

	fs := http.FileServer(http.Dir("ui/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	//http.HandleFunc("/upost", postHandler.PostNew)
	//http.HandleFunc("/posts", postHandler.Posts)
	http.HandleFunc("/", index)
	//http.HandleFunc("/newevent",newEvnet)

	http.HandleFunc("/events", eventHandle.Events)
	http.HandleFunc("/newevent", eventHandle.EventNew)

	http.HandleFunc("/vote", voteHandler.Vote) //
	http.HandleFunc("/choseParty", voteHandler.Chose)
	//http.HandleFunc("/choseParty", voteHandler.choseParty)
	//http.HandleFunc("/choseParty", voteHandler.choseParty) //

	http.HandleFunc("/voters", login)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8181", nil)
}

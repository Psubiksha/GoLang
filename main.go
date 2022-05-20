package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	
)
var dis []User

type User struct {
	gorm.Model
	Id int `json:"id"`
	Name  string  `json:"name"`
	Location float `json:"location"`
	Gender string `json:"gender"`


}

type Like struct{
	gorm.Model
	Id int `json:"id"`
	Wholikes int `json:"who_likes"`
	WhoIsLiked int `json:"who_is_liked"`
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{id}/{name}", NewUser).Methods("POST")
	myRouter.HandleFunc("/likes/{id}", Likes).Methods("GET")
	myRouter.HandleFunc("/likes/{id}", Distance).Methods("GET")
	myRouter.HandleFunc("/user/{id}/{name}", Query).Methods("GET")
	myRouter.HandleFunc("/user/{id}/{name}", UpdateUser).Methods("PUT")
	
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {

	initialMigration()
	// Handle Subsequent requests
	dis=[]User{
		User{"id":1,"location":-19},
        User{"id":2,"location":-25},
        User{"id":3,"location":11},
        User{"id":4,"location":-5},
        User{"id":5,"location":-14},
        User{"id":6,"location":-19},
        User{"id":7,"location":-15},
        User{"id":8,"location":-7},
        User{"id":9,"location":28},
        User{"id":10,"location":-25},
        User{"id":11,"location":-9},
        User{"id":12,"location":7},
        User{"id":13,"location":-22},
        User{"id":14,"location":0},
        User{"id":15,"location":10},
        User{"id":16,"location":-1},
        User{"id":17,"location":-9},
        User{"id":18,"location":-7},
        User{"id":19,"location":-18},
        User{"id":20,"location":-4},
        User{"id":21,"location":-10},
        User{"id":22,"location":-5},
        User{"id":23,"location":-1},
        User{"id":24,"location":-25},
        User{"id":25,"location":24},
        User{"id":26,"location":-19},
        User{"id":27,"location":20},
        User{"id":28,"location":15},
        User{"id":29,"location":28},
        User{"id":30,"location":12},
        User{"id":31,"location":-22},
        User{"id":32,"location":9},
        User{"id":33,"location":6},
        User{"id":34,"location":-5},
        User{"id":35,"location":18},
        User{"id":36,"location":-30},
        User{"id":37,"location":5},
        User{"id":38,"location":21},
        User{"id":39,"location":-28},
        User{"id":40,"location":-29},
        User{"id":41,"location":26},
        User{"id":42,"location":21},
        User{"id":43,"location":-25},
        User{"id":44,"location":18},
        User{"id":45,"location":-20},
        User{"id":46,"location":-5},
        User{"id":47,"location":-30},
        User{"id":48,"location":2},
        User{"id":49,"location":-15},
        User{"id":50,"location":-5},
        User{"id":51,"location":3},
        User{"id":52,"location":-13},
        User{"id":53,"location":-2},
        User{"id":54,"location":-1},
        User{"id":55,"location":-2},
        User{"id":56,"location":4},
        User{"id":57,"location":11},
        User{"id":58,"location":-3},
        User{"id":59,"location":-3},
        User{"id":60,"location":26},
        User{"id":61,"location":-16},
        User{"id":62,"location":-4},
        User{"id":63,"location":-3},
        User{"id":64,"location":-19},
        User{"id":65,"location":18},
        User{"id":66,"location":10},
        User{"id":67,"location":-7},
        User{"id":68,"location":-12},
        User{"id":69,"location":25},
        User{"id":70,"location":-4},
        User{"id":71,"location":-10},
        User{"id":72,"location":-28},
        User{"id":73,"location":-3},
        User{"id":74,"location":-15},
        User{"id":75,"location":23},
        User{"id":76,"location":13},
        User{"id":77,"location":28},
        User{"id":78,"location":-30},
        User{"id":79,"location":16},
        User{"id":80,"location":-5},
        User{"id":81,"location":-11},
        User{"id":82,"location":-9},
        User{"id":83,"location":-27},
        User{"id":84,"location":3},
        User{"id":85,"location":-30},
        User{"id":86,"location":4},
        User{"id":87,"location":-8},
        User{"id":88,"location":-7},
        User{"id":89,"location":20},
        User{"id":90,"location":-9},
        User{"id":91,"location":22},
        User{"id":92,"location":-6},
        User{"id":93,"location":-26},
        User{"id":94,"location":16},
        User{"id":95,"location":6},
        User{"id":96,"location":26},
        User{"id":97,"location":0},
        User{"id":98,"location":2},
        User{"id":99,"location":-24},
        User{"id":100,"location":-11}
	}


	handleRequests()
}
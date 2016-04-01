package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var database *sql.DB

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID    int    "json:id"
	Name  string "json:username"
	Email string "json:email"
	First string "json:first"
	Last  string "json:last"
}

func UserCreate(w http.ResponseWriter, r *http.Request) {

	NewUser := User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")
	output, err := json.Marshal(NewUser)

	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}
	
	“sql := "INSERT INTO users set user_nickname='" + NewUser.Name + "', user_first='" + NewUser.First + "', user_last='" + NewUser.Last + "', user_email='" + NewUser.Email + "'";
	
	q, err := database.Exec(sql)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(q)
}

func UserRetrieve(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Pragma", "no-cache");
	rows, _ := database.Query("select * from users LIMIT 10")
	
	Response := Users{}
	
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)
		
		Response.Users = append(Response.Users, user)
	}
	
	output,_ := json.Marshal(Response)
	fmt.Fprintln(w, string(output))
}

type API struct {
	Message string "json:message"
}

func Hello(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	name := urlParams["user"]
	HelloMessage := "Hello, " + name

	message := API{HelloMessage}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("Something went wrong.")
	}

	fmt.Fprintf(w, string(output))
}

func main() {
	
	db, err := sql.Open("mysql", "xizheye@/social_network")

	http.ListenAndServe(":8080", nil)
}

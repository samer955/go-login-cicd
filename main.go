package main

import (
	"fmt"
	"go-login/user"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Error   string
	Success string
}

func main() {

	http.HandleFunc("/", handle)

	fmt.Println("Starting webserver on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		getLoginForm(w, nil)
	case "/":
		getRegistrationForm(w, nil)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in":
		signInUser(w, r)
	}
}

func getRegistrationForm(w http.ResponseWriter, data any) {
	templating(w, "sign-up.html", data)
}

func getLoginForm(w http.ResponseWriter, data any) {
	templating(w, "sign-in.html", data)
}

func signUpUser(w http.ResponseWriter, r *http.Request) {

	newUser := getUser(r)
	err := user.Userservice.CreateUser(newUser)
	if err != nil {
		message := Message{Error: err.Error()}
		getRegistrationForm(w, message)
		return
	}

	success := "New user created: " + newUser.Email
	message := Message{Success: success}

	sendEmail(newUser.Email)
	getRegistrationForm(w, message)

}

func getUser(r *http.Request) user.User {

	username := r.FormValue("username")
	password := r.FormValue("password")

	newUser := user.User{
		Email:    username,
		Password: password,
	}

	return newUser
}

func signInUser(w http.ResponseWriter, r *http.Request) {

	loginUser := getUser(r)
	ok := user.Userservice.VerifyUser(loginUser)

	if !ok {
		errorMsg := Message{Error: "User or password doesnt match or not exists"}
		getLoginForm(w, errorMsg)
		return
	}
	successMsg := Message{Success: "User successfully logged-in"}
	getLoginForm(w, successMsg)

}

func templating(w http.ResponseWriter, temp string, data any) {
	t, err := template.ParseFiles(temp)
	if err != nil {
		log.Fatal("Error parsing template: ", temp)
	}
	t.ExecuteTemplate(w, temp, data)
}

func sendEmail(username string) {

	fmt.Println("Email sent to: " + username)

}

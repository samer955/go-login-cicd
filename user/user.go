package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var userDB = map[string]authUser{}

type User struct {
	Email    string
	Password string
}

type authUser struct {
	email        string
	hashpassword string
}

var Userservice userService

type userService struct {
}

func (userService) CreateUser(user User) error {

	_, ok := userDB[user.Email]
	if ok {
		fmt.Println("User already exists")
		return errors.New("user already exists")
	}

	hashedPsw, err := getHashedPassword(user.Password)
	if err != nil {
		return err
	}

	newUser := authUser{
		email:        user.Email,
		hashpassword: hashedPsw,
	}

	userDB[user.Email] = newUser
	fmt.Println("new user created")
	return nil
}

func (userService) VerifyUser(user User) bool {

	authUser, ok := userDB[user.Email]
	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(authUser.hashpassword),
		[]byte(user.Password))

	return err == nil
}

func getHashedPassword(password string) (string, error) {
	passw, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return "", err
	}
	return string(passw), nil
}

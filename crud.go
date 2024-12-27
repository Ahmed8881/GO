package main
import (
	"fmt"
	"net/http"
	"strconv"
)
type User struct {
	ID   int
	Name string
	Age int
}
var users []User
var nextID = 1
func createUser(name string, age int) User {
    user := User{
        ID:   nextID,
        Name: name,
        Age:  age,
    }
    users = append(users, user)
    nextID++
    return user
}
func getUsers() []User {
    return users
}
func getUserByID(id int) (User, bool) {
    for _, user := range users {
        if user.ID == id {
            return user, true
        }
    }
    return User{}, false
}

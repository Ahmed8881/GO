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
func updateUser(id int, name string, age int) bool {
    for i, user := range users {
        if user.ID == id {
            users[i].Name = name
            users[i].Age = age
            return true
        }
    }
    return false
}
func deleteUser(id int) bool {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return true
        }
    }
    return false
}
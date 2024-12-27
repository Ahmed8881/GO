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
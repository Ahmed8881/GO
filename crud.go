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

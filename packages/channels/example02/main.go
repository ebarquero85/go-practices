package main

import (
	"fmt"
)

type Server struct {
	users map[string]string
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]string),
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func main() {

}

func sendMessage(messagech chan<- string) {
	messagech <- "Hello!"
}

func readMessage(messagech <-chan string) {
	msg := <-messagech
	fmt.Println(msg)
}

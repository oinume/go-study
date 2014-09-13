package main

import "fmt"

func main() {
	session := NewSession()
	session.Open()
}

type Session struct {
	Opened bool
	Url    string
}

func (s *Session) Open() {
	s.Opened = true
	fmt.Println("Open() called")
}

func (s *Session) Close() {
	s.Opened = false
	fmt.Println("Close() called")
}

func NewSession() *Session {
	s := &Session{false, ""}
	defer s.Close()
	return s
}

package main

type Session struct {
	Url string
}

func (s *Session) Open() bool {
	return true
}

func main() {
	s := Session{"https://github.com/"}
	s.Open()
}

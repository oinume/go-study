package main

import "fmt"

// Database
type Database struct {
}
func (d *Database) NewSession() *Session {
	return &Session{}
}

// Session
type Session struct {
}
func (s *Session) Connect() string {
	return "Session.Connect()"
}

// ReadOnlyDatabase
type ReadOnlyDatabase Database
func (d *ReadOnlyDatabase) NewSession() *ReadOnlySession {
	return &ReadOnlySession{}
}

// ReadOnlySession
type ReadOnlySession Session
func (s *ReadOnlySession) Connect() string {
	return "ReadOnlySession.Connect()"
}

func main() {
	var database *Database = &Database{}
	var session *Session = database.NewSession()
	fmt.Println(session.Connect())

	var roDataBase *ReadOnlyDatabase = &ReadOnlyDatabase{}
	// ReadOnlySession -> Session型にキャスト
	var roSession *Session = (*Session)(roDataBase.NewSession())
	// ReadOnlySession.Connect()ではなくSession.Connect()が呼ばれる
	fmt.Println(roSession.Connect())

	// ReadOnlySession型
	var ros *ReadOnlySession = roDataBase.NewSession()
	// この場合はReadOnlySession.Connect()が呼ばれる
	fmt.Println(ros.Connect())
}

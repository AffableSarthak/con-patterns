package main

import "fmt"

type (
	Plsql struct {
		name string
	}

	Mongo struct {
		name string
	}

	DbMethods interface {
		Manipute() string
	}
)

func (pl Plsql) Manipute() string {
	return "some plsql db related stuff"
}

func (md *Mongo) manipulate() string {
	return "some mongo related stuff"
}

func createOrUpdateData(db DbMethods) {
	db.Manipute()
}

func main() {
	// Instatiate with plsql
	// Instantiate with mongo
	a := 2
	ps := Plsql{
		name: "s",
	}
	fmt.Println(a, ps)
}

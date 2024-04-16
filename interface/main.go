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

func (pl plsql) Manipute() string {
	return "some plsql db related stuff"
}

func (md *mongo) manipulate() string {
	return "some mongo related stuff"
}

func createOrUpdateData(db DbMethods) {
	db.Manipute()
}

func main() {
	// Instatiate with plsql
	// Instantiate with mongo
	a := 2
	ps := plsql{
		name: "s",
	}
	fmt.Println(a, ps)
}

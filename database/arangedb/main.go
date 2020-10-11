package main

import (
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func main() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://10.23.116.54:8529"},
	})
	if err != nil {
		fmt.Println(err)
	}
	// Create a client
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Open "examples_books" database
	db, err := c.Database(nil, "examples_books")
	if err != nil {
		fmt.Println(err)
	}

	// Open "books" collection
	col, err := db.Collection(nil, "books")
	if err != nil {
		fmt.Println(err)
	}

	// Create document
	book := Book{
		Title:   "ArangoDB Cookbook",
		NoPages: 257,
	}
	meta, err := col.CreateDocument(nil, book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Created document in collection '%s' in database '%s'\n meta :%+v", col.Name(), db.Name(), meta)
}

type Book struct {
	Title   string
	NoPages int
}

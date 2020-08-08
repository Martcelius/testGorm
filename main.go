package main

import (
	"flag"
	"log"
	"testGorm/contact"
	"testGorm/customer"
	"testGorm/dbconn"

	"net/http"
)

var (
	httpaddr = flag.String("http", ":8000", "http listen address")
)

func main() {
	dbconn.DbConn()
	defer dbconn.Db.Close()

	// dbconn.Db.DropTableIfExists(&customer.Customer{}, &contact.Contact{})
	dbconn.Db.AutoMigrate(&customer.Customer{}, &contact.Contact{})
	dbconn.Db.Model(&contact.Contact{}).AddForeignKey("cust_id", "customers(id)", "CASCADE", "CASCADE") // Foreign key need to define manually

	flag.Parse()

	log.Println("Server listen and serve on port: ", *httpaddr)
	err := http.ListenAndServe(*httpaddr, nil)

	if err != nil {
		log.Fatal("Error server", err)
	}
}

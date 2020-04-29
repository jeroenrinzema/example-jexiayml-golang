package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/gorilla/mux"
)

func user(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(fakeUser(vars["id"])))
}

func users(w http.ResponseWriter, r *http.Request) {
	length := int(gofakeit.Float32Range(2, 10))
	result := "["

	for i := 0; i < length; i++ {
		result += fakeUser(gofakeit.UUID()) + ","
	}

	result = result[:len(result)-1] + "]"

	w.Write([]byte(result))
}

func fakeUser(id string) string {
	template := `{
	"id": "%s",
	"first_name": "%s",
	"last_name": "%s",
	"age": "%s",
	"color": "%s",
	"created_at": "%d",
	"updated_at": "%d"
}`

	firstName := gofakeit.FirstName()
	lastName := gofakeit.LastName()
	age := int(gofakeit.Price(1, 80))
	color := gofakeit.HexColor()
	createdAt := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()).Unix()
	updatedAt := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()).Unix()

	return fmt.Sprintf(template, id, firstName, lastName, age, color, createdAt, updatedAt)
}

func product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(fakeProduct(vars["id"])))
}

func products(w http.ResponseWriter, r *http.Request) {
	length := int(gofakeit.Float32Range(5, 20))
	result := "["

	for i := 0; i < length; i++ {
		result += fakeProduct(gofakeit.UUID()) + ","
	}

	result = result[:len(result)-1] + "]"

	w.Write([]byte(result))
}

func fakeProduct(id string) string {
	template := `{
	"id": "%s",
	"name": "%s",
	"description": "%s",
	"price": "%d",
	"currency": "%s",
	"color": "%s",
	"created_at": "%d",
	"updated_at": "%d"
}`

	name := gofakeit.BeerName()
	description := gofakeit.Paragraph(1, 1, 50, "\n")
	price := int(gofakeit.Price(1, 40))
	currency := gofakeit.Currency().Short
	color := gofakeit.HexColor()
	createdAt := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()).Unix()
	updatedAt := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()).Unix()

	return fmt.Sprintf(template, id, name, description, price, currency, color, createdAt, updatedAt)
}

func main() {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", user)
	r.HandleFunc("/users", users)
	r.HandleFunc("/product/{id}", product)
	r.HandleFunc("/products", products)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Starting Server")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

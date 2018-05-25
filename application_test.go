package bibo

import (
	"testing"
	"fmt"
	"net/http"
)

func TestDefaultConfiguration(t *testing.T) {
	c := DefaultConfiguration()
	fmt.Println(c.getHost())
	fmt.Println(c.Charset)
}

func TestApplication_ListenAndServe(t *testing.T) {
	app := Default()
	app.routing.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("woca")
		fmt.Println("i am in home page")
	}, "home")

	app.routing.Post("/post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("post method")
	}, "post")
	app.ListenAndServe()
}
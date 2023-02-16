package main

import (
	"image/png"
	"log"
	"net/http"
	"os"
	"github.com/deepakrkris/IdentityCon/lib"
	"strings"
	"fmt"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	name := ""
	if param := q.Get("name"); param != "" {
		name = param
	}

	email := ""
	if param := q.Get("email"); param != "" {
		email = param
	}

	email_name := email[ : strings.Index(email, "@")]
	mail_server := email [ strings.Index(email, "@") : ]
	count_of_name_chars := 0
	for _, c := range strings.Split(name, "") {
		count_of_name_chars += strings.Count(email_name, c)
	}

    fmt.Println("identical email factor ", count_of_name_chars)

	areacode := ""
	if param := q.Get("areacode"); param != "" {
		areacode = param
	}

	params := map[string]string {
		"name" : name,
		"mail" : mail_server,
		"email_name": email_name,
		"areacode": areacode,
		"nameEmailFactor": fmt.Sprintf("%d", count_of_name_chars),
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, lib.GenerateIdenticon(params, 256, 256))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("Starting on port :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler{}))
}

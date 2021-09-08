package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"bytes"
	"net"
	"errors"
	"os"
)

var port = "8081"
var addr = "0.0.0.0"

func main() {
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	if len(os.Args) > 2 {
		addr = os.Args[2]
	} else {
		var err error
		addr, err = getIP()
		if err != nil {
			fmt.Println("No address can be determined, please specify when running.")
			return
		}
	}

	fmt.Printf("Running on %s:%s\n", addr, port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request from", r.RemoteAddr)
		t, err := template.New("index").ParseFiles("index.html")
		if err != nil {
			panic(err)
		}
		var resp bytes.Buffer
		err = t.ExecuteTemplate(&resp, "IP", addr)
		if err != nil {
			panic(err)
		}
		w.Write(resp.Bytes())
	})

	http.HandleFunc("/icon.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "icon.png")
	})

	http.HandleFunc("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "manifest.json")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// getIP returns the first non-loopback IP address in the interfaces.
func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil{
		return "", err
	}
	for _, a := range addrs {
		ipnet, ok := a.(*net.IPNet)
		// removed !ipnet.IP.IsPrivate() due to requiring go 1.17
		if ok && !ipnet.IP.IsLoopback() {
			return ipnet.IP.String(), nil
		}
	}
	return "", errors.New("No address")
}

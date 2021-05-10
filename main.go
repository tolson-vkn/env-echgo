package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type echgoHandler struct {
	Echgos []byte
}

func errorHandler(res http.ResponseWriter, req *http.Request, status int) {
	res.WriteHeader(status)

	notFound := make(map[string]string)
	notFound["message"] = "Hey what're you doing use the default route: /"
	nf, err := json.Marshal(notFound)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}

	if status == http.StatusNotFound {
		res.Write(nf)
	}
}

func (h echgoHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	header := res.Header()
	header.Set("Content-Type", "application/json")

	log.Printf("%s %s %s User-Agent: %s\n", req.Method, req.Proto, req.URL.Path, req.Header["User-Agent"])

	if req.URL.Path != "/" {
		errorHandler(res, req, http.StatusNotFound)
		return
	}

	res.WriteHeader(200)
	res.Write(h.Echgos)
}

func registerEchgos() map[string]string {
	registered := make(map[string]string)

	for _, s := range os.Environ() {

		split := strings.Split(s, "=")

		if strings.HasPrefix(split[0], "ECHGO") {

			eKey := strings.Split(split[0], "ECHGO_")
			if len(eKey) != 2 {
				log.Printf("Envar [%s] malformed, skipping\n", s)
				continue
			}

			key := eKey[1]
			key = strings.ToLower(key)
			if key == "message" {
				log.Printf("Key [%s] is reserverd, skipping\n", key)
				continue
			}

			value := split[1]
			registered[key] = value

			log.Printf("Processed envar [%s] as key: [%s], value: [%s]\n", s, key, value)
		}
	}

	return registered
}

func main() {
	registered := registerEchgos()
	registered["message"] = "Hello from Ecgho Server"

	echgos, err := json.Marshal(registered)
	if err != nil {
		log.Fatalf("err: %s\n", err)
	}
	log.Printf("Registered: %s", echgos)

	mux := http.NewServeMux()

	mux.Handle("/", &echgoHandler{Echgos: echgos})
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

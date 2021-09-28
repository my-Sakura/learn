package main

import (
	"fmt"
	"log"
	"net/http"

	"./lrucache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

// example: localhost:9999/_sakuracache/[Group]/[key]
// localhost:9999/_sakuracache/scores/Tom
func main() {
	lrucache.RegistGroup("scores", 2<<10, lrucache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("unhit cache")
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}

		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:9999"
	peers := lrucache.NewHTTPPool(addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}

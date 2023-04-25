package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"log"

	"github.com/gorilla/mux"
)

type ScanResult struct {
	SensitiveInfo map[string][]string `json:"sensitiveInfo"`
}

func startScan(w http.ResponseWriter, r *http.Request) {
	urls := r.FormValue("urls")
	urlList := strings.Split(urls, ",")

	var wg sync.WaitGroup
	info := make(map[string][]string)

	for _, url := range urlList {
		wg.Add(1)
		go func(targetURL string) {
			defer wg.Done()
			results := make(chan string)

			go func() {
				findSensitiveInfo(targetURL, results)
				close(results)
			}()

			tempInfo := []string{}
			for result := range results {
				tempInfo = append(tempInfo, result)
			}
			info[targetURL] = tempInfo
		}(url)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ScanResult{SensitiveInfo: info})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/startScan", startScan).Methods(http.MethodPost)


	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}

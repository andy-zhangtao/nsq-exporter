package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func generateRequestURL() (url, host string) {

	host = strings.Split(_url, "://")[1]
	url = fmt.Sprintf("%s/stats?format=json", _url)
	return
}

func requestStats(url string) (stats nsqdStats, err error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&stats)
	return
}

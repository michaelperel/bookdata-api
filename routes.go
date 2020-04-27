package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func timer(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	start := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		log.Printf("%s took %v\n", r.RequestURI, time.Since(start))
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	limit, err := getLimitParam(r)
	skip, err := getSkipParam(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
		return
	}
	data := books.GetAllBooks(limit, skip)
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "error marshalling data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return
}

func getLimitParam(r *http.Request) (int, error) {
	limit := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("limit")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return limit, err
		}
		limit = val
	}
	return limit, nil
}

func getSkipParam(r *http.Request) (int, error) {
	skip := 0
	queryParams := r.URL.Query()
	l := queryParams.Get("skip")
	if l != "" {
		val, err := strconv.Atoi(l)
		if err != nil {
			return skip, err
		}
		skip = val
	}
	return skip, nil
}

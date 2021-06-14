package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tohast/keyvalue"
)

const (
	storePath = "/store/"
)

func Serve(port string, store keyvalue.Store) error {
	mux := http.NewServeMux()
	mux.HandleFunc(storePath, NewStoreHandler(store))
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func NewStoreHandler(store keyvalue.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len(storePath):]
		if len(key) == 0 {
			http.Error(w, "no key in the path", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case "PUT":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			store.Put(keyvalue.Key(key), body)
		case "GET":
			v := store.Get(keyvalue.Key(key))
			if v == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Length", strconv.Itoa(len(v)))
			w.Header().Set("Content-Type", "application/octet-stream")
			_, err := io.Copy(w, bytes.NewReader(v))
			if err != nil {
				// if the client got HTTP Status, it can retry request
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "DELETE":
			store.Delete(keyvalue.Key(key))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

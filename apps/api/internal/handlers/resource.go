package handlers

import (
    "net/http"
)

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Resource data"))
}


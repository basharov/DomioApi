package types

import "net/http"

type Route struct {
    Name            string
    Method          string
    Pattern         string
    HandlerFunc     http.HandlerFunc
    CheckAccessFunc CheckAccessFunc
}

type CheckAccessFunc func(w http.ResponseWriter) bool

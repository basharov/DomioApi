package get_user_cards_handler

import (
    "net/http"
)

func CheckAccessFunc(req *http.Request) bool {
    return true
}
package get_user_card_handler

import "net/http"

type Data struct {
    Title string
}

func DataGetterFunc(req *http.Request) interface{} {
    return Data{Title:"hello there"}
}
package get_user_cards_handler

import (
    "net/http"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/db"
)

func GetUserCardsHandler(w http.ResponseWriter, req *http.Request) {
    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    userEmail := userProfile.Email
    userCards, _ := domiodb.GetCards(userEmail)

    responses.ReturnObjectResponse(w, userCards)

    defer req.Body.Close()
}
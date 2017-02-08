package login_user_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/responses"
    "domio_api/components/requests"
    //"log"
)

type UserLoggedinObject struct {
    Email       string  `json:"email"`
    Id          string  `json:"id"`
    TokenString string  `json:"token"`
}

func LoginUserHandler(w http.ResponseWriter, req *http.Request) {
    var emailAndPasswordPair domiodb.EmailAndPasswordPair

    err := requests.DecodeJsonRequestBody(req, &emailAndPasswordPair)

    if err != nil {
        responses.ReturnErrorResponse(w, domioerrors.IncorrectJSONInputError)
        return
    }

    if (emailAndPasswordPair.IsValid() != true) {
        responses.ReturnErrorResponse(w, domioerrors.PayloadValidationError)
        return
    }

    userClaims, tokenString, loginError := domiodb.LoginUser(emailAndPasswordPair)

    if (loginError != nil) {
        responses.ReturnErrorResponseWithCustomCode(w, domioerrors.WrongEmailOrPassword, http.StatusUnauthorized)
        return
    }

    responses.ReturnObjectResponse(w, UserLoggedinObject{Email:userClaims.Subject, Id:userClaims.Id, TokenString:tokenString})

    defer req.Body.Close()
}
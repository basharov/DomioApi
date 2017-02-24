package update_domain_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/components/requests"
    "log"
    "github.com/gorilla/mux"
)

func UpdateDomainHandler(w http.ResponseWriter, req *http.Request, data *interface{}) {

    var domainToEdit domiodb.DomainToEdit
    var updatedDomain domiodb.DomainJson

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != nil) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    err := requests.DecodeJsonRequestBody(req, &domainToEdit)

    if err != nil {
        log.Print(err)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    requestVars := mux.Vars(req)
    domainName := requestVars["name"]

    log.Print(userProfile)
    log.Print(domainName)

    domainUpdateError := domiodb.UpdateDomain(domainName, domainToEdit)

    if (domainUpdateError != nil) {
        log.Print(domainUpdateError)
        responses.ReturnErrorResponse(w, domioerrors.JsonDecodeError)
        return
    }

    responses.ReturnObjectResponse(w, updatedDomain)

    defer req.Body.Close()
}
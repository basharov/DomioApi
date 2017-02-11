package delete_user_handler

import (
    "net/http"
    "domio_api/db"
    domioerrors  "domio_api/errors"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/messages"
    "log"
    "github.com/stripe/stripe-go"
    "github.com/stripe/stripe-go/customer"
)

func DeleteUserHandler(w http.ResponseWriter, req *http.Request) {

    userProfile, verifyTokenError := tokens.VerifyTokenString(req.Header.Get("Authorization"))

    if (verifyTokenError != domioerrors.DomioError{}) {
        responses.ReturnErrorResponse(w, domioerrors.JwtTokenParseError)
        return
    }

    deletedUser, deleteError := domiodb.DeleteUser(userProfile.Email)

    if (deleteError != nil) {
        log.Println(deleteError)
        responses.ReturnErrorResponse(w, deleteError)
        return
    }

    deleteStripeCustomer(deletedUser.Id)

    responses.ReturnObjectResponse(w, messages.UserDeleted)

    defer req.Body.Close()
}

func deleteStripeCustomer(stripe_customer_id string) (*stripe.Customer, error) {

    stripe.Key = "sk_test_83T7gLMq9VQ4YLmWwBylJMS7"

    c, err := customer.Del(stripe_customer_id)
    return c, err
}
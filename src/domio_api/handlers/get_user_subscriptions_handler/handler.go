package get_user_subscriptions_handler

import (
    "net/http"
    "domio_api/components/tokens"
    "domio_api/components/responses"
    "domio_api/external_api/stripe/subscription"
)

func GetUserSubscriptionsHandler(w http.ResponseWriter, req *http.Request, userProfile *tokens.UserTokenWithClaims, isAccessGranted bool) {

    userSubscriptions, _ := stripe_subscription_adapter.GetUserSubscriptions(userProfile.Id)

    responses.ReturnObjectResponse(w, userSubscriptions)
}
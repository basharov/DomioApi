package routes

import (
    "net/http"
    "domio/handlers/create_user_handler"
    "domio/handlers/login_user_handler"
    "domio/handlers/verify_token_handler"
    "domio/handlers/get_available_domains_handler"
    "domio/handlers/get_user_domains_handler"
    "domio/handlers/create_domain_handler"
    "domio/handlers/get_domain_info_handler"
    "domio/handlers/create_subscription_handler"
    "domio/handlers/create_card_handler"
    "domio/handlers/delete_domain_handler"
    "domio/handlers/delete_subscription_handler"
    "domio/handlers/get_user_subscriptions_handler"
    "domio/handlers/show_status_handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
    Route{
        "ShowStatus",
        "GET",
        "/",
        show_status_handler.ShowStatusHandler,
    },
    Route{
        "CreateUser",
        "POST",
        "/users",
        create_user_handler.CreateUserHandler,
    },
    Route{
        "LoginUser",
        "POST",
        "/users/login",
        login_user_handler.LoginUser,
    },
    Route{
        "VerifyToken",
        "POST",
        "/tokens/verify",
        verify_token_handler.VerifyTokenHandler,
    },
    Route{
        "GetAvailableDomains",
        "GET",
        "/domains/available",
        get_available_domains_handler.GetAvailableDomainsHandler,
    },
    Route{
        "GetUserDomains",
        "GET",
        "/domains/user",
        get_user_domains_handler.GetUserDomainsHandler,
    },
    Route{
        "CreateDomain",
        "POST",
        "/domains",
        create_domain_handler.CreateDomainHandler,
    },
    Route{
        "DeleteDomain",
        "DELETE",
        "/domains/{name}",
        delete_domain_handler.DeleteDomainHandler,
    },
    Route{
        "GetUserSubscriptions",
        "GET",
        "/subscriptions",
        get_user_subscriptions_handler.GetUserSubscriptionsHandler,
    },
    Route{
        "CreateSubscription",
        "POST",
        "/subscriptions",
        create_subscription_handler.CreateSubscriptionHandler,
    },
    Route{
        "DeleteSubscription",
        "DELETE",
        "/subscriptions/{id}",
        delete_subscription_handler.DeleteSubscriptionHandler,
    },
    Route{
        "CreateCard",
        "POST",
        "/cards",
        create_card_handler.CreateCardHandler,
    },
    Route{
        "GetDomainInfo",
        "GET",
        "/domains/{name}",
        get_domain_info_handler.GetDomainInfoHandler,
    },
}
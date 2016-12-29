package get_user_domains_handler

import (
    "testing"
    . "github.com/franela/goblin"
    "domio/db"
)

func TestGetUserDomainsHandler(t *testing.T) {
    g := Goblin(t)
    g.Describe("GetUserDomainsHandler tests", func() {
        g.It("Should login and get user domains", func(done Done) {
            go func() {
                var virtualUser = VirtualUser{Email:"jack@gmail.com", Password:"jack@gmail.com"}
                virtualUser.Login()
                domainsResponse, _ := GetUserDomainsAs(&virtualUser)

                var domainsList []domiodb.Rental
                domainsResponse.JSON(&domainsList)

                g.Assert(domainsResponse.StatusCode).Equal(200)
                g.Assert(len(domainsList)).Equal(4)

                g.Assert(domainsList[0]).Equal(domiodb.Rental{Name:"jack200.com", Owner:"jack@gmail.com", Price:200})

                done()
            }()
        })
    })
}
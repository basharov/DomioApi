package create_domain_handler

import (
    "testing"
    . "github.com/franela/goblin"
    "domio/db"
    "domio/errors"
)

func TestCreateDomain(t *testing.T) {
    g := Goblin(t)
    g.Describe("Create Domain tests", func() {
        g.It("Should create a new domain succesfully", func(done Done) {
            go func() {
                var jack = VirtualUser{Email:"jack@gmail.com", Password:"jack@gmail.com"}
                jack.Login()

                domainToCreate := domiodb.Rental{Name:"jack.com", Price:1000}

                newDomainResponse, _ := CreateDomainAs(&jack, domainToCreate)

                var newDomain domiodb.Rental
                newDomainResponse.JSON(&newDomain)

                g.Assert(newDomainResponse.StatusCode).Equal(200)
                g.Assert(newDomain.Name).Equal("jack.com")
                g.Assert(newDomain.Price).Equal(1000)

                done()
            }()
        })

        g.It("Should reject domain creation for unauthorized user", func(done Done) {
            go func() {

                domainToCreate := domiodb.Rental{Name:"jack.com", Price:1000}
                errorResponse, _ := CreateDomainAs(&AnonymousUser, domainToCreate)

                var errorResponseObject errors.DomioError
                errorResponse.JSON(&errorResponseObject)

                g.Assert(errorResponse.StatusCode).Equal(401)

                g.Assert(errorResponseObject).Equal(errors.JwtTokenParseError)

                done()
            }()
        })
    })
}
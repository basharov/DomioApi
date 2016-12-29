package verify_token_handler

import (
    "testing"
    . "github.com/franela/goblin"
    "gopkg.in/h2non/gentleman.v1/plugins/body"
    "domio/errors"
)

func TestVerifyToken(t *testing.T) {
    g := Goblin(t)
    g.Describe("VerifyToken tests:", func() {
        g.It("Verify empty token", func(done Done) {
            go func() {
                var req = cli.Request()
                req.Path("/tokens/verify")
                req.Method("POST")

                res, _ := req.Send()

                var jsonResp errors.DomioError
                res.JSON(&jsonResp)

                g.Assert(res.StatusCode).Equal(401)

                g.Assert(jsonResp).Equal(errors.JsonDecodeError)
                done()
            }()
        })

        g.It("Verify wrong token", func(done Done) {
            go func() {
                var req = cli.Request()
                req.Path("/tokens/verify")
                req.Method("POST")

                wrongToken := Token{"abrakadabra"}
                req.Use(body.JSON(wrongToken))

                res, _ := req.Send()

                var jsonResp errors.DomioError
                res.JSON(&jsonResp)

                g.Assert(res.StatusCode).Equal(401)

                g.Assert(jsonResp).Equal(errors.JwtTokenParseError)

                done()
            }()
        })

        g.It("Login with correct email and password and then verify correct token", func(done Done) {
            go func() {
                // Login
                var req = cli.Request()
                req.Path("/users/login")
                req.Method("POST")

                data := map[string]string{"email": "jack@gmail.com", "password":"jack@gmail.com"}
                req.Use(body.JSON(data))

                loginResponse, _ := req.Send()

                var jsonResp UserTokenResponseBody
                _ = loginResponse.JSON(&jsonResp)

                g.Assert(loginResponse.StatusCode).Equal(200)

                g.Assert(jsonResp.Email).Equal("jack@gmail.com")

                // Verify token
                var tokenRequest = cli.Request()
                req.Path("/tokens/verify")
                req.Method("POST")

                tokenRequest.Path("/tokens/verify")

                correctToken := Token{jsonResp.Token}
                tokenRequest.Use(body.JSON(correctToken))

                tokenVerifyResponse, _ := tokenRequest.Send()

                var tokenResponse UserTokenResponseBody
                _ = tokenVerifyResponse.JSON(&tokenResponse)

                g.Assert(tokenVerifyResponse.StatusCode).Equal(200)

                g.Assert(jsonResp.Email).Equal("jack@gmail.com")

                done()
            }()
        })
    })
}
// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package client

import (
	"fmt"
	"strconv"

	securedservice "goa.design/goa/examples/security/gen/secured_service"
)

// BuildSigninPayload builds the payload for the secured_service signin
// endpoint from CLI flags.
func BuildSigninPayload(securedServiceSigninUsername string, securedServiceSigninPassword string) (*securedservice.SigninPayload, error) {
	var username string
	{
		username = securedServiceSigninUsername
	}
	var password string
	{
		password = securedServiceSigninPassword
	}
	payload := &securedservice.SigninPayload{
		Username: username,
		Password: password,
	}
	return payload, nil
}

// BuildSecurePayload builds the payload for the secured_service secure
// endpoint from CLI flags.
func BuildSecurePayload(securedServiceSecureFail string, securedServiceSecureToken string) (*securedservice.SecurePayload, error) {
	var err error
	var fail *bool
	{
		if securedServiceSecureFail != "" {
			val, err := strconv.ParseBool(securedServiceSecureFail)
			fail = &val
			if err != nil {
				err = fmt.Errorf("invalid value for fail, must be BOOL")
			}
		}
	}
	var token *string
	{
		if securedServiceSecureToken != "" {
			token = &securedServiceSecureToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &securedservice.SecurePayload{
		Fail:  fail,
		Token: token,
	}
	return payload, nil
}

// BuildDoublySecurePayload builds the payload for the secured_service
// doubly_secure endpoint from CLI flags.
func BuildDoublySecurePayload(securedServiceDoublySecureKey string, securedServiceDoublySecureToken string) (*securedservice.DoublySecurePayload, error) {
	var key *string
	{
		if securedServiceDoublySecureKey != "" {
			key = &securedServiceDoublySecureKey
		}
	}
	var token *string
	{
		if securedServiceDoublySecureToken != "" {
			token = &securedServiceDoublySecureToken
		}
	}
	payload := &securedservice.DoublySecurePayload{
		Key:   key,
		Token: token,
	}
	return payload, nil
}

// BuildAlsoDoublySecurePayload builds the payload for the secured_service
// also_doubly_secure endpoint from CLI flags.
func BuildAlsoDoublySecurePayload(securedServiceAlsoDoublySecureKey string, securedServiceAlsoDoublySecureOauthToken string, securedServiceAlsoDoublySecureToken string, securedServiceAlsoDoublySecureUsername string, securedServiceAlsoDoublySecurePassword string) (*securedservice.AlsoDoublySecurePayload, error) {
	var key *string
	{
		if securedServiceAlsoDoublySecureKey != "" {
			key = &securedServiceAlsoDoublySecureKey
		}
	}
	var oauthToken *string
	{
		if securedServiceAlsoDoublySecureOauthToken != "" {
			oauthToken = &securedServiceAlsoDoublySecureOauthToken
		}
	}
	var token *string
	{
		if securedServiceAlsoDoublySecureToken != "" {
			token = &securedServiceAlsoDoublySecureToken
		}
	}
	var username *string
	{
		if securedServiceAlsoDoublySecureUsername != "" {
			username = &securedServiceAlsoDoublySecureUsername
		}
	}
	var password *string
	{
		if securedServiceAlsoDoublySecurePassword != "" {
			password = &securedServiceAlsoDoublySecurePassword
		}
	}
	payload := &securedservice.AlsoDoublySecurePayload{
		Key:        key,
		OauthToken: oauthToken,
		Token:      token,
		Username:   username,
		Password:   password,
	}
	return payload, nil
}

// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogauth

import (
	"errors"
	"fmt"
	"github.com/arelate/gogauthurls"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strings"
)

const (
	reCaptchaError       = "reCAPTCHA present on the page"
	secondStepCodePrompt = "Please enter 2FA code (check your inbox):"
)

func authToken(client *http.Client) (token string, error error) {

	req, err := http.NewRequest(http.MethodGet, gogauthurls.AuthHost().String(), nil)
	gogauthurls.AddAuthHostDefaultHeaders(req)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	// check for captcha presence
	if querySelector(doc, scriptReCaptchaSelector) != nil {
		// TODO: Write how to add cookie from the browser to allow user to unblock themselves
		return "", errors.New(reCaptchaError)
	}

	input := querySelector(doc, inputLoginTokenSelector)
	return attrVal(input, "value"), nil
}

func secondStepAuth(client *http.Client, body io.ReadCloser, sfa func(string) (string, error)) error {

	doc, err := html.Parse(body)
	if err != nil {
		return err
	}

	input := querySelector(doc, inputSecondStepAuthTokenSelector)
	token := attrVal(input, "value")

	for token != "" {

		code := ""
		if sfa != nil {
			for len(code) != 4 {
				code, err = sfa(secondStepCodePrompt)
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("2FA token is requied, cannot get it with nil callback")
		}

		data := secondStepData(code, token)

		req, _ := http.NewRequest(http.MethodPost, gogauthurls.LoginTwoStep().String(), strings.NewReader(data))
		gogauthurls.AddLoginHostDefaultHeaders(req)
		gogauthurls.SetLoginFormHeaders(req, gogauthurls.LoginTwoStep())

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		doc, err = html.Parse(resp.Body)
		if err != nil {
			return err
		}

		input = querySelector(doc, inputSecondStepAuthTokenSelector)
		token = attrVal(input, "value")

		resp.Body.Close()
	}

	return nil
}

/*
Login to GOG.com for account formData queries using username and password

Overall flow is:
- Get auth token from the page (this would check for reCaptcha as well)
- Post username, password and token for check
- Check for presence of second step auth token
- (Optional) Post 4 digit second step auth code
*/
func Login(client *http.Client, username, password string, sfa func(string) (string, error)) error {

	token, err := authToken(client)
	if err != nil {
		return err
	}

	data := loginData(username, password, token)

	req, err := http.NewRequest(http.MethodPost, gogauthurls.LoginCheck().String(), strings.NewReader(data))
	if err != nil {
		return err
	}
	gogauthurls.AddLoginHostDefaultHeaders(req)
	// GOG.com redirects initial auth request from authHost to loginHost.
	gogauthurls.SetLoginFormHeaders(req, gogauthurls.LoginHost())

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if err := secondStepAuth(client, resp.Body, sfa); err != nil {
		return err
	}

	return resp.Body.Close()
}
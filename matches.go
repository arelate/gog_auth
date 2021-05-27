// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth

import (
	"github.com/arelate/gog_auth_urls"
	"golang.org/x/net/html"
	"strings"
)

func inputLoginToken(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "input" &&
		attrVal(n, "name") == "login[_token]"
}

func inputSecondStepAuthToken(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "input" &&
		attrVal(n, "name") == "second_step_authentication[_token]"
}

func scriptReCaptcha(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "script" &&
		strings.HasPrefix(attrVal(n, "src"), gog_auth_urls.ReCaptcha().String())
}

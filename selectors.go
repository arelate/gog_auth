// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogauth

import (
	"github.com/arelate/gogauthurls"
	"golang.org/x/net/html"
	"strings"
)

func inputLoginTokenSelector(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "input" &&
		attrVal(n, "name") == "login[_token]"
}

func inputSecondStepAuthTokenSelector(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "input" &&
		attrVal(n, "name") == "second_step_authentication[_token]"
}

func scriptReCaptchaSelector(n *html.Node) bool {
	return n != nil &&
		n.Type == html.ElementNode &&
		n.Data == "script" &&
		strings.HasPrefix(attrVal(n, "src"), gogauthurls.ReCaptcha().String())
}

// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/arelate/gog_auth_urls"
	"github.com/arelate/gog_types"
)

func LoggedIn(client *http.Client) (bool, error) {

	resp, err := client.Get(gog_auth_urls.UserData().String())

	if err != nil {
		return false, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var ud gog_types.UserData

	err = json.Unmarshal(respBody, &ud)
	if err != nil {
		return false, err
	}

	err = resp.Body.Close()
	if err != nil {
		return ud.IsLoggedIn, err
	}

	return ud.IsLoggedIn, nil
}

// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogauth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/arelate/gogauthurls"
	"github.com/arelate/gogtypes"
)

func LoggedIn(client *http.Client) (bool, error) {

	resp, err := client.Get(gogauthurls.UserData().String())

	if err != nil {
		return false, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var ud gogtypes.UserData

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

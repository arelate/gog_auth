# gog_auth

Implementation of the GOG.com website (as in "not Galaxy") authentication flow, including 2FA. Captcha is detected, not supported.

## GOG.com authentication details

Exported functions:

### Login

Implements authentication flow for a given username, password and a callback to get second factor auth token (you'll get this in mail on login attempt if you have it set up) 

### LoggedIn 

Returns whether a user is logged in.

## Persistent authentication

You'll need to pass http.Client reference with a cookieJar that would hold GOG.com authentication
cookies. Client implementations might differ, however general flow should look something like the
following:

- Load persistent cookies
- Create a http.CookieJar and add those cookies
- Create a http.Client with that cookieJar
- Check if the gog_auth.LoggedIn is true
    - If it is: interact with AccountProducts, Wishlist, etc. types that require authentication
    - If it isn't: Login with a username, password, optional 2FA
- At the end of the session - save persistent cookies from the cookieJar
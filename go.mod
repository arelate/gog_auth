module github.com/arelate/gog_auth

go 1.16

require (
	github.com/arelate/gog_auth_urls v0.1.2-alpha
	github.com/arelate/gog_types v0.1.6-alpha
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
)

replace (
	github.com/arelate/gog_auth_urls => ../gog_auth_urls
	github.com/arelate/gog_types => ../gog_types
)

package api

import "github.com/jub0bs/cors"

var CORSMiddleware *cors.Middleware

func init() {
	var err error
	// If you have a project that could use client-side API access
	// to hstspreload.org, feel free to send a pull request
	// to add your Web origin on GitHub:
	// https://github.com/chromium/hstspreload.org/edit/master/api/cors.go
	CORSMiddleware, err = cors.NewMiddleware(cors.Config{
		Origins: []string{
			"http://localhost:*",
			"https://localhost:*",
			"https://mozilla.github.io:*",
			"https://observatory.mozilla.org:*",
			"https://a.ncsccs.com:*",
			"https://chksite.com:*",
		},
	})
	if err != nil { // The CORS config is invalid.
		panic(err)
	}
}

package main

import (
	"net/http"
	"regexp"

	"github.com/facebookgo/grace/gracehttp"
	"gopkg.in/labstack/echo.v1"

	"xojoc.pw/useragent"
)

func main() {
	e := echo.New()
	e.Get("/", func(c *echo.Context) error {
		ua := useragent.Parse(c.Request().Header.Get("User-Agent"))

		switch ua.OS {
		case "Android":
			match, _ := regexp.MatchString("YaBrowser", ua.Original)
			if match {
				return c.Redirect(http.StatusFound, "intent://install#Intent;scheme=get;package=com.menu.joker;end")
			}
			return c.Redirect(http.StatusFound, "market://details?id=com.menu.joker")
		case "iOS":
			return c.Redirect(http.StatusFound, "itms-apps://itunes.apple.com/us/app/apple-store/id1086543332?l=tr&ls=1&mt=8")
		default:
			return c.Redirect(http.StatusFound, "https://joker.menu")
		}
	})
	server := e.Server(":1323")

	// HTTP2 is currently enabled by default in echo.New(). To override TLS handshake errors
	// you will need to override the TLSConfig for the server so it does not attempt to valudate
	// the connection using TLS as required by HTTP2
	server.TLSConfig = nil

	// Graceful shutdown
	gracehttp.Serve(server)
}

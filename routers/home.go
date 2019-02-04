package routers

import "gopkg.in/macaron.v1"

func Home(c *macaron.Context) {
	c.Write([]byte("Welcome to API v1"))
}

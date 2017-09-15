package midware

import (
	"strings"

	"master/data"
	"master/data/user"

	"master/util"

	"fly"
)

func CheckLogin(c *fly.Context) {
	path := c.Request.URL.Path

	whiteList := []string{
		"/login",
		"/static/help.html",
		"/register",
		"/static/js",
		"/static/style",
		"/static/fake",
		"/static/statsTMPL.html",
		"/stats",
		"/taskqueue",
	}

	for _, url := range whiteList {
		if strings.HasPrefix(path, url) {
			return
		}
	}

	token, err := c.Request.Cookie(data.Token)
	if err != nil || token == nil {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	uid, err := c.Request.Cookie(data.Uid)
	if err != nil || token == nil {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}

	u, ok := user.GetUserByUid(uid.Value)

	if !ok {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}

	if util.Hash(u.Name, u.Password) == token.Value {
		c.Put("user", u)
		return
	} else {
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
}

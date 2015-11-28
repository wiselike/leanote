package member

import (
	"github.com/revel/revel"
)

// 帐户信息

type MemberUser struct {
	MemberBaseController
}

func (c MemberUser) Username() revel.Result {
	c.SetUserInfo()
	c.SetLocale()
	c.RenderArgs["title"] = "Username"
	return c.RenderTemplate("member/user/username.html")
}

func (c MemberUser) Email() revel.Result {
	c.SetUserInfo()
	c.SetLocale()
	c.RenderArgs["title"] = "Email"
	return c.RenderTemplate("member/user/email.html")
}

func (c MemberUser) Password() revel.Result {
	c.SetUserInfo()
	c.SetLocale()
	c.RenderArgs["title"] = "Password"
	return c.RenderTemplate("member/user/password.html")
}

func (c MemberUser) Avatar() revel.Result {
	c.SetUserInfo()
	c.SetLocale()
	c.RenderArgs["title"] = "Avatar"
	c.RenderArgs["globalConfigs"] = configService.GetGlobalConfigForUser()
	return c.RenderTemplate("member/user/avatar.html")
}

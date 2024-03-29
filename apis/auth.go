package apis

import (
	"github.com/gin-gonic/gin"
	"go-sample/models"
	"go-sample/pkg"
	"net/http"
)

func Login(c *gin.Context) {
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")

	if hasSession := pkg.HasSession(c); hasSession == true {
		c.String(200, "用户已登陆")
		return
	}

	user := models.UserDetailByName(name)

	if err := pkg.Compare(user.Password, password); err != nil {
		c.String(401, "密码错误")
		return
	}

	pkg.SaveAuthSession(c, user.ID)

	c.String(200, "登录成功")
}

func Logout(c *gin.Context) {
	if hasSession := pkg.HasSession(c); hasSession == false {
		c.String(401, "用户未登陆")
		return
	}
	pkg.ClearAuthSession(c)
	c.String(200, "退出成功")
}

func Register(c *gin.Context) {
	var user models.User
	user.Name = c.Request.FormValue("name")
	user.Email = c.Request.FormValue("email")

	if hasSession := pkg.HasSession(c); hasSession == true {
		c.String(200, "用户已登陆")
		return
	}

	if existUser := models.UserDetailByName(user.Name); existUser.ID != 0 {
		c.String(200, "用户名已存在")
		return
	}

	if c.Request.FormValue("password") != c.Request.FormValue("password_confirmation") {
		c.String(200, "密码不一致")
		return
	}

	if pwd, err := pkg.Encrypt(c.Request.FormValue("password")); err == nil {
		user.Password = pwd
	}

	models.AddUser(&user)

	pkg.SaveAuthSession(c, user.ID)

	c.String(200, "注册成功")

}

func Me(c *gin.Context) {
	currentUser := c.MustGet("userId").(uint)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": currentUser,
	})
}

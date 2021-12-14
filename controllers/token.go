package controllers

import (
	"PrometheusAlert/models"
	beegojwt "PrometheusAlert/pkg/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

type TokenAuthController struct {
	beego.Controller
}

type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// GetToken 发放token
func (c *TokenAuthController) GetToken() {
	appKey := c.GetString("app_key")
	appSecret := c.GetString("app_secret")

	if appKey == "" || appSecret == "" {
		c.Data["json"] = response{
			Code: 400,
			Msg:  "签名错误",
		}
		c.ServeJSON()
		return
	}

	//获取用户
	user, err := models.GetUserByAppKey(appKey, appSecret)
	if err != nil || user == nil {
		c.Data["json"] = response{
			Code: 400,
			Msg:  "用户未找到",
		}
		c.ServeJSON()
		return
	}
	logs.Info("用户信息获取成功, 用户名: " + user.Name)

	expire := time.Now().Unix() + 3600*24*7
	tokenString := ""
	et := beegojwt.EasyToken{
		Username: user.AppKey + "_" + user.Name,
		Expires:  expire, //Segundos
	}
	tokenString, _ = et.GetToken()

	data := make(map[string]interface{})
	data["name"] = user.Name
	data["token"] = tokenString
	data["expire"] = expire
	c.Data["json"] = response{
		Code: 200,
		Msg:  "请求成功",
		Data: data,
	}
	c.ServeJSON()
	return
}

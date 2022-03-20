package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

type LoginUser struct {
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}
type KvPair struct {
	Key   string
	Value string
}

var LoginInfoKeys = []string{KwUserName, KwUsedFlow, KwUsedTime, KwBalance, KwIp}

var LoginInfoFields = map[string]string{
	KwUserName: XpathSpanUserName,
	KwUsedFlow: XpathSpanUsedFlow,
	KwUsedTime: XpathSpanUsedTime,
	KwBalance:  XpathSpanBalance,
	KwIp:       XpathSpanIp,
}

var LoginInfo = map[string]*KvPair{
	KwUserName: {Key: LabelUserName},
	KwUsedFlow: {Key: LabelUsedFlow},
	KwUsedTime: {Key: LabelUsedTime},
	KwBalance:  {Key: LabelBalance},
	KwIp:       {Key: LabelIp},
}

func PrintLoginInfo() {
	for _, k := range LoginInfoKeys {
		pair := LoginInfo[k]
		fmt.Printf("%s:%v\n", pair.Key, pair.Value)
	}
}

var Account LoginUser

func ValidCheck() bool {
	success, err := WebDriver.FindElements(selenium.ByXPATH, XpathLoginWaySuccess)
	Handle(err)
	login, err := WebDriver.FindElements(selenium.ByXPATH, XpathLoginWayAccount)
	Handle(err)
	if len(login) == 0 && len(success) == 0 {
		// Neither not-login nor login-success
		return false
	}
	return true
}

func Login() {
	inputUsername, err := WebDriver.FindElement(selenium.ByXPATH, XpathInputUserName)
	Handle(err)
	inputPassword, err := WebDriver.FindElement(selenium.ByXPATH, XpathInputPassword)
	Handle(err)
	Handle(inputUsername.Clear())
	Handle(inputUsername.SendKeys(Account.Username))
	Handle(inputPassword.Clear())
	Handle(inputPassword.SendKeys(Account.Password))
	rememberClick, err := WebDriver.FindElement(selenium.ByXPATH, XpathRememberUser)
	Handle(err)
	sel, err := rememberClick.IsSelected()
	if !sel {
		Handle(rememberClick.Click())
	}
	loginButton, err := WebDriver.FindElement(selenium.ByXPATH, XpathLoginButton)
	Handle(err)
	Handle(loginButton.Click())
}

// Status 获取登录状态
func Status() bool {
	// 检测 login-form 中是否存在 login-way success
	elems, err := WebDriver.FindElements(selenium.ByXPATH, XpathLoginWaySuccess)
	Handle(err)
	if len(elems) == 0 {
		return false
	}
	form := elems[0]
	for _, k := range LoginInfoKeys {
		xpath := LoginInfoFields[k]
		elems, err := form.FindElements(selenium.ByXPATH, xpath)
		Handle(err)
		if len(elems) > 0 {
			text, err := elems[0].Text()
			Handle(err)
			LoginInfo[k].Value = text
		}
	}
	return true
}

func Logout() {
	elems, err := WebDriver.FindElements(selenium.ByXPATH, XpathLogoutButton)
	Handle(err)
	if len(elems) == 0 {
		return
	}
	Handle(elems[0].Click())
}

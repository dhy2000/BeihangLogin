package main

// DefaultGateUrl 认证网关地址
const DefaultGateUrl = "https://gw.buaa.edu.cn"

// 认证页面元素选取 Xpath

// XpathLoginForm 认证登录组件整体
const XpathLoginForm = "//*[@class=\"section-content\"]/*[@class=\"module login-container\"]/div[@class=\"login-form\"]"

// Not Login

// XpathLoginWayAccount 用户登录 div
const XpathLoginWayAccount = XpathLoginForm + "/div[@class=\"login-way account\"]"

// XpathLoginInputForm 登录表单
const XpathLoginInputForm = XpathLoginWayAccount + "/form[@id=\"login-form\"]"

// XpathLoginButton 登录按钮
const XpathLoginButton = XpathLoginInputForm + "/div[@class=\"row btn-group\"]/button[@id=\"login\"]"

const XpathRememberUser = XpathLoginInputForm + "/div[@class=\"row func\"]/div[@class=\"remember\"]/label[@for=\"remember\"]/input[@id=\"remember\"]"

// 登录表单字段
const (
	KwInputUsername    = "username"
	XpathInputUserName = XpathLoginInputForm + "/div[@class=\"row input\"]/input[@id=\"" + KwInputUsername + "\"]"

	KwInputPassword    = "password"
	XpathInputPassword = XpathLoginInputForm + "/div[@class=\"row input\"]/input[@id=\"" + KwInputPassword + "\"]"
)

// Login Success

// XpathLoginWaySuccess 登录成功 div
const XpathLoginWaySuccess = XpathLoginForm + "/div[@class=\"login-way success\"]"

// XpathLoginRowInfos 登录用户信息表格
const XpathLoginRowInfos = XpathLoginWaySuccess + "/form[@id=\"login-form\"]/div[@class=\"row\"]/div[@class=\"row info\"]"

// XpathLogoutButton 注销按钮
const XpathLogoutButton = XpathLoginWaySuccess + "/div[@class=\"row btn-group\"]/button[@id=\"logout-dm\"]"

// 当前登录用户信息字段
const (
	KwUserName        = "user_name"
	LabelUserName     = "用户名"
	XpathSpanUserName = XpathLoginRowInfos + "/span[@id=\"" + KwUserName + "\"]"

	KwUsedFlow        = "used_flow"
	LabelUsedFlow     = "已用流量"
	XpathSpanUsedFlow = XpathLoginRowInfos + "/span[@id=\"" + KwUsedFlow + "\"]"

	KwUsedTime        = "used_time"
	LabelUsedTime     = "已用时长"
	XpathSpanUsedTime = XpathLoginRowInfos + "/span[@id=\"" + KwUsedTime + "\"]"

	KwBalance        = "balance"
	LabelBalance     = "帐户余额"
	XpathSpanBalance = XpathLoginRowInfos + "/span[@id=\"" + KwBalance + "\"]"

	KwIp        = "ip"
	LabelIp     = "IP地址"
	XpathSpanIp = XpathLoginRowInfos + "/span[@id=\"" + KwIp + "\"]"
)

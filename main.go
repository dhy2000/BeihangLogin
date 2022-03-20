package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"github.com/tebeka/selenium/firefox"
	"log"
	"os"
	"time"
)

func Handle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func WaitRefresh() {
	time.Sleep(time.Millisecond * 100)
}

const DriverPort = 4444

var (
	DriverPath  string
	BrowserName string
	GateUrl     string
	AccountPath string
)

var Operation string

func loadArgs() {
	var help bool
	pflag.StringVarP(&DriverPath, "driver", "d", "chromedriver", "Path to your chromedriver/geckodriver")
	pflag.StringVarP(&BrowserName, "browser", "b", "chrome", "Type of your browser (chrome/firefox)")
	pflag.StringVarP(&GateUrl, "gate", "g", DefaultGateUrl, "Url of the net gate.")
	pflag.StringVarP(&AccountPath, "account", "a", "account", "Path to account file.")
	pflag.BoolVarP(&help, "help", "h", false, "show help")
	pflag.Parse()
	if help {
		fmt.Println("Usage: beihang-login [login|status|logout] [options...]")
		fmt.Println("Available options: ")
		pflag.PrintDefaults()
		os.Exit(0)
	}
	args := pflag.Args()
	if len(args) == 0 {
		log.Fatalln("Operation required: login | status | logout")
	}
	Operation = args[0]
}

func readAccount() {
	viper.SetConfigFile(AccountPath)
	viper.SetConfigType("yaml")
	Handle(viper.ReadInConfig())

	Handle(viper.Unmarshal(&Account))
}

var DriverService *selenium.Service
var WebDriver selenium.WebDriver

func main() {
	loadArgs()
	// Start Service

	caps := selenium.Capabilities{"browserName": BrowserName}
	switch BrowserName {
	case "chrome":
		service, err := selenium.NewChromeDriverService(DriverPath, DriverPort)
		Handle(err)
		DriverService = service
		caps.AddChrome(chrome.Capabilities{
			Args: []string{"--headless"},
		})
	case "firefox":
		service, err := selenium.NewGeckoDriverService(DriverPath, DriverPort)
		Handle(err)
		DriverService = service
		caps.AddFirefox(firefox.Capabilities{
			Args: []string{"-headless"},
		})
	default:
		log.Fatalln("Wrong driver type: expected 'chrome' or 'firefox'")
	}
	defer func(DriverService *selenium.Service) {
		_ = DriverService.Stop()
	}(DriverService)

	// Start WebDriver
	var err error
	if WebDriver, err = selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", DriverPort)); err != nil {
		log.Fatalln(err)
	} else {
		defer func(WebDriver selenium.WebDriver) {
			_ = WebDriver.Close()
		}(WebDriver)
	}

	Handle(WebDriver.Get(GateUrl))
	Handle(WebDriver.Refresh())
	WaitRefresh()

	valid := ValidCheck()
	if !valid {
		html, err := WebDriver.PageSource()
		Handle(err)
		Handle(os.WriteFile(fmt.Sprintf("err-page-%v.html", time.Now().Format("2006-01-02-15-04-05")), []byte(html), 0644))
		log.Fatalln("认证页面打开失败，请稍后再试")
	}
	status := Status()
	switch Operation {
	case "login":
		readAccount()
		if status {
			log.Println("already login, logout now...")
			Logout()
			Handle(WebDriver.Refresh())
			WaitRefresh()
		}
		Login()
		Handle(WebDriver.Refresh())
		WaitRefresh()
		if Status() {
			fmt.Println("登录成功")
			os.Exit(0)
		} else {
			fmt.Println("登录失败")
			os.Exit(1)
		}
	case "status":
		if status {
			PrintLoginInfo()
		} else {
			fmt.Println("没有用户登录")
		}
	case "logout":
		if status {
			Logout()
			fmt.Println("注销成功")
			os.Exit(0)
		}
	default:
		log.Fatalln("Wrong operation type: expected login | status | logout.")
	}
}

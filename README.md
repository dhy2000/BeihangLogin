# Beihang Login

Yet another (not very stable) beihang login based on `selenium` and your web browser.

**Note**: Only chrome is supported now. There are still some bugs with firefox.

## Requirements

One of the follows:

- Chrome/Chromium web browser and [ChromeDriver](https://chromedriver.chromium.org/downloads)
- Firefox web browser and [GeckoDriver](https://github.com/mozilla/geckodriver/releases)

## Usage

Type `./beihang-login -h` and you will get the help message below.

    Usage: beihang-login [login|status|logout] [options...]
    Available options:
    -a, --account string   Path to account file. (default "account")
    -b, --browser string   Type of your browser (chrome/firefox) (default "chrome")
    -d, --driver string    Path to your chromedriver/geckodriver (default "chromedriver")
    -g, --gate string      Url of the net gate. (default "https://gw.buaa.edu.cn")
    -h, --help             show help

Download web driver and copy it into `PATH` (or use `-d` argument to specify the full path of driver).

Copy `account.example` to `account` and fill in your username and password.

## Build

```shell
go mod download
go build -o beihang-login
```

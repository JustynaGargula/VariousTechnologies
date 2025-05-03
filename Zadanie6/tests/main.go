package main

import (
	"log"

	"tests/testCases"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--ignore-certificate-errors",
			"--disable-web-security",
			"--allow-insecure-localhost",
		},
	}
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, "http://localhost:9515") // port chromedrivera
	if err != nil {
		log.Fatalf("Nie można połączyć się z WebDriverem: %s", err)
	}
	defer wd.Quit()

	testCases.SetWebDriver(wd)
	testCases.Test1()
	testCases.Test2()
	testCases.Test3()
	testCases.Test4()
	testCases.Test5()
	testCases.Test6()
	testCases.Test7()
	testCases.Test8()
	testCases.Test9()
	testCases.Test10()
	testCases.Test11()

}

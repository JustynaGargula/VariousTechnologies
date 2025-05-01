package main

import (
	"fmt"
	"log"

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

	if err := wd.Get("http://localhost:1323"); err != nil {
        log.Fatalf("Nie można otworzyć strony: %s", err)
    }

	fmt.Println("Test 1. Pobranie tytułu strony")
	title, err := wd.Title()
    if err != nil {
        log.Fatalf("Nie można pobrać tytułu strony: %s", err)
    }
	expected := ""
    if title == expected {
        fmt.Println("✅ Test zaliczony! Tytuł to:", title)
    } else {
        fmt.Printf("❌ Test niezaliczony. Oczekiwano '%s', otrzymano '%s'\n", expected, title)
    }
}
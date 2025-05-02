package testCases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tebeka/selenium"
)

var wd selenium.WebDriver
var newProductID int

func SetWebDriver(webDriver selenium.WebDriver) {
	wd = webDriver
}

func Test1() {
	fmt.Println("Test 1. Otworzenie strony głównej")
	if err := wd.Get("http://localhost:1323"); err != nil {
		log.Fatalf("Nie można otworzyć strony: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Println("✅ Test zaliczony! Udało się otworzyć stronę")
	}
}

func Test2() {
	fmt.Println("Test 2. Pobranie tytułu strony")
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

func Test3() {
	fmt.Println("Test 3. Sprawdzenie zawartości strony głównej")
	content, err := wd.FindElement(selenium.ByCSSSelector, "h1")
	if err != nil {
		log.Fatalf("Nie znaleziono nagłówka h1: %s", err)
	}

	expected := "Hello, World!"
	header, error := content.Text()
	if error != nil {
		log.Fatalf("Nie można odczytać nagłówka h1: %s", err)
	}

	if header == expected {
		fmt.Println("✅ Test zaliczony! Nagłówek to:", header)
	} else {
		fmt.Printf("❌ Test niezaliczony. Oczekiwano '%s', otrzymano '%s'\n", expected, header)
	}
}

func Test4() {
	fmt.Println("Test 4. Otworzenie /products")
	if err := wd.Get("http://localhost:1323/products"); err != nil {
		log.Fatalf("Nie można otworzyć strony: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Println("✅ Test zaliczony! Udało się otworzyć stronę")
	}
}

func Test5() {
	fmt.Println("Test 5. Odczytanie produktów")
	content, err := wd.PageSource()
	if err != nil {
		log.Fatalf("Nie znaleziono nagłówka h1: %s", err)
	}

	errorMessage := "No products found"
	if content != errorMessage {
		fmt.Println("✅ Test zaliczony! Poprawnie odczytano stronę z produktami.")
	} else {
		fmt.Printf("❌ Test niezaliczony. Otrzymano wiadomość '%s'\n", errorMessage)
	}
}

func Test6() {
	fmt.Println("Test 6. Dodanie produktu")

	product := map[string]interface{}{
		"Name":  "Checked notebook",
		"Price": 10,
	}
	body, err := json.Marshal(product)
	if err != nil {
		log.Fatalf("Nie udało się zserializować JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:1323/products", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Błąd podczas żądania POST: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		log.Fatalf("Nieprawidłowy status: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalf("Nie udało się sparsować odpowiedzi JSON: %v", err)
	}

	if result["name"] != "Checked notebook" {
		fmt.Printf("❌ Test niezaliczony. Zły produkt w odpowiedzi: %v", result)
	} else {
		idFloat, _ := result["ID"].(float64)
		newProductID = int(idFloat)
		fmt.Printf("✅ Test zaliczony! Poprawnie dodano nowy produkt z ID: %d. \n", newProductID)
	}
}

func Test7() {
	fmt.Println("Test 7. Szczegóły produktu")
	url := "http://localhost:1323/products/" + fmt.Sprint(newProductID)
	if err := wd.Get(url); err != nil {
		log.Fatalf("Nie można otworzyć strony szczegółów produktu: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Printf("✅ Test zaliczony! Udało się otworzyć stronę szczegółów produktu z ID: %d. \n", newProductID)
	}
}

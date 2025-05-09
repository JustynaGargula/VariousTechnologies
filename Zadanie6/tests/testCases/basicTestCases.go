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
var newProductName = "Checked notebook"
var newProductPrice = 10.0
var newCartID int
var newCartItemName = "product 1"
var newCartItemPrice = 99.9

func SetWebDriver(webDriver selenium.WebDriver) {
	wd = webDriver
}

func Test1() {
	fmt.Println("\nTest 1. Otworzenie strony głównej")
	if err := wd.Get("http://localhost:1323"); err != nil {
		log.Fatalf("Nie można otworzyć strony: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Println("✅ Test zaliczony! Udało się otworzyć stronę")
	}
}

func Test2() {
	fmt.Println("\nTest 2. Pobranie tytułu strony")
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
	fmt.Println("\nTest 3. Sprawdzenie zawartości strony głównej")
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
	fmt.Println("\nTest 4. Otworzenie /products")
	if err := wd.Get("http://localhost:1323/products"); err != nil {
		log.Fatalf("Nie można otworzyć strony: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Println("✅ Test zaliczony! Udało się otworzyć stronę")
	}
}

func Test5() {
	fmt.Println("\nTest 5. Odczytanie produktów")
	content, err := wd.PageSource()
	if err != nil {
		log.Fatalf("Nie znaleziono nic na stronie z produktami: %s", err)
	}

	errorMessage := "No products found"
	if content != errorMessage {
		fmt.Println("✅ Test zaliczony! Poprawnie odczytano stronę z produktami.")
	} else {
		fmt.Printf("❌ Test niezaliczony. Otrzymano wiadomość '%s'\n", errorMessage)
	}
}

func Test6() {
	fmt.Println("\nTest 6. Dodanie produktu")

	product := map[string]interface{}{
		"Name":  newProductName,
		"Price": newProductPrice,
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
	fmt.Println("\nTest 7. Otworzenie podstrony z szczegółami produktu")
	url := "http://localhost:1323/products/" + fmt.Sprint(newProductID)
	if err := wd.Get(url); err != nil {
		log.Fatalf("Nie można otworzyć strony szczegółów produktu: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Printf("✅ Test zaliczony! Udało się otworzyć stronę szczegółów produktu z ID: %d. \n", newProductID)
	}
}

func Test8() {
	fmt.Println("\nTest 8. Odebranie szczegółów produktu")
	content, err := wd.PageSource()
	if err != nil {
		log.Fatalf("Nie znaleziono nic na stronie: %s", err)
	}

	errorMessage := "Product not found"
	if content != errorMessage {
		fmt.Println("✅ Test zaliczony! Poprawnie odebrano szczegóły produktu.")
	} else {
		fmt.Printf("❌ Test niezaliczony. Otrzymano wiadomość '%s'\n", errorMessage)
	}
}

func Test9() {
	fmt.Println("\nTest 9. Sprawdzenie szczegółów produktu")
	content, err := wd.FindElement(selenium.ByCSSSelector, "pre")
	if err != nil {
		log.Fatalf("Nie znaleziono nic na stronie: %s", err)
	}

	contentText, _ := content.Text()
	var product map[string]interface{}
	err = json.Unmarshal([]byte(contentText), &product)
	if err != nil {
		log.Fatalf("Błąd parsowania JSON: %v", err)
	}
	if product["name"] == newProductName {
		fmt.Println("✅ 1. Test zaliczony! Poprawnie odczytano nazwę produktu.")
	} else {
		fmt.Printf("❌ 1. Test niezaliczony: błędna nazwa produktu. Oczekiwano: %s, otrzymano: %s.\n", newProductName, product["name"])
	}
	if product["price"] == newProductPrice {
		fmt.Println("✅ 2. Test zaliczony! Poprawnie odczytano cenę produktu.")
	} else {
		fmt.Printf("❌ 2. Test niezaliczony: błędna cena produktu. Oczekiwano: %f, otrzymano: %f.\n", newProductPrice, product["price"])
	}

}

func Test10() {
	fmt.Println("\nTest 10. Otworzenie podstrony szczegółów usuniętego produktu")
	url := "http://localhost:1323/products/1"
	if err := wd.Get(url); err != nil {
		log.Fatalf("Nie można otworzyć strony szczegółów produktu: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Printf("✅ Test zaliczony! Udało się otworzyć stronę szczegółów produktu z ID: 1\n")
	}
}

func Test11() {
	fmt.Println("\nTest 11. Pobranie szczegółów usuniętego produktu")
	content, err := wd.FindElement(selenium.ByTagName, "pre")
	if err != nil {
		log.Fatalf("Nie znaleziono nic na stronie: %s", err)
	}
	contentText, _ := content.Text()

	errorMessage := "Product not found"
	if contentText == errorMessage {
		fmt.Printf("✅ Test zaliczony! Poprawnie odebrano wiadomość '%s'\n", errorMessage)
	} else {
		fmt.Printf("❌ Test niezaliczony. Oczekiwano '%s', otrzymano '%s'\n", errorMessage, contentText)
	}
}

func Test12() {
	fmt.Println("\nTest 12. Dodanie nowego koszyka")
	cart := map[string]interface{}{
		"items": []map[string]interface{}{
			{
				"name":     newCartItemName,
				"price":    newCartItemPrice,
				"quantity": 1,
			},
			{
				"name":     "product 2",
				"price":    11.99,
				"quantity": 2,
			},
		},
	}

	body, err := json.Marshal(cart)
	if err != nil {
		log.Fatalf("Nie udało się zserializować JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:1323/cart", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Request failed: %v", err)
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

	if result["ID"] == 0 {
		fmt.Printf("❌ Test niezaliczony. Koszyk nie został poprawnego ID: %v. \n", result["ID"])
	} else {
		idFloat, _ := result["ID"].(float64)
		newCartID = int(idFloat)
		fmt.Printf("✅ Test zaliczony! Poprawnie dodano nowy produkt z ID: %d. \n", newCartID)
	}
}

func Test13() {
	fmt.Println("\nTest 13. Otworzenie strony z koszykiem")
	url := "http://localhost:1323/cart/" + fmt.Sprint(newCartID)
	err := wd.Get(url)
	if err != nil {
		log.Fatalf("Nie można otworzyć strony z koszykiem: %s", err)
		fmt.Println("❌ Test niezaliczony.")
	} else {
		fmt.Printf("✅ Test zaliczony! Udało się otworzyć stronę z koszykiem z ID: %d\n", newCartID)
	}
}

func Test14() {
	fmt.Println("\nTest 14. Wyświetlenie koszyka")
	content, err := wd.PageSource()
	if err != nil {
		log.Fatalf("Nie znaleziono nic na stronie: %s", err)
	}

	errorMessage := "Cart not found"
	if content != errorMessage {
		fmt.Println("✅ Test zaliczony! Poprawnie odebrano koszyk.")
	} else {
		fmt.Printf("❌ Test niezaliczony. Otrzymano wiadomość '%s'\n", errorMessage)
	}
}

func Test15() {
	fmt.Println("\nTest 15. Poprawna zawartość koszyka")
	content, err := wd.FindElement(selenium.ByCSSSelector, "pre")
	if err != nil {
		log.Fatalf("Nie znaleziono nic na stronie: %s", err)
	}

	contentText, _ := content.Text()
	var cart map[string]interface{}
	err = json.Unmarshal([]byte(contentText), &cart)
	if err != nil {
		log.Fatalf("Błąd parsowania JSON: %v", err)
	}

	items := cart["items"].([]interface{})

	if len(items) <= 0 {
		log.Fatalln("Koszyk nie zawiera żadnych produktów!")
	}

	firstItem := items[0].(map[string]interface{})
	name := firstItem["name"].(string)

	if name == newCartItemName {
		fmt.Println("✅ 1. Test zaliczony! Poprawnie odczytano nazwę produktu w koszyku.")
	} else {
		fmt.Printf("❌ 1. Test niezaliczony: błędna nazwa produktu w koszyku. Oczekiwano: %s, otrzymano: %s.\n", newCartItemName, name)
	}

	price := firstItem["price"]
	if price == newCartItemPrice {
		fmt.Println("✅ 2. Test zaliczony! Poprawnie odczytano cenę produktu w koszyku.")
	} else {
		fmt.Printf("❌ 2. Test niezaliczony: błędna cena produktu w koszyku. Oczekiwano: %f, otrzymano: %f.\n", newCartItemPrice, price)
	}
}

func Test16() {
	fmt.Println("\nTest 16. Podanie nieprawidłowego ID koszyka")
	url := "http://localhost:1323/cart/-1"
	err := wd.Get(url)
	if err != nil {
		log.Fatalf("Nie można otworzyć strony z koszykiem: %s", err)
	}

	contentRaw, _ := wd.FindElement(selenium.ByCSSSelector, "pre")
	content, _ := contentRaw.Text()
	if content != "Invalid ID format" {
		fmt.Println("❌ Test niezaliczony. Nie wykryto nieprawidłowego ID")
	} else {
		fmt.Printf("✅ Test zaliczony! Wykryto nieprawidłowe ID. Otrzymano wiadomość: %s\n", content)
	}
}

func Test17() {
	fmt.Println("\nTest 17. Podanie nieprawidłowego ID produktu")
	url := "http://localhost:1323/products/-1"
	err := wd.Get(url)
	if err != nil {
		log.Fatalf("Nie można otworzyć strony z produktem: %s", err)
	}

	contentRaw, _ := wd.FindElement(selenium.ByCSSSelector, "pre")
	content, _ := contentRaw.Text()
	if content != "Invalid ID format" {
		fmt.Println("❌ Test niezaliczony. Nie wykryto nieprawidłowego ID")
	} else {
		fmt.Printf("✅ Test zaliczony! Wykryto nieprawidłowe ID. Otrzymano wiadomość: %s\n", content)
	}
}

func Test18() {
	fmt.Println("\nTest 18. Próba otwarcia koszyka bez podania jego ID (z użyciem endpointu dla POST)")
	url := "http://localhost:1323/cart"
	err := wd.Get(url)
	if err != nil {
		log.Fatalf("Nie można otworzyć strony z koszykiem: %s", err)
	}

	contentRaw, _ := wd.FindElement(selenium.ByCSSSelector, "pre")
	content, _ := contentRaw.Text()
	var messageJson map[string]interface{}
	err = json.Unmarshal([]byte(content), &messageJson)
	if err != nil {
		log.Fatalf("Błąd parsowania JSON: %v", err)
	}

	if messageJson["message"] != "Method Not Allowed" {
		fmt.Println("❌ Test niezaliczony. Nie wykryto nieprawidłowego użycia endpointu")
	} else {
		fmt.Printf("✅ Test zaliczony! Wykryto nieprawidłowe użycie endpointu. Otrzymano wiadomość: %s\n", messageJson["message"])
	}
}

func Test19() {
	fmt.Println("\nTest 19. Dodanie produktu bez ceny")

	product := map[string]interface{}{
		"Name": newProductName,
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
		fmt.Printf("✅ Test zaliczony! Nie dodano produktu. Status odpowiedzi to %d\n", resp.StatusCode)
	} else {
		fmt.Printf("❌ Test niezaliczony. Dodano produkt z nieistniejącą cenę. Status odpowiedzi to %d\n", resp.StatusCode)
	}
}

func Test20() {
	fmt.Println("\nTest 20. Dodanie koszyka z produktani bez ceny")
	cart := map[string]interface{}{
		"items": []map[string]interface{}{
			{
				"name":     newCartItemName,
				"quantity": 1,
			},
			{
				"name":     "product 2",
				"quantity": 2,
			},
		},
	}

	body, err := json.Marshal(cart)
	if err != nil {
		log.Fatalf("Nie udało się zserializować JSON: %v", err)
	}

	resp, err := http.Post("http://localhost:1323/cart", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		fmt.Printf("✅ Test zaliczony! Nie dodano koszyka. Status odpowiedzi to %d\n", resp.StatusCode)
	} else {
		fmt.Printf("❌ Test niezaliczony. Dodano koszyk z nieistniejącymi cenami. Status odpowiedzi to %d\n", resp.StatusCode)
	}
}

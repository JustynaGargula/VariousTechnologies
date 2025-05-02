#!/bin/bash

# Uruchom chromedriver w tle i zapisz jego PID
./chromedriver-win64/chromedriver --port=9515 &
CHROMEDRIVER_PID=$!

# Czekaj chwilę na uruchomienie chromedrivera
sleep 2

# Uruchom backend (server.go) w tle i zapisz jego PID
cd ../../Zadanie4
go run server.go &
BACKEND_PID=$!

# Czekaj chwilę na uruchomienie backendu
sleep 2

# Funkcja sprzątająca
cleanup() {
  echo "Zatrzymywanie procesów..."
  kill $CHROMEDRIVER_PID
  kill $BACKEND_PID
}

# Zarejestruj cleanup na wyjście (normalne lub Ctrl+C)
trap cleanup EXIT

# Przejdź do katalogu z testami i uruchom je
cd ../Zadanie6/tests
go run main.go

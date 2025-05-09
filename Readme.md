# Zadania ebiznes

## Lista zadań

- [Zadanie 1 Docker](#zadanie-1-docker)
- [Zadanie 2 Scala](#zadanie-2-scala)
- [Zadanie 3 Kotlin](#zadanie-3-kotlin)
- [Zadanie 4 Go](#zadanie-4-go)
- [Zadanie 5 React](#zadanie-5-react)
- [Zadanie 6 Testy](#zadanie-6-testy)
- [Zadanie 7 Sonar](#zadanie-7-sonar)
- [Zadanie 8 Oauth2](#zadanie-8-oauth2)

## Zadanie 1 Docker

### Link do hub.docker.com:

https://hub.docker.com/repository/docker/justyna72/ebiznes/general
(obraz ma tag 'zadanie1')

### Opis Zadania

- 3.0 obraz ubuntu z Pythonem w wersji 3.10
- 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem
- 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC
  SQLite w ramach projektu na Gradle (build.gradle)
- 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
  przez CMD oraz gradle
- 5.0 dodać konfigurację docker-compose

[zbiorczy commit](https://github.com/JustynaGargula/VariousTechnologies/commit/19dcaac8acee28bc572d2bb9b9115d7c54820008)

Obraz należy wysłać na hub.docker.com, a link do obrazu należy dodać w
README na githubie.

## Zadanie 2 Scala

### Opis zadania

Należy stworzyć aplikację na frameworku Play w Scali 3.

- 3.0 Należy stworzyć kontroler do Produktów [commit ten + poprawki w wersji na 3.5](https://github.com/JustynaGargula/VariousTechnologies/commit/7d11bbba56bc0632b1da21393e5ed4bfbe0edf45)
- 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/a4e58457c4066609bb697d355e01ab2c2345897b)
- 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
  zgodnie z CRUD [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/3512efd8450940e91df07e7b5e6e1c6c12758609)
- 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
  skrypt uruchamiający aplikację via ngrok
- 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all,
show by id (get), update (put), delete (delete), add (post).

Kod w folderze: Zadanie2/ProductController

## Zadanie 3 Kotlin

- 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/00ba0fcbc03e8cbb0f7be04069ac930d50af2c47)
- 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota) [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/db75ed340f1c11f04775e1f44a4dbe37ab55f72a)
- 4.0 Zwróci listę kategorii na określone żądanie użytkownika [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/cec2cd346359ada456deb466e1464b958d561e6a)
- 4.5 Zwróci listę produktów wg żądanej kategorii [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/c1b74a089be7b62d57fc90fd0b367a93d4c6f39d)
- 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger, Webex

Kod w folderze: Zadanie3/MessageRedirect

## Zadanie 4 Go

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do
stworzenia 5 modeli, gdzie pomiędzy dwoma musi być relacja. Należy
zaimplementować proste endpointy do dodawania oraz wyświetlania danych
za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym
jednak pozostać przy sqlite.

- 3.0 Należy stworzyć aplikację we frameworku echo w j. Go, która będzie
  miała kontroler Produktów zgodny z CRUD [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/2f504bf2e7ec8f1be7bca72c3a90fa202ea84628)
- 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz
  wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast
  listy) [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/cfeefff2e26be705485748453d6b37112b4a5adf)
- 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint
- 4.5 Należy stworzyć model kategorii i dodać relację między kategorią,
  a produktem
- 5.0 pogrupować zapytania w gorm’owe scope'y

## Zadanie 5 React

Należy stworzyć aplikację kliencką wykorzystując bibliotekę React.js.
W ramach projektu należy stworzyć trzy komponenty: Produkty, Koszyk
oraz Płatności. Koszyk oraz Płatności powinny wysyłać do aplikacji
serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach
z aplikacji serwerowej. Aplikacja serwera w jednym z trzech języków:
Kotlin, Scala, Go. Dane pomiędzy wszystkimi komponentami powinny być
przesyłane za pomocą React hooks.

- 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz
  Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w
  Produktach powinniśmy pobierać dane o produktach z aplikacji
  serwerowej; [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/c576454d3ba4dbe27a348fad36df49e2a82f8502)
- 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/d1919877a02b474ddd59ead39cc00ebc854d1c67)
- 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za
  pomocą React hooks [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/4a11ffe5e46fa475cd255dc69da26cd2d17a207b)
- 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz
  kliencką na dockerze via docker-compose [commit (w następnym commitie jeszcze drobne poprawki zostały dodane)](https://github.com/JustynaGargula/VariousTechnologies/commit/653460750d0f473fa90beb29913aa92d4054a0f0)
- 5.0 Należy wykorzystać axios’a oraz dodać nagłówki pod CORS

## Zadanie 6 Testy

Należy stworzyć 20 przypadków testowych w jednym z rozwiązań:

- Cypress JS (JS)
- Selenium (Kotlin, Python, Java, JS, Go, Scala)

Testy mają w sumie zawierać minimum 50 asercji (3.5). Mają również
uruchamiać się na platformie Browserstack (5.0). Proszę pamiętać o
stworzeniu darmowego konta via https://education.github.com/pack.

- 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium
  (Kotlin, Python, Java, JS, Go, Scala) [this commit + previous ones](https://github.com/JustynaGargula/VariousTechnologies/commit/f636830de81aba80e9d19e59fa3f65fbda53250a)
- 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50
  asercji [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/37ee40f9c92bc8626b56102230258bfeb89365c7)
- 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego
  projektu z minimum 50 asercjami
- 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z
  minimum jednym scenariuszem negatywnym per endpoint
- 5.0 Należy uruchomić testy funkcjonalne na Browserstacku

## Zadanie 7 Sonar

Należy dodać projekt aplikacji klienckiej oraz serwerowej (jeden
branch, dwa repozytoria) do Sonara w wersji chmurowej
(https://sonarcloud.io/). Należy poprawić aplikacje uzyskując 0 bugów,
0 zapaszków, 0 podatności, 0 błędów bezpieczeństwa. Dodatkowo należy
dodać widżety sonarowe do README w repozytorium dane projektu z
wynikami.

- 3.0 Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w
  hookach gita
- 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod
  aplikacji serwerowej)
- 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod
  aplikacji serwerowej)
- 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa
  w kodzie w Sonarze (kod aplikacji serwerowej)
- 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie
  aplikacji klienckiej

## Zadanie 8 Oauth2

Należy skonfigurować klienta Oauth2 (4.0). Dane o użytkowniku wraz z
tokenem powinny być przechowywane po stronie bazy serwera, a nowy
token (inny niż ten od dostawcy) powinien zostać wysłany do klienta
(React). Można zastosować mechanizm sesji lub inny dowolny (5.0).
Zabronione jest tworzenie klientów bezpośrednio po stronie React'a
wyłączając z komunikacji aplikację serwerową, np. wykorzystując auth0.

Prawidłowa komunikacja: react-sewer-dostawca-serwer(via return
uri)-react.

- 3.0 logowanie przez aplikację serwerową (bez Oauth2)
- 3.5 rejestracja przez aplikację serwerową (bez Oauth2)
- 4.0 logowanie via Google OAuth2
- 4.5 logowanie via Facebook lub Github OAuth2
- 5.0 zapisywanie danych logowania OAuth2 po stronie serwera

# Zadania ebiznes
## Lista zadań
* [Zadanie 1 Docker](#zadanie-1-docker)
* [Zadanie 2 Scala](#zadanie-2-scala)

## Zadanie 1 Docker
### Link do hub.docker.com:
https://hub.docker.com/repository/docker/justyna72/ebiznes/general
(obraz ma tag 'zadanie1')
### Opis Zadania
* 3.0 obraz ubuntu z Pythonem w wersji 3.10
* 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem
* 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC
SQLite w ramach projektu na Gradle (build.gradle)
* 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle
* 5.0 dodać konfigurację docker-compose

Obraz należy wysłać na hub.docker.com, a link do obrazu należy dodać w
README na githubie.

## Zadanie 2 Scala

### Opis zadania

Należy stworzyć aplikację na frameworku Play w Scali 3.

* 3.0 Należy stworzyć kontroler do Produktów [commit ten + poprawki w wersji na 3.5](https://github.com/JustynaGargula/VariousTechnologies/commit/7d11bbba56bc0632b1da21393e5ed4bfbe0edf45)
* 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/a4e58457c4066609bb697d355e01ab2c2345897b)
* 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD [commit](https://github.com/JustynaGargula/VariousTechnologies/commit/3512efd8450940e91df07e7b5e6e1c6c12758609)
* 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok
* 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

Kontrolery mogą bazować na listach zamiast baz danych. CRUD: show all,
show by id (get), update (put), delete (delete), add (post).

Kod w folderze: Zadanie2/ProductController

## Zadanie 3

* 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord
* 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota)
* 4.0 Zwróci listę kategorii na określone żądanie użytkownika
* 4.5 Zwróci listę produktów wg żądanej kategorii
* 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger,
Webex

## Zadanie 4 Go

Należy stworzyć projekt w echo w Go. Należy wykorzystać gorm do
stworzenia 5 modeli, gdzie pomiędzy dwoma musi być relacja. Należy
zaimplementować proste endpointy do dodawania oraz wyświetlania danych
za pomocą modeli. Jako bazę danych można wybrać dowolną, sugerowałbym
jednak pozostać przy sqlite.

* 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie
miała kontroler Produktów zgodny z CRUD
* 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz
wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast
listy)
* 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint
* 4.5 Należy stworzyć model kategorii i dodać relację między kategorią,
a produktem
* 5.0 pogrupować zapytania w gorm’owe scope'y

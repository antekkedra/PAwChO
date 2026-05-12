- komentarze do kodu
- wrzucic wszystko na githuba
- zdjecia komend i działania

# Zadanie 1
## Opis aplikacji
Opracowana aplikacja została napisana w języku Go
Podczas uruchomienia aplikacja:
- wypisuje datę i godzinę startu,
- wyświetla dane autora,
- uruchamia serwer HTTP na porcie 8080,
- wysyła zapytania HTTP do zewnętrznego API,
- loguje przebieg działania aplikacji na standardowe wyjście.

## Użyte komendy
### Zbudowanie obrazu kontenera
```bash
docker build -t weather-app .
```
### Uruchomienie kontenera na podstawie zbudowanego obrazu
```bash
docker run -dp 8080:8080 --name weather-container weather-app
```
### Uzyskanie informacji z logów aplikacji
```bash
docker logs weather-container
```
### Sprawdzenie liczby warstw oraz rozmiaru obrazu
```bash
docker image history weather-app
```
```bash
docker image ls weather-app
```

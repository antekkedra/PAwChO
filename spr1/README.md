# Sprawozdanie nr 1
## Opis Dockerfile
Użycie najnowszego obrazu systemu Ubuntu jako bazy.
```dockerfile
FROM ubuntu:latest
```
Metadane o autorze
```dockerfile
LABEL author="Kędra s99569@pollub.edu.pl"
```

Aktualizacja systemu, instalacja Apache i czyszczenie cache.
```dockerfile
RUN apt update && \
    apt upgrade -y && \
    apt-get install -y apache2 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
```
Kopiowanie pliku index do katalogu domyślnego Apache.
```dockerfile
COPY index.html /var/www/html/
```
Nasłuchiwanie w porcie 80.
```dockerfile
EXPOSE 80
```
Uruchomienie Apache w trybie foreground żeby kontener nie zakończył działania
```dockerfile
CMD ["apache2ctl", "-D", "FOREGROUND"]
```
## Warstwy

Wynik polecenia `docker history web100`

```bash
IMAGE          CREATED          CREATED BY                                      SIZE      COMMENT
ccac04123f57   20 minutes ago   /bin/sh -c #(nop)  CMD ["apache2ctl" "-D" "F…   0B
4c8626ab1640   20 minutes ago   /bin/sh -c #(nop)  EXPOSE 80                    0B
a2deca0be878   20 minutes ago   /bin/sh -c #(nop) COPY file:8ca4971ce4c56c2f…   4.1kB
b0d7d4b62a80   21 minutes ago   /bin/sh -c apt update && apt upgrade -y && a…   133MB
13b6fa13d5ea   27 minutes ago   /bin/sh -c #(nop)  LABEL author=Kędra s99569…   0B
d1e2e92c075e   5 weeks ago      /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B
<missing>      5 weeks ago      /bin/sh -c #(nop) ADD file:1ae27d2ef43693611…   85.7MB
<missing>      5 weeks ago      /bin/sh -c #(nop)  LABEL org.opencontainers.…   0B
<missing>      5 weeks ago      /bin/sh -c #(nop)  LABEL org.opencontainers.…   0B
<missing>      5 weeks ago      /bin/sh -c #(nop)  ARG LAUNCHPAD_BUILD_ARCH     0B
<missing>      5 weeks ago      /bin/sh -c #(nop)  ARG RELEASE                  0B
```

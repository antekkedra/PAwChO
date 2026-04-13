# Sprawozdanie 2
## Opis Dockerfile

```dockerfile
COPY <<EOF /app/index.template.html
<h1>Sprawozdanie 2</h1>
<p>IP: {{IP}}</p>
<p>Hostname: {{HOST}}</p>
<p>Version: {{VERSION}}</p>
EOF
```

## Użyte komendy
### Polecenie użyte do budowy obrazu
```bash
docker build --build-arg VERSION=1.0 -t sprawozdanie2 .
```
### Polecenie uruchamiające serwer
```bash
docker run -dp 8080:80 --name spr2 sprawozdanie2
```
### Polecenie potwierdzające poprawne funkcjonowanie
```bash
docker inspect --format='{{.State.Health.Status}}' spr2
```
### Polecenie potwierdzające że aplikacja realizuje wymaganą funkcjonalność
```bash
curl http://localhost:8080
```

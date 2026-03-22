FROM ubuntu:latest

LABEL author="Kędra s99569@pollub.edu.pl"

RUN apt update && \
apt upgrade -y && \
apt-get install -y apache2 && \
apt-get clean && \
rm -rf /var/lib/apt/lists/*

COPY index.html /var/www/html/

EXPOSE 80

CMD ["apache2ctl", "-D", "FOREGROUND"]

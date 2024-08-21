FROM debian
MAINTAINER José Domingo Muñoz "josedom24@gmail.com"

RUN apt-get update && apt-get install -y apache2 && apt-get clean && rm -rf /var/lib/apt/lists/*

EXPOSE 80
ADD public_html/index.html /var/www/html/
ADD 000-default.conf /etc/apache2/sites-available/

ENV SERVER_NAME www.example.com

ENTRYPOINT ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]


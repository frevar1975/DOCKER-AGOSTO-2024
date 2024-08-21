FROM debian

MAINTAINER José Domingo Muñoz "josedom24@gmail.com"

RUN apt-get update && apt-get install -y apache2 && apt-get clean && rm -rf /var/lib/apt/lists/*

EXPOSE 80
ADD public_html/index.html /var/www/html/

CMD ["/usr/sbin/apache2ctl", "-D", "FOREGROUND"]


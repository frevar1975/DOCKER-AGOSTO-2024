FROM debian

MAINTAINER José Domingo Muñoz "josedom24@gmail.com"

RUN apt-get update && apt-get install -y apache2 && apt-get clean && rm -rf /var/lib/apt/lists/*

EXPOSE 80
ADD public_html/index.html /var/www/html/
ADD script.sh /usr/local/bin/

ENV SERVER_NAME www.example.com

CMD ["script.sh"]


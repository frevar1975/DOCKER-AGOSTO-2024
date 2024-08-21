#!/bin/bash

sed -i "s/#ServerName www.example.com/ServerName $SERVER_NAME/g" /etc/apache2/sites-available/000-default.conf 
apache2ctl -D FOREGROUND
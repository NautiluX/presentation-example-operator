#!/bin/bash


cat /templates/before.html > /usr/share/nginx/html/index.html
cat /config/slides.md >> /usr/share/nginx/html/index.html
cat /templates/after.html >> /usr/share/nginx/html/index.html

exec nginx -g 'daemon off;'

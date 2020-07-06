FROM alpine:3.10.3 as builder

# Add build dependencies
RUN apk add --update gcc musl-dev make libressl-dev file zlib-dev \
    libtool m4 automake autoconf bison flex git

# Clone repository and add diff
WORKDIR /opt/
RUN git clone https://bitbucket.org/tildeslash/monit.git
COPY monitctf.diff /opt/monit/

# Apply challenge diff, bootstrap and build monit 
WORKDIR /opt/monit/
RUN git checkout 10e823422b6edfe4081cfb2042bda9a229682753 \
    && git apply monitctf.diff && ./bootstrap \
    && ./configure --prefix=/opt/monit \
    --without-pam --without-ssl --without-largefiles \
    && make && make install

FROM alpine:3.10.3

ENV PATH=$PATH:/home/monit/bin \
    USER=monit

# Create user and group
RUN addgroup "$USER" \
    && adduser -D -g "" -h "/home/${USER}" -G "$USER" "$USER" 

# Copy binary and config file
COPY --from=builder  /opt/monit/bin/monit /home/monit/bin/
COPY --chown=monit:monit monitrc /home/monit/.monitrc

# Fix permissions on config file
RUN chmod 400 /home/monit/.monitrc

EXPOSE 2812
USER monit
CMD ["monit", "-I", "-B"]

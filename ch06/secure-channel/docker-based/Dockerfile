FROM alpine



RUN apk update && \
  apk add --no-cache openssl && \
  rm -rf /var/cache/apk/*
# cd /Users/growth/drinnovations/mywork_jmd/gogrpc/gooreillygrpcuprunning/samples/ch06/secure-channel/docker-based
# docker build -t myopenssl-alpine .
# docker run -v $PWD:/tls -ti myopenssl-alpine sh
#/ # pwd
#/
#/ # cd tls
#/tls # touch jmd
#/tls # openssl
#OpenSSL>
#OpenSSL> genrsa -out server.key 2048
#
#OpenSSL> req -new -x509 -sha256 -key server.key -out server.crt -days 3650
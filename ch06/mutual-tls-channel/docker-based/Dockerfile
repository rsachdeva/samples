FROM alpine

RUN apk update && \
  apk add --no-cache openssl && \
  rm -rf /var/cache/apk/*

# cd /Users/growth/drinnovations/mywork_jmd/gogrpc/gooreillygrpcuprunning/samples/ch06/mutual-tls-channel/docker-based
# docker build -t myopenssl-alpine .
# docker run -v $PWD:/tls -ti myopenssl-alpine sh
# Inside containmer
#/ # pwd
#/
#/ # cd tls
#/tls # touch jmd
#/tls # openssl

# Generate private RSA key


## Generate CA and self-signed certificates
# openssl genrsa -aes256 -out ca.key 4096
# openssl req -new -x509 -sha256 -days 730 -key ca.key -out ca.crt
# openssl x509 -noout -text -in ca.crt    // We can verify the root certificate using this command
#
# Generate server key and certificate - # using ca.crt and ca.key in last command below
# openssl genrsa -out server.key 2048
# openssl req -new -sha256 -key server.key -out server.csr
# openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 1 -out server.crt
#
# Generate client key and certificate - # using ca.crt and ca.key in last command below
# openssl genrsa -out client.key 2048
# openssl req -new -key client.key -out client.csr
# openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 2 -out client.crt

#A challenge password []:Something
#An optional company name []:LLC
#Enter pass phrase for ca.key: drinnovations
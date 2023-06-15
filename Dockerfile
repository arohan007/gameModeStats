FROM bitnami/golang:1.20.3

WORKDIR /app

#RUN apt-get install -y curl

#ARG cert_location=/usr/local/share/pers/ca-certificates

## Get certificate from "github.com"
#RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
## Get certificate from "proxy.golang.org"
#RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
## Update certificates
#RUN update-ca-certificates

# Verify-Peer=false - temp fix for certificate error for apt
#RUN apt-get -o "Acquire::https::Verify-Peer=false" update && apt-get -o "Acquire::https::Verify-Peer=false" install -y \
#    ca-certificates

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -o /main

EXPOSE 8080 8080

CMD ["/main"]
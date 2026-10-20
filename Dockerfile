FROM golang:latest AS builder
RUN apt-get update && apt-get install -y openssl && rm -rf /var/lib/apt/lists/*
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN curl -sSL https://github.com/zerodayz/dancing-gopher/archive/release.tar.gz | tar -v -C /app -xz
RUN cd dancing-gopher-release && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o web web.go
RUN mkdir -p /certs && \
    openssl ecparam -genkey -name secp384r1 -out /certs/tls.key && \
    openssl req -new -x509 -sha256 -key /certs/tls.key -out /certs/tls.crt -days 3650 -subj "/CN=dancing-gopher"

FROM scratch
LABEL maintainer="Robin Cernin <cerninr@gmail.com>"
COPY --from=builder /app/dancing-gopher-release /app
COPY --from=builder /certs /certs
WORKDIR /app
EXPOSE 8080 8443
CMD ["./web"]

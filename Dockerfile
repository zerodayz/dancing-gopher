FROM golang:latest as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN curl -sSL https://github.com/zerodayz/dancing-gopher/archive/release.tar.gz \
              | tar -v -C /app -xz
RUN cd dancing-gopher-release && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o web web.go
FROM scratch
LABEL maintainer="Robin Cernin <cerninr@gmail.com>"
COPY --from=builder /app/dancing-gopher-release /app
WORKDIR /app
CMD ["./web"]
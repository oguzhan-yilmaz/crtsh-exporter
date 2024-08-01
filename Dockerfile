ARG GOLANG_VERSION=1.22.0

ARG GOOS=linux
ARG GOARCH=amd64

ARG COMMIT
ARG VERSION

FROM docker.io/golang:${GOLANG_VERSION} as build

WORKDIR /crtsh-exporter

COPY go.* ./
COPY main.go .
COPY collector ./collector

ARG GOOS
ARG GOARCH

ARG VERSION
ARG COMMIT

RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
    go build \
    -ldflags "-X main.OSVersion=${VERSION} -X main.GitCommit=${COMMIT}" \
    -a -installsuffix cgo \
    -o /go/bin/exporter \
    ./main.go

FROM gcr.io/distroless/static-debian12:latest

LABEL org.opencontainers.image.description "Prometheus Exporter for crt.sh"
LABEL org.opencontainers.image.source https://github.com/DazWilkin/crtsh-exporter

COPY --from=build /go/bin/exporter /

EXPOSE 8080

ENTRYPOINT ["/exporter"]
CMD ["--endpoint=0.0.0.0:8080","--hosts=","--path=/metrics"]
FROM golang:1.13 AS builder

LABEL maintainer="Eddy Snow <eddysnow@gmail.com>"

WORKDIR /app

COPY fe-api/ .

RUN go get -u github.com/gin-gonic/gin

RUN go build .

FROM gcr.io/distroless/base-debian10

COPY --from=builder /app /

EXPOSE 8080

ENTRYPOINT [ "./app" ]
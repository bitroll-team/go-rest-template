# build

FROM golang:1.20.10-alpine3.17 AS build

RUN apk add upx
WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /opt/app
RUN upx /opt/app

# runner

FROM alpine:3.18.4

RUN adduser -S --disabled-password -h / -s /sbin/nologin -G nobody -u 3001 runner
USER runner

EXPOSE 8080
COPY --from=build /opt/app /opt/app
CMD /opt/app

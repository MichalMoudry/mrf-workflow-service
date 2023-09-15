FROM golang:1.21-alpine AS build
WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /build ./cmd/main.go

FROM alpine:3.17 AS release
WORKDIR /workflow-service

COPY config.json .
COPY --from=build /build ./build

EXPOSE 8443

ENTRYPOINT [ "/workflow-service/build" ]
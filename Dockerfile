# syntax=docker/dockerfile:1
FROM golang:1.22.5 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/terminal ./cmd/terminal.go

FROM alpine:3.20.2

RUN apk update && apk add --no-cache ca-certificates=20240705-r0 tzdata=2024a-r1

WORKDIR /src
COPY --from=build /bin/terminal /bin/terminal
RUN touch .env

CMD ["/bin/terminal"]
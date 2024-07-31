# syntax=docker/dockerfile:1
FROM golang:1.22.5 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/terminal ./main.go

FROM alpine:3.20.2

RUN apk add --no-cache ca-certificates=20240705-r0

WORKDIR /src
COPY --from=build /bin/terminal /bin/terminal
RUN touch .env

CMD ["/bin/terminal"]
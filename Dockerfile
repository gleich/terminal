# syntax=docker/dockerfile:1
FROM golang:1.24.5 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/terminal ./cmd/terminal.go

FROM alpine:3.20.2

WORKDIR /

RUN apk update && apk add --no-cache ca-certificates=20250619-r0 tzdata=2025b-r0 wget=1.24.5-r0 openssh=9.7_p1-r5

COPY --from=build /bin/terminal /bin/terminal
COPY --from=build /src/website/build ./website/build

CMD ["/bin/terminal"]

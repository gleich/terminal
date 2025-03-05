# syntax=docker/dockerfile:1
FROM golang:1.24.1 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/terminal ./cmd/terminal.go

FROM alpine:3.20.2

WORKDIR /bin

RUN apk update && apk add --no-cache ca-certificates=20241121-r1 tzdata=2025a-r0 wget=1.24.5-r0 openssh=9.7_p1-r5

RUN wget -q https://github.com/sorenisanerd/gotty/releases/download/v1.5.0/gotty_v1.5.0_linux_amd64.tar.gz \
    && tar xvfz gotty_v1.5.0_linux_amd64.tar.gz

COPY --from=build /bin/terminal /bin/terminal
COPY --from=build /src/entrypoint.sh /src/entrypoint.sh
COPY --from=build /src/gotty/.gotty /root/.gotty
COPY --from=build /src/gotty/index.html /src/gotty/index.html

CMD ["sh", "/src/entrypoint.sh"]

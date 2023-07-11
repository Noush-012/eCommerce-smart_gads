# build stage
FROM golang:1.20.4-alpine3.18 AS build-stage

# Maintainer info
LABEL maintainer="Noushad Ibrahim <noushadibrahim012@gmail.com>"

WORKDIR /home/app
COPY . .
RUN apk update && apk add --no-cache git
RUN go mod download
RUN go build -v -o /home/build/api ./cmd/api

# Final stage
FROM alpine:3.18

# Maintainer info
LABEL maintainer="Noushad Ibrahim <noushadibrahim012@gmail.com>"

WORKDIR /home/app
COPY --from=build-stage /home/build/api ./api
COPY .env .

CMD ["./api"]


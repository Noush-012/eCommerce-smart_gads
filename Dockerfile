# build stage
FROM golang:1.20.4-alpine3.18 AS build-stage

# Maintainer info
LABEL maintainer="Noushad Ibrahim <noushadibrahim012@gmail.com>"

WORKDIR /home/app
COPY . .
RUN apk update && apk add --no-cache git
RUN go mod download
RUN go build -v -o /home/build/api ./cmd/api

# Copy HTML files
# COPY ./views ./views

# Final stage
FROM gcr.io/distroless/static-debian11

# Maintainer info
LABEL maintainer="Noushad Ibrahim <noushadibrahim012@gmail.com>"

# WORKDIR /home/app
COPY --from=build-stage /home/build/api /api
COPY --from=build-stage /home/app/views /views
COPY .env /.env

CMD ["/api"]


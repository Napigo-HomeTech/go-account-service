FROM golang:1.17-alpine as build

WORKDIR /app
RUN apk update && apk add --no-cache git

COPY . .
COPY ./.env ./.env

# Install dependencies
RUN go mod download

# Build the app 
RUN go build -o /main



FROM alpine:3.10
RUN apk add --no-cache tzdata ca-certificates


COPY --from=build . .
COPY --from=build /app/.env ./.env

CMD ["/main"]

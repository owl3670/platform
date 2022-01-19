
##
## Build
##
FROM golang:1.17-alpine AS build

WORKDIR /app/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /server

##
## Deploy
##
FROM alpine

RUN mkdir /app
COPY --from=build /server /app

ENTRYPOINT ["/app/server"]
CMD ["ls", "app"]
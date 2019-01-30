# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add curl git gcc musl-dev
COPY . $GOPATH/app
WORKDIR $GOPATH/app
RUN go mod download
RUN go build -o main

# final stage
FROM alpine
COPY --from=build-env /go/app/main ./main
EXPOSE 3000
ENTRYPOINT ./main
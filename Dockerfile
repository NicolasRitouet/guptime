# build stage
FROM golang:alpine AS build-env
ARG Version
RUN apk --no-cache add curl git gcc musl-dev
COPY . $GOPATH/app
WORKDIR $GOPATH/app
RUN go mod download
RUN go build -o main -ldflags "-X main.Version=$Version"

# final stages
FROM alpine
WORKDIR /app
# need to add certificates to fetch data via https
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/app/main /app/main
RUN ls
EXPOSE 3000
ENTRYPOINT ./main
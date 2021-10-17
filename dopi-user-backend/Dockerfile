FROM golang:1.17.2 AS build
WORKDIR /src
COPY . .
RUN GOOS=linux GOARCH=arm64 go build main.go

FROM arm64v8/alpine
COPY --from=build /src/main /main
COPY .env .env
ENTRYPOINT [ "sh", "-c", "/main" ]
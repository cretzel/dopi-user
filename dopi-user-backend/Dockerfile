FROM golang:1.15.2-alpine AS build
RUN echo "foo"
WORKDIR /src
COPY . .
RUN GOOS=linux GOARCH=arm GOARM=5 go build main.go

FROM arm64v8/alpine
COPY --from=build /src/main /main
COPY .env .env
ENTRYPOINT [ "sh", "-c", "/main" ]
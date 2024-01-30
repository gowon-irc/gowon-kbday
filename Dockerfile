FROM golang:alpine as build-env
COPY . /src
WORKDIR /src
RUN go build -o gowon-kbday

FROM alpine:3.19.1
WORKDIR /app
COPY --from=build-env /src/gowon-kbday /app/
ENTRYPOINT ["./gowon-kbday"]

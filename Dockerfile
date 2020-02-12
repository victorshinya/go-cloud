# Build process
FROM golang as build

WORKDIR /go/src/github.com/victorshinya/go-cloud
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Expose service
FROM alpine

COPY --from=build /go/src/github.com/victorshinya/go-cloud .
EXPOSE 3000
ENTRYPOINT [ "./app" ]

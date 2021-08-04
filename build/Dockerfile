# Build process
FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]

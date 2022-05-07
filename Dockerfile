FROM golang

WORKDIR /app


COPY . .

RUN go get . && go build -o app main.go

EXPOSE 8080

CMD ["./app"]
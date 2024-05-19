FROM golang:1.22
LABEL authors="cuga"
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main", "test_file.txt"]
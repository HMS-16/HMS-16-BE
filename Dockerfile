FROM golang:1.19

WORKDIR /HMS-16-BE
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080
CMD ["go", "run", "main.go"]

FROM golang:1.25-alpine
#FROM golang

WORKDIR /usr/src/app


COPY . .

RUN go build cmd/backend/main.go 



# Copy go modules
#COPY go.mod go.sum ./
#RUN go mod tidy

EXPOSE 8080
ENTRYPOINT ["./main"]
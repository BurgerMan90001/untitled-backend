FROM golang:1.25-alpine

WORKDIR /usr/src/app
COPY . . 


# Copy go modules
#COPY go.mod go.sum ./
#RUN go mod tidy


COPY cmd/backend/main.go ./

RUN go build ./main.go
EXPOSE 8080

ENTRYPOINT ["./main"]
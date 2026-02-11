#FROM golang:1.25-alpine
FROM golang

WORKDIR /usr/src/app
COPY . . 


# Copy go modules
#COPY go.mod go.sum ./
#RUN go mod tidy


#COPY cmd/backend/main.go ./

RUN go build cmd/backend/main.go 


ENTRYPOINT ["./main"]
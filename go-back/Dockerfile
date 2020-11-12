FROM golang:1.15.4-alpine3.12
RUN apk update && apk upgrade && apk add git
RUN go get github.com/lib/pq
WORKDIR /app
COPY . . 
RUN go build härpäke.go availability.go candidates.go tokens.go votes.go votings.go authorisation.go voterapi.go
CMD ["./härpäke" ,"start"]

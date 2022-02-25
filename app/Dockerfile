FROM golang:latest as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/api ./main.go

FROM scratch as prod
COPY --from=builder /go/bin/api /go/bin/api
COPY ./config/environments/prod.yaml /secrets/prod.yaml
CMD ["/go/bin/api", "-e", "prod"]

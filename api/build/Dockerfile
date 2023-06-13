FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM gcr.io/distroless/base

WORKDIR /app

COPY --from=builder /app/configs/config.yml /app/configs/config.yml
COPY --from=builder /app/server /app/server

CMD ["/app/server"]
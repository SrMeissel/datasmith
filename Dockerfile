FROM golang:1.27 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o datasmith .

# ================

FROM scratch
COPY --from=builder /app/datasmith /datasmith

ENTRYPOINT ["/datasmith"]

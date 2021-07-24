FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main ./cmd/server

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main /

ENV isDocker="true"
ENV GO_ENV="development"
ENV PORT="8080"

ENTRYPOINT ["./main"]

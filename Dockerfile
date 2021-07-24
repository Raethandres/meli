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
ENV PORT="8110"
ENV MONGO_DATABASE_R1="test2-shard-00-00.jjsxj.mongodb.net:27017"
ENV MONGO_DATABASE_R2="test2-shard-00-01.jjsxj.mongodb.net:27017"
ENV MONGO_DATABASE_R3="test2-shard-00-02.jjsxj.mongodb.net:27017"
ENV MONGO_DATABASE="fire"
ENV USER_MONGO="andres"
ENV PW_MONGO="LM51e9I64WaP0lc0"

ENTRYPOINT ["./main"]

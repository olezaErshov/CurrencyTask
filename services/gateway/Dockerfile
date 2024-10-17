FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app

RUN go build -o  gateway_service ./cmd

RUN chmod +x gateway_service


FROM alpine:3

WORKDIR /app

COPY --from=build /app/gateway_service .

COPY --from=build /app/cmd ./config


EXPOSE 8000

ENTRYPOINT ["./gateway_service"]
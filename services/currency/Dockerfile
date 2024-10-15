FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app

RUN go build -o  currency_service ./cmd

RUN chmod +x currency_service


FROM alpine:3

WORKDIR /app

COPY --from=build /app/currency_service .

COPY --from=build /app/cmd ./config

COPY --from=build /app/database/migrations ./migrations


EXPOSE 8001

ENTRYPOINT ["./currency_service"]
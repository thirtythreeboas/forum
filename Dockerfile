FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o redditclone ./cmd/redditclone

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/redditclone .

CMD ["./forum"]
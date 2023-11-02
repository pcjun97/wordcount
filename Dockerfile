FROM golang:1.20 AS build-stage
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o wordcount ./cmd/wordcount

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=build-stage /app/wordcount /app/wordcount
USER nonroot:nonroot
ENTRYPOINT ["/app/wordcount"]

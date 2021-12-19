FROM registry.trendyol.com/platform/base/image/golang:1.17 as build

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go get github.com/vektra/mockery/.../
RUN go get github.com/swaggo/swag/cmd/swag
#Adding changed files last for hitting docker layer cache
COPY . .
RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/buybox-score-service

# Switch to a small base image
FROM scratch

# Get the TLS CA certificates from the build container, they're not provided by busybox.
COPY --from=build /etc/ssl/certs /etc/ssl/certs

# copy app to bin directory, and set it as entrypoint
WORKDIR /app
COPY --from=build /app/buybox-score-service /app/buybox-score-service
COPY --from=build /app/resources/ /app/resources

EXPOSE 8082

ENTRYPOINT ["/app/buybox-score-service"]

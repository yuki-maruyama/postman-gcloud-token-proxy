FROM golang:1.23.0 AS builder
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o postman-gcloud-token-proxy .

FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine AS runner
COPY --from=builder /app/postman-gcloud-token-proxy /app/postman-gcloud-token-proxy
EXPOSE 8080
CMD ["/app/postman-gcloud-token-proxy"]

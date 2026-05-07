# syntax=docker/dockerfile:1

FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY main.go ./
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-s -w" -o /out/pulse-svc .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /out/pulse-svc /pulse-svc
EXPOSE 9302
ENTRYPOINT ["/pulse-svc"]

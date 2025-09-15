# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.25.1 AS build
ARG TARGETARCH
ARG TARGETOS
WORKDIR /src

COPY go.mod .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 \
    GOARCH=$TARGETARCH \
    GOOS=$TARGETOS \
    go build -trimpath -ldflags="-s -w" -o /app/health

FROM gcr.io/distroless/base:latest
WORKDIR /app
COPY --from=build /app ./
ENTRYPOINT ["./health"]
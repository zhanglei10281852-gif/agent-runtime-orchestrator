FROM --platform=$BUILDPLATFORM golang:1.26.1-bookworm AS build
ARG TARGETOS
ARG TARGETARCH
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -trimpath -ldflags='-s -w' -o /out/agent-runtime-orchestrator ./cmd/server \
    && mkdir -p /out/data

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/agent-runtime-orchestrator /app/agent-runtime-orchestrator
COPY --from=build --chown=65532:65532 /out/data /data
ENV HTTP_ADDR=:8080 DATABASE_PATH=/data/agent-runtime.db BUSINESS_TIMEZONE=UTC
EXPOSE 8080
ENTRYPOINT ["/app/agent-runtime-orchestrator"]

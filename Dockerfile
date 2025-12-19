# syntax=docker/dockerfile:1

################################################################################
# Build stage
ARG GO_VERSION=1.25.5
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS build
WORKDIR /src

# Download dependencies
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

ARG TARGETARCH

# Build the application
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/server

################################################################################
# Final stage
FROM alpine:latest AS final

# Install runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user with host UID/GID
ARG UID=1000
ARG GID=1000
RUN addgroup -g "${GID}" appgroup && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "/app" \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    --ingroup appgroup \
    appuser

# Set working directory
WORKDIR /app

# Create data directory with correct ownership
RUN mkdir -p /app/data /app/tutorials && chown -R appuser:appgroup /app

# Copy the executable
COPY --from=build /bin/server /app/server

# Copy tutorials directory
COPY --chown=appuser:appuser tutorials/ /app/tutorials/

# Switch to non-root user
USER appuser

# Set environment variables
ENV PORT=8080
ENV TUTORIALS_DIR=/app/tutorials
ENV DATA_DIR=/app/data

EXPOSE 8080

ENTRYPOINT ["/app/server"]

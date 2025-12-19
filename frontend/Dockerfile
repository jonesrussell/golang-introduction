# syntax=docker/dockerfile:1

ARG NODE_VERSION=22-alpine

################################################################################
# Development stage - for local development with hot reload
FROM node:${NODE_VERSION} AS dev

# Create user with host UID/GID for proper file permissions
ARG UID=1000
ARG GID=1000
RUN deluser --remove-home node 2>/dev/null || true && \
    addgroup -g "${GID}" appgroup && \
    adduser -u "${UID}" -G appgroup -s /bin/sh -D appuser

WORKDIR /app
RUN chown appuser:appgroup /app

# Copy package files first for better caching
COPY --chown=appuser:appgroup package.json package-lock.json ./

# Switch to non-root user before npm install
USER appuser

# Install all dependencies (including devDependencies)
RUN npm ci

# Copy source files
COPY --chown=appuser:appgroup . .

# Expose Vite dev server port
EXPOSE 5173

# Run dev server with host binding for Docker
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]

################################################################################
# Build stage - for production builds
FROM node:${NODE_VERSION} AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build-only

################################################################################
# Production stage - serve built files
FROM nginx:alpine AS production

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]

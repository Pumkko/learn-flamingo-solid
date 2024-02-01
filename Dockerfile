# Example Dockerfile for Flamingo/Go based Projects
FROM node:20-alpine as nodebuild
WORKDIR /usr/src/app
COPY ./frontend /usr/src/app
RUN npm ci && npm run build

# Builder
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache ca-certificates tzdata git && update-ca-certificates
COPY . /app
RUN cd /app && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o klanikApp .

# Final image
FROM scratch

# add timezone data and ssl root certificates
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# add artifacts
ADD config/config.yml /usr/src/config/config.yml

# add binary
COPY --from=builder /app/klanikApp /usr/src/klanikApp
COPY --from=nodebuild /usr/src/app/dist /usr/src/frontend/dist


ENTRYPOINT ["/usr/src/klanikApp"]
CMD ["serve", "-a", ":8080"]
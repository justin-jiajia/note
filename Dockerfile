# Stage 1: Build Node.js application
FROM node:22 AS node-builder
WORKDIR /app/front
COPY . /app
RUN yarn install && yarn build

# Stage 2: Build Go application
FROM golang:1.23 AS go-builder
WORKDIR /app
COPY --from=node-builder /app /app
RUN CGO_ENABLED=0 go build -v -ldflags '-w -s' .

# Stage 3: Run the application
FROM alpine:3.14
WORKDIR /app
COPY --from=go-builder /app/note /app/note
EXPOSE 8080
CMD ["./note"]

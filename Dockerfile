# Stage 1: Build Node.js application
FROM node:22 AS node-builder
WORKDIR /app
COPY ./front /app
RUN yarn install && yarn build

# Stage 2: Build Go application
FROM golang:1.23 AS go-builder
WORKDIR /app
COPY . /app
RUN go build . -o note

# Stage 3: Run the application
FROM alpine:3.14
WORKDIR /app
COPY --from=go-builder /app/note /app/note
COPY --from=node-builder /app/build /app/front
EXPOSE 8080
CMD ["./note"]

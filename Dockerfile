# Step 1
FROM golang:1.19 as dev

RUN mkdir /work
WORKDIR /work

COPY . .
# Build Apk
RUN go mod init perqara
RUN go mod tidy
# Build swagger
RUN CGO_ENABLED=0 GOOS=linux go build -o ./perqara
# Step 2
FROM alpine:latest
COPY --from=dev /work/perqara .
RUN chmod +x /perqara
RUN apk add curl
CMD ["./perqara"]
## Build
FROM golang:latest AS build

# RUN mkdir /app
# ADD . /app
WORKDIR /cmd
# RUN CGO_ENABLED=0 GOOS=linux go build -o /Users/ronniezad/dev/nyumba-go-grpc-project/go-grpc-api-gateway/cmd

# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./.

COPY cmd/grpcsrv-gateway /cmd
COPY cmd/.env /cmd

## Deploy
FROM alpine:latest AS production

WORKDIR /

COPY --from=build /cmd .

EXPOSE 3000

RUN chmod +x ./grpcsrv-gateway

CMD ["./grpcsrv-gateway"]

## Build
FROM golang:latest AS build

WORKDIR /cmd
# RUN CGO_ENABLED=0 GOOS=linux go build -o /Users/ronniezad/dev/nyumba-go-grpc-project/go-grpc-api-gateway/cmd

COPY cmd/grpcsrv-application /cmd
COPY cmd/.env /cmd

## Deploy
FROM alpine:latest AS production

WORKDIR /

COPY --from=build /cmd .

EXPOSE 5001

RUN chmod +x ./grpcsrv-application

CMD ["./grpcsrv-application"]
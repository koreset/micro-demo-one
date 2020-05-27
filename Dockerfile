FROM golang:alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build


RUN go build -o api-1 .


FROM alpine:latest
RUN mkdir /app
WORKDIR /app


COPY --from=builder /build/api-1 /app/
ENV API_PORT=8085
EXPOSE $API_PORT
CMD ["./api-1"]
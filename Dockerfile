FROM golang:1.21-alpine

RUN apk add libc6-compat gcompat

WORKDIR /app

COPY build/* ./

EXPOSE 8080

CMD [ "/app/stack", "serve" ]
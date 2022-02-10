FROM golang:1.16-alpine AS build
WORKDIR /app
ADD . /app
RUN apk add build-base
RUN echo "Starting Build" && \
    CC=$(which musl-gcc) go build -buildmode=pie -trimpath --ldflags '-w -linkmode external -extldflags "-static"' && \
    echo "Completed Build" 

FROM scratch

WORKDIR /app

COPY --from=build /app/hellohttp /app/hellohttp

# Set the PORT you want to expose
ARG PORT "8080"
ENV HTTP_PORT "$PORT"

# Enable access logs
ENV ENABLE_ACCESS_LOGGING "YES"

# A friendly string
ENV HELLO_STRING "HELLO"

EXPOSE $PORT

CMD ["/app/hellohttp"] 



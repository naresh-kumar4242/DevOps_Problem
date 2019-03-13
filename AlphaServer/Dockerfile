# Build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o goapp

# Final stage - Serves binary only 
FROM alpine
WORKDIR /app
COPY --from=build-env /src/goapp /app/
EXPOSE 9001
ENTRYPOINT ./goapp

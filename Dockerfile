FROM golang:1.20-bullseye as builder

ARG GIT_COMMIT
ARG BUILD_TIME
ARG VERSION

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags "-X 'github.com/farovictor/MongoDbExtractor/src/cmd.GitCommit=${GIT_COMMIT}' -X 'github.com/farovictor/MongoDbExtractor/src/cmd.BuildTime=${BUILD_TIME}' -X 'github.com/farovictor/MongoDbExtractor/src/cmd.Version=${VERSION}' -s -w" -o mongoextract -a -installsuffix cgo ./extractor/src && \
    chmod a+x mongoextract

FROM alpine:latest

COPY --from=builder /app/mongoextract /usr/bin/mongoextract

ENTRYPOINT [ "mongoextract" ]

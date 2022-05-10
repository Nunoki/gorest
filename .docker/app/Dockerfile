# build an image of golang with air; which watches for
# file changes, and the re-compiles and re-runs the app.
# we are using vanilla golang, because the alpine version
# of golang doesn't have git installed
FROM golang:1.17
RUN go install github.com/cosmtrek/air@latest

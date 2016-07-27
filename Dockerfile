FROM golang:1.7

ENV GOOS=linux
ENV GOARCH=arm

COPY . /app
WORKDIR /app
CMD /app/go.sh
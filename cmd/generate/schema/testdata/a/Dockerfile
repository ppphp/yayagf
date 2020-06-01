
FROM golang as back

ENV GOPROXY=https://goproxy.io
ENV GOSUMDB=off
ENV GOPRIVATE=gitlab.papegames.com/*
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /main

ADD go.mod ./
ADD go.sum ./
RUN go mod download

ADD app ./app/
ADD main.go ./
RUN go build -o /main/main

FROM scratch
WORKDIR /main
COPY --from=back /main/main .

CMD ["/main/main"]


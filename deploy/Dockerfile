FROM golang:alpine AS lrsccomp
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn
WORKDIR /home
COPY . .
RUN go mod tidy && cd src && go build -o lrsc && mv /home/src/lrsc /home/bin

FROM lrsccomp
WORKDIR /home
COPY --from=lrsccomp /home/bin/lrsc .
EXPOSE 8080
ENTRYPOINT ["./lrsc"]
FROM golang:latest AS builder
COPY . /app/backend
WORKDIR /app/backend
RUN go get github.com/labstack/echo/...
# linux/amd64用のバイナリファイルを生成する。ビルドするディレクトリはカレントディレクトリ以下で、生成したビルドファイルは/main.exeになる。
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 1323

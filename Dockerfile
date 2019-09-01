from golang as builder

workdir /app
copy . .

run go get -d ./... && \
    go install ./...

run CGO_ENABLED=0 GOOS=linux go build -a -o tokens . && \
  go get github.com/cortesi/modd/cmd/modd

cmd ["modd"]

from golang:alpine

workdir /app
copy --from=builder /app/tokens .

cmd ["/app/tokens"]


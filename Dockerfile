from golang as builder

workdir /app
copy . .

run go get -d ./... && \
    go install ./... && \
    env GO111MODULE=on go get github.com/cortesi/modd/cmd/modd

run CGO_ENABLED=0 GOOS=linux go build -a -o tokens .

cmd ["modd"]

from golang:alpine

workdir /app
copy --from=builder /app/tokens .

cmd ["/app/tokens"]


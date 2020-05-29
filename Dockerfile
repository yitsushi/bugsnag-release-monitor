from alpine:3.11 as builder
run apk add --no-cache go
copy . /code
workdir /code
run go build ./cmd/bugsnag-release-monitor

from alpine:3.11
run apk --no-cache add ca-certificates
copy --from=builder /code/bugsnag-release-monitor /github-pr-creator
entrypoint ["/github-pr-creator"]

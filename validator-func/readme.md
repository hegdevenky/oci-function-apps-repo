```shell
echo -n '{"inputString": "hello","operation": "UPPER", "dryRun": true}' | fn invoke dlcdep-5738-test-app string-util-func --content-type application/json
```

```dockerfile
FROM fnproject/go:1.19-dev as build-stage
WORKDIR /function
WORKDIR /go/src/func/
ENV GO111MODULE=on
COPY . .
RUN go mod tidy
RUN go build -o func -v
FROM fnproject/go:1.19
WORKDIR /function
COPY --from=build-stage /go/src/func/func /function/
ENTRYPOINT ["./func"]
```
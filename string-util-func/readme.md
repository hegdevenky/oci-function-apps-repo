```shell
echo -n '{"inputString": "hello","operation": "UPPER", "dryRun": true}' | fn invoke <application-name> string-util-func --content-type application/json
```
// dlcdep-5738-test-app
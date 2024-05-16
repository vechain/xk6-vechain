To install build the k6 binary:

```bash
CGO_ENABLED=1 xk6 build --with github.com/darrenvechain/xk6-vechain=.
```

To run the tests:

```bash
./k6 run test.js
```

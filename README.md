# xk6-vechain

This repository contains the xk6 extension for interacting with the VeChain blockchain.

## Usage

To build the executable:
```bash
CGO_ENABLED=1 xk6 build --with github.com/darrenvechain/xk6-vechain=.
```

Start thor solo:
```
thor solo
```

To run the tests:

```bash
./k6 run test.js
```

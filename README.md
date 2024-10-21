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

**Optional**: Start Grafana + InfluxDB:

```bash
docker compose up -d --wait
```

To run the tests:

```bash
./k6 run --out influxdb=http://localhost:8086/k6 test.js
```

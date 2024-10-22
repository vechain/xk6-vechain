# xk6-vechain

This repository contains the xk6 extension for interacting with the VeChain blockchain.

## Usage

To build the executable:

```bash
CGO_ENABLED=1 xk6 build --with github.com/darrenvechain/xk6-vechain=. --with github.com/grafana/xk6-dashboard@latest --with github.com/grafana/xk6-output-influxdb@latest
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
K6_INFLUXDB_ORGANIZATION=vechain \
K6_INFLUXDB_BUCKET=vechain \
K6_INFLUXDB_TOKEN=admin-token \
./k6 run -o xk6-influxdb=http://localhost:8086 test.js
```

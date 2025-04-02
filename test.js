import vechain from "k6/x/vechain";
import http from "k6/http";
import {sleep, check} from "k6";

const stage = [
    { target: 33, duration: '10m' },
    { target: 42, duration: '5m' },
    { target: 36, duration: '5m' },
    { target: 20, duration: '10m' },
    { target: 15, duration: '5m' },
]

const stages = []
for (let i = 0; i < 10_000; i++) {
    for (let j = 0; j < stage.length; j++) {
        const st = stage[j]
        stages.push(st)
    }
}

export const options = {
    scenarios: {
        contacts: {
            executor: 'ramping-arrival-rate',
            // Start with `startRate` transactions per block. Eg set this to 10 to achieve 10 txs per block.
            startRate: 36,
            // Set the time unit to 10 seconds (ie. 1 block)
            timeUnit: '10s',
            // Pre-allocate necessary VUs.
            preAllocatedVUs: 50,
            maxVUs: 50,
            stages,
        },
    },
    tags: {
        test_name: "vechain-toolchain",
        test_run_id: 'v2.6.0'
    }
};

const url = "http://localhost:8669";

const thor = vechain.Client({
    url: url,
    mnemonic:
        "denial kitchen pet squirrel other broom bar gas better priority spoil cross",
    accounts: 100,
});

export default function (setup) {
    const rawTx = thor.newToolchainTransaction(setup.contracts[0]);
    const body = {
        raw: rawTx
    };

    const res = http.post(`${url}/transactions`, JSON.stringify(body));

    check(res, {
        "is status 200": (r) => r.status === 200,
    });

    sleep(5);
}

export function setup() {
    console.log("Setting up test");
    const contracts = thor.deployToolchain(1);
    return {contracts};
}

export function teardown() {
    console.log("Tearing down test");
}

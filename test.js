import vechain from "k6/x/vechain";
import http from "k6/http";
import {check} from "k6";
import {sleep} from "k6";

export const options = {
    scenarios: {
        contacts: {
            executor: 'ramping-arrival-rate',
            // Start with `startRate` transactions per block. Eg set this to 10 to achieve 10 txs per block.
            startRate: 10,
            // Set the time unit to 10 seconds (ie. 1 block)
            timeUnit: '10s',
            // Pre-allocate necessary VUs.
            preAllocatedVUs: 100,
            maxVUs: 100,
            stages: [
                { target: 25, duration: '30s' },
                { target: 35, duration: '30s' },
                { target: 60, duration: '1m' },
                { target: 60, duration: '30s' },
                { target: 10, duration: '30s' },
            ],
        },
    },
    tags: {
        test_name: "vechain-toolchain",
        test_run_id: new Date().toISOString(),
    }
};

const url = "http://localhost:8669";

const thor = vechain.Client({
    url: url,
    mnemonic:
        "denial kitchen pet squirrel other broom bar gas better priority spoil cross",
    accounts: 1000,
});

export default function (setup) {
    const rawTx = thor.newToolchainTransaction(setup.contracts[0]);
    const body = {
        raw: `0x${rawTx}`,
    };

    const res = http.post(`${url}/transactions`, JSON.stringify(body));

    check(res, {
        "is status 200": (r) => r.status === 200,
    });

    sleep(5);
}

export function setup() {
    console.log("Setting up test");
    const tenThousandVET = "21E19E0C9BAB2400000";
    thor.fund(10, tenThousandVET);
    const contracts = thor.deployToolchain(1);
    return {contracts};
}

export function teardown() {
    console.log("Tearing down test");
}

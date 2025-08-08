import { check, sleep } from "k6";
import http from "k6/http";
import vechain from "k6/x/vechain";

// ramping-arrival-rate
// const stages = [
//     { target: 44, duration: "10s" },
//     { target: 44, duration: "10m" }
// ]

export const options = {
    scenarios: {
        // ramping-arrival-rate
        // contacts_ramping: {
        //     executor: 'ramping-arrival-rate',
        //     // Start with `startRate` transactions per block. Eg set this to 10 to achieve 10 txs per block.
        //     startRate: 36,
        //     // Set the time unit to 10 seconds (ie. 1 block)
        //     timeUnit: '10s',
        //     // Pre-allocate necessary VUs.
        //     preAllocatedVUs: 50,
        //     maxVUs: 50,
        //     stages,
        //     gracefulStop: '30s',
        // },
        
        // constant-arrival-rate
        contacts: {
            executor: 'constant-arrival-rate',
            // Maintain constant rate
            rate: 22,
            timeUnit: '10s',
            // Pre-allocate necessary VUs.
            preAllocatedVUs: 50,
            maxVUs: 50,
            // No time limit - runs until you stop it
            duration: '0s',
            gracefulStop: '30s',
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

    console.log("sending tx");

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

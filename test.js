import vechain from "k6/x/vechain";
import http from "k6/http";
import {sleep, check} from "k6";

export const options = {
    scenarios: {
        contacts: {
            executor: 'ramping-arrival-rate',
            // Start with `startRate` transactions per block. Eg set this to 10 to achieve 10 txs per block.
            startRate: 22,
            // Set the time unit to 10 seconds (ie. 1 block)
            timeUnit: '10s',
            // Pre-allocate necessary VUs.
            preAllocatedVUs: 100,
            maxVUs: 100,
            stages: [
                { target: 22, duration: '1m' }, //~50% utilization for 1 minute
                { target: 33, duration: '1m' }, //~75% utilization for 1 minute
                { target: 44, duration: '1m' }, //~100% utilization for 1 minute
                { target: 33, duration: '1m' }, //~75% utilization for 1 minute
            ],
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

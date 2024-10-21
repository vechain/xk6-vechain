import vechain from "k6/x/vechain";
import http from "k6/http";
import {check} from "k6";

export const options = {
    scenarios: {
        constant_arrival_rate: {
            executor: "constant-arrival-rate",
            // How long the test lasts
            duration: "1m",
            // 60 transactions per block
            rate: 60,
            // 10 seconds -> 1 block
            timeUnit: "10s",
            preAllocatedVUs: 10,
            maxVUs: 10,
        },
    },
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

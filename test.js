import vechain from 'k6/x/vechain'
import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 1,
    iterations: 1,
}

const url = "http://localhost:8669"

const thor = vechain.Client({
    url: url,
    mnemonic: "denial kitchen pet squirrel other broom bar gas better priority spoil cross",
})

export default function (setup) {
    const rawTx = thor.newToolchainTransaction(setup.contracts[0])
    const body = {
        raw: `0x${rawTx}`,
    }

    const res = http.post(`${url}/transactions`, JSON.stringify(body))

    check(res, {
        'is status 200': (r) => r.status === 200,
    });

    const receipt = thor.waitForTx(res.json("id"))

    check(res, {
        'tx is successful': (r) => receipt.reverted === false,
    });
}

export function setup() {
    console.log("Setting up test");
    const contracts = thor.deployToolchain(1)
    return { contracts }
}

export function teardown() {
    console.log("Tearing down test");
}

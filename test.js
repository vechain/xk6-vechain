import transactions from 'k6/x/vechain/transactions';

export const options = {
    vus: 1,
    iterations: 10
}

export default function () {
    // log each key on account
    console.log(transactions.buildTransaction());
}

export function setup() {
    console.log("Setting up test")
}

export function teardown() {
    console.log("Tearing down test");
}

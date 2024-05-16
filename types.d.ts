declare module 'k6/x/vechain/'{
  module 'accounts'{
    type Accounts = {
      newAddress: () => string
      newPrivateKey: () => string
    }
    export default Accounts;
  }

  module 'transactions'{
    type Transactions = {
      signTransaction: (tx: string, privateKey: string) => string
      buildTransaction: () => string
    }
    export default Transactions;
  }
}

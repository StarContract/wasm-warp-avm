# 🐿 Warp contracts - Go template

> ⚠️ This template is using previous version of Warp SDK API (`1.0.1`). It will soon be updated to the `1.2.x` version.

Following repository is a template for writing SmartWeave contracts in Go and building them into WASM binaries which can be then processed by Warp SDK.

It contains an example implementation of a PST contract - which you can use as a base for implementing your own contract.
If you are not familiar with the concept of Profit Sharing Tokens, check out a [tutorial](https://academy.warp.cc/docs/pst/introduction/intro) for writing your first PST contract in our Warp Academy.

- [Installation](#-installation)
- [Code structure](#-code-structure)
- [Writing contract](#-writing-contract)
  - [Accessing JavaScript imports](#accessing-javascript-imports)
  - [Foreign contract read](#foreign-contract-read)
- [Build](#-build)
- [Tests](#-tests)
- [Deploy](#-deploy)
- [Using SDK](#-using-sdk)

## 📦 Installation

You will need:

- Go 1.17 (https://go.dev/doc/install).  
  💡**NOTE** Go 1.18 is currently _NOT_ supported.  
  In case of issues on MacOS - make sure that these env. variables are set properly

```bash
export GOPATH=$HOME/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOBIN
export PATH=$PATH:$GOROOT/bin
```

- [tinygo](https://tinygo.org/getting-started/install/) (a compiler that produces much smaller binaries (up to 10x) than the default Go compiler)
- [easyjson](https://github.com/mailru/easyjson#install) (a code generation tool for handling json marshalling/unmarshalling)
- [Node.js](https://nodejs.org/en/download/) version 16.5 or above
- [yarn](https://yarnpkg.com/getting-started/install) installed

To install all Node.js dependencies run the following command:

```bash
yarn install
```

## 🔍 Code structure

- [deploy/](deploy) - contains deployment scripts for localhost/testnet/mainnet and contract's initial state definition
- [.out/](.out) - generated by `tinygo` compiler during build process - contains compiled wasm binary
- [src/](src) - contains the source code of the contract
  - [common/](src/common) - package for commons code, that will be reused between contracts (
    handles the low-level WASM-JS communication for Go). Contains `SwContract` interface
    that has to implemented by contract developers.  
    🔥Do Not Edit🔥, unless you really know what you're doing :-)
  - [common_types/](src/common_types) - package for commons structs. Have to be a separate package,
    otherwise `easyjson` throws error during code generation
  - [impl/](src/impl) - package that contains implementation of the given contract (i.e. implementation of the `SwContract`
    interface)
    - [impl/contract.go](src/impl/contract.go) - contains the logic of the contract, e.g. mapping from actions to functions,
      that will be called
    - [impl/actions.go](src/impl/actions.go) - contains the logic of all the contract's functions.
  - [types/](src/types) - contains definition of all the types used by the contract (state, actions schemas, etc)
  - [main.go](src/main.go) - entry point for the contract. Initialises contract's module and registers it on host.
- [tests/](tests) - contains integration tests written in Jest

## 🧑‍💻 Writing contract

If you want to edit contract's code and create your own implementation you can do it by following these steps:

1. Edit `init-state.json` by adding the initial state for your contract - [deploy/state/init-state.json](deploy/state/init-state.json)
2. Modify the state definition of the contract - [src/types/types.go](src/types/types.go#L7).  
    Add schemas which should describe input and output types for your actions - [src/types/types.go](src/types/types.go#L16).  
   Each time you change the definitions in the file above, run

```bash
yarn generate
```

3. Edit/add actions which user will be able to call while interacting with the contract - [src/impl/actions](src/impl/actions.go).
4. Add above action functions to the pattern matching in `Handle` function in [src/impl/contract.go](src/impl/contract.go#L16)

### Accessing JavaScript imports

An example of how to access imports can be found here: [impl/contract.go](src/impl/contract.go#L17)

### Foreign contract read

An example of how to read other contract state can be found here: [impl/actions.go](src/impl/actions.go#L56)

## 👷 Build

Compile your contract to WASM binary by running following command:

```bash
yarn build
```

This command will both generate the code for the json marshalling/unmarshalling and compile the contract with `tinygo` compiler.

## 🧪 Tests

Write tests for your contract (we will use Jest library for testing) - you can find a template in the [tests/](tests) folder.
Run tests with

```bash
yarn test
```

## 📜 Deploy

Deploy your contract to one of the networks (mainnet/Warp public testnet/localhost) by running following command (`network`: `mainnet` | `testnet` | `local`)

```bash
yarn deploy:[network]
```

💡**NOTE**: If you want to deploy your contract locally you need to run Arlocal by typing following command:

```bash
npx arlocal
```

💡**NOTE**: When using mainnet please put your wallet key in [deploy/mainnet/.secrets/wallet-mainnet.json](deploy/mainnet/.secrets/wallet-mainnet.json). `.secrets` folder has been added to `.gitignore` so your key is kept securely.

You can view deploy script code [here](deploy/scripts/deploy.js).

## 🟥 Using SDK

Optionally - you can run one of the scripts which uses Warp SDK to interact with the contract. Using SDKs' methods works exactly the same as in case of a regular JS contract.

💡**NOTE** You will need to have a file with the wallet key and a file with the contract id to run these scripts. If you do not have them please run a [deploy](#-deploy) script.

1. `read` - reads contract state, check out the code in [deploy/scripts/read-contract-state.js](deploy/scripts/read-contract-state.js)

```bash
    npm run read:[network]
```

2. `balance` - get balance for a wallet address, check out the code in [deploy/scripts/interact-balance.js](deploy/scripts/interact-balance.js)

```bash
    npm run balance:[network]
```

3. `transfer` - transfer specific amount of tokens to the indicated wallet, check out the code in [deploy/scripts/interact-transfer.js](deploy/scripts/interact-transfer.js)

##🧑‍💻ImgContract

This is a custom contract example that includes methods for uploading and cropping images.
Here are the usage steps:

1.
```bash
yarn generate
```
2.
```bash
yarn build
```
3.
```bash
npx arlocal
```
4.
```
yarn deploy:local
```
5.
```
npm run read:local
```
6.
```
node deploy/local/interact-upLoad-local.js
```
7.
```
node deploy/local/interact-cropImg-local.js
```

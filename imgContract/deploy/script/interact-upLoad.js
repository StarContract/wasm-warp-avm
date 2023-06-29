const { loadWallet } = require('./utils/load-wallet');
const { connectArweave } = require('./utils/connect-arweave');
const { connectContract } = require('./utils/connect-contract');
const { contractTxId } = require('./utils/contract-tx-id');
const { mineBlock } = require('./utils/mine-block');
const fs = require('fs');

//import { interactWrite } from 'smartweave'
const { interactWrite } = require('smartweave');
module.exports.interactUpLoad = async function (host, port, protocol, target, walletJwk) {
  const arweave = connectArweave(host, port, protocol);
  const wallet = await loadWallet(arweave, walletJwk, target, true);
  const txId = contractTxId(target);
  const img = await connectContract(arweave, wallet, txId);

  const filePath = '/root/workspace/goProject/wrap_wasm/warp-wasm-templates/go/deploy/scripts/utils/rabbit.jpg';
  const base64String = readFileAsBase64(filePath);
  //console.log(base64String);


  const upLoadId = await img.bundleInteraction({
    function: 'upLoad',
    //imageData: base64String,
    imageData: 'rabbit.jpg',
    name: 'rabbit'
  }
  );

  await mineBlock(arweave);

  const state = await img.readState();

  console.log('Updated state:', state);
  //console.log('Contract tx id', txId);

  if (target == 'testnet') {
    console.log(`Check upLoadId interaction at https://sonar.warp.cc/#/app/interaction/${upLoadId}?network=testnet`);
  } else {
    console.log('upLoad tx id', upLoadId);
  }
};

function readFileAsBase64(filePath) {
  const fileData = fs.readFileSync(filePath);
  const base64String = fileData.toString('base64');
  return base64String;
}




  // // 构造 upLoad 的交互数据
  // const action = {
  //   Function: "upLoad", // 调用的函数名称，可以是 "upLoad" 或 "cropImg"
  //   imageData: "rabbit.jpg",
  //   name: "rabbit"
  // };
  // const actionBytes = Buffer.from(JSON.stringify(action));

  // // 调用合约的 Handle 函数进行交互
  // const upLoadId = await img.writeInteraction(action, actionBytes);

  // if (error) {
  //   console.log("Error:", error);
  // } else {
  //   console.log("Updated State:", state);
  //   console.log("Result:", result);
  // }


  // await mineBlock(arweave);

  // console.log('Updated state:', state);
  // console.log('Contract tx id', upLoadId);

//   if (target == 'testnet') {
//     console.log(`Check upLoadId interaction at https://sonar.warp.cc/#/app/interaction/${upLoadId}?network=testnet`);
//   } else {
//     console.log('upLoad tx id', upLoadId);
//   }
//};




// const contract = new ImgContract(); // 创建 ImgContract 实例

// 构造 Action 对象
// const action = {
//   function: "upLoad", // 调用的函数名称，可以是 "upLoad" 或 "cropImg"
//   imageData: "rabbit.jpg",
//   name: "rabbit"
// };

// // 将 Action 对象转换为 JSON 字符串
// const actionJson = JSON.stringify(action);

// // 调用 ImgContract 的 Handle 方法执行交互交易
// const [result, actionResult, error] = img.Handle(action, actionJson);
// // if (error) {
// //   // 处理错误
// // }
// await mineBlock(arweave);

// // 处理结果
// // result 是函数调用返回的状态对象
// // actionResult 是交互交易的结果
// // ...

// const transactionId = result.TransactionId(); // 获取交易 ID
// const updatedState = result.State(); // 获取更新的状态

// // // 打印交易 ID 和更新的状态
// console.log("Transaction ID:", transactionId);
// console.log("Updated State:", updatedState);
//};
// import Web3 from "web3";

// let web3;

// export const loadWeb3 = async () => {
//   try {
//     if (window.ethereum) {
//       web3 = new Web3(window.ethereum);
//       const accounts = await window.ethereum.request({
//         method: "eth_requestAccounts",
//       });
//       console.log("Connected accounts:", accounts);
//       return { web3, accounts, error: null };
//     } else {
//       throw new Error(
//         "Non-Ethereum browser detected. Please install MetaMask."
//       );
//     }
//   } catch (error) {
//     console.error("MetaMask Connection Error:", error.message);
//     return { web3: null, accounts: [], error: error.message };
//   }
// };

// export const getBalance = async (address) => {
//   if (!web3) return "0";
//   const balance = await web3.eth.getBalance(address);
//   return web3.utils.fromWei(balance, "ether"); // Convert Wei to Ether
// };

// export const getAccounts = async () => {
//   const accounts = await web3.eth.getAccounts();
//   return accounts;
// };
import Web3 from "web3";
import { ethers } from 'ethers';
let web3;
let selectedAccount = null;

// Fungsi untuk menghubungkan ke MetaMask dan mendapatkan semua akun
export const loadWeb3 = async () => {
  try {
    if (window.ethereum) {
      web3 = new Web3(window.ethereum);
      const accounts = await window.ethereum.request({
        method: "eth_requestAccounts",
      });

      selectedAccount = accounts[0]; // Akun pertama sebagai default
      console.log("Connected accounts:", accounts);

      return { web3, accounts, error: null };
    } else {
      throw new Error(
        "Non-Ethereum browser detected. Please install MetaMask."
      );
    }
  } catch (error) {
    console.error("MetaMask Connection Error:", error.message);
    return { web3: null, accounts: [], error: error.message };
  }
};

// Fungsi untuk mendapatkan saldo akun
export const getBalance = async (address) => {
  try {
    if (!web3)
      throw new Error(
        "Web3 is not initialized. Please connect to MetaMask first."
      );
    const balance = await web3.eth.getBalance(address);
    return { balance: web3.utils.fromWei(balance, "ether"), error: null };
  } catch (error) {
    console.error("Balance Fetch Error:", error.message);
    return { balance: "0", error: error.message };
  }
};

// Fungsi untuk mendapatkan informasi jaringan
export const getNetwork = async () => {
  try {
    if (!web3)
      throw new Error(
        "Web3 is not initialized. Please connect to MetaMask first."
      );
    const networkId = await web3.eth.net.getId();
    const chainId = await web3.eth.getChainId();
    return { networkId, chainId, error: null };
  } catch (error) {
    console.error("Network Fetch Error:", error.message);
    return { networkId: null, chainId: null, error: error.message };
  }
};

// Fungsi untuk disconnect dari MetaMask
export const disconnectMetaMask = async () => {
  try {
    selectedAccount = null;
    return { success: true, error: null };
  } catch (error) {
    console.error("Disconnect Error:", error.message);
    return { success: false, error: error.message };
  }
};

// Event listener untuk perubahan akun atau jaringan
export const listenForAccountChange = (callback) => {
  if (window.ethereum) {
    window.ethereum.on("accountsChanged", (accounts) => {
      console.log("Account changed:", accounts);
      callback(accounts);
    });

    window.ethereum.on("chainChanged", (chainId) => {
      console.log("Chain changed:", chainId);
      callback([]);
    });
  }
};

export async function issueDegree(degreeHash, sekolah, issueDate) {
  if (!window.ethereum) throw new Error("MetaMask tidak terdeteksi!");
  
  const provider = new ethers.providers.Web3Provider(window.ethereum);
  const signer = provider.getSigner();
  const contract = new ethers.Contract(
    IjazahContract.address,
    IjazahContract.abi,
    signer
  );

  const tx = await contract.issueDegree(degreeHash, sekolah, issueDate);
  await tx.wait();
  return tx.hash;
}
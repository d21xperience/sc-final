import Web3 from "web3";

let web3;

export const loadWeb3 = async () => {
  if (window.ethereum) {
    web3 = new Web3(window.ethereum);
    // try {
    // await window.ethereum.enable() //deprecated
    window.ethereum
      .request({ method: "eth_requestAccounts" })
      .then(function (accounts) {
        // Pengguna memberikan akses ke akun mereka
        console.log(accounts);
        return web3;
      })
      .catch(function (error) {
        // Pengguna menolak memberikan akses ke akun mereka
        console.error(error);
      });
    // } catch (error) {
    //   console.log("user denied account access");
    // }
  } else if (window.web3) {
    web3 = new Web3(window.web3.currentProvider);
    return web3;
  } else {
    console.error(
      "Non-Ethereum browser detected, Yout should consider trying Netmask"
    );
  }
};

export const getAccounts = async () => {
  const accounts = await web3.eth.getAccounts();
  return accounts;
};

export const getBalance = async (address) => {
  const balance = await web3.eth.getBalance(address);
  return web3.utils.fromWei(balance, "ether"); //Mengkonversi balance dari eth ke wei
};

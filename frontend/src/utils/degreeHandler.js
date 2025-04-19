// File: src/utils/degreeHandler.js

import { BrowserProvider, Contract, keccak256, AbiCoder } from "ethers";
import contractABI from "@/VerifikasiIjazahABI.json";

const contractAddress = "0x700b6A60ce7EaaEA56F065753d8dcB9653dbAD35";

export async function prepareDegreeIssue(degreeData, sekolah, ipfsUrl, transcript) {
  if (!window.ethereum) throw new Error("Instal MetaMask!");

  const provider = new BrowserProvider(window.ethereum);
  const signer = await provider.getSigner(); // perlu await di ethers v6

  const contract = new Contract(contractAddress, contractABI, signer);

  const abiCoder = AbiCoder.defaultAbiCoder();
  const encodedData = abiCoder.encode(
    ["string", "string", "string[]", "uint8[]"],
    [degreeData, sekolah, transcript.subjects, transcript.grades]
  );

  const degreeHash = keccak256(encodedData);

  const gasEstimate = await contract.issueDegree.estimateGas(
    degreeHash,
    sekolah,
    Math.floor(Date.now() / 1000),
    ipfsUrl,
    transcript.subjects,
    transcript.grades
  );

  return { degreeHash, gasEstimate };
}

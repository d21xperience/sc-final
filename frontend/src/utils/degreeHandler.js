// File: src/utils/degreeHandler.js
import { ethers } from 'ethers';
import contractABI from '@/VerifikasiIjazahABI.json';

const contractAddress = "0x700b6A60ce7EaaEA56F065753d8dcB9653dbAD35"; // Alamat smart contract

export async function prepareDegreeIssue(
  degreeData, 
  sekolah, 
  ipfsUrl, 
  transcript
) {
  // 1. Hubungkan ke MetaMask
  if (!window.ethereum) throw new Error("Instal MetaMask!");
  const provider = new ethers.providers.Web3Provider(window.ethereum);
  const signer = provider.getSigner();
  
  // 2. Buat instance contract
  const contract = new ethers.Contract(contractAddress, contractABI, signer);

  // 3. Hitung hash ijazah gabungkan semua data)
  const degreeHash = ethers.utils.keccak256(
    ethers.utils.defaultAbiCoder.encode(
      ["string", "string", "string[]", "uint8[]"],
      [degreeData, sekolah, transcript.subjects, transcript.grades]
    )
  );

  // 4. Estimasi gas fee
  const gasEstimate = await contract.estimateGas.issueDegree(
    degreeHash,
    sekolah,
    Math.floor(Date.now() / 1000), // Unix timestamp
    ipfsUrl,
    transcript.subjects,
    transcript.grades
  );

  return { degreeHash, gasEstimate };
}
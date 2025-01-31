// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract VervalIjazah{
    struct Ijazah {
        string nisn;
        string nama;
        string nomor;
        uint16 tahunLulus;
        uint8 nilaiTranskrip;
        string cidURL;
        bool valid;
    }
    
    mapping(string => Ijazah) private ijazahMapping;
    address private owner;

    modifier onlyOwner() {
        require(msg.sender == owner, "Hanya pemilik kontrak yang bisa menambahkan");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    function add(string memory _nisn, string memory _nama, uint16 _tahunLulus, string memory _nomor, uint8 _nilaiTranskrip, string memory _cidURL) public onlyOwner{
        require(!ijazahMapping[_nomor].valid, "Ijazah sudah ada");

        Ijazah memory newIjazah = Ijazah ({
            nisn: _nisn,
            nama: _nama,
            tahunLulus: _tahunLulus,
            nomor: _nomor,
            nilaiTranskrip: _nilaiTranskrip,
            cidURL: _cidURL,
            valid: true
        });

        ijazahMapping[_nomor] = newIjazah;       
    }
    
    function validasiIjazah(string memory _nomor) public view returns (bool){
        return ijazahMapping[_nomor].valid;
    }

    function get(string memory _nomor) public view returns (string memory,string memory, uint, string memory, uint32, bool){
        Ijazah memory ijazah = ijazahMapping[_nomor];
        require(ijazah.valid, "Ijazah tidak valid");
        return (ijazah.nisn, ijazah.nama, ijazah.tahunLulus, ijazah.nomor, ijazah.nilaiTranskrip,ijazah.valid );
    }
}
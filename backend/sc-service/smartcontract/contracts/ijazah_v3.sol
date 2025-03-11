// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VerifikasiIjazah {
    address public owner;

    struct Subject {
        bytes32 name; 
        uint8 grade;
    }

    struct Degree {
        bytes32 degreeHash;
        bytes32 sekolah; 
        uint256 issueDate;
        bytes32 ipfsUrl; 
        bool exists; // Flag untuk menandai keberadaan ijazah
    }

    modifier isOwner() {
        require(owner == msg.sender, "Hanya pemilik yang bisa melakukan ini.");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    mapping(bytes32 => Degree) public degrees;
    mapping(bytes32 => mapping(uint256 => Subject)) public transcripts; // Menggunakan mapping untuk transkrip

    event DegreeIssued(
        bytes32 indexed degreeHash,
        bytes32 sekolah,
        uint256 issueDate,
        bytes32 ipfsUrl
    );
    event DegreeDeleted(bytes32 indexed degreeHash);

    function issueDegree(
        bytes32 _degreeHash,
        bytes32 _sekolah,
        uint256 _issueDate,
        bytes32 _ipfsUrl,
        bytes32[] calldata _mataPelajaran,
        uint8[] calldata _nilai
    ) public isOwner {
        require(!degrees[_degreeHash].exists, "Ijazah sudah terdaftar.");
        require(
            _mataPelajaran.length == _nilai.length,
            "Jumlah mata pelajaran dan nilai harus sama."
        );

        degrees[_degreeHash] = Degree(
            _degreeHash,
            _sekolah,
            _issueDate,
            _ipfsUrl,
            true
        );

        for (uint256 i = 0; i < _mataPelajaran.length; i++) {
            transcripts[_degreeHash][i] = Subject(_mataPelajaran[i], _nilai[i]);
        }

        emit DegreeIssued(_degreeHash, _sekolah, _issueDate, _ipfsUrl);
    }

    function deleteDegree(bytes32 _degreeHash) public isOwner {
        require(degrees[_degreeHash].exists, "Ijazah tidak ditemukan.");

        degrees[_degreeHash].exists = false; // Tandai sebagai dihapus
        emit DegreeDeleted(_degreeHash);
    }

    function getDegree(
        bytes32 _degreeHash
    )
        public
        view
        returns (bytes32, uint256, bytes32, bytes32[] memory, uint8[] memory)
    {
        Degree storage degree = degrees[_degreeHash];
        require(degree.exists, "Ijazah tidak ditemukan.");

        uint256 length = 0;
        while (transcripts[_degreeHash][length].name != bytes32(0)) {
            length++;
        }

        bytes32[] memory mataPelajaran = new bytes32[](length);
        uint8[] memory nilai = new uint8[](length);

        for (uint256 i = 0; i < length; i++) {
            mataPelajaran[i] = transcripts[_degreeHash][i].name;
            nilai[i] = transcripts[_degreeHash][i].grade;
        }

        return (
            degree.sekolah,
            degree.issueDate,
            degree.ipfsUrl,
            mataPelajaran,
            nilai
        );
    }
}

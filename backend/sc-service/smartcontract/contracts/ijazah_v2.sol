// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract VerifikasiIjazah {
    address public owner;

    struct Subject {
        string name;
        uint8 grade;
    }

    struct Degree {
        bytes32 degreeHash;
        string sekolah;
        uint256 issueDate;
        string ipfsUrl;
        Subject[] transcript; // Transkrip langsung dalam struct Degree
    }

    modifier isOwner() {
        require(owner == msg.sender, "Hanya pemilik yang bisa melakukan ini.");
        _;
    }

    constructor() {
        owner = msg.sender;
    }

    mapping(bytes32 => Degree) public degrees;
    bytes32[] private degreesList;

    event DegreeIssued(
        bytes32 indexed degreeHash,
        string sekolah,
        uint256 issueDate,
        string ipfsUrl
    );
    event DegreeDeleted(bytes32 indexed degreeHash);

    // Mengeluarkan ijazah + transkrip dalam satu fungsi
    function issueDegree(
        bytes32 _degreeHash,
        string memory _sekolah,
        uint256 _issueDate,
        string memory _ipfsUrl,
        string[] memory _mataPelajaran,
        uint8[] memory _nilai
    ) public isOwner {
        require(
            degrees[_degreeHash].degreeHash == bytes32(0),
            "Ijazah sudah terdaftar."
        );
        require(
            _mataPelajaran.length == _nilai.length,
            "Jumlah mata pelajaran dan nilai harus sama."
        );

        Degree storage newDegree = degrees[_degreeHash];
        newDegree.degreeHash = _degreeHash;
        newDegree.sekolah = _sekolah;
        newDegree.issueDate = _issueDate;
        newDegree.ipfsUrl = _ipfsUrl;

        for (uint256 i = 0; i < _mataPelajaran.length; i++) {
            newDegree.transcript.push(Subject(_mataPelajaran[i], _nilai[i]));
        }

        degreesList.push(_degreeHash);

        emit DegreeIssued(_degreeHash, _sekolah, _issueDate, _ipfsUrl);
    }

    // Menghapus ijazah dan transkrip sekaligus
    function deleteDegree(bytes32 _degreeHash) public isOwner {
        require(
            degrees[_degreeHash].degreeHash != bytes32(0),
            "Ijazah tidak ditemukan."
        );

        delete degrees[_degreeHash];
        _removeFromList(degreesList, _degreeHash);

        emit DegreeDeleted(_degreeHash);
    }

    // Mendapatkan detail ijazah + transkrip sekaligus
    function getDegree(
        bytes32 _degreeHash
    )
        public
        view
        returns (
            string memory,
            uint256,
            string memory,
            string[] memory,
            uint8[] memory
        )
    {
        Degree storage degree = degrees[_degreeHash];
        require(degree.degreeHash != bytes32(0), "Ijazah tidak ditemukan.");

        uint256 length = degree.transcript.length;
        string[] memory mataPelajaran = new string[](length);
        uint8[] memory nilai = new uint8[](length);

        for (uint256 i = 0; i < length; i++) {
            mataPelajaran[i] = degree.transcript[i].name;
            nilai[i] = degree.transcript[i].grade;
        }

        return (
            degree.sekolah,
            degree.issueDate,
            degree.ipfsUrl,
            mataPelajaran,
            nilai
        );
    }

    // Mendapatkan daftar semua ijazah
    function listDegrees() public view returns (bytes32[] memory) {
        return degreesList;
    }

    // Fungsi internal untuk menghapus elemen dari array
    function _removeFromList(bytes32[] storage list, bytes32 value) internal {
        for (uint256 i = 0; i < list.length; i++) {
            if (list[i] == value) {
                list[i] = list[list.length - 1]; // Swap dengan elemen terakhir
                list.pop();
                break;
            }
        }
    }
}

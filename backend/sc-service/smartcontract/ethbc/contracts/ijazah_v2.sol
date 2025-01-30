// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DegreeVerification {
    struct Degree {
        bytes32 degreeHash;   // Hash dari ijazah
        string sekolah;       // Sekolah pengeluar ijazah
        uint256 issueDate;    // Tanggal pengeluaran ijazah
    }

    struct Transcript {
        string[] mataPelajaran;  // Nama mata kuliah
        uint8[] nilai;        // Nilai (0-100)
    }

    mapping(bytes32 => Degree) public degrees;
    mapping(bytes32 => Transcript) private transcripts; // Menyimpan transkrip nilai berdasarkan degreeHash

    event DegreeIssued(bytes32 indexed degreeHash, string sekolah, uint256 issueDate);
    event TranscriptAdded(bytes32 indexed degreeHash);

    // Fungsi untuk mengeluarkan ijazah
    function issueDegree(bytes32 _degreeHash, string memory _sekolah, uint256 _issueDate) public {
        require(degrees[_degreeHash].degreeHash == 0, "Ijazah sudah terdaftar.");

        degrees[_degreeHash] = Degree({
            degreeHash: _degreeHash,
            sekolah: _sekolah,
            issueDate: _issueDate
        });

        emit DegreeIssued(_degreeHash, _sekolah, _issueDate);
    }

    // Fungsi untuk menambahkan transkrip nilai
    function addTranscript(bytes32 _degreeHash, string[] memory _mataPelajaran, uint8[] memory _nilai) public {
        require(degrees[_degreeHash].degreeHash != 0, "Ijazah tidak ditemukan.");
        require(_mataPelajaran.length == _nilai.length, "Jumlah mata pelajaran dan nilai harus sama.");

        transcripts[_degreeHash] = Transcript({
            mataPelajaran: _mataPelajaran,
            nilai: _nilai
        });

        emit TranscriptAdded(_degreeHash);
    }

    // Fungsi untuk mendapatkan transkrip nilai
    function getTranscript(bytes32 _degreeHash) public view returns (string[] memory, uint8[] memory) {
        require(transcripts[_degreeHash].mataPelajaran.length > 0, "Transkrip tidak ditemukan.");

        Transcript memory transcript = transcripts[_degreeHash];
        return (transcript.mataPelajaran, transcript.nilai);
    }

    // Fungsi untuk memverifikasi ijazah
    function verifyDegree(bytes32 _degreeHash) public view returns (string memory, uint256) {
        Degree memory degree = degrees[_degreeHash];
        require(degree.degreeHash != 0, "Ijazah tidak ditemukan.");

        return (degree.sekolah, degree.issueDate);
    }
}

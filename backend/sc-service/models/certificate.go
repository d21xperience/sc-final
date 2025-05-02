package models

import (
	"time"

	"github.com/google/uuid"
)

type Certificate struct {
	CertificateID   uint      `gorm:"primaryKey" json:"certificate_id"`
	StudentID       uint      `gorm:"index" json:"student_id"`
	CertificateHash string    `gorm:"size:64;unique" json:"certificate_hash"`
	IPFSCID         string    `gorm:"size:64" json:"ipfs_cid"`
	IssueDate       time.Time `json:"issue_date"`
	IssuerName      string    `gorm:"size:100" json:"issuer_name"`
	NilaiRata       string    `gorm:"size:100" json:"nilai_rata"`
	BlockchainType  string    `gorm:"size:50" json:"blockchain_type"`
	TransactionHash string    `gorm:"size:64" json:"transaction_hash"`
	ContractAddress string    `gorm:"size:64" json:"contract_address"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Verifications []Verification `gorm:"foreignKey:CertificateID" json:"verifications"`
	IPFSMetadata  IPFSMetadata   `gorm:"foreignKey:IPFSCID;references:IPFSCID" json:"ipfs_metadata"`
}

type Verification struct {
	VerificationID   uint      `gorm:"primaryKey" json:"verification_id"`
	CertificateID    uint      `gorm:"index" json:"certificate_id"`
	VerifiedBy       string    `gorm:"size:100" json:"verified_by"`
	VerificationDate time.Time `gorm:"autoCreateTime" json:"verification_date"`
	Result           bool      `json:"result"`
	Remarks          string    `gorm:"type:text" json:"remarks"`
}

type IjazahBc struct {
	ID                          uuid.UUID `gorm:"primaryKey;type:uuid"` // atau varchar sesuai kebutuhan
	PesertaDidikID              uuid.UUID `gorm:"not null"`
	Nama                        string
	NIS                         string
	NISN                        string
	NPSN                        string
	NomorIjazah                 string
	TempatLahir                 string
	TanggalLahir                *time.Time
	NamaOrtuwali                string
	PaketKeahlian               string
	KabupatenKota               string
	Provinsi                    string
	ProgramKeahlian             string
	SekolahPenyelenggaraUjianUS string
	SekolahPenyelenggaraUjianUN string
	AsalSekolah                 string
	TempatIjazah                string
	TanggalIjazah               *time.Time
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

func (IjazahBc) TableName() string {
	return "ijazah_bc"
}

type DegreeData struct {
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	IjazahID       uuid.UUID `gorm:"not null"`
	Ijazah         IjazahBc  `gorm:"foreignKey:IjazahID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DegreeHash     string
	TxHash         string
	IpfsURL        string
	BcType         string
	LinkBcExplorer string
	// UrlBCExplorerEther  string `gorm:"column:url_bc_explorerEther"`
	// UrlBCExplorerQuorum string `gorm:"column:url_bc_explorerQuorum"`
	// UrlBCExplorerFabric string `gorm:"column:url_bc_explorerFabric"`
	SekolahId     uuid.UUID
	TahunAjaranId int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (DegreeData) TableName() string {
	return "degree_data"
}

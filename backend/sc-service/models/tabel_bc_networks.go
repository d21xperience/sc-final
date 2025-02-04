package models

import (
	"time"
)

type BlockchainNetwork struct {
	NetworkID      uint   `gorm:"primaryKey;autoIncrement" json:"network_id"`
	NetworkName    string `gorm:"size:100" json:"network_name"`
	BlockchainType string `gorm:"size:50" json:"blockchain_type"` // public atau private
	Description    string `gorm:"type:text" json:"description"`
	Unit           string `gorm:"size:10" json:"unit"`            // ETH, USD, etc
	BlockExplorer  string `gorm:"size:255" json:"block_explorer"` // URL block explorer
	RPCURL         string `gorm:"size:255" json:"rpc_url"`
	ChainID        int    `json:"chain_id"`
	// kolom Activate digunakan untuk menampilkan jaringan pada saat pemilihan jaringan
	Activate bool `gorm:"default:false" json:"activate"`
	// kolom NetworkAvailable digunakan jika logic bisnis sudah dibuat, saat ini baru tersedia ethereum, quorum dan hyperledger fabric
	NetworkAvailable bool `gorm:"default:false" json:"network_available"`
	// Digunakan untuk menampilkan jaringan secara default
	Applicable bool      `gorm:"default:false" json:"applicable"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// Relationships
	EthereumQuorumSettings    EthereumQuorumSetting    `gorm:"foreignKey:NetworkID"`
	HyperledgerFabricSettings HyperledgerFabricSetting `gorm:"foreignKey:NetworkID"`
	Nodes                     []Node                   `gorm:"foreignKey:NetworkID"`
	Participants              []NetworkParticipant     `gorm:"foreignKey:NetworkID"`
	Account                   []Account                `gorm:"foreignKey:NetworkID" json:"accounts"`
}

type EthereumQuorumSetting struct {
	SettingID uint `gorm:"primaryKey;autoIncrement" json:"setting_id"`
	NetworkID uint `gorm:"index" json:"network_id"`
	// RpcURL    string  `gorm:"size:255" json:"rpc_url"`
	ChainID   int     `json:"chain_id"`
	GasPrice  float64 `gorm:"type:numeric(20,8)" json:"gas_price"`
	BlockTime int     `json:"block_time"`
	Consensus string  `gorm:"size:50" json:"consensus"`
}

type HyperledgerFabricSetting struct {
	SettingID   uint   `gorm:"primaryKey;autoIncrement" json:"setting_id"`
	NetworkID   uint   `gorm:"index" json:"network_id"`
	MSPID       string `gorm:"size:100" json:"msp_id"`
	ChannelName string `gorm:"size:100" json:"channel_name"`
	OrdererURL  string `gorm:"size:255" json:"orderer_url"`
	PeerURL     string `gorm:"size:255" json:"peer_url"`
	CAURL       string `gorm:"size:255" json:"ca_url"`
	TLSEnabled  bool   `json:"tls_enabled"`
}

type Node struct {
	NodeID    uint   `gorm:"primaryKey;autoIncrement" json:"node_id"`
	NetworkID uint   `gorm:"index" json:"network_id"`
	NodeName  string `gorm:"size:100" json:"node_name"`
	NodeURL   string `gorm:"size:255;unique" json:"node_url"`
	NodeRole  string `gorm:"size:50" json:"node_role"`
	Status    string `gorm:"size:20" json:"status"`
}

type NetworkParticipant struct {
	ParticipantID    uint      `gorm:"primaryKey;autoIncrement" json:"participant_id"`
	NetworkID        uint      `gorm:"index" json:"network_id"`
	OrganizationName string    `gorm:"size:100" json:"organization_name"`
	Role             string    `gorm:"size:50" json:"role"`
	PublicKey        string    `gorm:"type:text" json:"public_key"`
	JoinedAt         time.Time `gorm:"autoCreateTime" json:"joined_at"`
}
type Block struct {
	BlockID          uint          `gorm:"primaryKey;autoIncrement" json:"block_id"`
	PreviousHash     string        `gorm:"size:64" json:"previous_hash"`
	CurrentHash      string        `gorm:"size:64;unique" json:"current_hash"`
	Timestamp        time.Time     `gorm:"autoCreateTime" json:"timestamp"`
	Nonce            int           `json:"nonce"`
	TransactionCount int           `json:"transaction_count"`
	Transactions     []Transaction `gorm:"foreignKey:BlockID;constraint:OnDelete:CASCADE" json:"transactions"`
}

// type Transaction struct {
// 	TransactionID   string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"transaction_id"`
// 	BlockID         uint      `gorm:"index" json:"block_id"`
// 	SenderAddress   string    `gorm:"size:64" json:"sender_address"`
// 	ReceiverAddress string    `gorm:"size:64" json:"receiver_address"`
// 	Amount          float64   `gorm:"type:numeric(20,8)" json:"amount"`
// 	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
// 	TransactionHash string    `gorm:"size:64;unique" json:"transaction_hash"`
// 	Metadata        string    `gorm:"type:jsonb" json:"metadata,omitempty"`
// }

type AuditLog struct {
	LogID     uint      `gorm:"primaryKey;autoIncrement" json:"log_id"`
	Action    string    `gorm:"size:50" json:"action"`
	Details   string    `gorm:"type:jsonb" json:"details"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UserID    uint      `json:"user_id"`
}

// type User struct {
// 	UserID      uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
// 	Username    string    `gorm:"size:50;unique" json:"username"`
// 	Email       string    `gorm:"size:100;unique" json:"email"`
// 	Role        string    `gorm:"size:100" json:"role"`
// 	TenantID    string    `gorm:"size:100" json:"tenant_id"`
// 	Phone       string    `gorm:"size:100" json:"phone"`
// 	AlamatJalan string    `gorm:"size:100" json:"alamat_jalan"`
// 	Desa        string    `gorm:"size:100" json:"desa"`
// 	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
// }
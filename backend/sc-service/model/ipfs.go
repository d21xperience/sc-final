package model

import "time"

type IPFSMetadata struct {
	MetadataID uint      `gorm:"primaryKey" json:"metadata_id"`
	IPFSCID    string    `gorm:"size:64;unique" json:"ipfs_cid"`
	Size       int       `json:"size"`
	FileName   string    `gorm:"size:255" json:"file_name"`
	FileType   string    `gorm:"size:50" json:"file_type"`
	UploadedAt time.Time `gorm:"autoCreateTime" json:"uploaded_at"`
}

type IPFSNode struct {
	NodeID    uint      `gorm:"primaryKey" json:"node_id"`
	NodeName  string    `gorm:"size:100" json:"node_name"`
	Host      string    `gorm:"size:255" json:"host"`
	Port      int       `json:"port"`
	Protocol  string    `gorm:"size:10" json:"protocol"`
	Status    string    `gorm:"size:20" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Files []IPFSFile `gorm:"foreignKey:NodeID" json:"files"`
}

type IPFSFile struct {
	FileID     uint      `gorm:"primaryKey" json:"file_id"`
	FileName   string    `gorm:"size:255" json:"file_name"`
	FileType   string    `gorm:"size:50" json:"file_type"`
	FileSize   int64     `json:"file_size"`
	IPFSCID    string    `gorm:"size:64;unique" json:"ipfs_cid"`
	NodeID     uint      `gorm:"index" json:"node_id"`
	UploadedAt time.Time `gorm:"autoCreateTime" json:"uploaded_at"`
	PinStatus  string    `gorm:"size:20" json:"pin_status"`
	Metadata   string    `gorm:"type:jsonb" json:"metadata"`

	Transactions []IPFSTransaction `gorm:"foreignKey:FileID" json:"transactions"`
}

type IPFSTransaction struct {
	TransactionID   uint      `gorm:"primaryKey" json:"transaction_id"`
	FileID          uint      `gorm:"index" json:"file_id"`
	TransactionType string    `gorm:"size:50" json:"transaction_type"` // upload/download
	Status          string    `gorm:"size:20" json:"status"`
	Timestamp       time.Time `gorm:"autoCreateTime" json:"timestamp"`
	ErrorMessage    string    `gorm:"type:text" json:"error_message"`
}

type IPFSPinningService struct {
	ServiceID   uint      `gorm:"primaryKey" json:"service_id"`
	ServiceName string    `gorm:"size:100" json:"service_name"`
	APIKey      string    `gorm:"type:text" json:"api_key"`
	EndpointURL string    `gorm:"size:255" json:"endpoint_url"`
	Status      string    `gorm:"size:20" json:"status"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

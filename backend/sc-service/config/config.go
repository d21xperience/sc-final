package config

import (
	"fmt"
	"log"
	"sc-service/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbName   = "postgres"
)

var DB *gorm.DB

func InitDatabase() {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	var err error
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Terjadi error dg error sbb: ", err)
	}
	DB.AutoMigrate(
		&model.BlockchainNetwork{},
		&model.EthereumQuorumSetting{},
		&model.HyperledgerFabricSetting{},
		&model.Node{},
		&model.NetworkParticipant{},
		&model.Block{},
		&model.Transaction{},
		&model.AuditLog{},
		&model.Account{},
	)
	DB.AutoMigrate(&model.IPFSMetadata{}, &model.IPFSNode{}, &model.IPFSFile{}, &model.IPFSTransaction{}, &model.IPFSPinningService{})

	DB.AutoMigrate(&model.Student{}, &model.Certificate{})
	log.Println("Database connected successfully!")
}

//

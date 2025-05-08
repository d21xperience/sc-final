package worker

import (
	"auth_service/models"
	"auth_service/services"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func StartSekolahInitWorker(rdb *redis.Client) {
	ctx := context.Background()

	for {
		val, err := rdb.BLPop(ctx, 0*time.Second, "sekolah_init_queue").Result()
		if err != nil {
			log.Printf("Gagal ambil task dari queue: %v", err)
			continue
		}

		var model *models.SekolahTenant
		if err := json.Unmarshal([]byte(val[1]), &model); err != nil {
			log.Printf("Gagal parsing task: %v", err)
			continue
		}

		if err := initSekolahService(model); err != nil {
			log.Printf("Gagal eksekusi initSekolahService: %v", err)
		} else {
			log.Println("initSekolahService berhasil dijalankan.")
		}
	}
}

func initSekolahService(sekolahModel *models.SekolahTenant) error {
	sekolahClient, err := services.NewSekolahServiceClient()
	if err != nil {
		return err
	}

	if err := sekolahClient.RegistrasiSekolah(sekolahModel); err != nil {
		return err
	}

	if err := sekolahClient.CreateSekolah(sekolahModel); err != nil {
		return err
	}

	return nil
}

package queue

import (
	"auth_service/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisEnqueue struct {
	rdb *redis.Client
}

func NewRedisEnqueue(rdb *redis.Client) *RedisEnqueue {
	return &RedisEnqueue{
		rdb: rdb,
	}
}

var ctx = context.Background()

func (r *RedisEnqueue) EnqueueInitSekolahTask(sekolahModel models.SekolahTenant) error {
	payload, err := json.Marshal(sekolahModel)
	if err != nil {
		return err
	}

	err = r.rdb.RPush(ctx, "sekolah_init_queue", payload).Err()
	if err != nil {
		return err
	}

	fmt.Println("Task ditambahkan ke queue")
	return nil
}
func (r *RedisEnqueue) EnqueueInitSCTask(sekolahModel models.SekolahTenant) error {
	payload, err := json.Marshal(sekolahModel)
	if err != nil {
		return err
	}

	err = r.rdb.RPush(ctx, "sc_init_queue", payload).Err()
	if err != nil {
		return err
	}

	fmt.Println("Task ditambahkan ke queue")
	return nil
}

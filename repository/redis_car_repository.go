package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aabdullahgungor/go-restapi-redis/model"
	"github.com/go-redis/redis"
)

var (
	ErrCarNotFound = errors.New("FromRepository - car not found")
	Ctx            = context.TODO()
)

type RedisCarRepository struct {
	connectionPool *redis.Client
}

func NewRedisCarRepository() *RedisCarRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Check to connect to Redis server
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Could not connect to redis server:", err)
	}
	fmt.Println("Connected to Redis server:", pong)

	return &RedisCarRepository{
		connectionPool: client,
	}
}

func (r *RedisCarRepository) GetAllCars() ([]model.Car, error) {

}

func (r *RedisCarRepository) GetCarById(id string) (model.Car, error) {

}

func (r *RedisCarRepository) CreateCar(car *model.Car) error {

}

func (r *RedisCarRepository) EditCar(car *model.Car) error {

}

func (r *RedisCarRepository) DeleteCar(id string) error {

}

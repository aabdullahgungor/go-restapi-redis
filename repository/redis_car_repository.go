package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/aabdullahgungor/go-restapi-redis/model"
	"github.com/go-redis/redis"
)

var (
	ErrCarNotFound  = errors.New("fromRepository - car not found")
	ErrCarNotCreate = errors.New("fromRepository - car not create")
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
	IDKEY := strconv.Itoa(car.Id)
	jsonCar, errJson := json.Marshal(&car)
	if errJson != nil {
		fmt.Println("fromRepository", errJson)
		return ErrCarNotCreate
	}
	errRedis := r.connectionPool.Set(IDKEY, jsonCar, 0).Err()
	if errRedis != nil {
		fmt.Println("fromRepository", errRedis)
		return ErrCarNotCreate
	}
	log.Printf("\ndisplay the ids of the newly inserted car: %v", car.Id)
	return nil
}

func (r *RedisCarRepository) EditCar(car *model.Car) error {

}

func (r *RedisCarRepository) DeleteCar(id string) error {

}

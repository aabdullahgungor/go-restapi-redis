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
	ErrCarNotFound   = errors.New("fromRepository - car not found")
	ErrCarNotCreate  = errors.New("fromRepository - car not create")
	ErrCarsNotGet    = errors.New("fromRepository - cars couldn't get")
	ErrCarNotDeleted = errors.New("fromRepository - car not deleted")
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
	var cars []model.Car
	var car model.Car

	// Get all the keys matching
	keys, errRedis := r.connectionPool.Keys("*").Result()
	if errRedis != nil {
		fmt.Println("fromRepository", errRedis)
		return []model.Car{}, ErrCarsNotGet
	}

	for key := range keys {
		jsonCar, err := r.connectionPool.Get(keys[key]).Result()
		if err != nil {
			fmt.Println("Data read error:", err)
		}
		// Marshal JSON data to Car
		errJson := json.Unmarshal([]byte(jsonCar), &car)
		if errJson != nil {
			fmt.Println("fromRepository", errJson)
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (r *RedisCarRepository) GetCarById(id string) (model.Car, error) {
	IDKEY := id
	var car model.Car

	jsonCar, errRedis := r.connectionPool.Get(IDKEY).Result()
	if errRedis != nil {
		fmt.Println("Data read error:", errRedis)
		return model.Car{}, ErrCarNotFound
	}
	// Marshal JSON data to Car
	errJson := json.Unmarshal([]byte(jsonCar), &car)
	if errJson != nil {
		fmt.Println("fromRepository", errJson)
	}

	return car, nil
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
	log.Printf("\nUpdated the car with id number: %v", car.Id)
	return nil
}

func (r *RedisCarRepository) DeleteCar(id string) error {
	IDKEY := id
	_, errRedis := r.connectionPool.Del(IDKEY).Result()
	if errRedis != nil {
		fmt.Println("Data deleted error:", errRedis)
		return ErrCarNotDeleted
	}
	log.Printf("\nDeleted the car with id number: " + IDKEY)
	return nil
}

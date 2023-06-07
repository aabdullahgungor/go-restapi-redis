package service

import (
	"errors"

	"github.com/aabdullahgungor/go-restapi-redis/model"
	"github.com/aabdullahgungor/go-restapi-redis/repository"
)

var (
	ErrIDIsNotValid    = errors.New("id is not valid")
	ErrBrandIsNotEmpty = errors.New("brand is not empty")
	ErrCarNotFound     = errors.New("car cannot be found")
)

type DefaultCarService struct {
	carRepo repository.ICarRepository
}

func NewDefaultCarService(cRepo repository.ICarRepository) *DefaultCarService {
	return &DefaultCarService{
		carRepo: cRepo,
	}
}

func (d *DefaultCarService) GetAll() ([]model.Car, error) {
	return d.carRepo.GetAllCars()
}

func (d *DefaultCarService) GetById(id string) (model.Car, error) {

	car, err := d.carRepo.GetCarById(id)

	if err != nil {
		return model.Car{}, ErrCarNotFound
	}

	return car, nil
}

func (d *DefaultCarService) Create(car *model.Car) error {

	if car.Brand == "" {
		return ErrBrandIsNotEmpty
	}

	return d.carRepo.CreateCar(car)
}

func (d *DefaultCarService) Edit(car *model.Car) error {

	if car.Id <= 0 {
		return ErrIDIsNotValid
	}
	if car.Brand == "" {
		return ErrBrandIsNotEmpty
	}

	err := d.carRepo.EditCar(car)

	if err != nil {
		return ErrCarNotFound
	}

	return nil
}

func (d *DefaultCarService) Delete(id string) error {

	err := d.carRepo.DeleteCar(id)

	if err != nil {
		return ErrCarNotFound
	}

	return nil
}

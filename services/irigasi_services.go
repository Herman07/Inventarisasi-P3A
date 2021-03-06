package services

import (
	"fmt"
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
)

type IrigasiDataService struct {
	IrigasiDataRepository repository.IrigasiDataRepository
}

func NewIrigasiDataService(repository repository.IrigasiDataRepository) *IrigasiDataService {
	return &IrigasiDataService{
		IrigasiDataRepository: repository,
	}
}

func (s *IrigasiDataService) Create(request request.TeknikIrigasiReq) (*models.TeknikIrigasi, error) {
	entity := models.TeknikIrigasi{
		Operasi:      request.Operasi,
		Partisipatif: request.Partisipatif,
		ID:           request.ID,
	}
	data, err := s.IrigasiDataRepository.Create(entity)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *IrigasiDataService) UpdateById(id string, dto request.TeknikIrigasiReq) (*models.TeknikIrigasi, error) {
	entity := models.TeknikIrigasi{
		ID:           id,
		Operasi:      dto.Operasi,
		Partisipatif: dto.Partisipatif,
	}
	fmt.Println("Service Data", entity)
	data, err := s.IrigasiDataRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *IrigasiDataService) FindById(id string) (*models.TeknikIrigasi, error) {
	data, err := s.IrigasiDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *IrigasiDataService) Delete(id string) error {
	entity := models.TeknikIrigasi{
		ID: id,
	}
	err := s.IrigasiDataRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}

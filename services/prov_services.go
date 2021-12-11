package services

import (
	"fmt"
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/repository"
	"Inventarisasi-P3A/request"
	"strings"
)

type ProvDataService struct {
	ProvDataRepository repository.ProvDataRepository
}

func NewProvDataService(repository repository.ProvDataRepository) *ProvDataService {
	return &ProvDataService{
		ProvDataRepository: repository,
	}
}

func (s *ProvDataService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.MasterDataProvinsi, err error) {
	recordTotal, err = s.ProvDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.ProvDataRepository.CountWhere("or", map[string]interface{}{

			"nama_prov LIKE ?": "%" + searchValue + "%" ,

		})

		data, err = s.ProvDataRepository.FindAllWhere("or", orderType, "nama_prov", limit, offset, map[string]interface{}{
			"nama_prov LIKE ?": "%" + searchValue + "%",
			"id_prov LIKE ?": "%" + searchValue + "%",

		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.ProvDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.ProvDataRepository.FindAllWhere("or", orderType, "id_prov", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}

func (s *ProvDataService) Create(request request.ProvinsiReq) (*models.MasterDataProvinsi, error) {
	entity := models.MasterDataProvinsi{
		Provinsi:  request.Nama,
		ID: request.ID,
	}
	fmt.Println("Isinya ", entity)
	data, err := s.ProvDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *ProvDataService) UpdateById(id string, dto request.ProvinsiReq) (*models.MasterDataProvinsi, error) {
	entity := models.MasterDataProvinsi{
		ID:        id,
		Provinsi:  dto.Nama,
	}

	data, err := s.ProvDataRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *ProvDataService) FindById(id string) (*models.MasterDataProvinsi, error) {
	data, err := s.ProvDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *ProvDataService) Find(id string) (*models.MasterDataProvinsi, error) {
	data, err := s.ProvDataRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *ProvDataService) Delete(id string) error {
	entity := models.MasterDataProvinsi{
		ID: id,
	}
	err := s.ProvDataRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
package services

import (
	"fmt"
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
	"strings"
)

type MasterDataService struct {
	MasterDataRepository repository.MasterDataRepository
}

func NewMasterDataService(repository repository.MasterDataRepository) *MasterDataService {
	return &MasterDataService{
		MasterDataRepository: repository,
	}
}

func (s *MasterDataService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.MasterDataProvinsi, err error) {
	recordTotal, err = s.MasterDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.MasterDataRepository.CountWhere("or", map[string]interface{}{

			"nama_prov LIKE ?": "%" + searchValue + "%" ,

		})

		data, err = s.MasterDataRepository.FindAllWhere("or", orderType, "nama_prov", limit, offset, map[string]interface{}{
			"nama_prov LIKE ?": "%" + searchValue + "%",
			"id_prov LIKE ?": "%" + searchValue + "%",

		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.MasterDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.MasterDataRepository.FindAllWhere("or", orderType, "id_prov", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}

func (s *MasterDataService) Create(request request.ProvinsiReq) (*models.MasterDataProvinsi, error) {
	entity := models.MasterDataProvinsi{
		Provinsi:  request.Nama,
		ID: request.ID,
	}
	fmt.Println("Isinya ", entity)
	data, err := s.MasterDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
package services

import (
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/repository"
	"Inventarisasi-P3A/request"
	"strings"
)

type KabDataService struct {
	KabDataRepository repository.KabDataRepository
}

func NewKabDataService(repository repository.KabDataRepository) *KabDataService {
	return &KabDataService{
		KabDataRepository: repository,
	}
}

func (s *KabDataService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.MasterDataKab, err error) {
	recordTotal, err = s.KabDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.KabDataRepository.CountWhere("or", map[string]interface{}{

			"nama_kab LIKE ?": "%" + searchValue + "%" ,

		})

		data, err = s.KabDataRepository.FindAllWhere("or", orderType, "nama_kab", limit, offset, map[string]interface{}{
			"nama_kab LIKE ?": "%" + searchValue + "%",
			"id_kab LIKE ?": "%" + searchValue + "%",

		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.KabDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.KabDataRepository.FindAllWhere("or", orderType, "id_kab", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}

func (s *KabDataService) Create(request request.KabReq) (*models.MasterDataKab, error) {
	entity := models.MasterDataKab{
		Kabupaten:  request.Nama,
		IDProv: request.IDProv,
		ID: request.ID,
	}
	data, err := s.KabDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *KabDataService) UpdateById(id string, dto request.KabReq) (*models.MasterDataKab, error) {
	entity := models.MasterDataKab{
		ID:        id,
		Kabupaten:  dto.Nama,
	}

	data, err := s.KabDataRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *KabDataService) FindById(id string) (*models.MasterDataKab, error) {
	data, err := s.KabDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *KabDataService) Delete(id string) error {
	entity := models.MasterDataKab{
		ID: id,
	}
	err := s.KabDataRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (s *KabDataService) Find(id string) (*[]models.MasterDataKab, error) {
	data, err := s.KabDataRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return data, err
}
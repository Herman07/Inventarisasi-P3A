package services

import (
	"Inventarisasi-P3A/models"
	"Inventarisasi-P3A/repository"
	"Inventarisasi-P3A/request"
	"strings"
)

type KecDataService struct {
	KecDataRepository repository.KecDataRepository
}

func NewKecDataService(repository repository.KecDataRepository) *KecDataService {
	return &KecDataService{
		KecDataRepository: repository,
	}
}

func (s *KecDataService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.MasterDataKec, err error) {
	recordTotal, err = s.KecDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.KecDataRepository.CountWhere("or", map[string]interface{}{

			"nama_kecamatan LIKE ?": "%" + searchValue + "%" ,

		})

		data, err = s.KecDataRepository.FindAllWhere("or", orderType, "nama_kecamatan", limit, offset, map[string]interface{}{
			"nama_kecamatan LIKE ?": "%" + searchValue + "%",
			"id_kec LIKE ?": "%" + searchValue + "%",

		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.KecDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.KecDataRepository.FindAllWhere("or", orderType, "id_kec", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}

func (s *KecDataService) Create(request request.KecReq) (*models.MasterDataKec, error) {
	entity := models.MasterDataKec{
		Kecamatan:  request.Nama,
		IDKab: request.IDKab,
		ID: request.ID,
	}
	data, err := s.KecDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *KecDataService) UpdateById(id string, dto request.KecReq) (*models.MasterDataKec, error) {
	entity := models.MasterDataKec{
		ID:        id,
		Kecamatan:  dto.Nama,
	}

	data, err := s.KecDataRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *KecDataService) FindById(id string) (*models.MasterDataKec, error) {
	data, err := s.KecDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *KecDataService) Delete(id string) error {
	entity := models.MasterDataKec{
		ID: id,
	}
	err := s.KecDataRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *KecDataService) Find(id string) (*[]models.MasterDataKec, error) {
	data, err := s.KecDataRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return data, err
}
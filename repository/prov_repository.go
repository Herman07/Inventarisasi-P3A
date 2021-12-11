package repository

import (
	"gorm.io/gorm"
	"Inventarisasi-P3A/models"
)

type ProvDataRepository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.MasterDataProvinsi, error)
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	Create(entity models.MasterDataProvinsi) (*models.MasterDataProvinsi, error)
	UpdateById(entity models.MasterDataProvinsi)(*models.MasterDataProvinsi, error)
	Delete(models.MasterDataProvinsi) error
	FindById(id string) (*models.MasterDataProvinsi, error)
	FindByID(id string) (*models.MasterDataProvinsi, error)
	DbInstance() *gorm.DB
}


type provdataRepository struct {
	*gorm.DB
}

func NewProvDataRepository(db *gorm.DB) ProvDataRepository {
	return &provdataRepository{
		DB: db,
	}
}
func (r *provdataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.MasterDataProvinsi, error) {
	var entity []models.MasterDataProvinsi
	res := r.DB.Table("provinsi").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

	for k, v := range keyVal {
		switch operation {
		case "and":
			res = res.Where(k, v)
		case "or":
			res = res.Or(k, v)
		}
	}
	err := res.Find(&entity).Error
	return entity, err

}
func (r provdataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("provinsi").Count(&count).Error
	return count, err
}

func (r provdataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.MasterDataProvinsi{})
	for k, v := range keyVal {
		switch operation {
		case "and":
			q = q.Where(k, v)
		case "or":
			q = q.Or(k, v)
		}
	}

	err := q.Count(&count).Error
	return count, err
}

func (r provdataRepository) Create(entity models.MasterDataProvinsi) (*models.MasterDataProvinsi, error) {
	err := r.DB.Table("provinsi").Create(&entity).Error
	return &entity, err
}

func (r provdataRepository) UpdateById(entity models.MasterDataProvinsi)(*models.MasterDataProvinsi, error){
	err := r.DB.Model(&models.MasterDataProvinsi{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r provdataRepository) FindById(id string) (*models.MasterDataProvinsi, error) {
	var entity models.MasterDataProvinsi
	err := r.DB.Table("provinsi").Where("id_prov = ?", id).First(&entity).Error
	return &entity, err
}
func (r provdataRepository) FindByID(id string) (*models.MasterDataProvinsi, error) {
	var entity models.MasterDataProvinsi
	err := r.DB.Table("provinsi").Where("id_prov = ?", id).First(&entity).Error
	return &entity, err
}

func (r provdataRepository) Delete(entity models.MasterDataProvinsi) error {
	return r.DB.Table("provinsi").Delete(&entity).Error
}

func (r *provdataRepository) DbInstance() *gorm.DB {
	return r.DB
}


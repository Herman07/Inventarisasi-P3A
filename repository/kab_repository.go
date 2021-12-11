package repository

import (
	"gorm.io/gorm"
	"Inventarisasi-P3A/models"
)

type KabDataRepository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.MasterDataKab, error)
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	Create(entity models.MasterDataKab) (*models.MasterDataKab, error)
	UpdateById(entity models.MasterDataKab)(*models.MasterDataKab, error)
	FindById(id string) (*models.MasterDataKab, error)
	Delete(models.MasterDataKab) error
	FindByID(id string) (*[]models.MasterDataKab, error)
	DbInstance() *gorm.DB
}


type kabdataRepository struct {
	*gorm.DB
}

func NewKabDataRepository(db *gorm.DB) KabDataRepository {
	return &kabdataRepository{
		DB: db,
	}
}
func (r *kabdataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.MasterDataKab, error) {
	var entity []models.MasterDataKab
	res := r.DB.Table("kabupaten").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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
func (r kabdataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("kabupaten").Count(&count).Error
	return count, err
}

func (r kabdataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.MasterDataKab{})
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

func (r kabdataRepository) Create(entity models.MasterDataKab) (*models.MasterDataKab, error) {
	err := r.DB.Table("kabupaten").Create(&entity).Error
	return &entity, err
}

func (r kabdataRepository) UpdateById(entity models.MasterDataKab)(*models.MasterDataKab, error){
	err := r.DB.Model(&models.MasterDataKab{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}
func (r kabdataRepository) FindById(id string) (*models.MasterDataKab, error) {
	var entity models.MasterDataKab
	err := r.DB.Table("kabupaten").Where("id_kab = ?", id).First(&entity).Error
	return &entity, err
}

func (r kabdataRepository) Delete(entity models.MasterDataKab) error {
	return r.DB.Table("kabupaten").Delete(&entity).Error
}
func (r kabdataRepository) FindByID(id string) (*[]models.MasterDataKab, error) {
	var entity []models.MasterDataKab
	err := r.DB.Table("kabupaten").Where("id_prov_fk = ?", id).Find(&entity).Error
	return &entity, err
}
func (r *kabdataRepository) DbInstance() *gorm.DB {
	return r.DB
}


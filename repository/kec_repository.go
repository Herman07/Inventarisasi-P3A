package repository

import (
	"gorm.io/gorm"
	"Inventarisasi-P3A/models"
)

type KecDataRepository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.MasterDataKec, error)
	Count() (int64, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	Create(entity models.MasterDataKec) (*models.MasterDataKec, error)
	UpdateById(entity models.MasterDataKec)(*models.MasterDataKec, error)
	FindById(id string) (*models.MasterDataKec, error)
	Delete(models.MasterDataKec) error
	FindByID(id string) (*[]models.MasterDataKec, error)
	DbInstance() *gorm.DB
}


type kecdataRepository struct {
	*gorm.DB
}

func NewKecDataRepository(db *gorm.DB) KecDataRepository {
	return &kecdataRepository{
		DB: db,
	}
}
func (r *kecdataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.MasterDataKec, error) {
	var entity []models.MasterDataKec
	res := r.DB.Table("kecamatan").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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
func (r kecdataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("kecamatan").Count(&count).Error
	return count, err
}

func (r kecdataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.MasterDataKec{})
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

func (r kecdataRepository) Create(entity models.MasterDataKec) (*models.MasterDataKec, error) {
	err := r.DB.Table("kecamatan").Create(&entity).Error
	return &entity, err
}

func (r kecdataRepository) UpdateById(entity models.MasterDataKec)(*models.MasterDataKec, error){
	err := r.DB.Model(&models.MasterDataKec{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r kecdataRepository) FindById(id string) (*models.MasterDataKec, error) {
	var entity models.MasterDataKec
	err := r.DB.Table("kecamatan").Where("id_kec = ?", id).First(&entity).Error
	return &entity, err
}

func (r kecdataRepository) Delete(entity models.MasterDataKec) error {
	return r.DB.Table("kecamatan").Delete(&entity).Error
}
func (r kecdataRepository) FindByID(id string) (*[]models.MasterDataKec, error) {
	var entity []models.MasterDataKec
	err := r.DB.Table("kecamatan").Where("id_kab_fk = ?", id).Find(&entity).Error
	return &entity, err
}

func (r *kecdataRepository) DbInstance() *gorm.DB {
	return r.DB
}


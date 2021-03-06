package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type P3Repository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Inventaris, error)
	Count() (int64, error)
	GetData(dataReq models.Inventaris)(*models.Inventaris, error)
	FindById(id string) (*models.Inventaris, error)
	UpdateById(models.Inventaris)(*models.Inventaris, error)
	Delete(models.Inventaris) error
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	Create(entity models.Inventaris) (*models.Inventaris, error)
	CreateStatusLegal(entity models.StatusLegal) (*models.StatusLegal, error)
	CreatePengurus(entity models.Pengurus) (*models.Pengurus, error)
	CreateIrigasi(entity models.TeknikIrigasi) (*models.TeknikIrigasi, error)
	CreatePertanian(entity models.TeknikPertanian) (*models.TeknikPertanian, error)
	DbInstance() *gorm.DB
}


type p3Repository struct {
	*gorm.DB
}

func NewP3Repository(db *gorm.DB) P3Repository {
	return &p3Repository{
		DB: db,
	}
}

func (r *p3Repository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.Inventaris, error) {
	var entity []models.Inventaris
	res := r.DB.Table("data_p3a").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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

func (r p3Repository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("data_p3a").Count(&count).Error
	return count, err
}
func (r *p3Repository)GetData(dataReq models.Inventaris)(*models.Inventaris,error) {
	data := models.Inventaris{}
	err := r.DB.Table("inventaris").Find(&data).Error

	return &data, err

}
func (r p3Repository) FindById(id string) (*models.Inventaris, error) {
	var entity models.Inventaris
	err := r.DB.Table("inventaris").Where("id = ?", id).First(&entity).Error
	return &entity, err
}

func (r p3Repository) UpdateById(entity models.Inventaris)(*models.Inventaris, error){
	err := r.DB.Model(&models.Inventaris{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r p3Repository) Delete(entity models.Inventaris) error {
	return r.DB.Table("inventaris").Delete(&entity).Error
}

func (r p3Repository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.Inventaris{})
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

func (r p3Repository) Create(entity models.Inventaris) (*models.Inventaris, error) {
	err := r.DB.Table("inventaris").Create(&entity).Error
	return &entity, err
}

func (r p3Repository) CreatePengurus(entity models.Pengurus) (*models.Pengurus, error) {
	err := r.DB.Table("kepengurusan").Create(&entity).Error
	return &entity, err
}
func (r p3Repository) CreateStatusLegal(entity models.StatusLegal) (*models.StatusLegal, error) {
	err := r.DB.Table("status_legal").Create(&entity).Error
	return &entity, err
}
func (r p3Repository) CreateIrigasi(entity models.TeknikIrigasi) (*models.TeknikIrigasi, error) {
	err := r.DB.Table("teknik_irigasi").Create(&entity).Error
	return &entity, err
}
func (r p3Repository) CreatePertanian(entity models.TeknikPertanian) (*models.TeknikPertanian, error) {
	err := r.DB.Table("teknik_pertanian").Create(&entity).Error
	return &entity, err
}


func (r *p3Repository) DbInstance() *gorm.DB {
	return r.DB
}


package repository

import (
	"Inventarisasi-P3A/models"
	"gorm.io/gorm"
)

type UsersDataRepository interface {
	Create(entity models.User) (*models.User, error)
	CreateAkun(entity models.Akun) (*models.Akun, error)
	FindById(id string) (*models.UserUpdate, error)
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.User, error)
	Count() (int64, error)
	DeleteUser(entity models.User) error
	UpdateUser(entity models.User)(*models.User, error)
	UpdateAkun(entity models.Akun)(*models.Akun, error)
	UpdateFoto(entity models.Akun)(*models.Akun, error)
	DbInstance() *gorm.DB
}


type usersdataRepository struct {
	*gorm.DB
}

func NewUsersDataRepository(db *gorm.DB) UsersDataRepository {
	return &usersdataRepository{
		DB: db,
	}
}
func (r usersdataRepository) Create(entity models.User) (*models.User, error) {
	err := r.DB.Table("user").Create(&entity).Error
	return &entity, err
}
func (r usersdataRepository) CreateAkun(entity models.Akun) (*models.Akun, error) {
	err := r.DB.Table("user_group").Create(&entity).Error
	return &entity, err
}
func (r usersdataRepository) FindById(id string) (*models.UserUpdate, error) {
	var entity models.UserUpdate
	err := r.DB.Table("user a").
		Select("a.*,b.*").
		Joins("LEFT JOIN user_group b on b.id_usergroup = a.id_user").Where("id_user = ?", id).Find(&entity).Error
	return &entity, err
}
func (r *usersdataRepository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.User, error) {
	var entity []models.User
	res := r.DB.Table("user").Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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
func (r usersdataRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Table("user").Count(&count).Error
	return count, err
}

func (r usersdataRepository) CountWhere(operation string, keyVal map[string]interface{}) (int64, error) {
	var count int64
	q := r.DB.Model(&models.User{})
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

func (r usersdataRepository) DeleteUser(entity models.User) error {
	return r.DB.Table("user").Delete(&entity).Error
}

func (r usersdataRepository) UpdateUser(entity models.User)(*models.User, error){
	err := r.DB.Table("user").Where("id_user = ?",entity.ID).Updates(&entity).Error
	return &entity, err
}

func (r usersdataRepository) UpdateAkun(entity models.Akun)(*models.Akun, error){
	err := r.DB.Table("user_group").Where("id_usergroup = ?",entity.ID).Updates(&entity).Error
	return &entity, err
}

func (r usersdataRepository) UpdateFoto(entity models.Akun)(*models.Akun, error){
	err := r.DB.Table("user_group").Where("id_usergroup = ?",entity.ID).Updates(&entity).Error
	return &entity, err
}
func (r *usersdataRepository) DbInstance() *gorm.DB {
	return r.DB
}


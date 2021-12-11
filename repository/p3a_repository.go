package repository

import (
	"gorm.io/gorm"
	"Inventarisasi-P3A/models"
)

type P3Repository interface {
	FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.P3AModels, error)
	Count() (int64, error)
	GetData(dataReq models.Inventaris)(*models.Inventaris, error)
	FindById(id string) (*models.P3AModels, error)
	UpdateById(models.Inventaris)(*models.Inventaris, error)
	Delete(models.Inventaris) error
	CountWhere(operation string, keyVal map[string]interface{}) (int64, error)
	Create(entity models.Inventaris) (*models.Inventaris, error)
	CreateStatusLegal(entity models.StatusLegal) (*models.StatusLegal, error)
	CreatePengurus(entity models.Pengurus) (*models.Pengurus, error)
	CreateIrigasi(entity models.TeknikIrigasi) (*models.TeknikIrigasi, error)
	CreatePertanian(entity models.TeknikPertanian) (*models.TeknikPertanian, error)
	DbInstance() *gorm.DB
	UpdateStatusLegal(entity models.StatusLegal)(*models.StatusLegal, error)
	UpdatePengurus(entity models.Pengurus)(*models.Pengurus, error)
	UpdateIrigasi(models.TeknikIrigasi)(*models.TeknikIrigasi, error)
	UpdatePertanian(pertanian models.TeknikPertanian)(*models.TeknikPertanian, error)
	ExportExcel()(*[]models.P3AModels, error)
	DeleteStatusLegal(id string) error
	DeleteIrigasi(id string) error
	DeletePengurusan(id string) error
	DeletePertanian(id string) error
}


type p3Repository struct {
	*gorm.DB
}

func NewP3Repository(db *gorm.DB) P3Repository {
	return &p3Repository{
		DB: db,
	}
}

func (r *p3Repository) FindAllWhere(operation string, orderType string, orderBy string, limit int, offset int, keyVal map[string]interface{}) ([]models.P3AModels, error) {
	var entity []models.P3AModels
	res := r.DB.Table("data_p3a a").
	Select("a.*,b.*,c.*,d.*,e.*,f.*,g.*,h.*").
		Joins("LEFT JOIN status_legal b on b.id_status_legal = a.id_status_legal_fk").
		Joins("LEFT JOIN kepengurusan c on c.id_kepengurusan = a.id_kepengurusan_fk").
		Joins("LEFT JOIN teknik_irigasi d on d.id_t_irigasi = a.id_teknik_irigasi_fk").
		Joins("LEFT JOIN teknik_pertanian e on e.id_t_pertanian = a.id_teknik_pertanian_fk").
		Joins("LEFT JOIN provinsi f on f.id_prov = a.id_prov_fk").
		Joins("LEFT JOIN kabupaten g on g.id_kab = a.id_kab_fk").
		Joins("LEFT JOIN kecamatan h on h.id_kec = a.id_kec_fk").
		Order(orderBy + " " + orderType).Limit(limit).Offset(offset)

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
	err := r.DB.Table("data_p3a").Find(&data).Error

	return &data, err

}
func (r p3Repository) FindById(id string) (*models.P3AModels, error) {
	var entity models.P3AModels
	err := r.DB.Table("data_p3a a").
		Select("a.*,b.*,c.*,d.*,e.*,f.*,g.*,h.*").
		Joins("LEFT JOIN status_legal b on b.id_status_legal = a.id_status_legal_fk").
		Joins("LEFT JOIN kepengurusan c on c.id_kepengurusan = a.id_kepengurusan_fk").
		Joins("LEFT JOIN teknik_irigasi d on d.id_t_irigasi = a.id_teknik_irigasi_fk").
		Joins("LEFT JOIN teknik_pertanian e on e.id_t_pertanian = a.id_teknik_pertanian_fk").
		Joins("LEFT JOIN provinsi f on f.id_prov = a.id_prov_fk").
		Joins("LEFT JOIN kabupaten g on g.id_kab = a.id_kab_fk").
		Joins("LEFT JOIN kecamatan h on h.id_kec = a.id_kec_fk").
		Where("id_p3a = ?", id).Find(&entity).Error
	return &entity, err
}

func (r p3Repository) UpdateById(entity models.Inventaris)(*models.Inventaris, error){
	err := r.DB.Model(&models.Inventaris{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r p3Repository) Delete(entity models.Inventaris) error {
	return r.DB.Table("data_p3a").Delete(&entity).Error
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
	err := r.DB.Table("data_p3a").Create(&entity).Error
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

func (r p3Repository) UpdateStatusLegal(entity models.StatusLegal)(*models.StatusLegal, error){
	err := r.DB.Model(&models.StatusLegal{ID: entity.ID}).UpdateColumns(&entity).Error
	return &entity, err
}

func (r p3Repository) UpdatePengurus(entity models.Pengurus)(*models.Pengurus, error){
	err := r.DB.Model(&models.Pengurus{ID: entity.ID}).UpdateColumns(&entity).Error
	return &entity, err
}

func (r p3Repository) UpdateIrigasi(entity models.TeknikIrigasi)(*models.TeknikIrigasi, error){
	err := r.DB.Model(&models.TeknikIrigasi{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}


func (r p3Repository) UpdatePertanian(entity models.TeknikPertanian)(*models.TeknikPertanian, error){
	err := r.DB.Model(&models.TeknikPertanian{ID: entity.ID}).Updates(&entity).Error
	return &entity, err
}

func (r p3Repository) ExportExcel()(*[]models.P3AModels, error) {
	var entity []models.P3AModels
	err := r.DB.Table("data_p3a a").
		Select("a.*,b.*,c.*,d.*,e.*,f.*,g.*,h.*").
		Joins("LEFT JOIN status_legal b on b.id_status_legal = a.id_status_legal_fk").
		Joins("LEFT JOIN kepengurusan c on c.id_kepengurusan = a.id_kepengurusan_fk").
		Joins("LEFT JOIN teknik_irigasi d on d.id_t_irigasi = a.id_teknik_irigasi_fk").
		Joins("LEFT JOIN teknik_pertanian e on e.id_t_pertanian = a.id_teknik_pertanian_fk").
		Joins("LEFT JOIN provinsi f on f.id_prov = a.id_prov_fk").
		Joins("LEFT JOIN kabupaten g on g.id_kab = a.id_kab_fk").
		Joins("LEFT JOIN kecamatan h on h.id_kec = a.id_kec_fk").
		Find(&entity).Error
	return &entity, err
}


func (r p3Repository) DeleteStatusLegal(id string) error {
	entity := models.StatusLegal{
		ID: id,
	}
	return r.DB.Table("status_legal").Delete(&entity).Error
}


func (r p3Repository) DeleteIrigasi(id string) error {
	entity := models.TeknikIrigasi{
		ID: id,
	}
	return r.DB.Table("teknik_irigasi").Delete(&entity).Error
}

func (r p3Repository) DeletePengurusan(id string) error {
	entity := models.Pengurus{
		ID: id,
	}
	return r.DB.Table("kepengurusan").Delete(&entity).Error
}


func (r p3Repository) DeletePertanian(id string) error {
	entity := models.TeknikPertanian{
		ID: id,
	}
	return r.DB.Table("teknik_pertanian").Delete(&entity).Error
}

func (r *p3Repository) DbInstance() *gorm.DB {
	return r.DB
}


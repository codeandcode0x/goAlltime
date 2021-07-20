package model

import (
	"gin-scaffold/db"
	"time"

	"gorm.io/gorm"
)

// base model
type BaseModel struct {
	ID        uint64 `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// base dao
type BaseDAO interface {
	Create(entity interface{}) error
	Update(entity interface{}, uid uint64) (int64, error)
	Delete(entity interface{}, uid uint64) (int64, error)
	FindAll(entity interface{}, opts ...DAOOption) error
	FindByKeys(entity interface{}, keys map[string]interface{}) error
	FindByPages(entity interface{}, currentPage, pageSize int) error
	FindByPagesWithKeys(entity interface{}, keys map[string]interface{}, currentPage, pageSize int, opts ...DAOOption) error
	SearchByPagesWithKeys(entity interface{}, keys, keyOpts map[string]interface{}, currentPage, pageSize int, opts ...DAOOption) error
	Count(entity interface{}, count *int64) error
	CountWithKeys(entity interface{}, count *int64, keys, keyOpts map[string]interface{}, opts ...DAOOption) error
}

type DAOOption struct {
	Select string
	Order  string
	Where  map[string]interface{}
	Limit  int
}

// create
func (u *BaseModel) Create(entity interface{}) error {
	result := db.Conn.Model(entity).Create(entity)
	return result.Error
}

// update
func (u *BaseModel) Update(entity interface{}, uid uint64) (int64, error) {
	result := db.Conn.Model(entity).Where("id = ?", uid).Save(entity)
	return result.RowsAffected, result.Error
}

// delete
func (u *BaseModel) Delete(entity interface{}, uid uint64) (int64, error) {
	result := db.Conn.Model(entity).Where("id = ?", uid).Delete(entity)
	return result.RowsAffected, result.Error
}

// find all
func (u *BaseModel) FindAll(entity interface{}, opts ...DAOOption) error {
	/**
	// 反射获取类型
	modelType := reflect.ValueOf(entity).Type()
	log.Println("model ........", modelType)
	// switch type 获取类型
	switch entity.(type) {
	case *Order:
		break
	}
	*/
	tx := db.Conn.Model(entity)
	if len(opts) > 0 {
		beginCustomTx(tx, opts)
	}

	tx.Find(entity)
	return tx.Error
}

// find by id
func (u *BaseModel) FindByKeys(entity interface{}, keys map[string]interface{}) error {
	result := db.Conn.Model(entity).Where(keys).Find(entity)
	return result.Error
}

// find by pages
func (u *BaseModel) FindByPages(entity interface{}, currentPage, pageSize int) error {
	result := Paginate(currentPage, pageSize).Model(entity).Find(entity)
	return result.Error
}

// find by pages with keys
func (u *BaseModel) FindByPagesWithKeys(entity interface{},
	keys map[string]interface{},
	currentPage, pageSize int,
	opts ...DAOOption) error {
	// tx begiin
	tx := Paginate(currentPage, pageSize).Model(entity)
	beginCustomTx(tx, opts)
	result := tx.Where(keys).Find(entity)
	return result.Error
}

// find by pages with keys
func (u *BaseModel) Count(entity interface{}, count *int64) error {
	result := db.Conn.Model(entity).Count(count)
	return result.Error
}

// find by pages with keys
func (u *BaseModel) CountWithKeys(entity interface{},
	count *int64,
	keys, keyOpts map[string]interface{},
	opts ...DAOOption) error {
	// tx begin
	tx := db.Conn.Model(entity)
	beginCustomTx(tx, opts)
	searchCustomTx(tx, keys, keyOpts)
	tx.Count(count)
	return tx.Error
}

// search by pages
func (u *BaseModel) SearchByPagesWithKeys(entity interface{},
	keys, keyOpts map[string]interface{},
	currentPage, pageSize int,
	opts ...DAOOption) error {
	// tx begin
	tx := Paginate(currentPage, pageSize).Model(entity)
	beginCustomTx(tx, opts)
	searchCustomTx(tx, keys, keyOpts)
	tx.Find(entity)
	return tx.Error
}

// search custom tx
func searchCustomTx(tx *gorm.DB, keys, keyOpts map[string]interface{}) {
	compareOpt := map[string]string{">": ">", "<": "<", ">=": ">=", "<=": "<=", "=": "="}
	for k, v := range keys {
		usingLike := true
		// set tx
		if _, exist := keyOpts[k]; exist {
			tx.Where(k+keyOpts[k].(string)+" ?", v.(string))
			if _, ok := compareOpt[keyOpts[k].(string)]; ok {
				usingLike = false
			}
		}
		// no using like
		if usingLike {
			tx.Where(k+" LIKE ?", "%"+v.(string)+"%")
		}
	}
}

// 分页 paginate
func Paginate(currentPage, pageSize int) *gorm.DB {
	offset := (currentPage - 1) * pageSize
	return db.Conn.Offset(offset).Limit(pageSize)
}

// begin custom tx
func beginCustomTx(tx *gorm.DB, opts []DAOOption) {
	for _, opt := range opts {
		// set tx order
		if opt.Order != "" {
			tx.Order(opt.Order)
		}
		// set tx where
		if opt.Where != nil {
			tx.Where(opt.Where)
		}
		// set tx select
		if opt.Select != "" {
			tx.Select(opt.Select)
		}
		// set tx limit
		if opt.Limit > 0 {
			tx.Limit(int(opt.Limit))
		}
	}
}

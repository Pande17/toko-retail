package modelfunc

import (
	"fmt"
	"projek/toko-retail/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Diskon struct {
	model.Diskon
}

func (kd *Diskon) CreateDiskon(db *gorm.DB) error {
	err := db.Model(Diskon{}).Create(&kd).Error
	if err != nil {
		return err
	}

	return nil
}

func (kd *Diskon) GetAll(db *gorm.DB) ([]Diskon, error) {
	res := []Diskon{}

	err := db.
		Model(Diskon{}).
		Find(&res).
		Error

	if err != nil {
		return []Diskon{}, err
	}

	return res, nil
}

func (kd *Diskon) GetByCode(db *gorm.DB, kodediskon string) (Diskon, error) {
    res := Diskon{}
    
    logrus.Println("Executing query with kode_diskon:", kodediskon)
    
    // Check database connection
    sqlDB, err := db.DB()
    if err != nil {
        logrus.Println("Error getting database connection:", err)
        return Diskon{}, err
    }

    if err := sqlDB.Ping(); err != nil {
        logrus.Println("Database connection failed:", err)
        return Diskon{}, err
    }

    err = db.
        Where("kode_diskon = ?", kodediskon).
        First(&res).
        Error

    if err != nil {
        logrus.Println("Error executing query:", err)
        if err == gorm.ErrRecordNotFound {
            return Diskon{}, fmt.Errorf("record not found")
        }
        return Diskon{}, err
    }

    logrus.Println("Query successful, found record:", res)
    return res, nil
}



func (kd *Diskon) GetByID(db *gorm.DB) (Diskon, error) {
	res := Diskon{}

	err := db.
		Model(Diskon{}).
		Where("id = ?", kd.ID).
		Take(&res).
		Error

	if err != nil {
		return Diskon{}, err
	}

	return res, nil
}

func (kd *Diskon) Delete(db *gorm.DB) error {
	err := db.
		Model(&Diskon{}).
		Where("id = ?", kd.ID).
		Delete(&kd).
		Error

	if err != nil {
		return err
	}

	return nil
}
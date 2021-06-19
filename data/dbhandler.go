package data

import (
	"fmt"
	"mohammadinasab-dev/grpctask/configuration"
	"mohammadinasab-dev/grpctask/logwrapper"
	"mohammadinasab-dev/grpctask/protos"

	"github.com/jinzhu/gorm"
)

type MySqlDB struct {
	DB     *gorm.DB
	STDLog *logwrapper.STDLog
}

func CreateMySQLDBConnection(config configuration.Config, STDLog *logwrapper.STDLog) (*MySqlDB, error) {
	connstring := fmt.Sprintf(config.DBUsername + ":" + config.DBPassword + "@" + config.DBAddress + "/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(config.DBDriver, connstring)
	if err != nil {
		STDLog.WarningLogger.Fatalln(err)
		return nil, err
	}
	db.AutoMigrate(&Product{})
	return &MySqlDB{
		DB:     db,
		STDLog: STDLog,
	}, nil

}

// func CreateMyOtherBConnection(config configuration.Config) (*MyOtherDB, error){...}

func (ms *MySqlDB) DBGetProduct(in *protos.ProductRequest) (*Product, error) { // think about (ms *MySqlDB)
	ms.STDLog.ErrorLogger.Println("i'm here in DBGetProduct")

	product := &Product{}
	// if result := ms.DB.Where("p_id = ?", in.Id).Where("currency = ?", in.Currency).First(&product); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
	if result := ms.DB.Where("p_id = ?", in.Id).First(&product); result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		return &Product{}, result.Error
	}
	return product, nil
}

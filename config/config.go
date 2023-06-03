package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
  var err error
  db, err = InitializeSQLite()

  if err != nil {
    return fmt.Errorf("error initializing sqlite: %v", err)
  }

	return nil
}

func GetLogger(p string) *Logger {
	return NewLogger(p)
}

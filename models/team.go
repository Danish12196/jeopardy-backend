package models

import "gorm.io/gorm"

type Team struct {
  gorm.Model
  GameID uint
  Name   string
  Score  int
}

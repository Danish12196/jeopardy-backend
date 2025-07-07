package models

import "gorm.io/gorm"

type Game struct {
  gorm.Model
  JoinCode   string `gorm:"unique;not null"`
  OwnerID    uint
  Status     string // "waiting", "active", "finished"
  StartedAt  *string
  FinishedAt *string
  Teams      []Team `gorm:"foreignKey:GameID"`
  WinnerID   *uint // ID of the winning team
}

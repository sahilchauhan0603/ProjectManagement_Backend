package models

import "time"

type Admin struct {
	AdminID   int64 `gorm:"primaryKey;autoIncrement"`
	AdminName string
	Username  string
	Password  string
	Uploader  []Uploader `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Project   []Project  `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
}
type Project struct {
	ProjectID   int64 `gorm:"primaryKey;autoIncrement"`
	Title       string
	Topic       string
	Approach    string
	Description string
	Supervisor  string
	AdminID     int64 `gorm:"index"`
	Proof       string
	Uploader    []Uploader `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
}
type Uploader struct {
	EnrollmentNo int64 `gorm:"primaryKey;autoIncrement"`
	FirstName    string
	LastName     string
	CollegeName  string
	ProjectID    int64  `gorm:"index"`
	AdminID      int64  `gorm:"index"`
	Email        string `gorm:"unique"`
	CreatedAt    time.Time
}

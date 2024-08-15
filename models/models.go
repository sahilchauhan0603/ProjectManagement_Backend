package models

import "time"

type ProjectAdmin struct {
	AdminID            int64 `gorm:"primaryKey;autoIncrement"`
	AdminName          string
	Username           string
	Password           string
	ProjectIdsToVerify string
	Uploader           []ProjectUploader `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Work               []ProjectWork     `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProjectWork struct {
	WorkID      int64 `gorm:"primaryKey;autoIncrement"`
	Title       string
	Type        string
	Approach    string
	Description string
	Supervisor  string
	AdminID     int64 `gorm:"index"`
	Proof       string
	Uploaders   []ProjectUploader `gorm:"foreignKey:WorkID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StartDate   time.Time
	EndDate     time.Time
	Category    string
}

type ProjectUploader struct {
	EnrollmentNo    int64 `gorm:"primaryKey;autoIncrement"`
	FirstName       string
	LastName        string
	InstitutionName string
	WorkID          int64  `gorm:"index"`
	AdminID         int64  `gorm:"index"`
	Email           string `gorm:"unique"`
	PhoneNo         string
	Department      string
	BatchStart      time.Time
	BatchEnd        time.Time
	GithubProfile   *string
	LinkedInProfile *string
}

type EmailRequest struct {
	Email string `json:"email"`
}

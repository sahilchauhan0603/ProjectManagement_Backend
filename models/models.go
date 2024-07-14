package models

import "time"

// Admin represents the Admin model
// @Description Admin model
type Admin struct {

	// ID of the admin
	// required: true
	AdminID int64 `json:"admin_id" gorm:"primaryKey;autoIncrement"`

	// Name of the Admin
	// required: true
	AdminName string `json:"admin_name"`

	// UerName of the Admin
	// required: true
	Username string `json:"username"`

	// Password of the Admin
	// required: true
	Password string `json:"password"`

	// Uploader associated with the Admin
	Uploader []Uploader `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Projects associated with the Admin
	Project []Project `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// CreatedAt timestamp
	CreatedAt time.Time
}

// Project represents the project model
// @Description Project model
type Project struct {
	// ProjectID is the primary key
	// required: true
	ProjectID int64 `json:"project_id" gorm:"primaryKey;autoIncrement"`

	// Title of the project
	// required: true
	Title string `json:"title"`

	// Topic of the project
	// required: true
	Topic string `json:"topic"`

	// Approach used in the project
	// required: true
	Approach string `json:"approach"`

	// Description of the project
	// required: true
	Description string `json:"description"`

	// Supervisor of the project
	// required: true
	Supervisor string `json:"supervisor"`

	// AdminID is the foreign key to Admin
	// required: true
	AdminID int64 `json:"admin_id" gorm:"index"`

	// Proof related to the project
	// required: true
	Proof string `json:"proof"`

	// Uploaders associated with the project
	Uploaders []Uploader `json:"uploaders" gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// CreatedAt timestamp
	CreatedAt time.Time `json:"created_at"`
}

// Uploader represents the uploader model
// @Description Uploader model
type Uploader struct {
	// EnrollmentNo is the primary key
	// required: true
	EnrollmentNo int64 `json:"enrollment_no" gorm:"primaryKey;autoIncrement"`

	// FirstName of the uploader
	// required: true
	FirstName string `json:"first_name"`

	// LastName of the uploader
	// required: true
	LastName string `json:"last_name"`

	// CollegeName of the uploader
	// required: true
	CollegeName string `json:"college_name"`

	// ProjectID is the foreign key to Project
	// required: true
	ProjectID int64 `json:"project_id" gorm:"index"`

	// AdminID is the foreign key to Admin
	// required: true
	AdminID int64 `json:"admin_id" gorm:"index"`

	// Email of the uploader
	// required: true
	Email string `json:"email" gorm:"unique"`

	// CreatedAt timestamp
	CreatedAt time.Time `json:"created_at"`
}

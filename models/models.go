package models

type Uploader struct {
	EnrollmentNo int64 `gorm:"primaryKey;autoIncrement"`
	FirstName    string
	LastName     string
	CollegeName  string
	ProjectID    int64 `gorm:"index"`
	AdminID      int64 `gorm:"index"`
}

type Project struct {
	ProjectID   int64 `gorm:"index"`
	Title       string
	Topic       string
	Approach    string
	Description string
	Supervisor  string
	AdminID     int64 `gorm:"index"`
	Proof       string
}

type Admin struct {
	AdminID   int64 `gorm:"index"`
	AdminName string
	Username  string
	Password  string
}

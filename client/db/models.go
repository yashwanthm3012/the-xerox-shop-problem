package db

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	Password  string
	PrintJobs []PrintJob `gorm:"foreignKey:UserID"` // One-to-many relationship
}

type PrintJob struct {
	ID         uint `gorm:"primaryKey"`
	FileName   string
	UserID     uint // Foreign key
	User       User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Foreign key constraint
	PrintType  string
	NumPages   string
	Copies     int
	BWPages    string
	ColorPages string
}

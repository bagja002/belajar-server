package models

type Admin struct {
	IdAdmin uint		`gorm:"primaryKey;autoIncrement:true"`
	Nama    string
	NoTelpon int 
	Email string
	Password string
	Status string
}

type User struct {
	IdUser      uint 		`gorm:"primaryKey;autoIncrement:true"`
	NamaLengkap string
	NoTelpon    int
	Email       string
	Password    string
	Status string
}

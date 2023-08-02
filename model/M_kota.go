package model

type Kota struct {
	IdKota   int64  `gorm:"primarykey" json:"id"`
	NamaKota string `gorm:"varchar(255)" json:"nama_kota"`
}

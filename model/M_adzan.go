package model

type Adzan struct {
	Id      int64  `gorm:"primarykey" json:"id"`
	IDKota  int64  `gorm:"column:id_kota;not null" json:"id_kota"`
	Tanggal string `gorm:"type:date" json:"tgl_adzan"`
	Imsak   string `gorm:"varchar(255)" json:"imsak"`
	Shubuh  string `gorm:"varchar(255)" json:"shubuh"`
	Terbit  string `gorm:"varchar(255)" json:"terbit"`
	Duha    string `gorm:"varchar(255)" json:"duha"`
	Dzuhur  string `gorm:"varchar(255)" json:"dzuhur"`
	Ashar   string `gorm:"varchar(255)" json:"ashar"`
	Mahgrib string `gorm:"varchar(255)" json:"mahgrib"`
	Isya    string `gorm:"varchar(255)" json:"isya"`

	// Add Kota field to hold the related Kota data
	Kota Kota `gorm:"foreignKey:IDKota" json:"kota"`
}

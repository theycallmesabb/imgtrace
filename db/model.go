package db

type Image struct {
	FilePath string `gorm:"type:varchar(255)"`
	Hash     uint64 `gorm:"uniqueIndex"` // unique index on hash

}

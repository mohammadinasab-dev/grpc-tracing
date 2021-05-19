package data

type Product struct {
	PID      int `gorm:"primary_key"`
	Name     string
	Price    float32 `gorm:"NOT NULL; UNIQUE"`
	Currency string  `gorm:"NOT NULL; UNIQUE"`
}

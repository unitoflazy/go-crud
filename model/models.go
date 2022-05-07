package model

type Employee struct {
	ID   int    `json:"id" gorm:"primaryKey;not null;autoIncrement:true"`
	Name string `json:"name" gorm:"type:VARCHAR(200) NOT NULL"`
	Age  int    `json:"age" gorm:"type:INT NOT NULL; check:age >= 1"`
}

// gorm:"primaryKey;type:INT NOT NULL;autoIncrement:true"
// gorm:"type:VARCHAR(200) NOT NULL;check:age >= 1"
// gorm:"type:INT NOT NULL"

package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	ProductID uint  
	Name      string
	Review    string 
	Rating    int   
}

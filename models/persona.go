package models

type Persona struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Cedula    string `json:"cedula"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Telefono2 string `json:"telefono2"`
	Email     string `json:"email"`
	Email2    string `json:"email2"`
}

package avitoStartApp

type RequestName struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}

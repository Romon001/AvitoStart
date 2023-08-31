package avitoStartApp

type Segment struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}

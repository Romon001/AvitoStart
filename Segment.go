package avitoStartApp

type Segment struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"slug" binding:"required"`
}

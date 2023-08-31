package avitoStartApp

type UserSegmentPair struct {
	Id      int    `json:"-" db:"id"`
	Segment string `json:"segment" binding:"required"`
	User    string `json:"user" binding:"required"`
}

package avitoStartApp

type UserSegmentPair struct {
	Id      int    `json:"-" db:"id"`
	Segment string `json:"segmentName" db:"segmentName" binding:"required"`
	User    int    `json:"userId" db:"userId" binding:"required"`
}

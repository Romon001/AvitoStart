package avitoStartApp

type UserSegmentPair struct {
	Id      int    `json:"-" db:"id"`
	Segment string `json:"segmentName" db:"segmentname" binding:"required"`
	User    int    `json:"userId" db:"userid" binding:"required"`
}

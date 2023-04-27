package constform

type Sex int

const (
	Male   Sex = iota // 0
	Female            // 1
)

type Step int

const (
	FullInfo   Step = iota //0
	NoInfo                 //1
	NoHashtags             //2
	NoFilters              //3
	NoPhotos               //4
)

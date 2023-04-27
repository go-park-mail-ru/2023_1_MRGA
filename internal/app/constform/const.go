package constform

type Sex int

const (
	Male   Sex = iota // 0
	Female            // 1
)

type Step int

const (
	FullInfo Step = iota //0
	MainInfo             //1
	Hashtag              //2
	Filters              //3
	Photos               //4
)

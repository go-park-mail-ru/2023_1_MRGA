package dataStruct

type User struct {
	Id       uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userId" gorm:"primaryKey;unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	BirthDay string `json:"birthDay" sql:"type:date" gorm:"type:date"`
}

///Info

type UserInfo struct {
	Id          uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userInfoId" gorm:"primaryKey;unique"`
	UserId      uint   `json:"userId" gorm:"foreignKey;unique"`
	Name        string `json:"name"`
	CityId      uint   `json:"cityId" gorm:"foreignKey"`
	Sex         uint   `json:"sex" gorm:"foreignKey"`
	Description string `json:"description"`
	Zodiac      uint   `json:"zodiac" gorm:"foreignKey"`
	Job         uint   `json:"job" gorm:"foreignKey"`
	Education   uint   `json:"education" gorm:"foreignKey"`
}

type Sex struct {
	Id  uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"sexId" gorm:"primaryKey;unique"`
	Sex string `json:"sex"`
}
type UserPhoto struct {
	Id     uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userPhotoId" gorm:"primaryKey;unique"`
	UserId uint `json:"userId" gorm:"foreignKey"`
	Photo  uint `json:"photo"`
	Avatar bool `json:"avatar"`
}

type City struct {
	Id   uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"cityId" gorm:"primaryKey;unique"`
	City string `json:"city"`
}

type Zodiac struct {
	Id     uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"zodiacId" gorm:"primaryKey;unique"`
	Zodiac string `json:"zodiac"`
}

type Job struct {
	Id  uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"jobId" gorm:"primaryKey;unique"`
	Job string `json:"job"`
}

type Education struct {
	Id        uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"educationId" gorm:"primaryKey;unique"`
	Education string `json:"education"`
}

///Filter

type UserFilter struct {
	Id        uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userFilterId" gorm:"primaryKey;unique"`
	UserId    uint `json:"userId" gorm:"foreignKey;unique"`
	MinAge    int  `json:"minAge"`
	MaxAge    int  `json:"maxAge"`
	SearchSex uint `json:"sexId" gorm:"foreignKey"`
}

type UserReason struct {
	Id       uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userReasonId" gorm:"primaryKey;unique"`
	UserId   uint `json:"userId" gorm:"foreignKey"`
	ReasonId uint `json:"reasonId" gorm:"foreignKey"`
}

type Reason struct {
	Id     uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"reasonId" gorm:"primaryKey;unique"`
	Reason string `json:"reason"`
}

type UserHashtag struct {
	Id        uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userHashtag" gorm:"primaryKey;unique"`
	UserId    uint `json:"userId" gorm:"foreignKey"`
	HashtagId uint `json:"hashtagId" gorm:"foreignKey"`
}

type Hashtag struct {
	Id      uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"hashtagId" gorm:"primaryKey;unique"`
	Hashtag string `json:"hashtag"`
}

///Match

type UserReaction struct {
	Id         uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userReactionId" gorm:"primaryKey;unique"`
	UserId     uint `json:"userId" gorm:"foreignKey"`
	UserFromId uint `json:"userFromId" gorm:"foreignKey"`
	ReactionId uint `json:"reactionIdId" gorm:"foreignKey"`
}

type Reaction struct {
	Id       uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"reactionId" gorm:"primaryKey;unique"`
	Reaction string `json:"reaction"`
}

type UserHistory struct {
	Id            uint   `sql:"unique;type:uuid;primary_key;servicedefault:" json:"userHistoryId" gorm:"primaryKey;unique"`
	UserId        uint   `json:"userId" gorm:"foreignKey"`
	UserProfileId uint   `json:"userProfileId" gorm:"foreignKey"`
	ShowDate      string `json:"birthDay" sql:"type:date" gorm:"type:date"`
}

type Match struct {
	Id           uint `sql:"unique;type:uuid;primary_key;servicedefault:" json:"MarchId" gorm:"primaryKey;unique"`
	UserFirstId  uint `json:"userFirstId" gorm:"foreignKey"`
	UserSecondId uint `json:"userSecondId" gorm:"foreignKey"`
	Shown        bool `json:"shown"`
}

package models

type RecieveVansStruct struct {
	Name string `bson:"name"`
	Code string `bson:"code"`
	Desc string `bson:"desc"`
}

type ReturnVansStruct struct {
	Id        string   `bson:"_id"`
	Name      string   `bson:"name"`
	Code      string   `bson:"code"`
	Desc      string   `bson:"desc"`
	ImagePath []string `bson:"imagePath"`
}

type VansStruct struct {
	Name      string   `bson:"name"`
	Code      string   `bson:"code"`
	Desc      string   `bson:"desc"`
	ImagePath []string `bson:"imagePath"`
}

type ScheduleStruct struct {
	VanId       string `bson:"van_id"`
	Date        string `bson:"date"`
	Destination string `bson:"destination"`
}

type ReturnScheduleStruct struct {
	Id          string `bson:"_id"`
	VanId       string `bson:"van_id"`
	Date        string `bson:"date"`
	Destination string `bson:"destination"`
}

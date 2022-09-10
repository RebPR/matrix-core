package biz

type Article struct {
	ArticleId int32
	Uuid      string
	Status    int32
	Auth      int32
}

type Talk struct {
	TalkId int32
	Uuid   string
	Status int32
	Auth   int32
}

type ArticleStatistic struct {
	ArticleId int32
	Uuid      string
	Agree     int32
	View      int32
	Collect   int32
	Comment   int32
	Auth      int32
}

type TalkStatistic struct {
	TalkId  int32
	Uuid    string
	Agree   int32
	View    int32
	Collect int32
	Comment int32
	Auth    int32
}

type ArticleStatisticJudge struct {
	Agree   bool
	Collect bool
}

type TalkStatisticJudge struct {
	Agree   bool
	Collect bool
}

type ArticleDraft struct {
	Id     int32
	Status int32
	Uuid   string
}

type TalkDraft struct {
	Id     int32
	Status int32
	Uuid   string
}

type ArticleReview struct {
	Uuid string
	Id   int32
	Mode string
}

type TalkReview struct {
	Uuid string
	Id   int32
	Mode string
}

type LeaderBoard struct {
	Id    int32
	Agree int32
	Uuid  string
	Mode  string
}

type Collections struct {
	CollectionsId int32
	Uuid          string
	Article       int32
	Column        int32
	Talk          int32
	Auth          int32
}

type CollectionsDraft struct {
	Id     int32
	Status int32
	Uuid   string
}

type CollectionsReview struct {
	Uuid string
	Id   int32
	Mode string
}

type Column struct {
	ColumnId int32
	Uuid     string
	Auth     int32
}

type ColumnDraft struct {
	Id     int32
	Status int32
	Uuid   string
}

type ColumnReview struct {
	Uuid string
	Id   int32
	Mode string
}

type ColumnStatistic struct {
	ColumnId int32
	Uuid     string
	Agree    int32
	View     int32
	Collect  int32
	Auth     int32
}

type ColumnStatisticJudge struct {
	Agree   bool
	Collect bool
}

type Subscribe struct {
	ColumnId int32
	AuthorId string
	Uuid     string
	Status   int32
}

type News struct {
	Id     string
	Update string
	Title  string
	Author string
	Text   string
	Tags   string
	Cover  string
	Url    string
}

type ArticleSearch struct {
	Id     int32
	Total  int32
	Title  string
	Tags   string
	Uuid   string
	Text   string
	Cover  string
	Update string
}

type TalkSearch struct {
	Id     int32
	Total  int32
	Title  string
	Tags   string
	Uuid   string
	Text   string
	Cover  string
	Update string
}

type ColumnSearch struct {
	Id        int32
	Total     int32
	Name      string
	Tags      string
	Uuid      string
	Introduce string
	Cover     string
	Update    string
}

type CreationUser struct {
	Article     int32
	Column      int32
	Talk        int32
	Collections int32
	Collect     int32
	Subscribe   int32
}

type ImageReview struct {
	Id         int32
	CreationId int32
	Kind       string
	Uid        string
	CreateAt   string
	Uuid       string
	JobId      string
	Url        string
	Label      string
	Result     int32
	Category   string
	SubLabel   string
	Mode       string
	Score      int32
}

type TextReview struct {
	Id         int32
	CreationId int32
	CreateAt   string
	Title      string
	Kind       string
	JobId      string
	Label      string
	Result     int32
	Uuid       string
	Mode       string
	Section    string
}

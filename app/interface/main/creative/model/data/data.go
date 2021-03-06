package data

import (
	xtime "go-common/library/time"

	"go-common/app/interface/main/creative/model/medal"
)

// Stat info
type Stat struct {
	FanLast     int64             `json:"fan_last"`
	Fan         int64             `json:"fan"`
	DmLast      int64             `json:"dm_last"`
	Dm          int64             `json:"dm"`
	CommentLast int64             `json:"comment_last"`
	Comment     int64             `json:"comment"`
	Play        int64             `json:"play"`
	PlayLast    int64             `json:"play_last"`
	Fav         int64             `json:"fav"`
	FavLast     int64             `json:"fav_last"`
	Like        int64             `json:"like"`
	LikeLast    int64             `json:"like_last"`
	Share       int64             `json:"share"`
	ShareLast   int64             `json:"share_last"`
	Elec        int64             `json:"elec"`
	ElecLast    int64             `json:"elec_last"`
	Coin        int64             `json:"coin"`
	CoinLast    int64             `json:"coin_last"`
	Day30       map[string]int    `json:"30,omitempty"`
	Arcs        map[string][]*Arc `json:"arcs"`
}

// Arc Arcs info.
type Arc struct {
	Aid   int64  `json:"aid"`
	Title string `json:"title"`
	Click int64  `json:"click"`
}

// Tags Arcs info
type Tags struct {
	Tags []string `json:"tags"`
}

// CheckedTag tags with checked
type CheckedTag struct {
	Tag     string `json:"tag"`
	Checked int    `json:"checked"`
}

// AppStat arc stat.
type AppStat struct {
	Date string `json:"date"`
	Num  int64  `json:"num"`
}

// AppStatList for arc stat list.
type AppStatList struct {
	Danmu   []*AppStat `json:"danmu"`
	View    []*AppStat `json:"view"`
	Fans    []*AppStat `json:"fans"`
	Comment []*AppStat `json:"comment"`
	Show    int8       `json:"show"`
}

// ViewerTrend for up trend data.
type ViewerTrend struct {
	Tag map[int]string   `json:"tag"`
	Ty  map[string]int64 `json:"ty"`
}

// ViewerIncr for up increment data.
type ViewerIncr struct {
	Arcs      []*ArcInc      `json:"arc_inc"`
	TotalIncr int            `json:"total_inc"`
	TyRank    map[string]int `json:"type_rank"`
}

// ArcInc for archive increment data.
type ArcInc struct {
	AID     int64      `json:"aid"`
	Incr    int        `json:"incr"`
	Title   string     `json:"title"`
	DayTime int64      `json:"daytime"`
	PTime   xtime.Time `json:"ptime"`
}

// PeriodTip  period tip for data.
type PeriodTip struct {
	ModuleOne   string `json:"module_one"`
	ModuleTwo   string `json:"module_two"`
	ModuleThree string `json:"module_three"`
	ModuleFour  string `json:"module_four"`
}

// AppViewerIncr for up increment data.
type AppViewerIncr struct {
	DateKey   int64     `json:"date_key"`
	Arcs      []*ArcInc `json:"arc_inc"`
	TotalIncr int       `json:"total_inc"`
	TyRank    []*Rank   `json:"type_rank"`
}

// ThirtyDay for 30 days data.
type ThirtyDay struct {
	DateKey   int64 `json:"date_key"`
	TotalIncr int64 `json:"total_inc"`
}

// Rank type rank for up data.
type Rank struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

// CreatorDataShow for display archive/article data module.
type CreatorDataShow struct {
	Archive int `json:"archive"`
	Article int `json:"article"`
}

// AppFan for stat.
type AppFan struct {
	Summary    map[string]int64            `json:"summary"`
	RankMap    map[string]map[string]int32 `json:"-"`
	TypeList   map[string]int64            `json:"type_list"`
	TagList    []*Rank                     `json:"tag_list"`
	ViewerArea map[string]int64            `json:"viewer_area"`
	ViewerBase *ViewerBase                 `json:"viewer_base"`
}

//AppOverView for data overview.
type AppOverView struct {
	Stat         *Stat        `json:"stat"`
	AllArcIncr   []*ThirtyDay `json:"all_arc_inc"`
	SingleArcInc []*ArcInc    `json:"single_arc_inc"`
}

//VideoQuit customer play Retention Rate
type VideoQuit struct {
	Point    []int64 `json:"point"`
	Duration []int64 `json:"duration"`
}

//for fan manager top mids.
const (
	//Total ????????????-????????????
	Total = iota
	//Seven ????????????-7?????????
	Seven
	//Thirty ????????????-30?????????
	Thirty
	//Ninety ????????????-90?????????
	Ninety
	//PlayDuration ????????????
	PlayDuration = "video_play"
	//VideoAct ????????????
	VideoAct = "video_act"
	//DynamicAct ????????????
	DynamicAct = "dynamic_act"
)

// WebFan for stat.
type WebFan struct {
	RankMap   map[string]map[string]int32  `json:"-"`
	Summary   map[string]int32             `json:"summary"`
	RankList  map[string][]*RankInfo       `json:"rank_list"`
	RankMedal map[string][]*medal.FansRank `json:"rank_medal"`
	Source    map[string]int32             `json:"source"`
}

// FansAnalysis for medal fans.
type FansAnalysis struct {
	F map[string]int32 `family:"f"`
	T map[string]int32 `family:"t"`
	S map[string]int32 `family:"s"`
}

//FansAnalysisByF for f family
type FansAnalysisByF struct {
	All      int32 `family:"f" qualifier:"all" json:"total"`           //?????????
	Inc      int32 `family:"f" qualifier:"inc" json:"inc"`             //????????????
	Act      int32 `family:"f" qualifier:"act" json:"active"`          //????????????
	Mdl      int32 `family:"f" qualifier:"mdl" json:"medal"`           //??????????????????
	Elec     int32 `family:"f" qualifier:"elec" json:"elec"`           //????????????
	ActDiff  int32 `family:"f" qualifier:"act_diff" json:"act_diff"`   //????????????????????????
	MdlDiff  int32 `family:"f" qualifier:"mdl_diff" json:"mdl_diff"`   //??????????????????????????????
	ElecDiff int32 `family:"f" qualifier:"elec_diff" json:"elec_diff"` //??????????????????????????????
	Play     int32 `family:"f" qualifier:"v" json:"v"`                 //??????????????????*10000
	Dm       int32 `family:"f" qualifier:"dm" json:"dm"`               //??????????????????*10000
	Reply    int32 `family:"f" qualifier:"r" json:"r"`                 //??????????????????*10000
	Coin     int32 `family:"f" qualifier:"c" json:"c"`                 //??????????????????*10000
	//??????????????????
	InterRatio int32 `family:"f" qualifier:"inter" json:"inter"` //???????????????*10000
	PlayRatio  int32 `family:"f" qualifier:"vv" json:"vv"`       //???????????????*10000
	DanmuRatio int32 `family:"f" qualifier:"da" json:"da"`       //?????????????????????*10000
	ReplyRatio int32 `family:"f" qualifier:"re" json:"re"`       //?????????????????????*10000
	CoinRatio  int32 `family:"f" qualifier:"co" json:"co"`       //?????????????????????*10000
	FavRatio   int32 `family:"f" qualifier:"fv" json:"fv"`       //?????????????????????*10000
	ShareRatio int32 `family:"f" qualifier:"sh" json:"sh"`       //?????????????????????*10000
	LikeRatio  int32 `family:"f" qualifier:"lk" json:"lk"`       //?????????????????????*10000
	//3.4??????
	NewAct    int32 `family:"f" qualifier:"new_act" json:"new_act"`     //?????????????????????
	NewInter  int32 `family:"f" qualifier:"new_inter" json:"new_inter"` //?????????????????????
	NewReply  int32 `family:"f" qualifier:"new_re" json:"new_re"`       //??????
	NewDanmu  int32 `family:"f" qualifier:"new_da" json:"new_da"`       //??????
	NewCoin   int32 `family:"f" qualifier:"new_co" json:"new_co"`       //??????
	NewLike   int32 `family:"f" qualifier:"new_lk" json:"new_lk"`       //??????
	NewFav    int32 `family:"f" qualifier:"new_fv" json:"new_fv"`       //??????
	NewShare  int32 `family:"f" qualifier:"new_sh" json:"new_sh"`       //??????
	LiveDanmu int32 `family:"f" qualifier:"live_dm" json:"live_dm"`     //????????????
	LiveCoin  int32 `family:"f" qualifier:"live_coin" json:"live_coin"` //????????????
}

//FansAnalysisByT for family t
type FansAnalysisByT struct {
	PlayDuration  map[string]int32 `family:"t" qualifier:"dr" json:"dr"`   //??????????????????
	VidoeAction   map[string]int32 `family:"t" qualifier:"act" json:"act"` //?????????????????????
	DynamicAction map[string]int32 `family:"t" qualifier:"dy" json:"dy"`   //?????????????????????
}

//FansAnalysisByS for family s
type FansAnalysisByS struct { //????????????????????????
	Space   int32 `family:"s" qualifier:"pf1" json:"space"`   //?????????????????????
	Main    int32 `family:"s" qualifier:"pf2" json:"main"`    //???????????????
	Live    int32 `family:"s" qualifier:"pf4" json:"live"`    //??????
	Audio   int32 `family:"s" qualifier:"pf5" json:"audio"`   //	??????
	Article int32 `family:"s" qualifier:"pf6" json:"article"` //	??????
}

// RankInfo str
type RankInfo struct {
	MID      int64  `json:"mid"`
	Uname    string `json:"uname"`
	Photo    string `json:"photo"`
	Relation int    `json:"relation"`
}

// PlaySource for play soucre.
type PlaySource struct {
	PlayProportion map[string]int32 `json:"play_proportion"`
	PageSource     map[string]int32 `json:"page_source"`
}

// ArchivePlay for archive play.
type ArchivePlay struct {
	AID         int64  `json:"aid"`
	View        int32  `json:"view"`
	Rate        int32  `json:"rate"`
	CTime       int32  `json:"ctime"`
	Duration    int64  `json:"duration"`
	AvgDuration int64  `json:"avg_duration"`
	Title       string `json:"title"`
}

//ArchivePlayList for arc play list.
type ArchivePlayList struct {
	ArcPlayList []*ArchivePlay `json:"arc_play_list"`
}

// UpFansMedal for medal fans.
type UpFansMedal struct {
	MedalFans     int32 `family:"f" qualifier:"medal_fans" json:"medal_fans"`           //???????????????
	WearMedalFans int32 `family:"f" qualifier:"wear_medal_fans" json:"wear_medal_fans"` //???????????????
}

// Tip for base survey.
func Tip() (pt *PeriodTip) {
	pt = &PeriodTip{
		ModuleOne:   "???????????????12:00 a.m. ?????????????????????",
		ModuleTwo:   "??????12:00 a.m. ?????????????????????",
		ModuleThree: "?????????12:00 a.m. ?????????????????????",
		ModuleFour:  "???????????????12:00 a.m. ?????????????????????",
	}
	return
}

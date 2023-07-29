package model

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// 官网首页配置信息

type WebsiteNewsInfo struct {

	//Id       uint   `json:"id"`        //自增长ID
	//Type     int8   `json:"type"`      //（1:家长故事 2：孩子的故事 3：扶小鹰成长故事 4：公司新闻 5：行业新闻 6：亲子教育）
	//ImageUrl string `json:"image_url"` //图片url
	//VideoUrl string `json:"video_url"` //视频url
	//Title    string `json:"title"`     //标题
	//Source   string `json:"source"`    //来源
	//Adit     string `json:"adit"`      //责任编辑
	//Desc     string `json:"desc"`      //描述
	//Content  string `json:"content"`   //详情
	List []*JmWebsiteNews
	Pagination
}

//官网首页统计数据信息

type UserCount struct {
	ServiceNum  int64 `json:"service_num"`  //服务人数
	LearnTime   int64 `json:"learn_time"`   //学习时长
	ServiceArea int64 `json:"service_area"` //服务地区
}

// 产品列表信息

type ProductList struct {
	Id          int64  `json:"id"`          //课程id
	Title       string `json:"title"`       //课程名
	ImagePath   string `json:"imagePath"`   //头图
	Descp       string `json:"descp"`       //描述
	Cnt         int    `json:"cnt"`         //章节数
	Subscribers int64  `json:"subscribers"` //订阅人数
}

//会员课返回列表

type MemberList struct {
	Id        int64  `json:"id"`        //课程id
	Title     string `json:"title"`     //课程名
	ImagePath string `json:"imagePath"` //头图
	Descp     string `json:"descp"`     //描述
}

//会员课返回权益列表

type MemberClassList struct {
	Id        int64  `json:"id"`        //课程id
	Title     string `json:"title"`     //课程名
	ImagePath string `json:"imagePath"` //头图
	Descp     string `json:"descp"`     //描述
}

// 会员素养课信息

type CategoryList struct {
	List []Categorys
	Pagination
	//TeacherInfo []TeacherInfo `json:"teacherInfo"`
}

// 会员素养课列表
type Categorys struct {
	Id          int64   `json:"id"`          //课程id
	Title       string  `json:"title"`       //课程名
	ImagePath   string  `json:"imagePath"`   //头图
	Descp       string  `json:"descp"`       //描述
	CostPrice   float64 `json:"costPrice"`   //成本价
	SellPrice   float64 `json:"sellPrice"`   //销售价
	Subscribers int64   `json:"subscribers"` //订阅人数
	Name        string  `json:"name"`        //姓名
	HeadUrl     string  `json:"headUrl"`     //照片
}

// 会员素养课老师信息

type TeacherInfo struct {
	Name    string `json:"name"`    //姓名
	HeadUrl string `json:"headUrl"` //照片
}

// 课程信息

type CourseInfo struct {
	Id        int64   `json:"id"`        //课程id
	Title     string  `json:"title"`     //课程名
	ImagePath string  `json:"imagePath"` //头图
	Descp     string  `json:"descp"`     //描述
	CostPrice float64 `json:"costPrice"` //成本价
	SellPrice float64 `json:"sellPrice"` //销售价
	Name      string  `json:"name"`      //姓名
	HeadUrl   string  `json:"headUrl"`   //照片
}

// 分类栏
type CategoryInfo struct {
	Id   int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"` //产品分类名称

}

// 智慧父母成长课程

type ParentInfo struct {
	Id        int64   `json:"id"`        //课程id
	Title     string  `json:"title"`     //课程名
	ImagePath string  `json:"imagePath"` //头图
	Descp     string  `json:"descp"`     //描述
	CostPrice float64 `json:"costPrice"` //成本价
	SellPrice float64 `json:"sellPrice"` //销售价
	Name      string  `json:"name"`      //姓名
	HeadUrl   string  `json:"headUrl"`   //照片
}

// 获取训练营课程

type TrainCampList struct {
	Id        int64   `json:"id"`        //课程id
	Title     string  `json:"title"`     //课程名
	ImagePath string  `json:"imagePath"` //头图
	Descp     string  `json:"descp"`     //描述
	CostPrice float64 `json:"costPrice"` //成本价
	SellPrice float64 `json:"sellPrice"` //销售价
	Name      string  `json:"name"`      //姓名
	HeadUrl   string  `json:"headUrl"`   //照片
}

// 扶小鹰2信息
type SiteIndex struct {
	Id uint `json:"id"` //自增长ID
	//Type     int8   `json:"type"`      //（1:家长故事 2：孩子的故事 3：扶小鹰成长故事 4：公司新闻 5：行业新闻 6：亲子教育）
	ImageUrl string `json:"image_url"` //图片url
	VideoUrl string `json:"video_url"` //视频url
	Age      int    `json:"age"`       //年龄
	Name     string `json:"name"`      //姓名
	//Title    string `json:"title"`     //标题
	//Source   string `json:"source"`    //来源
	//Adit     string `json:"adit"`      //责任编辑
	Desc string `json:"desc"` //描述
	//Content  string `json:"content"`   //详情
}

//搜索查询

type SearchInfo struct {
	Id    uint   `json:"id"` //自增长ID
	Title string `json:"title"`
}

// 课程详情
type CourseDetailList struct {
	Id      uint   `json:"id"`      //自增长ID
	Content string `json:"content"` //课程url
	Type    int    `json:"type"`    //当本条数据对应的type≠1时，不展示,当type=1时展示。
}

//课程目录

type CourseCatalogueList struct {
	Id                uint   `json:"id"`                //章节id
	Title             string `json:"title"`             //标题
	SourceTotalSecond string `json:"sourceTotalSecond"` //时长
}

// 留言
type MessageList struct {
	Id         uint   `json:"id"`         //评论id
	Content    string `json:"content"`    //评论内容
	CreateTime int    `json:"createTime"` //评论时间
	PhotoUrl   string `json:"photoUrl"`   // "头像"
	Nickname   string `json:"nickname"`   //昵称
}

//二维码

type CodeList struct {
	Id      uint   `json:"id"`      //id
	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容

}

// 直播信息
type LiveInfo struct {
	Id         uint   `json:"id"`         //id
	Title      string `json:"title"`      //标题
	Info       string `json:"info"`       //简介
	PicPath    string `json:"picPath"`    //图片
	StartTime  int    `json:"startTime"`  //开始时间
	LiveStatus int    `json:"liveStatus"` //直播状态 直播状态，未开始=1，直播中=2，暂停中=3，已结束=4
	LivePerson int    `json:"livePerson"` //预约人数
}

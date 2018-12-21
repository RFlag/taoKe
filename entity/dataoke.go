package entity

import(
	"time"
)

type DTKData1 struct {
	Data DataInfo1
}

type DTKData2 struct {
	Data   DataInfo2
	Result []DTKResult2
}

type DTKData3 struct {
	Data   DataInfo2
	Result []DTKResult1
}

type DTKData4 struct {
	Data   DataInfo2
	Result DTKResult1
}

type DataInfo1 struct {
	ApiType    string     `json:"api_type"`
	UpdateTime string     `json:"update_time"`
	TotalNum   int        `json:"total_num"`
	ApiContent string     `json:'api_content'`
	Result     *[]DTKResult2 `json:"result"`
}

type DataInfo2 struct {
	ApiType    string `json:"api_type"`
	UpdateTime string `json:"update_time"`
	TotalNum   int    `json:"total_num"`
	ApiContent string `json:'api_content'`
}

type DTKResult struct {
	Id                int       `json:"ID" db:"id"`
	CreateAt          time.Time `json:"create_at" db:"create_at"`
	GoodsId           string    `json:"GoodsId" db:"goods_id"`
	Title             string    `json:"Title" db:"title"`
	DTitle            string    `json:"D_title" db:"d_title"`
	Pic               string    `json:"Pic" db:"pic"`
	Cid               int       `json:"Cid" db:"cid"`
	OrgPrice          float64   `json:"Org_Price" db:"org_price"`
	Price             float64   `json:"Price" db:"price"`
	IsTmall           int       `json:"IsTmall" db:"is_tmall"`
	SalesNum          int       `json:"Sales_num" db:"sales_num"`
	Dsr               float64   `json:"Dsr" db:"dsr"`
	SellerId          string    `json:"SellerID" db:"seller_id"`
	Commission        float64   `json:"Commission" db:"commission"` // 单品详情
	CommissionJihua   float64   `json:"Commission_jihua" db:"commission_jihua"`
	CommissionQueqiao float64   `json:"Commission_queqiao" db:"commission_queqiao"`
	JihuaLink         string    `json:"Jihua_link" db:"jihua_link"`
	JihuaShenhe       float64   `json:"Jihua_shenhe" db:"jihua_shenhe"`
	Introduce         string    `json:"Introduce" db:"introduce"`
	QuanId            string    `json:"Quan_id" db:"quan_id"`
	QuanPrice         float64   `json:"Quan_price" db:"quan_price"`
	QuanTime          string    `json:"Quan_time" db:"quan_time"`
	QuanSurplus       int       `json:"Quan_surplus" db:"quan_surplus"`
	QuanReceive       float64   `json:"Quan_receive" db:"quan_receive"`
	QuanCondition     string    `json:"Quan_condition" db:"quan_condition"`
	QuanMLink         *string   `json:"Quan_m_link" db:"quan_m_link"`
	QuanLink          string    `json:"Quan_link" db:"quan_link"`
	YongjinType       int       `json:"Yongjin_type" db:"yongjin_type"`
	QueSiteid         string    `json:"Que_siteid" db:"que_siteid"`
}

type DTKResult1 struct {
	Id                string  `json:"ID"`
	GoodsId           string  `json:"GoodsId"`
	Title             string  `json:"Title"`
	DTitle            string  `json:"D_title"`
	Pic               string  `json:"Pic"`
	Cid               string  `json:"Cid"`
	OrgPrice          string  `json:"Org_Price"`
	Price             float64 `json:"Price"`
	IsTmall           string  `json:"IsTmall"`
	SalesNum          string  `json:"Sales_num"`
	Dsr               string  `json:"Dsr"`
	SellerId          string  `json:"SellerID"`
	Commission        string  `json:"Commission"` // 单品详情
	CommissionJihua   string  `json:"Commission_jihua"`
	CommissionQueqiao string  `json:"Commission_queqiao"`
	JihuaLink         string  `json:"Jihua_link"`
	JihuaShenhe       string  `json:"Jihua_shenhe"`
	Introduce         string  `json:"Introduce"`
	QuanId            string  `json:"Quan_id"`
	QuanPrice         string  `json:"Quan_price"`
	QuanTime          string  `json:"Quan_time"`
	QuanSurplus       string  `json:"Quan_surplus"`
	QuanReceive       string  `json:"Quan_receive"`
	QuanCondition     string  `json:"Quan_condition"`
	QuanMLink         *string `json:"Quan_m_link"`
	QuanLink          string  `json:"Quan_link"`
	YongjinType       string  `json:"Yongjin_type"`
	QueSiteid         string  `json:"Que_siteid"`
}

type DTKResult2 struct {
	Id                int     `json:"ID"`
	GoodsId           string  `json:"GoodsId"`
	Title             string  `json:"Title"`
	DTitle            string  `json:"D_title"`
	Pic               string  `json:"Pic"`
	Cid               int     `json:"Cid"`
	OrgPrice          float64 `json:"Org_Price"`
	Price             float64 `json:"Price"`
	IsTmall           int     `json:"IsTmall"`
	SalesNum          int     `json:"Sales_num"`
	Dsr               float64 `json:"Dsr"`
	SellerId          string  `json:"SellerID"`
	Commission        float64 `json:"Commission"` // 单品详情
	CommissionJihua   float64 `json:"Commission_jihua"`
	CommissionQueqiao float64 `json:"Commission_queqiao"`
	JihuaLink         string  `json:"Jihua_link"`
	JihuaShenhe       float64 `json:"Jihua_shenhe"`
	Introduce         string  `json:"Introduce"`
	QuanId            string  `json:"Quan_id"`
	QuanPrice         float64 `json:"Quan_price"`
	QuanTime          string  `json:"Quan_time"`
	QuanSurplus       int     `json:"Quan_surplus"`
	QuanReceive       float64 `json:"Quan_receive"`
	QuanCondition     string  `json:"Quan_condition"`
	QuanMLink         *string `json:"Quan_m_link"`
	QuanLink          string  `json:"Quan_link"`
	YongjinType       int     `json:"Yongjin_type"`
	QueSiteid         string  `json:"Que_siteid"`
}

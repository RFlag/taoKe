package entity

type SHZData struct {
	Return     int
	CateId     *string
	Pagesize   int
	Rowscount  int
	Pagecount  int
	Pagination int
	Result     []SHZResult `json:"result"`
}

type SHZResult struct {
	Gid             string  `db:"gid" json:"gid"`//商品淘宝id
	Cate            string  `db:"cate" json:"cate"`//分类名称
	Site            string  `db:"site" json:"site"`//所属站点：tmall 天猫，taobao 淘宝
	Title           string  `db:"title" json:"title"`//商品标题
	Price           string  `db:"price" json:"price"`//券后价
	Ratio           string  `db:"ratio" json:"ratio"`//佣金比例
	Prime           string  `db:"prime" json:"prime"`//正常售价
	Thumb           string  `db:"thumb" json:"thumb"`//商品主图
	Url             string  `db:"url" json:"url"`//购买链接
	SubTitle        string  `db:"sub_title" json:"sub_title"` //商品简称
	LongPic         string  `db:"long_pic" json:"long_pic"`//商品长图
	Activity        string  `db:"activity" json:"activity"`//活动 (无, 聚划算, 淘抢购, 双十一)
	ActivityStime   string  `db:"activity_stime" json:"activity_stime"`
	ActivityEtime   string  `db:"activity_etime" json:"activity_etime"`
	IsBrand         string  `db:"is_brand" json:"is_brand"`//品牌状态（0未申请品牌，1待审，2已审）
	BrandName       string  `db:"brand_name" json:"brand_name"`//品牌名称
	Freight         string  `db:"freight" json:"freight"`//运费险（0：否，1：是）
	RatioType       string  `db:"ratio_type" json:"ratio_type"`//通用, 定向, 鹊桥
	Video           string  `db:"video" json:"video"`//视频地址
	VideoId         string  `db:"video_id" json:"video_id"`//视频id
	SellerId        string	`db:"seller_id" json:"seller_id"`
	PlanHigh        *string `db:"plan_high" json:"plan_high"`//高佣计划地址
	FinalSales      string  `db:"final_sales" json:"final_sales"`//最终销量
	Coupon          string  `db:"coupon" json:"coupon"`//是否有优惠券（0无，1有）
	CouponMoney     string  `db:"coupon_money" json:"coupon_money"`//优惠券金额
	CouponUrl       string  `db:"coupon_url" json:"coupon_url"` //优惠券链接
	CouponTotal     string  `db:"coupon_total" json:"coupon_total"`//优惠券总数
	CouponLatest    string  `db:"coupon_latest" json:"coupon_latest"`//当前已领优惠券数
	CouponExpire    string  `db:"coupon_expire" json:"coupon_expire"`//券有效期
	CouponCondition string  `db:"coupon_condition" json:"coupon_condition"`//券使用条件，满20可用
	IntroFoot       string  `db:"intro_foot" json:"intro_foot"`//推荐理由
	Timeline        string  `db:"timeline" json:"timeline"`//更新时间
	Stoptime        string  `db:"stoptime" json:"stoptime"`//结束时间
	Id              string  `db:"id" json:"id"`//分类ID
	ActivityId      string  `db:"activity_id" json:"activity_id"`//活动ID
	NewUrl          string  `db:"new_url" json:"new_url"`//二合一链接
	CouponStartTime string  `db:"coupon_start_time" json:"coupon_start_time"`//优惠券开始时间
}

type Sort struct {
	Return int             `json:"return"`
	Result []SHZSortResult `json:"result"`
}

type SHZSortResult struct {
	Id   string `json:"id"`
	Cate string `json:"cate"`
	Stat string `json:"stat`
}

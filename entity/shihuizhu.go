package entity

type Data struct{
	Return int
	CateId string
	Pagesize int
	Rowscount int
	Pagecount int
	Pagination int
	Result []Result5
}

type Result5 struct{
	Gid  string  //商品淘宝id
	Cate  string              //分类名称
	Site   string             //所属站点：tmall 天猫，taobao 淘宝
	Title  string  //商品标题
	Price  string          //券后价
	Ratio  string           //佣金比例
	Prime   string          //正常售价
	Thumb   string //商品主图
	Url string //购买链接
	SubTitle string  //商品简称
	LongPic  string	 //商品长图
	Activity string	 //活动 (无, 聚划算, 淘抢购, 双十一)
	IsBrand   string	 //品牌状态（0未申请品牌，1待审，2已审）
	BrandName string	 //品牌名称
	Freight  string //运费险（0：否，1：是）
	RatioType  string	 //通用, 定向, 鹊桥
	Video  string	 //视频地址
	PlanHigh *string      //高佣计划地址
	FinalSales string       //最终销量
	Coupon string      //是否有优惠券（0无，1有）
	CouponMoney string  //优惠券金额
	CouponUrl string //shop.m.taobao.com/shop/coupon.htm?seller_id=2097118490&activity_id=aff031b5be6e47dbac80f5be54defe5f",  //优惠券链接
	CouponTotal string //优惠券总数
	CouponLatest string   //当前已领优惠券数
	CouponExpire string    //券有效期
	CouponCondition string   //券使用条件，满20可用
	IntroFoot string //推荐理由
	Timeline string    //更新时间
	Stoptime  string     //结束时间
	Id  string   //分类ID
	NewUrl string   //二合一链接
	CouponStartTime string   //优惠券开始时间
}


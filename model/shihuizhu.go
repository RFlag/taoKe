package model

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"net/http"
	"project/dingdangke-dataoke/conf"
	"project/dingdangke-dataoke/entity"
	"project/ftgo/ftsql"
	"project/ftgo/safeclose"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// 商品分类
func GoodsSort() ([]entity.SHZSortResult,error) {
	err:=shzEmptyData("shihuizhu_goods_sort")
	if err != nil {
		return nil,err
	}
	
	data := new(entity.Sort)
	req, err := http.NewRequest("GET", "http://gateway.shihuizhu.net/open/cates", nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("APPID", conf.SHZAppId)
	header.Add("APPKEY", conf.SHZAppKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		return nil, err
	}

	sqlParam :=[]interface{}{}
	for _,vv:=range data.Result{
		sqlParam=append(sqlParam,vv.Id,vv.Cate,vv.Stat)
	}

	sqlStr:="(?,?,?)"
	for i:=0;i<len(data.Result)-1;i++{
		sqlStr+=",(?,?,?)"
	}
	result,err:=ftsql.DB.Exec(`INSERT INTO shihuizhu_goods_sort (id,cate,stat) values `+sqlStr,sqlParam...)
	if err != nil{
		return nil,err
	}
	if rows, rowErr := result.RowsAffected(); int(rows) != len(data.Result) || rowErr != nil {
		return nil,errors.New("insert order Error ")
	}
	return data.Result, nil
}


// 商品详情
func GoodsInfo()([]entity.SHZResult,error) {
	return nil,nil
}


// 商品列表
func GoodsList()(*entity.SHZData,error) {
	err:=shzEmptyData("shihuizhu_goods")
	if err != nil {
		return nil,err
	}

	data := new(entity.SHZData)
	req, err := http.NewRequest("GET", "http://gateway.shihuizhu.net/open/goods/0", nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("APPID", conf.SHZAppId)
	header.Add("APPKEY", conf.SHZAppKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		return nil, err
	}
	const n = 3
	pagenoBatch := [n][]int{}
	for i:=1;i<=data.Pagecount;i++{
		pagenoBatch[i%n]=append(pagenoBatch[i%n],i)
	}

	var wg sync.WaitGroup

	for _,pageno:= range pagenoBatch{
		func(pageno []int){
			wg.Add(1)
			safeclose.DoContext(func(ctx context.Context){
				defer wg.Done()
				for _, page := range pageno {
					select {
					default:
					case <-ctx.Done():
						return
					}
					data := new(entity.SHZData)
					req, err := http.NewRequest("GET", "http://gateway.shihuizhu.net/open/goods/0/"+strconv.Itoa(page), nil)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}
					header := req.Header
					header.Add("Accept", "application/json")
					header.Add("Content-Type", "application/json;charset=utf-8")
					header.Add("APPID", conf.SHZAppId)
					header.Add("APPKEY", conf.SHZAppKey)
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}
					bodyByte, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}
					err = json.Unmarshal(bodyByte, &data)
					if err != nil {
						print("@##@",page,string(bodyByte))
						log.Println("调用领券接口失败", err)
						continue
					}

					func(result []entity.SHZResult){
						safeclose.Do(func() {
							sqlParam := []interface{}{}
							for _, row := range result {
								if row.PlanHigh == nil {
									a := ""
									row.PlanHigh = &a
								}
								sqlParam = append(sqlParam, row.Gid, row.Cate, row.Site,row.Title,row.Price, row.Ratio,
									row.Prime, row.Thumb, row.SubTitle,  row.IntroFoot,row.LongPic,row.Activity,
									row.IsBrand, row.BrandName, row.Freight, row.RatioType,row.Video,row.VideoId,
									row.SellerId, row.Url,row.PlanHigh,row.FinalSales,row.Coupon, row.CouponMoney,
									row.CouponUrl, row.CouponTotal,row.CouponLatest, row.Timeline, row.Stoptime,
									row.ActivityStime,row.ActivityEtime, row.CouponExpire, row.CouponCondition,
									row.CouponStartTime,row.Id,row.ActivityId,row.NewUrl)
							}
							if len(data.Result) > 0 {
								tx := ftsql.DB.MustBegin()
								err := shzAddCommodity(tx, "shihuizhu_goods", len(data.Result), sqlParam)
								if err != nil {
									if !strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
										log.Print("插入数据库失败!", err)
									}
									tx.Rollback()
									return
								}
								tx.Commit()
							}
						})
					}(data.Result)

				}
			})
		}(pageno)
	}
	wg.Wait()

	return data, nil
}

// 今日疯抢榜
func CrazyToday()([]entity.SHZResult,error) {
	err:=shzEmptyData("shihuizhu_crazy_today")
	if err != nil {
		return nil,err
	}

	data := new(entity.SHZData)
	req, err := http.NewRequest("GET", "http://gateway.shihuizhu.net/open/today/0", nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("APPID", conf.SHZAppId)
	header.Add("APPKEY", conf.SHZAppKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		return nil, err
	}
	const n = 3
	pagenoBatch := [n][]int{}
	for i:=1;i<=data.Pagecount;i++{
		pagenoBatch[i%n]=append(pagenoBatch[i%n],i)
	}

	var wg sync.WaitGroup

	for _,pageno:= range pagenoBatch{
		func(pageno []int){
			wg.Add(1)
			safeclose.DoContext(func(ctx context.Context){
				defer wg.Done()
				for _, page := range pageno {
					select {
					default:
					case <-ctx.Done():
						return
					}
					data := new(entity.SHZData)
					req, err := http.NewRequest("GET", "http://gateway.shihuizhu.net/open/today/0/"+strconv.Itoa(page), nil)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}
					header := req.Header
					header.Add("Accept", "application/json")
					header.Add("Content-Type", "application/json;charset=utf-8")
					header.Add("APPID", conf.SHZAppId)
					header.Add("APPKEY", conf.SHZAppKey)
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}
					bodyByte, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}
					err = json.Unmarshal(bodyByte, &data)
					if err != nil {
						log.Println("调用领券接口失败", err)
						continue
					}

					func(result []entity.SHZResult){
						safeclose.Do(func() {
							sqlParam := []interface{}{}
							for _, row := range result {
								if row.PlanHigh == nil {
									a := ""
									row.PlanHigh = &a
								}
								sqlParam = append(sqlParam, row.Gid, row.Cate, row.Site,row.Title,row.Price, row.Ratio,
									row.Prime, row.Thumb, row.SubTitle,  row.IntroFoot,row.LongPic,row.Activity,
									row.IsBrand, row.BrandName, row.Freight, row.RatioType,row.Video,row.VideoId,
									row.SellerId, row.Url,row.PlanHigh,row.FinalSales,row.Coupon, row.CouponMoney,
									row.CouponUrl, row.CouponTotal,row.CouponLatest, row.Timeline, row.Stoptime,
									row.ActivityStime,row.ActivityEtime, row.CouponExpire, row.CouponCondition,
									row.CouponStartTime,row.Id,row.ActivityId,row.NewUrl)
							}
							if len(data.Result) > 0 {
								tx := ftsql.DB.MustBegin()
								err := shzAddCommodity(tx, "shihuizhu_crazy_today", len(data.Result), sqlParam)
								if err != nil {
									if !strings.HasPrefix(err.Error(), "Error 1062: Duplicate entry") {
										log.Print("插入数据库失败", err)
									}
									tx.Rollback()
									return
								}
								tx.Commit()
							}
						})
					}(data.Result)

				}
			})
		}(pageno)
	}
	wg.Wait()

	return data.Result, nil
}


func shzAddCommodity(tx sqlx.Ext, sqlTable string, sub int, sqlParam []interface{}) error {

	sqlStr := "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	for i := 0; i < sub-1; i++ {
		sqlStr += ",(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	}

	result, err := tx.Exec(`
		INSERT INTO `+sqlTable+`
		(
			gid,
			cate,
			site,
			title,
			price,
			ratio,
			prime,
			thumb,
			sub_title,
			intro_foot,
			long_pic,
			activity,
			is_brand,
			brand_name,
			freight,
			ratio_type,
			video,
			Video_id,
			seller_id,
			url,
			plan_high,
			final_sales,
			coupon,
			coupon_money,
			coupon_url,
			coupon_total,
			coupon_latest,
			timeline,
			stoptime,
			activity_stime,
			activity_etime,
			coupon_expire,
			coupon_condition,
			coupon_start_time,
			id,
			activity_id,
			new_url
			) VALUES `+sqlStr, sqlParam...)
	if err != nil {
		return err
	}
	if rows, rowErr := result.RowsAffected(); int(rows) != sub || rowErr != nil {
		return errors.New("insert order Error ")
	}

	return nil
}

// 清空
func shzEmptyData(sqlTable string) error  {
	result, err := ftsql.DB.Exec(`DELETE from `+sqlTable)
	if err != nil {
		return err
	}
	if rows, rowsErr := result.RowsAffected(); rows == 0 || rowsErr != nil {
		return errors.WithMessage(rowsErr,"delete "+sqlTable+" error")
	}
	return nil
}



// 视频专区
func VideoZone(cateId,page,psize int) ([]entity.SHZResult,int,error){
	result:=[]entity.SHZResult{}
	total:=0

	err:=ftsql.DB.QueryRow(`select count(1) from shihuizhu_goods where id=? and video!=""`,cateId).Scan(&total)
	if err!=nil {
		return nil,0,err
	}

	err=ftsql.DB.Select(&result,`select * from shihuizhu_goods where id=? and video!="" limit ?,?`,cateId,(page-1)*psize,psize)
	if err !=nil{
		return nil,0,err
	}
	return result,total,nil
}

// 排行榜
func Leaderboard(cateId,page,psize int) ([]entity.SHZResult,int,error){
	result:=[]entity.SHZResult{}
	total:=0

	err:=ftsql.DB.QueryRow(`select count(1) from shihuizhu_crazy_today where id=?`,cateId).Scan(&total)
	if err!=nil {
		return nil,0,err
	}

	err=ftsql.DB.Select(&result,`select * from shihuizhu_crazy_today where id=? limit ?,?`,cateId,(page-1)*psize,psize)
	if err !=nil{
		return nil,0,err
	}
	return result,total,nil
}

// 头条
func Headline() {

}

// 咚咚抢
func Grab() {

}

// 半价
func HalfPrice(end time.Time,page,psize int)([]entity.SHZResult,int,error) {
	result:=[]entity.SHZResult{}
	total:=0
	en:=end.Unix()
	err:=ftsql.DB.QueryRow(`select count(1) from shihuizhu_goods where price<=prime/2 and stoptime>=?`,en).Scan(&total)
	if err!=nil {
		return nil,0,err
	}

	err=ftsql.DB.Select(&result,`select * from shihuizhu_goods where price<=prime/2 and stoptime>=? order by stoptime limit ?,?`,en,(page-1)*psize,psize)
	if err !=nil{
		return nil,0,err
	}
	return result,total,nil
}

// 9.9
func NinePNine(cateId,page,psize int)([]entity.SHZResult,int,error) {
	result:=[]entity.SHZResult{}
	total:=0

	err:=ftsql.DB.QueryRow(`select count(1) from shihuizhu_goods where id=? and price<=20`,cateId).Scan(&total)
	if err!=nil {
		return nil,0,err
	}

	err=ftsql.DB.Select(&result,`select * from shihuizhu_goods where id=? and price<=20 limit ?,?`,cateId,(page-1)*psize,psize)
	if err !=nil{
		return nil,0,err
	}
	return result,total,nil
}

// 为您推荐
func Recommend(page,psize int)([]entity.SHZResult,int,error) {
	result:=[]entity.SHZResult{}
	total:=0

	err:=ftsql.DB.QueryRow(`select count(1) from shihuizhu_goods `).Scan(&total)
	if err!=nil {
		return nil,0,err
	}

	err=ftsql.DB.Select(&result,`select * from shihuizhu_goods order by final_sales`,(page-1)*psize,psize)
	if err !=nil{
		return nil,0,err
	}
	return result,total,nil
}
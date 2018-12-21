package model

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"log"
	"strings"
	"math"

	"project/dingdangke-dataoke/conf"
	"project/dingdangke-dataoke/entity"

	"project/ftgo/ftsql"
	"project/ftgo/safeclose"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func WebsiteSpecial() (*entity.DTKData1, error) {
	data := new(entity.DTKData1)
	req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=goodsLink/www&type=www_quan&appkey="+conf.Appkey+"&v=2", nil)
	if err != nil {
		return nil, err
	}
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
	return data, nil
}

func QqSpecial() (*entity.DTKData1, error) {
	data := new(entity.DTKData1)
	req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=goodsLink/qq&type=qq_quan&appkey="+conf.Appkey+"&v=2", nil)
	if err != nil {
		return nil, err
	}
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
	return data, nil
}

func DrawCoupon() (*entity.DTKData2, error) {
	err:=emptyData("dataoke_lingquan")
	if err != nil {
		return nil,err
	}

	data := new(entity.DTKData2)
	// resultId:=[]int{}
	req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=Port/index&type=total&appkey="+conf.Appkey+"&v=2", nil)
	if err != nil {
		return nil, errors.WithMessage(err, "调用淘客领券接口失败")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithMessage(err, "调用淘客领券接口失败")
	}
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithMessage(err, "调用淘客领券接口失败")
	}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		return nil, errors.WithMessage(err, "调用淘客领券接口失败")
	}

	total := math.Ceil(float64(data.Data.TotalNum)/50)
	const n = 3
	pagenoBatch := [n][]int{}
	for i := 1; i <= int(total); i++ {
		pagenoBatch[i%n] = append(pagenoBatch[i%n], i)
	}

	var wg sync.WaitGroup

	for _, pageno := range pagenoBatch {
		wg.Add(1)
		func(pageno []int) {
			safeclose.DoContext(func(ctx context.Context) {
				defer wg.Done()
				for _, page := range pageno {
					select {
					default:
					case <-ctx.Done():
						return
					}
					data := new(entity.DTKData2)
					req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=Port/index&type=total&appkey="+conf.Appkey+"&v=2&page="+strconv.Itoa(page), nil)
					if err != nil {
						log.Println("调用领券接口失败",err)
						continue
					}
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						log.Println("调用领券接口失败",err)
						continue
					}
					bodyByte, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Println("调用领券接口失败",err)
						continue
					}
					err = json.Unmarshal(bodyByte, &data)
					if err != nil {
						log.Println("调用领券接口反序列化失败","data:",page,string(bodyByte),err)
						continue
					}

					func(result []entity.DTKResult2) {
						safeclose.Do(func() {
							sqlParam := []interface{}{}
							for _, row := range result {
								if row.QuanMLink == nil {
									a := ""
									row.QuanMLink = &a
								}
								sqlParam = append(sqlParam, row.Id, row.GoodsId, row.Title, row.DTitle,
									row.Pic, row.Cid, row.OrgPrice, row.Price, row.IsTmall,
									row.SalesNum, row.Dsr, row.SellerId, row.Commission,
									row.CommissionJihua, row.CommissionQueqiao, row.JihuaLink,
									row.JihuaShenhe, row.Introduce, row.QuanId, row.QuanPrice,
									row.QuanTime, row.QuanSurplus, row.QuanReceive, row.QuanCondition,
									row.QuanMLink, row.QuanLink, row.YongjinType, row.QueSiteid)
							}
							if len(data.Result) > 0 {
								tx := ftsql.DB.MustBegin()
								err := addCommodity(tx, "dataoke_lingquan", len(data.Result), sqlParam)
								if err != nil {
									if !strings.HasPrefix(err.Error(),"Error 1062: Duplicate entry"){
										log.Print("插入数据库失败",err)
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

func DrawCouponList(page,psize int) ([]entity.DTKResult,int,error)  {
	result:=[]entity.DTKResult{}
	total := 0
	err := ftsql.DB.QueryRow(`select count(1) from dataoke_lingquan`,).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	err =ftsql.DB.Select(&result,`select * from dataoke_lingquan limit ?,?`,(page-1)*psize,psize)
	if err != nil {
		return nil,0,err
	}
	return result,total,nil
}

func Top100Popularity() (*entity.DTKData3, error) {
	err:=emptyData("dataoke_top")
	if err != nil {
		return nil,err
	}

	data := new(entity.DTKData3)
	req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=Port/index&type=top100&appkey="+conf.Appkey+"&v=2", nil)
	if err != nil {
		return nil, err
	}
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

	sqlParam := []interface{}{}
	for _, row := range data.Result {
		if row.QuanMLink == nil {
			a := ""
			row.QuanMLink = &a
		}
		sqlParam = append(sqlParam, row.Id, row.GoodsId, row.Title, row.DTitle,
			row.Pic, row.Cid, row.OrgPrice, row.Price, row.IsTmall,
			row.SalesNum, row.Dsr, row.SellerId, row.Commission,
			row.CommissionJihua, row.CommissionQueqiao, row.JihuaLink,
			row.JihuaShenhe, row.Introduce, row.QuanId, row.QuanPrice,
			row.QuanTime, row.QuanSurplus, row.QuanReceive, row.QuanCondition,
			row.QuanMLink, row.QuanLink, row.YongjinType, row.QueSiteid)
	}
	if len(data.Result) > 0 {
		tx := ftsql.DB.MustBegin()
		err := addCommodity(tx, "dataoke_top", len(data.Result), sqlParam)
		if err != nil {
			if !strings.HasPrefix(err.Error(),"Error 1062: Duplicate entry"){
				log.Print("插入数据库失败",err)
			}
			tx.Rollback()
			return data,nil
		}
		tx.Commit()
	}

	return data, nil
}

func Top100PopularityList(page,psize int)([]entity.DTKResult,int,error)  {
	result:=[]entity.DTKResult{}
	total := 0
	err := ftsql.DB.QueryRow(`select count(1) from dataoke_top`,).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	err =ftsql.DB.Select(&result,`select * from dataoke_top limit ?,?`,(page-1)*psize,psize)
	if err != nil {
		return nil,0,err
	}
	return result,total,nil
}

func RealtimeAmount() (*entity.DTKData3, error) {
	err:=emptyData("dataoke_paoliang")
	if err != nil {
		return nil,err
	}

	data := new(entity.DTKData3)
	req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=Port/index&type=paoliang&appkey="+conf.Appkey+"&v=2", nil)
	if err != nil {
		return nil, err
	}
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

	sqlParam := []interface{}{}
	for _, row := range data.Result {
		if row.QuanMLink == nil {
			a := ""
			row.QuanMLink = &a
		}
		sqlParam = append(sqlParam, row.Id, row.GoodsId, row.Title, row.DTitle,
			row.Pic, row.Cid, row.OrgPrice, row.Price, row.IsTmall,
			row.SalesNum, row.Dsr, row.SellerId, row.Commission,
			row.CommissionJihua, row.CommissionQueqiao, row.JihuaLink,
			row.JihuaShenhe, row.Introduce, row.QuanId, row.QuanPrice,
			row.QuanTime, row.QuanSurplus, row.QuanReceive, row.QuanCondition,
			row.QuanMLink, row.QuanLink, row.YongjinType, row.QueSiteid)
	}
	if len(data.Result) > 0 {
		tx := ftsql.DB.MustBegin()
		err := addCommodity(tx, "dataoke_paoliang", len(data.Result), sqlParam)
		if err != nil {
			if !strings.HasPrefix(err.Error(),"Error 1062: Duplicate entry"){
				log.Print("插入数据库失败",err)
			}
			tx.Rollback()
			return data,nil
		}
		tx.Commit()
	}
	return data, nil
}

func RealtimeAmountList(page,psize int)([]entity.DTKResult,int,error)  {
	result:=[]entity.DTKResult{}
	total := 0
	err := ftsql.DB.QueryRow(`select count(1) from dataoke_paoliang`,).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	err =ftsql.DB.Select(&result,`select * from dataoke_paoliang limit ?,?`,(page-1)*psize,psize)
	if err != nil {
		return nil,0,err
	}
	return result,total,nil
}

func Goods(id int) (*entity.DTKData4, error) {
	data := new(entity.DTKData4)
	req, err := http.NewRequest("GET", "http://api.dataoke.com/index.php?r=port/index&appkey="+conf.Appkey+"&v=2&id="+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}
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
	return data, nil
}

func addCommodity(tx sqlx.Ext, sqlTable string, sub int, sqlParam []interface{}) error {

	sqlStr := "(?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	for i := 0; i < sub-1; i++ {
		sqlStr += ",(?,now(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	}

	result, err := tx.Exec(`
		INSERT INTO `+sqlTable+`
		(
			id,
			create_at,
			goods_id,
			title,
			d_title,
			pic,
			cid,
			org_price,
			price,
			is_tmall,
			sales_num,
			dsr,
			seller_id,
			commission,
			commission_jihua,
			commission_queqiao,
			jihua_link,
			jihua_shenhe,
			introduce,
			quan_id,
			quan_price,
			quan_time,
			quan_receive,
			quan_surplus,
			quan_condition,
			quan_m_link,
			quan_link,
			yongjin_type,
			que_siteid
			) VALUES `+sqlStr, sqlParam...)
	if err != nil {
		return err
	}
	if rows, rowErr := result.RowsAffected(); int(rows) != sub || rowErr != nil {
		return errors.New("insert order Error ")
	}

	return nil
}


func emptyData(sqlTable string) error  {
	result, err := ftsql.DB.Exec(`DELETE from `+sqlTable)
   if err != nil {
	   return err
   }
   if rows, rowsErr := result.RowsAffected(); rows == 0 || rowsErr != nil {
	   return errors.WithMessage(rowsErr,"delete "+sqlTable+" error")
   }
   return nil
}

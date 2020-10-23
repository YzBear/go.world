package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/robfig/cron/v3"
	"strconv"
	"strings"
	"time"
)
var stopCh chan interface{}
func main() {
	stopCh=make(chan interface{})
	//xiuluoyunTask()
	//测试获取指定商品
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		xiuluoyunTargetTask("https://www.xiuluohost.com/product/zj-nat","NAT-JS-XZ-A1")
	})
	c.Start()
	defer c.Stop()
	select {
		case<-stopCh:
			fmt.Println("抢单成功，定时任务已停止")
	}
}

type Goods struct {
	title string
	price string
	status string
	href string
}
func (goods *Goods) toString()  {
	fmt.Printf("%s----价格为：%s----商品状态：%s----购买链接：%s\n",goods.title,goods.price,goods.status,goods.href)
}
//这是爬所有的商品
func xiuluoyunTask()  {
	const url = "https://www.xiuluohost.com"
	const userAgent="Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36"
	//开启一个爬虫
	c:=colly.NewCollector()
	c.UserAgent=userAgent
	//下面是详细的展示
	taskC:=c.Clone()
	taskC.Async=true
	taskC.Limit(&colly.LimitRule{
		DomainRegexp: "",
		DomainGlob: "(*.xiuhuohost.com",
		Delay: 10*time.Second,
		RandomDelay: 0,
		Parallelism: 1,
	})
	//然后第一次访问主页，获取到产品的信息
	c.OnHTML("ul[class='dropdown-menu']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, element *colly.HTMLElement) {
			goodsType:=element.ChildText("a")
			href:=element.ChildAttr("a","href")
			ctx := colly.NewContext()
			ctx.Put("goodsType",goodsType)
			ctx.Put("href",href)
			//通过context上下文将c采集到的数据传递给采集器2
			taskC.Request("GET",url+href,nil,ctx,nil)
		})
	})
	//子采集器开始采集页面数据
	taskC.OnHTML("div[class='pricing-tables-wrap']", func(e *colly.HTMLElement) {
		goodsType:=e.Request.Ctx.Get("goodsType")
		//href:=e.Request.Ctx.Get("href")
		fmt.Printf("===========%s============\n",goodsType)
		e.ForEach("div[class='pricing-table']>div[class='pricing-table-inner']", func(i int, element *colly.HTMLElement) {
			//这里就是价格和库存了
			title:=element.ChildText("div[class='pricing-table-main']>div[class='pricing-table-header mb-24 pb-24']>div[class='pricing-table-title h4 mt-0 mb-16']")
			price := element.ChildText("div[class='pricing-table-main']>div[class='pricing-table-header mb-24 pb-24']>div[class='pricing-table-price']>span[class='pricing-table-price-amount h1']")
			buttonText:=element.ChildText("div[class='pricing-table-cta']>a[class='button button-secondary button-shadow button-block']")
			goodsHref:=element.ChildAttr("div[class='pricing-table-cta']>a[class='button button-secondary button-shadow button-block']","href")
			goods:=&Goods{title: title,price: price,status: buttonText,href: url+goodsHref}
			goods.toString()

		})
		fmt.Printf("%s over",goodsType)
	})
	err := c.Visit(url)
	if err != nil {
		fmt.Println("xiuluoyunTask err:",err.Error())
	}
	taskC.Wait()
}
//爬指定商品,但是只能爬父级分类
func xiuluoyunTargetTask(parentUrl string,targetTitle string)  {
	const url = "https://www.xiuluohost.com"
	c:=colly.NewCollector()
	c.OnHTML("div[class='pricing-tables-wrap']", func(e *colly.HTMLElement) {
		e.ForEach("div[class='pricing-table']>div[class='pricing-table-inner']", func(i int, element *colly.HTMLElement) {
			//这里就是价格和库存了
			title:=element.ChildText("div[class='pricing-table-main']>div[class='pricing-table-header mb-24 pb-24']>div[class='pricing-table-title h4 mt-0 mb-16']")
			if title==targetTitle {
				price := element.ChildText("div[class='pricing-table-main']>div[class='pricing-table-header mb-24 pb-24']>div[class='pricing-table-price']>span[class='pricing-table-price-amount h1']")
				buttonText:=element.ChildText("div[class='pricing-table-cta']>a[class='button button-secondary button-shadow button-block']")
				goodsHref:=element.ChildAttr("div[class='pricing-table-cta']>a[class='button button-secondary button-shadow button-block']","href")
				goods:=&Goods{title: title,price: price,status: buttonText,href: url+goodsHref}
				goods.toString()
				goodsId:=string(goods.href[strings.LastIndex(goods.href,"/")+1:])
				if goodsId!="-1" && goods.status=="立即购买" {
					xiuluoyunBuy(goodsId)
				}
			}
		})
	})
	err := c.Visit(parentUrl)
	if err != nil {
		fmt.Println("xiuluoyunTargetTask err:",err.Error())
	}
}
//登录修罗云
func xiuluoyunBuy(productId string)  {
	const url = "https://www.xiuluohost.com/dashboard/submit_confirm"
	const cookie="referer_url=%2F; upv2=20201023%2C10; ci_session=g84guu89etgjl9bnb0oigdg85vherjuk; uid=e3a5fPCJRoMrBdLwf3Z6gV1owwUCfxiTc4%2FKng6cka%2F1vAEFB%2F%2BX0CUg4AaD8J8A%2F; upw=4baddkxAC%2FFcE6VT3GfizD2Csr5jvUhf%2BSh4Oc%2FIrFyMbrvZb1TP%2FRnzypwBT0J%2BRrFaZ%2FXHN32YMZxMEeQ"
	const password="antake666666"
	fmt.Println("开始下单")
	c:=colly.NewCollector()
	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Headers.Set("Cookie", cookie)
	})
	c.OnResponse(func(resp *colly.Response) {
		result, err := zhToUnicode(resp.Body)
		if err != nil {
			fmt.Println("解析结果错误:",err.Error())
			return
		}
		fmt.Println(string(result))
		if strings.Contains(string(result),"成功") || strings.Contains(string(result),"参数错误") || strings.Contains(string(result),"余额不足") || strings.Contains(string(result),"订单已创建，前去付款"){
			close(stopCh)
		}
	})
	requestData:=map[string]string{
		"os":"centos7",
		"os_passwd":password,
		"product_id":productId,
		"coupon":"",
		"ex_coupon":"0",
		"ips_num":"1",
		"remark":"",
		"confirmtos":"on",
		"billing_cycle":"30",
		"payment":"alipay",//这里修改支付方式
	}
	//alipay
	err := c.Post(url,requestData)
	if err != nil {
		fmt.Println("下单错误")
		return
	}
	fmt.Println("下单结束")
}
func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/keepchen/go-sail/v3/sail"
	"nav-server/app/admin/http/api/ack"
	"nav-server/app/admin/http/api/req"
	"nav-server/pkg/constants"
	"nav-server/pkg/models"
	"time"
)

type indexSvc struct{}

var Index = &indexSvc{}

func (*indexSvc) List(c *gin.Context) {
	var (
		form req.NavListReq
		resp ack.NavListAck
		//loggerSvc   = sail.LogTrace(c).GetLogger()
		_, cancel = context.WithTimeout(context.Background(), time.Second*5)
	)
	defer cancel()
	if err := c.ShouldBind(&form); err != nil {
		sail.Response(c).Wrap(constants.ErrRequestParamsInvalid, resp, err.Error()).Send()
		return
	}
	if code, err := form.Validator(); err != nil {
		sail.Response(c).Wrap(code, resp, err.Error()).Send()
		return
	}

	//使用内存中的数据而不是每次实时查询数据库
	//
	//内存中的数据使用计划任务定时更新
	resp.List = navListFromMem
	sail.Response(c).Data(resp)
}

var navListFromMem = make([]ack.NavCategory, 0)

func (*indexSvc) SaveNavCategoryToMem() {
	var (
		categories []models.Category
		sites      []models.Site
	)
	sail.GetDBR().Order("sort_index asc").Find(&categories)
	sail.GetDBR().Order("sort_index asc").Find(&sites)
	var navCategories = make([]ack.NavCategory, 0, len(categories))
	for _, category := range categories {
		navCategory := ack.NavCategory{
			Identity:  category.Identity,
			Name:      category.Name,
			Icon:      category.Icon,
			SortIndex: category.SortIndex,
			Sites:     make([]ack.NavSite, 0),
		}
		for _, site := range sites {
			if site.CategoryIdentity == category.Identity {
				navCategory.Sites = append(navCategory.Sites, ack.NavSite{
					CategoryIdentity: site.CategoryIdentity,
					Identity:         site.Identity,
					Name:             site.Name,
					Description:      site.Description,
					URL:              site.URL,
					Icon:             site.Icon,
					SortIndex:        site.SortIndex,
				})
			}
		}

		navCategories = append(navCategories, navCategory)
	}

	navListFromMem = navCategories
}

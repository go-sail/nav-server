package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/keepchen/go-sail/v3/sail"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nav-server/app/admin/http/api/ack"
	"nav-server/app/admin/http/api/req"
	"nav-server/pkg/constants"
	"nav-server/pkg/models"
	"nav-server/pkg/utils"
	"time"
)

type navSvc struct{}

var Nav = &navSvc{}

func (*navSvc) List(c *gin.Context) {
	var (
		form req.NavListReq
		resp ack.NavListAck
		//loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	var (
		categories []models.Category
		sites      []models.Site
	)
	sail.GetDBR().WithContext(ctx).Order("sort_index asc").Find(&categories)
	sail.GetDBR().WithContext(ctx).Order("sort_index asc").Find(&sites)
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

	resp.List = navCategories
	sail.Response(c).Data(resp)
}

func (*navSvc) Categories(c *gin.Context) {
	var (
		form req.NavCategoryListReq
		resp ack.NavCategoryListAck
		//loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	var categories []models.Category
	sail.GetDBR().WithContext(ctx).Order("sort_index asc").Find(&categories)
	var navCategories = make([]ack.NavCategory, 0, len(categories))
	for _, category := range categories {
		navCategory := ack.NavCategory{
			Identity:  category.Identity,
			Name:      category.Name,
			Icon:      category.Icon,
			SortIndex: category.SortIndex,
			Sites:     make([]ack.NavSite, 0),
		}
		navCategories = append(navCategories, navCategory)
	}

	resp.List = navCategories
	sail.Response(c).Data(resp)
}

func (*navSvc) Sites(c *gin.Context) {
	var (
		form req.NavSiteListReq
		resp ack.NavSiteListAck
		//loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	var sites []models.Site
	sail.GetDBR().WithContext(ctx).
		Where(&models.Site{CategoryIdentity: form.CategoryIdentity}).
		Order("sort_index asc").
		Find(&sites)
	var navSites = make([]ack.NavSite, 0, len(sites))

	for _, site := range sites {
		navSites = append(navSites, ack.NavSite{
			CategoryIdentity: site.CategoryIdentity,
			Identity:         site.Identity,
			Name:             site.Name,
			Description:      site.Description,
			URL:              site.URL,
			Icon:             site.Icon,
			SortIndex:        site.SortIndex,
		})
	}

	resp.List = navSites
	sail.Response(c).Data(resp)
}

func (*navSvc) SaveCategory(c *gin.Context) {
	var (
		form req.NavCategorySaveReq
		resp ack.NavCategorySaveAck
		//loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	if len(form.Identity) == 0 {
		//创建
		var maxSortIndex int
		sail.GetDBW().WithContext(ctx).Model(&models.Category{}).Select("MAX(sort_index)").Scan(&maxSortIndex)
		sail.GetDBW().WithContext(ctx).Create(&models.Category{
			Identity:  utils.MakeIdentity(),
			Name:      form.Name,
			Icon:      form.Icon,
			SortIndex: maxSortIndex + 1,
		})
	} else {
		//更新
		sail.GetDBW().WithContext(ctx).Where(&models.Category{Identity: form.Identity}).
			Updates(&models.Category{
				Name: form.Name,
				Icon: form.Icon,
			})
	}
	sail.Response(c).Data(resp)
}

func (*navSvc) DeleteCategory(c *gin.Context) {
	var (
		form        req.NavCategoryDeleteReq
		resp        ack.NavCategoryDeleteAck
		loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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

	var category models.Category
	sail.GetDBR().WithContext(ctx).Where(&models.Category{Identity: form.Identity}).First(&category)
	if len(category.Identity) == 0 {
		sail.Response(c).Wrap(constants.ErrCategoryNotFound, resp).Send()
		return
	}

	err := sail.GetDBW().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Where(&models.Category{Identity: form.Identity}).Delete(&models.Category{}).Error
		if err != nil {
			return err
		}
		return tx.Where(&models.Site{CategoryIdentity: form.Identity}).Delete(&models.Site{}).Error
	})

	if err != nil {
		loggerSvc.Error("删除分类失败", zap.String("identity", form.Identity), zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
		return
	} else {
		//TODO 可能存在同一文件被多处引用，因此暂时不删除本地文件
		//removeLocalIcon(category.Icon)
	}

	sail.Response(c).Data(resp)
}

func (*navSvc) DeleteSite(c *gin.Context) {
	var (
		form        req.NavSiteDeleteReq
		resp        ack.NavSiteDeleteAck
		loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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

	var site models.Site
	sail.GetDBR().WithContext(ctx).Where(&models.Site{Identity: form.Identity}).First(&site)
	if len(site.Identity) == 0 {
		sail.Response(c).Wrap(constants.ErrSiteNotFound, resp).Send()
		return
	}

	err := sail.GetDBW().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Where(&models.Site{Identity: form.Identity}).Delete(&models.Site{}).Error
	})

	if err != nil {
		loggerSvc.Error("删除站点失败", zap.String("identity", form.Identity), zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
		return
	} else {
		//TODO 可能存在同一文件被多处引用，因此暂时不删除本地文件
		//removeLocalIcon(site.Icon)
	}

	sail.Response(c).Data(resp)
}

func (*navSvc) SaveSite(c *gin.Context) {
	var (
		form req.NavSiteSaveReq
		resp ack.NavSiteSaveAck
		//loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	if len(form.Identity) == 0 {
		//创建
		var maxSortIndex int
		sail.GetDBW().WithContext(ctx).Model(&models.Site{}).Select("MAX(sort_index)").Scan(&maxSortIndex)
		sail.GetDBW().WithContext(ctx).Create(&models.Site{
			CategoryIdentity: form.CategoryIdentity,
			Identity:         utils.MakeIdentity(),
			Name:             form.Name,
			Description:      form.Description,
			URL:              form.URL,
			Icon:             form.Icon,
			SortIndex:        maxSortIndex + 1,
		})
	} else {
		//更新
		sail.GetDBW().WithContext(ctx).Where(&models.Site{Identity: form.Identity}).
			Updates(&models.Site{
				Name:        form.Name,
				Description: form.Description,
				URL:         form.URL,
				Icon:        form.Icon,
			})
	}
	sail.Response(c).Data(resp)
}

func (*navSvc) SortCategories(c *gin.Context) {
	var (
		form        req.NavCategorySortedReq
		resp        ack.NavCategorySortedAck
		loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	err := sail.GetDBW().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for idx, identity := range form.Identities {
			err := tx.Where(&models.Category{Identity: identity}).Updates(&models.Category{
				SortIndex: idx + 1,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		loggerSvc.Error("保存分类排序失败", zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
	}

	sail.Response(c).Data(resp)
}

func (*navSvc) SortSites(c *gin.Context) {
	var (
		form        req.NavSiteSortedReq
		resp        ack.NavSiteSortedAck
		loggerSvc   = sail.LogTrace(c).GetLogger()
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
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
	err := sail.GetDBW().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for idx, identity := range form.Identities {
			err := tx.Where(&models.Site{CategoryIdentity: form.CategoryIdentity, Identity: identity}).
				Updates(&models.Site{
					SortIndex: idx + 1,
				}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		loggerSvc.Error("保存站点排序失败", zap.String("err", err.Error()))
		sail.Response(c).Wrap(constants.ErrInternalServerError, resp, err.Error()).Send()
	}

	sail.Response(c).Data(resp)
}

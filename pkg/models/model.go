package models

import (
	"fmt"
	"github.com/keepchen/go-sail/v3/sail"
	"gorm.io/gorm"
	"nav-server/app/admin/config"
	"nav-server/app/admin/http/api/ack"
	"nav-server/pkg/utils"
)

// AutoMigrate 自动同步表结构和数据迁移
func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&Category{},
		&Site{},
	)
	if err != nil {
		return err
	}

	//数据迁移

	//-初始化用户
	var user User
	sail.GetDBW().First(&user)
	if len(user.UID) == 0 {
		username := config.Get().Nav.InitUser.Username
		password := config.Get().Nav.InitUser.Password
		if len(username) == 0 || len(password) == 0 {
			panic("站点管理员账号和密码不能为空")
		}
		passwordEncrypt, err := sail.JWT().Encrypt(password)
		if err != nil {
			return err
		}
		err = db.Create(&User{
			UID:      utils.MakeUid(),
			Username: username,
			Password: passwordEncrypt,
		}).Error
		if err != nil {
			return err
		}
	}

	//-初始化导航数据
	var (
		categoryCount int64
		siteCount     int64
	)
	db.Model(&Category{}).Count(&categoryCount)
	db.Model(&Site{}).Count(&siteCount)

	//如果数据表不为空，则不写入初始数据
	if categoryCount != 0 || siteCount != 0 {
		return nil
	}

	list := DefaultNavList()

	return db.Transaction(func(tx *gorm.DB) error {
		for cateIndex, category := range list {
			cateIdentity := utils.MakeIdentity()
			err := tx.Create(&Category{
				Identity:  cateIdentity,
				Name:      category.Name,
				Icon:      category.Icon,
				SortIndex: cateIndex + 1,
			}).Error

			if err != nil {
				return err
			}

			for siteIndex, site := range category.Sites {
				siteIdentity := utils.MakeIdentity()
				err := tx.Create(&Site{
					CategoryIdentity: cateIdentity,
					Identity:         siteIdentity,
					Name:             site.Name,
					Description:      site.Description,
					Icon:             site.Icon,
					URL:              site.URL,
					SortIndex:        siteIndex + 1,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func DefaultNavList() []ack.NavCategory {
	return []ack.NavCategory{
		{
			Name: "热门常用",
			Icon: fmt.Sprintf("%s/icons/category/fire.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "GitHub",
					Description: "开源胜地，代码托管",
					Icon:        fmt.Sprintf("%s/icons/site/github.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://github.com",
				},
				{
					Name:        "Go-Sail",
					Description: "轻量的Go Web框架",
					Icon:        fmt.Sprintf("%s/icons/site/go-sail-logo.png", config.Get().Nav.IconEndpoint),
					URL:         "https://go-sail.dev",
				},
				{
					Name:        "StarDots",
					Description: "专业的图像托管平台",
					Icon:        fmt.Sprintf("%s/icons/site/stardots.io.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://stardots.io",
				},
			},
		},
		{
			Name: "AI智能",
			Icon: fmt.Sprintf("%s/icons/category/robot.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "ChatGPT",
					Description: "AI助手 by OpenAI",
					Icon:        fmt.Sprintf("%s/icons/site/chat.openai.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://chat.openai.com",
				},
				{
					Name:        "Grok",
					Description: "AI助手 by X",
					Icon:        fmt.Sprintf("%s/icons/site/grok.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://grok.com",
				},
				{
					Name:        "DeepSeek",
					Description: "深度求索",
					Icon:        fmt.Sprintf("%s/icons/site/deepseek.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://chat.deepseek.com",
				},
			},
		},
		{
			Name: "云服务",
			Icon: fmt.Sprintf("%s/icons/category/cloudy.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "Cloudflare",
					Description: "全球CDN和网络安全服务",
					Icon:        fmt.Sprintf("%s/icons/site/www.cloudflare.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.cloudflare.com",
				},
				{
					Name:        "腾讯云",
					Description: "腾讯云计算服务",
					Icon:        fmt.Sprintf("%s/icons/site/cloud.tencent.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://cloud.tencent.com",
				},
				{
					Name:        "Bunney.net",
					Description: "全球边缘平台",
					Icon:        fmt.Sprintf("%s/icons/site/bunny.net.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://bunny.net",
				},
			},
		},
		{
			Name: "开发工具",
			Icon: fmt.Sprintf("%s/icons/category/settings.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "Go-Sail",
					Description: "轻量的Go Web框架",
					Icon:        fmt.Sprintf("%s/icons/site/go-sail-logo.png", config.Get().Nav.IconEndpoint),
					URL:         "https://go-sail.dev",
				},
				{
					Name:        "Postman",
					Description: "API测试工具",
					Icon:        fmt.Sprintf("%s/icons/site/www.postman.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.postman.com",
				},
			},
		},
		{
			Name: "社区论坛",
			Icon: fmt.Sprintf("%s/icons/category/forum.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "GitHub",
					Description: "开源胜地，代码托管",
					Icon:        fmt.Sprintf("%s/icons/site/github.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://github.com",
				},
				{
					Name:        "V2EX",
					Description: "创意工作者社区",
					Icon:        fmt.Sprintf("%s/icons/site/v2ex.com.png", config.Get().Nav.IconEndpoint),
					URL:         "https://v2ex.com",
				},
				{
					Name:        "LinuxDo",
					Description: "Linux技术社区，Peace and Love",
					Icon:        fmt.Sprintf("%s/icons/site/linux.do.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://linux.do",
				},
			},
		},
		{
			Name: "设计工具",
			Icon: fmt.Sprintf("%s/icons/category/design.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "Figma",
					Description: "UI设计工具",
					Icon:        fmt.Sprintf("%s/icons/site/figma.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://figma.com",
				},
				{
					Name:        "Sketch",
					Description: "界面设计工具",
					Icon:        fmt.Sprintf("%s/icons/site/www.sketch.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.sketch.com",
				},
				{
					Name:        "幕布",
					Description: "笔记、思维导图APP",
					Icon:        fmt.Sprintf("%s/icons/site/mubu.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://mubu.com",
				},
			},
		},
		{
			Name: "财经投资",
			Icon: fmt.Sprintf("%s/icons/category/finance.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "币安",
					Description: "加密货币交易平台",
					Icon:        fmt.Sprintf("%s/icons/site/www.binance.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.binance.com",
				},
				{
					Name:        "OKX",
					Description: "数字资产交易服务平台",
					Icon:        fmt.Sprintf("%s/icons/site/www.okx.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.okx.com",
				},
				{
					Name:        "KuCoin",
					Description: "加密货币交易平台",
					Icon:        fmt.Sprintf("%s/icons/site/kucoin.com.png", config.Get().Nav.IconEndpoint),
					URL:         "https://kucoin.com",
				},
				{
					Name:        "雪球",
					Description: "聪明的投资者都在这里",
					Icon:        fmt.Sprintf("%s/icons/site/xueqiu.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://xueqiu.com",
				},
				{
					Name:        "富途牛牛",
					Description: "港美股交易软件",
					Icon:        fmt.Sprintf("%s/icons/site/www.futunn.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.futunn.com",
				},
				{
					Name:        "同花顺",
					Description: "专业股票软件及金融信息服务",
					Icon:        fmt.Sprintf("%s/icons/site/www.10jqka.com.cn.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.10jqka.com.cn",
				},
			},
		},
		{
			Name: "学习资源",
			Icon: fmt.Sprintf("%s/icons/category/books.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "MDN Web Docs",
					Description: "Web开发权威文档",
					Icon:        fmt.Sprintf("%s/icons/site/developer.mozilla.org.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://developer.mozilla.org",
				},
				{
					Name:        "W3Schools",
					Description: "Web技术教程",
					Icon:        fmt.Sprintf("%s/icons/site/www.w3schools.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.w3schools.com",
				},
			},
		},
		{
			Name: "在线工具",
			Icon: fmt.Sprintf("%s/icons/category/website.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "StarDots",
					Description: "专业的图像托管平台",
					Icon:        fmt.Sprintf("%s/icons/site/stardots.io.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://stardots.io",
				},
				{
					Name:        "Recompressor",
					Description: "最优图像优化",
					Icon:        fmt.Sprintf("%s/icons/site/recompressor.com.svg", config.Get().Nav.IconEndpoint),
					URL:         "https://zh.recompressor.com",
				},
				{
					Name:        "Excalidraw",
					Description: "虚拟协作白板工具",
					Icon:        fmt.Sprintf("%s/icons/site/excalidraw.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://excalidraw.com",
				},
				{
					Name:        "Random100",
					Description: "黑客风格的随机生成器",
					Icon:        fmt.Sprintf("%s/icons/site/random100.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://random100.com",
				},
			},
		},
		{
			Name: "娱乐休闲",
			Icon: fmt.Sprintf("%s/icons/category/game.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "Emu666",
					Description: "在线模拟器游戏",
					Icon:        fmt.Sprintf("%s/icons/site/www.emu666.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.emu666.com",
				},
				{
					Name:        "哔哩哔哩",
					Description: "弹幕视频网站",
					Icon:        fmt.Sprintf("%s/icons/site/bilibili.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://bilibili.com",
				},
			},
		},
		{
			Name: "办公协作",
			Icon: fmt.Sprintf("%s/icons/category/bag.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "飞书",
					Description: "企业协作平台",
					Icon:        fmt.Sprintf("%s/icons/site/www.feishu.cn.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://www.feishu.cn",
				},
				{
					Name:        "Slack",
					Description: "团队协作工具",
					Icon:        fmt.Sprintf("%s/icons/site/slack.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://slack.com",
				},
				{
					Name:        "Tower",
					Description: "为协作而设计",
					Icon:        fmt.Sprintf("%s/icons/site/tower.im.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://tower.im",
				},
			},
		},
		{
			Name: "博客圈子",
			Icon: fmt.Sprintf("%s/icons/category/blog.svg", config.Get().Nav.IconEndpoint),
			Sites: []ack.NavSite{
				{
					Name:        "十年之约",
					Description: "这是一个记录十年之约的网站",
					Icon:        fmt.Sprintf("%s/icons/site/foreverblog.cn.webp", config.Get().Nav.IconEndpoint),
					URL:         "https://foreverblog.cn",
				},
				{
					Name:        "BlogWe",
					Description: "致敬还在写博客的我们",
					Icon:        fmt.Sprintf("%s/icons/site/blogwe.com.ico", config.Get().Nav.IconEndpoint),
					URL:         "https://blogwe.com",
				},
				{
					Name:        "博友圈",
					Description: "博客收录与文章RSS聚合网站",
					Icon:        fmt.Sprintf("%s/icons/site/www.boyouquan.com.png", config.Get().Nav.IconEndpoint),
					URL:         "https://www.boyouquan.com",
				},
				{
					Name:        "博客录",
					Description: "博客收录展示平台",
					Icon:        fmt.Sprintf("%s/icons/site/boke.lu.png", config.Get().Nav.IconEndpoint),
					URL:         "https://boke.lu",
				},
			},
		},
	}
}

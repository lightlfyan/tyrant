package cfg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type CfgDataType struct {
	Url         string `yaml:"url"`
	MaxConnect  int    `yaml:"maxConnect"`
	TimeWaitMs  int    `yaml:"timeWaitMs"`
	TimeOutMs   int    `yaml:"timeOutMs"`
	ContentType string `yaml:"contentType"`
	ApiType     int    `yaml:"apitype"`
	Code        int    `yaml:"code"`
	Resp        string `yaml:"resp"`
}

var GetUrlCfg map[string]CfgDataType
var PostUrlCfg map[string]CfgDataType

var UrlChanMap map[string]chan int

func initChan() {
	UrlChanMap = make(map[string]chan int)
	for key, value := range GetUrlCfg {
		if value.MaxConnect > 0 {
			UrlChanMap[key] = make(chan int, value.MaxConnect)
		}
	}
}

func GetCfg(path, method string) *CfgDataType {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}
	UrlCfg := GetUrlCfg

	if method == "POST" {
		UrlCfg = PostUrlCfg
	}

	if method == "PATCH" {
		return &CfgDataType{
			MaxConnect:  0,
			TimeOutMs:   0,
			ContentType: "text/html",
			ApiType:     0,
			Code:        205,
			Resp:        ``,
		}
	}

	if method == "PUT" {
		return &CfgDataType{
			MaxConnect:  0,
			TimeOutMs:   0,
			ContentType: "text/html",
			ApiType:     0,
			Code:        205,
			Resp:        ``,
		}
	}

	if method == "DELETE" {
		return &CfgDataType{
			MaxConnect:  0,
			TimeOutMs:   0,
			ContentType: "text/html",
			ApiType:     0,
			Code:        204,
			Resp:        ``,
		}
	}

	c, ok := UrlCfg[path]
	if ok {
		return &c
	}

	// try
	path1 := "^" + path + "$"
	c, ok = UrlCfg[path1]
	if ok {
		return &c
	}

	for k, v := range UrlCfg {
		valid := regexp.MustCompile(k)
		check := valid.MatchString(path)
		//log.Println(path, k, check)

		if check {
			return &v
		}
	}

	return nil
}

type T struct {
	Config []CfgDataType `yaml:"config"`
}

func Load() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("config: " + dir + "/conf.json")
	}

	bins, _ := ioutil.ReadFile(dir + "/config.yaml")
	t := T{}

	err = yaml.Unmarshal(bins, &t)
	log.Println("config: ", t, err)
	return

	file, _ := os.Open(dir + "/conf.json")
	decoder := json.NewDecoder(file)
	GetUrlCfg = make(map[string]CfgDataType)
	err = decoder.Decode(&GetUrlCfg)
	if err != nil {
		fmt.Println("err: ", err)
	}
	initChan()
	file.Close()

	PostUrlCfg = make(map[string]CfgDataType)

	PostUrlCfg["^/auth/$"] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "username": "caozijie",
  "name": "曹子杰",
  "nick_name": "caozj",
  "avatar": "http://123.con/234.jpg",
  "account_id": "2c90a5924954f846014954f8926f0002",
  "org_id": "8a80809c5bf03b4a015bf0f19342000d",
  "id": "LKnguYMQ53szcU9znkLwic9ifff5a19d16d24eaf",
  "permissions": [
    {
      "id": "4cdc74c802bf43a4ba6306177261b2ba",
      "pid": "8900a8c3b5b84e89aa96d4c35a07c4ea",
      "portal_id": "PORTAL_SALES",
      "code": "MENU_PERM",
      "name": "权限新增",
      "type": "MENU",
      "url": "/**",
      "service": "/path/to/micro-service/operation",
      "level": "12",
      "sort": 10,
      "stage": "0"
    }
  ],
  "portal_id": "PORTAL_MGT"
}`,
	}

	GetUrlCfg[`^/auth/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "username": "caozijie",
  "name": "曹子杰",
  "nick_name": "caozj",
  "avatar": "http://123.con/234.jpg",
  "account_id": "2c90a5924954f846014954f8926f0002",
  "org_id": "8a80809c5bf03b4a015bf0f19342000d",
  "v10n": {
    "id": "666",
    "code": "123456",
    "account_id": "666",
    "account_name": "test@yupiao.com",
    "target": "sms",
    "mobile": "13000008888",
    "cur_pwd": "654321"
  },
  "id": "LKnguYMQ53szcU9znkLwic9ifff5a19d16d24eaf",
  "permissions": [
    {
      "id": "4cdc74c802bf43a4ba6306177261b2ba",
      "pid": "8900a8c3b5b84e89aa96d4c35a07c4ea",
      "portal_id": "PORTAL_SALES",
      "code": "MENU_PERM",
      "name": "权限新增",
      "type": "MENU",
      "url": "/**",
      "service": "/path/to/micro-service/operation",
      "level": "12",
      "sort": 10,
      "stage": "0"
    }
  ],
  "portal_id": "PORTAL_MGT"
}`,
	}

	PostUrlCfg[`^/accounts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "2c90a5924954f846014954f8926f0002",
  "username": "caozijie@wepiao.com",
  "name": "",
  "gender": 0,
  "password": "123456",
  "roles": [],
  "org_name": "上海研发",
  "email": ""
}`,
	}

	GetUrlCfg[`^/accounts/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "2c90a5924954f846014954f8926f0002",
  "username": "caozijie@wepiao.com",
  "name": "曹子杰",
  "gender": 0,
  "password": "123456",
  "mobile": "188166661666",
  "stage": "0",
  "org_id": "8a8080935bcd0e54015bcdc0a4a5003a",
  "roles": [
    {
      "id": "40288c0f5be700fe015be70ca24f0002",
      "name": "AC_ROOT",
      "portal_id": "PORTAL_SALES",
      "permissions": []
    }
  ],
  "org_name": "上海研发"
}`,
	}

	GetUrlCfg[`^/accounts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "2c90a5924954f846014954f8926f0002",
      "username": "caozijie@wepiao.com",
      "name": "曹子杰",
      "gender": 0,
      "password": "123456",
      "mobile": "188166661666",
      "stage": "0",
      "org_id": "8a8080935bcd0e54015bcdc0a4a5003a",
      "roles": [
        {
          "id": "40288c0f5be700fe015be70ca24f0002",
          "name": "AC_ROOT",
          "portal_id": "PORTAL_SALES",
          "permissions": []
        }
      ],
      "org_name": "上海研发"
    }
  ]
}`,
	}

	PostUrlCfg[`^/accounts/\w+/@roles/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "2c90a5924954f846014954f8926f0002",
  "username": "caozijie@wepiao.com",
  "name": "",
  "gender": 0,
  "password": "123456",
  "roles": [],
  "org_name": "上海研发"
}`,
	}

	GetUrlCfg[`^/accounts/\w+/@permissions/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "4cdc74c802bf43a4ba6306177261b2ba",
      "pid": "8900a8c3b5b84e89aa96d4c35a07c4ea",
      "portal_id": "PORTAL_SALES",
      "code": "MENU_PERM",
      "name": "权限新增",
      "type": "MENU",
      "url": "/**",
      "service": "/path/to/micro-service/operation",
      "level": "12",
      "sort": 10,
      "stage": "0",
      "created": "Hello, world!"
    }
  ]
}`,
	}

	PostUrlCfg[`^/orgs/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "8cae3fc620ee44078800f0a5e0706cd8",
  "hier_name": "项目运营部-城市运营中心"
}`,
	}

	GetUrlCfg[`^/orgs/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "8cae3fc620ee44078800f0a5e0706cd8",
  "pid": "0ae369a82daf47e5a05ce6fca55fdc70",
  "name": "上海研发部",
  "roles": [
    {
      "id": "40288c0f5be700fe015be70ca24f0002",
      "name": "AC_ROOT",
      "portal_id": "PORTAL_SALES"
    }
  ],
  "level": "4",
  "sort": "10",
  "stage": "0",
  "type": "ORG",
  "hier_name": "项目运营部-城市运营中心"
}`,
	}

	GetUrlCfg[`^/orgs/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "8cae3fc620ee44078800f0a5e0706cd8",
      "pid": "0ae369a82daf47e5a05ce6fca55fdc70",
      "name": "上海研发部",
      "roles": [
        {
          "id": "40288c0f5be700fe015be70ca24f0002",
          "name": "AC_ROOT",
          "portal_id": "PORTAL_SALES",
          "permissions": []
        }
      ],
      "level": "4",
      "sort": "10",
      "stage": "0",
      "type": "ORG",
      "hier_name": "项目运营部-城市运营中心"
    }
  ]
}`,
	}

	PostUrlCfg[`^/orgs/\w+/@roles/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "8cae3fc620ee44078800f0a5e0706cd8",
  "roles": [],
  "hier_name": "项目运营部-城市运营中心"
}`,
	}

	PostUrlCfg[`^/merts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "4028b8e55c0cae90015c0eee63d80000",
  "biz_prin_org_id": "8a80809c5bf03b4a015bf0f19342000d"
}`,
	}

	GetUrlCfg[`^/merts/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "4028b8e55c0cae90015c0eee63d80000",
  "name": "娱悦发行",
  "code": "SPL12345678",
  "biz_prin": "张三丰",
  "biz_prin_org_id": "8a80809c5bf03b4a015bf0f19342000d",
  "prin_name": "张三",
  "city_code": "1111",
  "mert_type": "0",
  "stage": "0",
  "prin_mobile": "18816661666",
  "prin_email": "zhangsan@gmail.com",
  "prin_idcard": "310105195010040034",
  "idcard_front": "/path/to/idcard_front.jpg",
  "idcard_back": "/path/to/idcard_back.jpg",
  "license_no": "12345",
  "license_img": "path/to/license_img.jpg",
  "started": "2017-05-18 17:05:09",
  "ended": "2017-10-18 17:05:09",
  "created": "2017-10-18 17:05:09",
  "creator": "李四"
}`,
	}

	GetUrlCfg[`^/merts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "4028b8e55c0cae90015c0eee63d80000",
      "name": "娱悦发行",
      "code": "SPL12345678",
      "biz_prin": "张三丰",
      "biz_prin_org_id": "8a80809c5bf03b4a015bf0f19342000d",
      "prin_name": "张三",
      "city_code": "1111",
      "mert_type": "0",
      "stage": "0",
      "prin_mobile": "18816661666",
      "prin_email": "zhangsan@gmail.com",
      "prin_idcard": "310105195010040034",
      "idcard_front": "/path/to/idcard_front.jpg",
      "idcard_back": "/path/to/idcard_back.jpg",
      "license_no": "12345",
      "license_img": "path/to/license_img.jpg",
      "started": "2017-05-18 17:05:09",
      "ended": "2017-10-18 17:05:09",
      "created": "2017-10-18 17:05:09",
      "creator": "李四"
    }
  ]
}`,
	}

	PostUrlCfg[`^/dists/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "4028b8e55c0cae90015c0eee63d80000",
  "biz_prin_org_id": "8a80809c5bf03b4a015bf0f19342000d"
}`,
	}

	GetUrlCfg[`^/dists/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "4028b8e55c0cae90015c0eee63d80000",
  "name": "周末去哪儿",
  "code": "AGC12345678",
  "biz_prin": "张三丰",
  "biz_prin_org_id": "8a80809c5bf03b4a015bf0f19342000d",
  "prin_name": "张三",
  "city_code": "1111",
  "mert_type": "0",
  "stage": "0",
  "prin_mobile": "18816661666",
  "prin_email": "zhangsan@gmail.com",
  "prin_idcard": "310105195010040034",
  "idcard_front": "path/to/idcard_front.jpg",
  "idcard_back": "/path/to/idcard_back.jpg",
  "license_no": "12345",
  "license_img": "path/to/license_img.jpg",
  "started": "2017-05-18 17:05:09",
  "ended": "2017-10-18 17:05:09",
  "created": "2017-10-18 17:05:09",
  "creator": "李四"
}`,
	}

	GetUrlCfg[`^/dists/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "4028b8e55c0cae90015c0eee63d80000",
      "name": "周末去哪儿",
      "code": "AGC12345678",
      "biz_prin": "张三丰",
      "biz_prin_org_id": "8a80809c5bf03b4a015bf0f19342000d",
      "prin_name": "张三",
      "city_code": "1111",
      "mert_type": "0",
      "stage": "0",
      "prin_mobile": "18816661666",
      "prin_email": "zhangsan@gmail.com",
      "prin_idcard": "310105195010040034",
      "idcard_front": "path/to/idcard_front.jpg",
      "idcard_back": "/path/to/idcard_back.jpg",
      "license_no": "12345",
      "license_img": "path/to/license_img.jpg",
      "started": "2017-05-18 17:05:09",
      "ended": "2017-10-18 17:05:09",
      "created": "2017-10-18 17:05:09",
      "creator": "李四"
    }
  ]
}`,
	}

	PostUrlCfg[`^/roles/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "40288c0f5be700fe015be70ca24f0002",
  "permissions": []
}`,
	}

	GetUrlCfg[`^/roles/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "40288c0f5be700fe015be70ca24f0002",
  "name": "AC_ROOT",
  "portal_id": "PORTAL_SALES",
  "permissions": [
    {
      "id": "4cdc74c802bf43a4ba6306177261b2ba",
      "portal_id": "PORTAL_SALES"
    }
  ]
}`,
	}

	GetUrlCfg[`^/roles/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "40288c0f5be700fe015be70ca24f0002",
      "name": "AC_ROOT",
      "portal_id": "PORTAL_SALES",
      "permissions": [
        {
          "id": "4cdc74c802bf43a4ba6306177261b2ba",
          "portal_id": "PORTAL_SALES"
        }
      ]
    }
  ]
}`,
	}

	PostUrlCfg[`^/roles/\w+/@permissions/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "id": {
      "type": "string",
      "description": "角色ID"
    },
    "name": {
      "type": "string",
      "description": "角色名"
    },
    "portal_id": {
      "type": "string",
      "description": "平台 (PORTAL_MGT: 电商平台, PORTAL_EC: 电商平台, PORTAL_SALES: 售票平台, PORTAL_TICKET: 检票平台, PORTAL_MERCHANT: 商户平台)"
    },
    "permissions": {
      "description": "权限集"
    },
    "created": {
      "type": "string",
      "description": "创建时间"
    },
    "modified": {
      "type": "string",
      "description": "修改时间"
    }
  }
}`,
	}

	GetUrlCfg[`^/roles/\w+/@accounts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "account_name": "caozijie",
      "account_id": "40288c0f5bdc1391015bdc16ee1d0000",
      "org_id": "8a80809c5bf03b4a015bf0f19342000d",
      "hier_name": "上海研发-产品4"
    }
  ]
}`,
	}

	PostUrlCfg[`^/permissions/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "id": "4cdc74c802bf43a4ba6306177261b2ba"
}`,
	}

	GetUrlCfg[`^/permissions/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "4cdc74c802bf43a4ba6306177261b2ba",
  "pid": "8900a8c3b5b84e89aa96d4c35a07c4ea",
  "portal_id": "PORTAL_SALES",
  "code": "MENU_PERM",
  "name": "权限新增",
  "type": "MENU",
  "url": "/**",
  "service": "/path/to/micro-service/operation",
  "level": "12",
  "sort": 10,
  "stage": "0",
  "created": "Hello, world!"
}`,
	}

	GetUrlCfg[`^/permissions/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "4cdc74c802bf43a4ba6306177261b2ba",
      "pid": "8900a8c3b5b84e89aa96d4c35a07c4ea",
      "portal_id": "PORTAL_SALES",
      "code": "MENU_PERM",
      "name": "权限新增",
      "type": "MENU",
      "url": "/**",
      "service": "/path/to/micro-service/operation",
      "level": "12",
      "sort": 10,
      "stage": "0",
      "created": "Hello, world!"
    }
  ]
}`,
	}

	GetUrlCfg[`^/permissions/\w+/@roles/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "4cdc74c802bf43a4ba6306177261b2ba",
      "pid": "8900a8c3b5b84e89aa96d4c35a07c4ea",
      "portal_id": "PORTAL_SALES",
      "code": "MENU_PERM",
      "name": "权限新增",
      "type": "MENU",
      "url": "/**",
      "service": "/path/to/micro-service/operation",
      "level": "12",
      "sort": 10,
      "stage": "0",
      "created": "Hello, world!"
    }
  ]
}`,
	}

	PostUrlCfg[`^/dicts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "properties": {
    "id": {
      "type": "string"
    },
    "group": {
      "type": "string"
    },
    "dict_key": {
      "type": "string"
    },
    "dict_value": {
      "type": "string"
    },
    "sort": {
      "type": "string"
    },
    "stage": {
      "type": "string"
    }
  }
}`,
	}

	GetUrlCfg[`^/dicts/\w+/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        200,
		Resp: `{
  "id": "192088df78dc89df78d6f7c6df78dsvb8f7df",
  "group": "pay_method",
  "dict_key": "paid",
  "dict_value": "30010",
  "sort": "0",
  "stage": "0"
}`,
	}

	GetUrlCfg[`^/dicts/$`] = CfgDataType{
		MaxConnect:  0,
		TimeOutMs:   0,
		ContentType: "application/json;charset=UTF-8",
		ApiType:     0,
		Code:        201,
		Resp: `{
  "info": {
    "page": 1,
    "per_page": 10,
    "search_time": 0,
    "total": 9
  },
  "list": [
    {
      "id": "192088df78dc89df78d6f7c6df78dsvb8f7df",
      "group": "pay_method",
      "dict_key": "paid",
      "dict_value": "30010",
      "sort": "0",
      "stage": "0"
    }
  ]
}`,
	}
}

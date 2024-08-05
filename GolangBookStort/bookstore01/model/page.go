package model

type Page struct {
	Books       []*Book //每页图书
	PageNo      int64   //当前页
	PageSize    int64   //每页显示树
	TotalPageNo int64   //计算
	TotalRecord int64   //查询数据库
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}

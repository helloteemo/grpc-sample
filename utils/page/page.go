package page

// Page 分页信息
type Page struct {
	PageNumber int `form:"page_number" json:"page_number"` // 分页信息，和page_size配合使用，页码，从1开始
	PageSize   int `form:"page_size" json:"page_size"`     // 分页信息，和page_number配合使用，页数，每页最大20
}

func (p *Page) SqlLimitAndOffset() (limit int, offset int) {
	return p.SqlLimit(), p.SqlOffset()
}

// SqlLimit 获取sql参数中的limit
func (p *Page) SqlLimit() int {
	return p.PageSize
}

// SqlOffset 获取sql参数中的offset
func (p *Page) SqlOffset() int {
	return (p.PageNumber - 1) * p.PageSize
}

// RedisZset 生成redis的start和end
func (p *Page) RedisZset() (start int, end int) {
	return p.RedisZsetStart(), p.RedisZsetEnd()
}

// RedisZsetStart 获取zset的start参数
func (p *Page) RedisZsetStart() int {
	return p.PageSize * (p.PageNumber - 1)
}

// RedisZsetEnd 获取zset的end参数
func (p *Page) RedisZsetEnd() int {
	return p.RedisZsetStart() + p.PageSize - 1
}

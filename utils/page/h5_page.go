package page

// H5Page 分页信息
type H5Page struct {
	Limit  int `form:"limit" json:"limit"`   // 分页信息，和offset配合使用,代表要查询N条数据
	Offset int `form:"offset" json:"offset"` // 分页信息，和limit配合使用,从0开始,代表要从第N行开始往后查
}

func (p *H5Page) SqlLimitAndOffset() (limit int, offset int) {
	return p.SqlLimit(), p.SqlOffset()
}

// SqlLimit 获取sql参数中的limit
func (p *H5Page) SqlLimit() int {
	return p.Limit
}

// SqlOffset 获取sql参数中的offset
func (p *H5Page) SqlOffset() int {
	return p.Offset
}

// RedisZset 生成redis的start和end
func (p *H5Page) RedisZset() (start int, end int) {
	return p.RedisZsetStart(), p.RedisZsetEnd()
}

// RedisZsetStart 获取zset的start参数
func (p *H5Page) RedisZsetStart() int {
	return p.Offset
}

// RedisZsetEnd 获取zset的end参数
func (p *H5Page) RedisZsetEnd() int {
	return p.RedisZsetStart() + p.Limit - 1
}

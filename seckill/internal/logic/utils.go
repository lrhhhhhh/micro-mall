package logic

// 以下常量必须和lua文件中定义的一致
// secKill constant
const (
	ActivityNotFound = 1001
	ActivityNotStart = 1002
	ActivityExpire   = 1003
	ActivitySoldOut  = 1004
	ExceedBuyLimit   = 1005
	CronTaskFatal    = 1006
	ProductNotEnough = 1007
	SeckillSuccess   = 1008
	SeckillRetry     = 1009

	ErrUnknown = -1
)

// secKill revert constant
const (
	RevertSuccess          = 3000
	ErrActivityNotFound    = 3001
	ErrActivityNotStart    = 3002
	ErrHistoryNotFound     = 3003
	ErrUserHistoryNotFound = 3004
	ErrNegativeQuantity    = 3005
)

var StatusMap = map[int]string{
	ActivityNotFound: "活动不存在",
	ActivityNotStart: "活动未开始",
	ActivityExpire:   "活动已经结束",
	ActivitySoldOut:  "商品已经售空",
	ExceedBuyLimit:   "超过购买限制",
	ProductNotEnough: "商品不足",
	CronTaskFatal:    "定时任务失效",
	SeckillSuccess:   "秒杀成功",
	SeckillRetry:     "秒杀失败，请重试",
	ErrUnknown:       "未知错误",

	RevertSuccess:          "回滚成功",
	ErrActivityNotFound:    "活动不存在",
	ErrActivityNotStart:    "活动未开始",
	ErrHistoryNotFound:     "活动历史记录不存在",
	ErrUserHistoryNotFound: "用户历史记录不存在",
	ErrNegativeQuantity:    "商品数量为负数！",
}

func GetMessage(code int) (message string) {
	var ok bool
	if message, ok = StatusMap[code]; !ok {
		message, _ = StatusMap[ErrUnknown]
	}
	return
}
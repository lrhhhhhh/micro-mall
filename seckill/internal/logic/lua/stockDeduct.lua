--[[
用到的redis键有：
(1) 命名格式：SeckillActivity:XXX, XXX是ActivityId；类型：hashmap；
用途：记录了秒杀活动的所有信息，key是属性，value是属性对应的值

(2) 命名格式：SeckillActivityXXXHistory, XXX是ActivityId；类型：hasmap；
用途：记录了用户在秒杀活动XXX中的购买记录，key是用户id，value是用户购买数量

错误码，int类型：
1001 秒杀活动不存在
1002 秒杀活动未开始
1003 秒杀活动已经结束
1004 商品已经卖光
1005 用户超过购买限制
1006 商品不足
1006 成功
--]]


local ActivityNotFound = 1001;
local ActivityNotStart = 1002;
local ActivityExpire   = 1003;
local ActivitySoldOut  = 1004;
local ExceedBuyLimit   = 1005;
local CronTaskFatal    = 1006;
local ProductNotEnough = 1007;
local SeckillSuccess   = 1008;
local SeckillRetry     = 1009;


local ActivityName = KEYS[1];
local buyCnt = tonumber(ARGV[2]);
local nowTime = tonumber(ARGV[3]);  -- timestamp
local probability = tonumber(ARGV[4])


-- check activity whether exists
if redis.call("exists", ActivityName) == 0 then
    return ActivityNotFound
end

-- check start and end
-- todo: what happen if result is not a int
if tonumber(redis.call("hget", ActivityName, "StartTime")) > nowTime then
    return ActivityNotStart
end
if tonumber(redis.call("hget", ActivityName, "EndTime")) < nowTime then
    return ActivityExpire
end

-- check buyCnt
if buyCnt > tonumber(redis.call("hget", ActivityName, "BuyLimit")) then
    return ExceedBuyLimit
end

-- check product left
local activityTotal = redis.call("hget", ActivityName, "Total");
activityTotal = tonumber(activityTotal)
if activityTotal < buyCnt then
    if activityTotal == 0 then
        return ActivitySoldOut
    end
    return  ProductNotEnough
end

-- check probability
local activityProbability = tonumber(redis.call("hget", ActivityName, "BuyProbability"))
if tonumber(probability) < activityProbability then
    return SeckillRetry
end

-- success
redis.call("hset", ActivityName, "Total", activityTotal - buyCnt)
-- redis.call("hset", HistoryName, uid, buyCnt + historyBuyCnt)
return SeckillSuccess
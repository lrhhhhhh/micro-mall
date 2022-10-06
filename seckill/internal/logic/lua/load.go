package lua

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"io/ioutil"
	"seckill/internal/svc"
)

// LoadByFilename load lua script to redis and return sha1
// 将lua脚本加载到redis中，用一个key记录lua值
// 服务启动时，访问一次redis，直接将sha1保存在本地config中（这样是选择避免每次请求都要访问redis读取lua的sha值）
func LoadByFilename(rdb *redis.Redis, filename, key string) string {
	//sha, err := rdb.Get(key)
	//if err == nil && sha != "" {
	//	return sha
	//}

	// todo: only one client load the lua script

	lua, err := ioutil.ReadFile("./internal/logic/lua/" + filename)
	if err != nil {
		panic(err)
	}

	sha, err := rdb.ScriptLoad(string(lua))
	if err != nil {
		panic(err)
	}

	err = rdb.Set(key, sha)
	if err != nil {
		panic(err)
	}
	return sha
}

func LoadLuaScript(svcCtx *svc.ServiceContext) {
	svcCtx.Config.StockDeductSha1 = LoadByFilename(svcCtx.Redis, "stockDeduct.lua", svcCtx.Config.StockDeductLua)
	logx.Infof(
		"load lua done: (key=%s, sha1=%s)",
		svcCtx.Config.StockDeductLua, svcCtx.Config.StockDeductSha1,
	)
}

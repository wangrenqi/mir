package srv

import (
	"github.com/robfig/cron"
	"mir/com"
)

// 定时任务模块
func InitScheduler() *cron.Cron {
	c := cron.New()
	c.AddFunc("@every 120s", SpawnMonster) // TODO test
	c.Start()
	return c
}

// 地图刷怪
func SpawnMonster() {
	maps := com.GetMaps()
	db := com.GetDB()
	for mapIndex := range *maps {
		respawnInfos := com.GetMapRespawnInfos(mapIndex, db)
		for _, respawnInfo := range respawnInfos {
			existedCount := com.GetMapExistedMonsterCount(mapIndex, respawnInfo.MonsterIndex)
			if existedCount < respawnInfo.Count {
				com.MapAddMonster(mapIndex, respawnInfo, respawnInfo.Count-existedCount, db)
			}
		}
	}
}

// 怪物视野

// 玩家状态状态

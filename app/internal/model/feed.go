package model

import (
	"tiktok/app/global"
)

//tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
//Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
//if categoryIdentity != "" {
//tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
//Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
//}

func (*VideoDaoStruct) GetVideoFeed(lastTime int32) ([]VideoInfo, bool) {

	var result []VideoInfo

	global.MysqlDB.Debug().Raw("SELECT `users`.`id` AS `UserID`,`users`.`name`, `videos`.`id` AS `VideoID`,"+
		"`videos`.`play_url`, `videos`.`cover_url`,`videos`.`publish_time`,`videos`.`title` "+
		"FROM `videos` INNER JOIN `users` ON `users`.`id` = `videos`.`author_id` "+
		"WHERE `videos`.`publish_time` < ? ORDER BY `videos`.`publish_time` DESC LIMIT 10", lastTime).Scan(&result)

	if result == nil {
		return nil, false
	}
	return result, true
}

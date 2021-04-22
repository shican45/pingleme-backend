//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	_ = Repo.DB.AutoMigrate(&User{}, &Class{}, &EvaluationItemGrade{}, &EvaluationTable{}, &EvaluationItem{}, &Homework{},
		&ScoringItem{}, &JobSubmission{}, &Partner{}, &Performance{}, &PersonalBlogGrade{}, &Role{}, &Permission{},
		&Team{}, &TeamBlogGrade{})
}

package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Content string
}

func CreateTask(content string) (res bool, msg string) {
	if content == "" {
		return false, "Task 不能为空。"
	}
	var task Task
	db.First(&task, "content = ?", content)
	if task.Content != "" {
		db.Model(&task).Update("content", content)
		return true, "Task 重复，已更新最后编辑时间。"
	}
	db.Create(&Task{Content: content})
	return true, "成功！"
}

func QueryTaskByUpdateTime(startTime string, endTime string) []Task {
	var tasks []Task
	db.Where("created_at BETWEEN ? AND ?", startTime, endTime).Find(&tasks)
	return tasks
}

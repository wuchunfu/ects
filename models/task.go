package models

import (
	"time"
)

const (
	TASK_STATUS_NORMAL   = "normal"
	TASK_STATUS_DISABLED = "disabled"
	TASK_MODE_SHELL      = "shell"
	TASK_MODE_REST       = "rest"
	TASK_MODE_MAIL       = "mail"
	TASK_MODE_HOOK       = "hook"
)

type (
	Task struct {
		Id          string    `json:"id" validate:"-" xorm:"not null pk comment('用户ID') CHAR(36)"`
		TeamId      string    `json:"team_id" validate:"required" xorm:"not null comment('团队ID') index CHAR(36)"`
		Name        string    `json:"name" validate:"required" xorm:"not null comment('名称') VARCHAR(255)"`
		Content     string    `json:"content" validate:"required" xorm:"not null comment('内容') TEXT"`
		Mode        string    `json:"mode" validate:"required" xorm:"not null comment('执行方式') VARCHAR(255)"`
		Description string    `json:"description" validate:"-" xorm:"comment('描述') VARCHAR(255)"`
		CreatedAt   time.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
		UpdatedAt   time.Time `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
	}
)

// 定义模型的数据表名称
func (task *Task) TableName() string {
	return "tasks"
}

// 创建任务
func (task *Task) Store() error {
	_, err := Engine.Insert(task)
	return err
}

// 更新任务
func (task *Task) Update() error {
	_, err := Engine.Id(task.Id).Update(task)
	return err
}

// 删除任务
func (task *Task) Destroy() error {
	_, err := Engine.Delete(task)
	return err
}

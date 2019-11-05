package models

import (
	"bubble/dblayer"
)

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// TableName 指定表名
func (b *Todo) TableName() string {
	return "todo"
}

// GetAllTodo 查询所有待办事项的列表
func GetAllTodo(todo *[]Todo) (err error) {
	if err = dblayer.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

// CreateATodo 创建一个待办事项
func CreateATodo(todo *Todo) (err error) {
	if err = dblayer.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

// GetATodo 创建一个代办事项
func GetATodo(todo *Todo, id string) (err error) {
	if err := dblayer.DB.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

// UpdateATodo 更新一个代办事项
func UpdateATodo(todo *Todo, id string) (err error) {
	dblayer.DB.Save(todo)
	return nil
}

// DeleteATodo 删除一个代办事项
func DeleteATodo(todo *Todo, id string) (err error) {
	dblayer.DB.Where("id = ?", id).Delete(todo)
	return nil
}

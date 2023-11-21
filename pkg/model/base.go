/*
 * @Author: huangpengju 15713716933@163.com
 * @Date: 2023-09-27 15:02:23
 * @LastEditors: huangpengju 15713716933@163.com
 * @LastEditTime: 2023-11-21 15:03:48
 * @FilePath: \chitchat4.0\pkg\model\base.go
 * @Description:
 *
 * Copyright (c) 2023 by huangpengju, All Rights Reserved.
 */
package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"` // 软删除 soft delete
}

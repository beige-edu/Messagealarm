package service

import (
	"PrometheusAlert/models"
	"encoding/json"
	"time"
)

type CommonMessageRes struct {
	errcode int
	errmsg  string
}

func CreateMessagePushLog(appkey, name, sendType, message, sendResult string) error {
	status := 0 //默认发送成功
	switch sendType {
	case "dd":
		var ddMsgRel CommonMessageRes
		json.Unmarshal([]byte(sendResult), &ddMsgRel)
		if ddMsgRel.errcode != 0 && ddMsgRel.errmsg != "ok" {
			status = 1
		}
	default:
	}
	insertData := &models.MessagePushLogs{
		SendType: sendType,
		Source: name,
		Content: message,
		Status: int8(status),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := models.InsertMessagePushLog(insertData)
	return err
}

/**
 * @Auth: Nuts
 * @Date: 2021/5/31 2:29 下午
 */
package aliyun

import (
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/golang/protobuf/proto"
	"time"
)

type AliLog struct {
	projectName string
	store       string
	client      *sls.Client
	topic       string
}

func (c AliLog) NewLogsClient(project, endpoint, store, topic, key, secret string) *AliLog {
	client := &sls.Client{
		Endpoint:        endpoint,
		AccessKeyID:     key,
		AccessKeySecret: secret,
	}
	return &AliLog{
		projectName: project,
		store:       store,
		client:      client,
		topic:       topic,
	}
}

func (c AliLog) Logs(level string, l map[string]interface{}) error {
	if c.client == nil {
		return nil
	}

	var logs []*sls.Log

	var content []*sls.LogContent
	if level == "" {
		level = "info"
	}
	content = append(content, &sls.LogContent{
		Key:   proto.String("level"),
		Value: proto.String(level),
	})

	for i, n := range l {
		content = append(content, &sls.LogContent{
			Key:   proto.String(i),
			Value: proto.String(fmt.Sprintf("%v", n)),
		})
	}

	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}
	logs = append(logs, log)

	loggroup := &sls.LogGroup{
		Topic: proto.String(c.topic),
		//Source: proto.String("10.230.201.117"),
		Logs: logs,
	}
	err := c.client.PutLogs(c.projectName, c.store, loggroup)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	return err
}

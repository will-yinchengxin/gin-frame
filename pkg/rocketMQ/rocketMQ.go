package rocketMQ

import (
	"encoding/json"
	"frame/global"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"fmt"
	"reflect"
	"context"
)

func RocketClose() {
	global.RocketConsumer.Shutdown()
	global.RocketProducer.Shutdown()
}

//---------------------------------------------------消费者-----------------------------------------------------------

type Consumer struct {
}

func (c Consumer) Start() {
	global.RocketConsumer.Subscribe(global.RocketSetting.GroupName, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (c consumer.ConsumeResult, err error) {
		// 执行消费的逻辑
		for i := range msgs {
			req := Data{}
			err = json.Unmarshal(msgs[i].Body, &req)
			if err != nil {
				// 日志的记录
				global.Logger.Fatalf("Consumer err, cause %s", err.Error())
				fmt.Println(err, req)
				continue
			} else {
				msg := string(msgs[i].Body)
				//fmt.Println(msg)
				fmt.Println(msg)
				global.Logger.Debugf("Consumer err, cause %s", err.Error())
			}
		}
		c = consumer.ConsumeSuccess
		return
	})
	// 日志记录 ....
	err := global.RocketConsumer.Start()
	if err != nil {
		return
	}
}

//---------------------------------------------------生产者-----------------------------------------------------------

type Producer struct {
}

func (p *Producer) Send() (err error){
	data := GetData()
	v := reflect.ValueOf(data)

	// 修改值必须是指针类型否则不可行
	var sendData []*primitive.Message
	if v.Kind() == reflect.Slice {
		l := v.Len()
		for i := 0; i < l; i++ {
			value := v.Index(i) // Value of item
			infoByte, _ := json.Marshal(value.Interface())
			pmsg := &primitive.Message{
				Topic: global.RocketSetting.Topic,
				Body:  infoByte,
			}
			sendData = append(sendData, pmsg)
		}

	} else {
		//记录日志 ...
		infoByte, _ := json.Marshal(data)
		pmsg := &primitive.Message{
			Topic: global.RocketSetting.Topic,
			Body:  infoByte,
		}
		sendData = append(sendData, pmsg)
	}
	ctx := context.TODO() // 这里可以做成外部传递的形式
	err = global.RocketProducer.SendAsync(ctx, func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			return
		}
	}, sendData...)
	return nil
}

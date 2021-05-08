package main

import (
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/identity/idgen"
	"github.com/tiptok/gocomm/pkg/broker"
	"github.com/tiptok/gocomm/pkg/broker/kafkax"
	"github.com/tiptok/gocomm/pkg/broker/local"
	"github.com/tiptok/gocomm/pkg/broker/models"
	_ "github.com/tiptok/gopp/orm/go-pg"
	go_pg "github.com/tiptok/gopp/orm/go-pg"
	"github.com/tiptok/gopp/pkg/constant"
	"github.com/tiptok/gopp/pkg/domain"
	"time"
)

func main() {
	var (
		err error
	)
	producer := broker.NewMessageProducer(
		models.WithKafkaHost(constant.KAFKA_HOSTS),
		models.WithMessageProduceRepository(local.NewPgMessageProduceRepository(go_pg.DB, nil)),
		models.WithVersion("0.10.2.1"),
	)
	err = autoPublish(producer)
	if err != nil {
		return
	}
	time.Sleep(time.Second * 30)
}

// 自动组装消息发送
func autoPublish(producer *kafkax.MessageDirector) error {
	messages := []interface{}{
		domain.Users{
			Id:         20200129,
			Name:       "user_2020",
			Phone:      "18860183058",
			Status:     1,
			Roles:      []int64{7, 8},
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	}
	err := producer.Publish(constant.TopicUserLogin, messages, models.WithMessageProduceRepository(local.NewPgMessageProduceRepository(go_pg.DB, nil)))
	if err != nil {
		return err
	}
	return nil
}

// 手动组装消息发送
func manualPublish(producer *kafkax.MessageDirector) error {
	messages := []*models.Message{
		&models.Message{Id: idgen.Next(), Topic: constant.TopicUserLogin, MsgTime: time.Now().Unix(), Value: common.JsonAssertString(
			domain.Users{
				Id:         20200129,
				Name:       "user_2020",
				Phone:      "18860183058",
				Status:     1,
				Roles:      []int64{7, 8},
				CreateTime: time.Now(),
				UpdateTime: time.Now(),
			},
		), FinishStatus: 0},
	}
	err := producer.PublishMessages(messages, models.WithMessageProduceRepository(local.NewPgMessageProduceRepository(go_pg.DB, nil)))
	return err
}

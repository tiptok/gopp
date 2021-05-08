package main

import (
	"fmt"
	"github.com/tiptok/gocomm/common"
	"github.com/tiptok/gocomm/pkg/broker/kafkax"
	"github.com/tiptok/gocomm/pkg/broker/local"
	"github.com/tiptok/gocomm/pkg/broker/models"
	"github.com/tiptok/gocomm/pkg/log"
	go_pg "github.com/tiptok/gopp/orm/go-pg"
	"github.com/tiptok/gopp/pkg/constant"
	"github.com/tiptok/gopp/pkg/domain"
)

func main() {
	store := local.NewRedisMessageStore("gopp", constant.REDIS_HOST+":"+constant.REDIS_PORT, "")

	saramaConsumer := kafkax.NewSaramaConsumer(
		constant.KAFKA_HOSTS, constant.ServiceName,
		models.WithVersion("0.10.2.1"),
		models.WithConsumeRetryOption(3, 30, store),
	)
	saramaConsumer.WithTopicHandler(constant.TopicUserLogin, UserLoginHandler)
	saramaConsumer.WithMessageReceiver(local.NewPgMessageReceiverRepository(go_pg.DB, nil)) // 持久化

	err := saramaConsumer.StartConsume()
	if err != nil {
		log.Error(err)
	}
}

var value = 0

func UserLoginHandler(message interface{}) error {
	msg, ok := message.(*models.Message)
	if ok {
		var user = &domain.Users{}
		common.JsonUnmarshal(string(msg.Value), user)
		log.Info("消费消息:", msg.Id, msg.Topic, msg.Value)
		log.Info("登录用户信息:", user.Id, user.Name)
	}
	if value < 2 {
		value++
		return fmt.Errorf("handler user login error ->> id: %v", msg.Id)
	}
	return nil
}

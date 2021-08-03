package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"niceBackend/common/global"
)

func Mongo() {
	mongoCfg := global.NICE_CONFIG.Mongo
	var err error
	clientOptions := options.Client().ApplyURI(mongoCfg.Addr)
	// 连接到MongoDB
	global.NICE_Mongo, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		global.NICE_LOG.Error("mongo connect ping failed, err:", zap.Any("err", err))
	}
	// 检查连接
	err = global.NICE_Mongo.Ping(context.TODO(), nil)
	if err != nil {
		global.NICE_LOG.Error("mongo connect ping failed, err:", zap.Any("err", err))
	}
}
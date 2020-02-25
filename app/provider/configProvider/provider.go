package configProvider

import (
	"github.com/koushamad/goro-app/config"
	catchConfigProvider "github.com/koushamad/goro-cache/provider/configProvider"
	"github.com/koushamad/goro-core/app/conf"
	dbConfigProvider "github.com/koushamad/goro-db/provider/configProvider"
	queueConfigProvider "github.com/koushamad/goro-queue/provider/configProvider"
)

func Load() {
	catchConfigProvider.Load()
	queueConfigProvider.Load()
	dbConfigProvider.Load()

	conf.Add("App", config.App)
	conf.Add("Log", config.Log)
	conf.Add("Cache", config.Cache)
	conf.Add("Redis", config.Redis)
	conf.Add("DB", config.Database)
	conf.Add("Queue", config.Queue)
	conf.Add("Rabbit", config.RabbitMQ)
}

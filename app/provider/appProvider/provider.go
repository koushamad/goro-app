package appProvider

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-cache/app/cache"
	"github.com/koushamad/goro-db/app/database/sql"
	"github.com/koushamad/goro-queue/app/queue"
)

func Boot(egn *gin.Engine) {
	queue.Boot()
	cache.Boot()
	sql.Boot()
}

func Kill() {
	sql.Kill()
}

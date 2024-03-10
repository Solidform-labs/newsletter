package fiber_storage

import (
	"sync"

	"github.com/Solidform-labs/newsletter/configs"
	"github.com/gofiber/storage/postgres/v2"
)

var storage *postgres.Storage

var storageOnce sync.Once

func Create(config configs.Config) *postgres.Storage {
	storageOnce.Do(func() {
		storage = postgres.New(postgres.Config{
			ConnectionURI: config.DbConnectionString,
			Reset:         config.FiberStorageReset,
		})
	})
	return storage
}

func GetStorage() *postgres.Storage {
	if storage == nil {
		Create(configs.GetConfig())
	}
	return storage
}

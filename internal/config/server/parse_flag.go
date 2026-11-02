package server

import (
	"flag"
	"time"
)

const defaultStorageFilePath = "metrics.json"

func parseByFlag() Server {
	addr := flag.String("a", "localhost:8080", "адрес эндопинта HTTP сервера")
	storeInterval := flag.Int("i", 300, "интервал сохранения метрик на диск (сек); 0 — синхронно")
	fileStoragePath := flag.String("f", defaultStorageFilePath, "путь к файлу хранилища")
	restore := flag.Bool("r", false, "загружать метрики из файла при старте")
	flag.Parse()

	return Server{
		HTTP: HTTPServer{Addr: *addr},
		FileStorage: FileStorage{
			StoreInterval:   time.Duration(*storeInterval) * time.Second,
			FileStoragePath: *fileStoragePath,
			Restore:         *restore,
		},
	}
}

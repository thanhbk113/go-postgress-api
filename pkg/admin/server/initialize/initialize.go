package initialize

import "thanhbk113/internal/config"

func Init() {
	config.Init()
	database()
}

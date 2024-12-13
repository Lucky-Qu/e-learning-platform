package loading

import "e-learning-platform/config"

func Loading() {
	config.LoadConfig("config/config.json")
}

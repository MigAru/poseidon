package logger

import (
	"github.com/MigAru/poseidon/internal/config"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogrus(cfg *config.Config) (*logrus.Logger, func(), error) {
	log := logrus.New()
	if !cfg.DebugMode {
		file, err := os.OpenFile("bot.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			return nil, func() {}, err
		}
		log.SetOutput(file)
		return log, func() { file.Close() }, nil
	}

	return log, func() {}, nil
}

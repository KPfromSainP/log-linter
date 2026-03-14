package testdata

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}(logger)
	password := "secret"
	apiKey := "secretKey"
	token := "secretToken"

	// Zap messages
	logger.Info("Starting server on port 8080")
	logger.Error("Failed to connect to database")
	logger.Info("запуск сервера")
	logger.Error("ошибка подключения к базе данных")
	logger.Info("server started!🚀")
	logger.Error("connection failed!!!")
	logger.Warn("warning: something went wrong...")
	logger.Info("user password: " + password)
	logger.Debug("api_key=" + apiKey)
	logger.Info("token: " + token)

	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")
	logger.Info("starting server")
	logger.Error("failed to connect to database")
	logger.Info("server started")
	logger.Error("connection failed")
	logger.Warn("something went wrong")
	logger.Info("server started")
	logger.Error("connection failed")
	logger.Warn("something went wrong")
	logger.Info("user authenticated successfully")
	logger.Debug("api request completed")
	logger.Info("token validated")
}

package testdata

import (
	"log/slog"
)

func main() {
	password := "secret"
	apiKey := "secretKey"
	token := "secretToken"

	// Slog messages
	slog.Info("Starting server on port 8080")
	slog.Error("Failed to connect to database")
	slog.Info("запуск сервера")
	slog.Error("ошибка подключения к базе данных")
	slog.Info("server started!🚀")
	slog.Error("connection failed!!!")
	slog.Warn("warning: something went wrong...")
	slog.Info("user password: " + password)
	slog.Debug("api_key=" + apiKey)
	slog.Info("token: " + token)

	slog.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
	slog.Info("starting server")
	slog.Error("failed to connect to database")
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
	slog.Info("server started")
	slog.Error("connection failed")
	slog.Warn("something went wrong")
	slog.Info("user authenticated successfully")
	slog.Debug("api request completed")
	slog.Info("token' validated")

}

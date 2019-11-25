package conf

import (
	"github.com/spf13/viper"
)

func init() {

	// Logger Defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.encoding", "console")
	viper.SetDefault("logger.color", true)
	viper.SetDefault("logger.dev_mode", true)
	viper.SetDefault("logger.disable_caller", false)
	viper.SetDefault("logger.disable_stacktrace", true)

	// Pidfile
	viper.SetDefault("pidfile", "")

	// Profiler config
	viper.SetDefault("profiler.enabled", false)
	viper.SetDefault("profiler.host", "")
	viper.SetDefault("profiler.port", "6060")

	// Server Configuration
	viper.SetDefault("server.host", "")
	viper.SetDefault("server.port", "8900")
	viper.SetDefault("server.log_requests", true)
	viper.SetDefault("server.profiler_enabled", false)
	viper.SetDefault("server.profiler_path", "/debug")

	// Database Settings
	viper.SetDefault("storage.type", "postgres")
	viper.SetDefault("storage.username", "postgres")
	viper.SetDefault("storage.password", "password")
	viper.SetDefault("storage.host", "postgres")
	viper.SetDefault("storage.port", 5432)
	viper.SetDefault("storage.database", "vendopunkto")
	viper.SetDefault("storage.sslmode", "disable")
	// viper.SetDefault("storage.max_connections", 80)

	// Monero Settings
	viper.SetDefault("monero.wallet_rpc_url", "")

}

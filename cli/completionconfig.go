package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var completionConfig *viper.Viper

func initCompletionConfig() {
	completionConfig = viper.New()

	completionConfig.SetConfigName("completion")
	completionConfig.AddConfigPath(viper.GetString("config-directory"))

	// Write a blank cache if no file is already there. Later you can use
	// configs.SaveConfig() to write new values.
	filename := filepath.Join(viper.GetString("config-directory"), "completion.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := os.WriteFile(filename, []byte("{}"), 0600); err != nil {
			panic(err)
		}
	}

	err := completionConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

package initializer

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadViper() {
	viper.SetConfigType("env")  // set the file type
	viper.SetConfigFile(".env") // set the file name and path
	err := viper.ReadInConfig() // read the config file
	if err != nil {             // handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

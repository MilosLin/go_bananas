package config_test

import (
	"fmt"

	"github.com/MilosLin/go_bananas/core/config"
)

func ExampleInstance() {
	ip := config.Instance().GetString("ip")
	fmt.Printf("ip:%s \n", ip)
}

package redis_keys

import (
	"fmt"

	"xueyigou_demo/internal/utils"
)

func Dymmy() string {
	return fmt.Sprintf("topstack:dummy:%s", utils.UUIDWithoutDash())
}

func AccountForAuthenticate(userID string) string {
	return fmt.Sprintf("topstack:authenticate:%s", userID)
}

func Session(sessionKey string) string {
	return fmt.Sprintf("topstack:session:%s", sessionKey)
}

func Config(configKey string) string {
	return fmt.Sprintf("topstack:config:%s", configKey)
}

package api

import (
	"errors"
	"fmt"
)

func GetWeaponInfo(name string) (string, error) {
	if name == "" {
		return "", errors.New("需要武器名称")
	}
	return fmt.Sprintf("%s的合成", name), nil
}

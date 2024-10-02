package api

import (
	"errors"
	"fmt"
)

func GetEquipmentInfo(name string, game int) (string, error) {
	if name == "" {
		return "", errors.New("需要装备名称")
	}
	return fmt.Sprintf("%s的信息", name), nil
}

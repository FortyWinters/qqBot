package mh

import (
	"errors"
	"fmt"
)

func GetMonsterInfo(name string) (string, error) {
	if name == "" {
		return "", errors.New("需要怪物名称")
	}
    return fmt.Sprintf("%s的弱点", name), nil
}

func GetWeaponInfo(name string) (string, error) {
		if name == "" {
		return "", errors.New("需要武器名称")
	}
    return fmt.Sprintf("%s的合成", name), nil
}

func GetEquipmentInfo(name string) (string, error) {
		if name == "" {
		return "", errors.New("需要装备名称")
	}
    return fmt.Sprintf("%s的面板", name), nil
}
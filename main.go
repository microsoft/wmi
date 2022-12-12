package main

import (
	"fmt"

	wmierrors "github.com/microsoft/wmi/pkg/errors"
)

func main() {

	err := wmierrors.NewWMIError(32027)
	fmt.Println(err)
	fmt.Println(err.Error())

	fmt.Println(wmierrors.IsWMIError(err))
}

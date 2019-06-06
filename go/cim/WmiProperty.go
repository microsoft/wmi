package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type WmiProperty struct {
}

// Name
func (c WmiProperty) Name() string {
	panic("not implemented")

}

// Value
func (c WmiProperty) Value() string {
	panic("not implemented")

}

// Type
func (c WmiProperty) Type() wmi.WmiType {
	panic("not implemented")

}

// Flags
func (c WmiProperty) Flags() wmi.PropertyFlags {
	panic("not implemented")
}

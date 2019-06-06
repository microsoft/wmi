package cim;

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type CimWmiProperty struct {

}


// Name
func (c CimWmiProperty)  Name() string{
	panic("not implemented")

}

// Value
func (c CimWmiProperty)  	Value() string{
	panic("not implemented")

}
// Type
func (c CimWmiProperty)  	Type() wmi.WmiType{
	panic("not implemented")

}

// Flags
func (c CimWmiProperty)  	Flags() wmi.WmiPropertyFlags{
	panic("not implemented")
}
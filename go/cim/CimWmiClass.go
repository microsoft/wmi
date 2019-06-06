package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type CimWmiClass struct {
}

// ClassName
func (c CimWmiClass) ClassName() string {
	panic("not implemented")
}

// SuperClassName
func (c CimWmiClass) SuperClassName() string {

	panic("not implemented")
}

// ServerName
func (c CimWmiClass) ServerName() string {
	panic("not implemented")

}

// Namespace
func (c CimWmiClass) Namespace() string {
	panic("not implemented")

}

// SuperClass
func (c CimWmiClass) SuperClass() *wmi.WmiClass {
	panic("not implemented")

}

// Properties
func (c CimWmiClass) Properties() []string {
	panic("not implemented")

}

// Qualifiers
func (c CimWmiClass) Qualifiers() []string {
	panic("not implemented")

}

// Methods
func (c CimWmiClass) Methods() []string {
	panic("not implemented")

}

// MethodParameters
func (c CimWmiClass) MethodParameters(methodName string) []string {
	panic("not implemented")

}
func (c CimWmiClass) InvokeMethod(methodName string, methodParams []string, inputOptions string) (error, string) {
	panic("not implemented")

}

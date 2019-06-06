package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type WmiClass struct {
}

// ClassName
func (c WmiClass) ClassName() string {
	panic("not implemented")
}

// SuperClassName
func (c WmiClass) SuperClassName() string {

	panic("not implemented")
}

// ServerName
func (c WmiClass) ServerName() string {
	panic("not implemented")

}

// Namespace
func (c WmiClass) Namespace() string {
	panic("not implemented")

}

// SuperClass
func (c WmiClass) SuperClass() *wmi.Class {
	panic("not implemented")

}

// Properties
func (c WmiClass) Properties() []string {
	panic("not implemented")

}

// Qualifiers
func (c WmiClass) Qualifiers() []string {
	panic("not implemented")

}

// Methods
func (c WmiClass) Methods() []string {
	panic("not implemented")

}

// MethodParameters
func (c WmiClass) MethodParameters(methodName string) []string {
	panic("not implemented")

}
func (c WmiClass) InvokeMethod(methodName string, methodParams []string, inputOptions string) (error, string) {
	panic("not implemented")

}

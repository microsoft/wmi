package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

// WmiInstanceManager
type WmiInstanceManager struct {
}

// ServerName
func (m WmiInstanceManager) ServerName() string { panic("not implemented") }

// Namespace
func (m WmiInstanceManager) Namespace() string { panic("not implemented") }

// Credentials
func (m WmiInstanceManager) Credentials() *wmi.Credentials { panic("not implemented") }

// EnumerateInstances
func (m WmiInstanceManager) EnumerateInstances(className string) (*[]wmi.Instance, error) {
	panic("not implemented")
}

// QueryInstances
func (m WmiInstanceManager) QueryInstances(query string) (*[]wmi.Instance, error) {
	panic("not implemented")
}

// CreateInstance
func (m WmiInstanceManager) QueryInstancesEx(query wmi.Query) (*[]wmi.Instance, error) {
	panic("not implemented")
}

// ClassName
func (m WmiInstanceManager) CreateInstance(className string, propertyValues map[string]string) (*wmi.Instance, error) {
	panic("not implemented")
}

// GetInstance
func (m WmiInstanceManager) GetInstance(className string, propertyValues map[string]string) (*wmi.Instance, error) {
	panic("not implemented")
}

// GetClass
func (m WmiInstanceManager) GetClass(className string) (*WmiClass, error) { panic("not implemented") }

// EnumerateClasses
func (m WmiInstanceManager) EnumerateClasses() (*[]WmiClass, error) { panic("not implemented") }

// GetInstancesFromPaths
func (m WmiInstanceManager) GetInstancesFromPaths(pathArray []string) (*[]wmi.Instance, error) {
	panic("not implemented")
}

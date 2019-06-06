package cim

import (
	"github.com/microsoft/wmicodegen/go/wmi"
)

type WmiInstance struct {
}

// GetInstance
func (c WmiInstance) GetInstance() string {
	panic("not implemented")

}

// GetProperty
func (c WmiInstance) GetProperty() string {
	panic("not implemented")

}

// ResetProperty
func (c WmiInstance) ResetProperty() string {
	panic("not implemented")

}

// Class
func (c WmiInstance) Class() *wmi.Class {
	panic("not implemented")

}

// Class
func (c WmiInstance) EmbeddedInstance() string {
	panic("not implemented")

}

// InstanceManager
func (c WmiInstance) InstanceManager() *wmi.InstanceManager {
	panic("not implemented")

}

// Equals
func (c WmiInstance) Equals(*wmi.Instance) bool {
	panic("not implemented")

}

// Refresh
func (c WmiInstance) Refresh() {
	panic("not implemented")

}

// Commit
func (c WmiInstance) Commit() {
	panic("not implemented")

}

// Modify
func (c WmiInstance) Modify() {
	panic("not implemented")

}

// Delete
func (c WmiInstance) Delete() {
	panic("not implemented")

}

// InstancePath
func (c WmiInstance) InstancePath() string {
	panic("not implemented")

}

// InvokeMethod
func (c WmiInstance) InvokeMethod(methodName string, propertyValues map[string]string) (string, error) {
	panic("not implemented")

}

// GetRelated
func (c WmiInstance) GetRelated(resultClassName string) *[]wmi.Instance {
	panic("not implemented")

}

// Rel
func (c WmiInstance) GetRelatedEx(resultClassName, associatedClassName, resultRole, sourceRole string) *[]wmi.Instance {
	panic("not implemented")

}

// GetAssociated
func (c WmiInstance) GetAssociated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]wmi.Instance {

	panic("not implemented")
}

// EnumerateReferencingInstances
func (c WmiInstance) EnumerateReferencingInstances(associatedClassName, sourceRole string) *[]wmi.Instance {
	panic("not implemented")

}

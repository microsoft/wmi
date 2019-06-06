package cim

import (
	 "github.com/microsoft/wmicodegen/go/wmi"
)

type CimWmiInstance struct {
}

// GetInstance
func (c CimWmiInstance)  GetInstance() string {
	panic("not implemented")

}
// GetProperty
func (c CimWmiInstance)  	GetProperty() string {
	panic("not implemented")

}
// ResetProperty
func (c CimWmiInstance)  	ResetProperty() string{
	panic("not implemented")

}
// Class
func (c CimWmiInstance)  	Class() *wmi.WmiClass{
	panic("not implemented")

}
// Class
func (c CimWmiInstance)  	Class() string{
	panic("not implemented")

}
// InstanceManager
func (c CimWmiInstance)  	InstanceManager() *wmi.WmiInstanceManager{
	panic("not implemented")

}
// Equals
func (c CimWmiInstance)  	Equals(*wmi.WmiInstance) bool{
	panic("not implemented")

}
// Refresh
func (c CimWmiInstance)  	Refresh(){
	panic("not implemented")

}
// Commit
func (c CimWmiInstance)  	Commit(){
	panic("not implemented")

}
// Modify
func (c CimWmiInstance)  	Modify(){
	panic("not implemented")

}
// Delete
func (c CimWmiInstance)  	Delete(){
	panic("not implemented")

}
// InstancePath
func (c CimWmiInstance)  	InstancePath() string{
	panic("not implemented")

}
// InvokeMethod
func (c CimWmiInstance)  	InvokeMethod(methodName string, propertyValues map[string]string ) (string, error){
	panic("not implemented")

}
// GetRelated
func (c CimWmiInstance)  	GetRelated(resultClassName string) *[]wmi.WmiInstance{
	panic("not implemented")

}
// Rel
func (c CimWmiInstance)  	GetRelated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]wmi.WmiInstance{
	panic("not implemented")

}
// GetAssociated
func (c CimWmiInstance)  	GetAssociated(resultClassName, associatedClassName, resultRole, sourceRole string) *[]wmi.WmiInstance{

	panic("not implemented")
}
// EnumerateReferencingInstances
func (c CimWmiInstance)  	EnumerateReferencingInstances(associatedClassName, sourceRole string) *[]wmi.WmiInstance{
	panic("not implemented")

}
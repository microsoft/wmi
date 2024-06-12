// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_ResourceType struct
type MSCluster_ResourceType struct {
	*MSCluster_LogicalElement

	//
	AdminExtensions []string

	//
	DeadlockTimeout uint32

	//
	DeleteRequiresAllNodes bool

	//
	DisplayName string

	//
	DllName string

	//
	DumpLogQuery []string

	//
	DumpPolicy uint64

	//
	DumpServices []string

	//
	EnabledEventLogs []string

	//
	IsAlivePollInterval uint32

	//
	LocalQuorumCapable bool

	//
	LooksAlivePollInterval uint32

	//
	MaximumMonitors uint32

	//
	PendingTimeout uint32

	//
	PrivateProperties MSCluster_Property

	//
	QuorumCapable bool

	//
	RequiredDependencyClasses []uint32

	//
	RequiredDependencyTypes []string

	//
	ResourceClass uint32

	//
	WprProfiles []string

	//
	WprStartAfter uint64
}

func NewMSCluster_ResourceTypeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_ResourceType, err error) {
	tmp, err := NewMSCluster_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceType{
		MSCluster_LogicalElement: tmp,
	}
	return
}

func NewMSCluster_ResourceTypeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_ResourceType, err error) {
	tmp, err := NewMSCluster_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_ResourceType{
		MSCluster_LogicalElement: tmp,
	}
	return
}

// SetAdminExtensions sets the value of AdminExtensions for the instance
func (instance *MSCluster_ResourceType) SetPropertyAdminExtensions(value []string) (err error) {
	return instance.SetProperty("AdminExtensions", (value))
}

// GetAdminExtensions gets the value of AdminExtensions for the instance
func (instance *MSCluster_ResourceType) GetPropertyAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("AdminExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDeadlockTimeout sets the value of DeadlockTimeout for the instance
func (instance *MSCluster_ResourceType) SetPropertyDeadlockTimeout(value uint32) (err error) {
	return instance.SetProperty("DeadlockTimeout", (value))
}

// GetDeadlockTimeout gets the value of DeadlockTimeout for the instance
func (instance *MSCluster_ResourceType) GetPropertyDeadlockTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeadlockTimeout")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetDeleteRequiresAllNodes sets the value of DeleteRequiresAllNodes for the instance
func (instance *MSCluster_ResourceType) SetPropertyDeleteRequiresAllNodes(value bool) (err error) {
	return instance.SetProperty("DeleteRequiresAllNodes", (value))
}

// GetDeleteRequiresAllNodes gets the value of DeleteRequiresAllNodes for the instance
func (instance *MSCluster_ResourceType) GetPropertyDeleteRequiresAllNodes() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteRequiresAllNodes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSCluster_ResourceType) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSCluster_ResourceType) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDllName sets the value of DllName for the instance
func (instance *MSCluster_ResourceType) SetPropertyDllName(value string) (err error) {
	return instance.SetProperty("DllName", (value))
}

// GetDllName gets the value of DllName for the instance
func (instance *MSCluster_ResourceType) GetPropertyDllName() (value string, err error) {
	retValue, err := instance.GetProperty("DllName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetDumpLogQuery sets the value of DumpLogQuery for the instance
func (instance *MSCluster_ResourceType) SetPropertyDumpLogQuery(value []string) (err error) {
	return instance.SetProperty("DumpLogQuery", (value))
}

// GetDumpLogQuery gets the value of DumpLogQuery for the instance
func (instance *MSCluster_ResourceType) GetPropertyDumpLogQuery() (value []string, err error) {
	retValue, err := instance.GetProperty("DumpLogQuery")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDumpPolicy sets the value of DumpPolicy for the instance
func (instance *MSCluster_ResourceType) SetPropertyDumpPolicy(value uint64) (err error) {
	return instance.SetProperty("DumpPolicy", (value))
}

// GetDumpPolicy gets the value of DumpPolicy for the instance
func (instance *MSCluster_ResourceType) GetPropertyDumpPolicy() (value uint64, err error) {
	retValue, err := instance.GetProperty("DumpPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDumpServices sets the value of DumpServices for the instance
func (instance *MSCluster_ResourceType) SetPropertyDumpServices(value []string) (err error) {
	return instance.SetProperty("DumpServices", (value))
}

// GetDumpServices gets the value of DumpServices for the instance
func (instance *MSCluster_ResourceType) GetPropertyDumpServices() (value []string, err error) {
	retValue, err := instance.GetProperty("DumpServices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetEnabledEventLogs sets the value of EnabledEventLogs for the instance
func (instance *MSCluster_ResourceType) SetPropertyEnabledEventLogs(value []string) (err error) {
	return instance.SetProperty("EnabledEventLogs", (value))
}

// GetEnabledEventLogs gets the value of EnabledEventLogs for the instance
func (instance *MSCluster_ResourceType) GetPropertyEnabledEventLogs() (value []string, err error) {
	retValue, err := instance.GetProperty("EnabledEventLogs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetIsAlivePollInterval sets the value of IsAlivePollInterval for the instance
func (instance *MSCluster_ResourceType) SetPropertyIsAlivePollInterval(value uint32) (err error) {
	return instance.SetProperty("IsAlivePollInterval", (value))
}

// GetIsAlivePollInterval gets the value of IsAlivePollInterval for the instance
func (instance *MSCluster_ResourceType) GetPropertyIsAlivePollInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsAlivePollInterval")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetLocalQuorumCapable sets the value of LocalQuorumCapable for the instance
func (instance *MSCluster_ResourceType) SetPropertyLocalQuorumCapable(value bool) (err error) {
	return instance.SetProperty("LocalQuorumCapable", (value))
}

// GetLocalQuorumCapable gets the value of LocalQuorumCapable for the instance
func (instance *MSCluster_ResourceType) GetPropertyLocalQuorumCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalQuorumCapable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetLooksAlivePollInterval sets the value of LooksAlivePollInterval for the instance
func (instance *MSCluster_ResourceType) SetPropertyLooksAlivePollInterval(value uint32) (err error) {
	return instance.SetProperty("LooksAlivePollInterval", (value))
}

// GetLooksAlivePollInterval gets the value of LooksAlivePollInterval for the instance
func (instance *MSCluster_ResourceType) GetPropertyLooksAlivePollInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("LooksAlivePollInterval")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetMaximumMonitors sets the value of MaximumMonitors for the instance
func (instance *MSCluster_ResourceType) SetPropertyMaximumMonitors(value uint32) (err error) {
	return instance.SetProperty("MaximumMonitors", (value))
}

// GetMaximumMonitors gets the value of MaximumMonitors for the instance
func (instance *MSCluster_ResourceType) GetPropertyMaximumMonitors() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumMonitors")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPendingTimeout sets the value of PendingTimeout for the instance
func (instance *MSCluster_ResourceType) SetPropertyPendingTimeout(value uint32) (err error) {
	return instance.SetProperty("PendingTimeout", (value))
}

// GetPendingTimeout gets the value of PendingTimeout for the instance
func (instance *MSCluster_ResourceType) GetPropertyPendingTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingTimeout")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPrivateProperties sets the value of PrivateProperties for the instance
func (instance *MSCluster_ResourceType) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", (value))
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_ResourceType) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
	retValue, err := instance.GetProperty("PrivateProperties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSCluster_Property)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSCluster_Property is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSCluster_Property(valuetmp)

	return
}

// SetQuorumCapable sets the value of QuorumCapable for the instance
func (instance *MSCluster_ResourceType) SetPropertyQuorumCapable(value bool) (err error) {
	return instance.SetProperty("QuorumCapable", (value))
}

// GetQuorumCapable gets the value of QuorumCapable for the instance
func (instance *MSCluster_ResourceType) GetPropertyQuorumCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("QuorumCapable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRequiredDependencyClasses sets the value of RequiredDependencyClasses for the instance
func (instance *MSCluster_ResourceType) SetPropertyRequiredDependencyClasses(value []uint32) (err error) {
	return instance.SetProperty("RequiredDependencyClasses", (value))
}

// GetRequiredDependencyClasses gets the value of RequiredDependencyClasses for the instance
func (instance *MSCluster_ResourceType) GetPropertyRequiredDependencyClasses() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RequiredDependencyClasses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetRequiredDependencyTypes sets the value of RequiredDependencyTypes for the instance
func (instance *MSCluster_ResourceType) SetPropertyRequiredDependencyTypes(value []string) (err error) {
	return instance.SetProperty("RequiredDependencyTypes", (value))
}

// GetRequiredDependencyTypes gets the value of RequiredDependencyTypes for the instance
func (instance *MSCluster_ResourceType) GetPropertyRequiredDependencyTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredDependencyTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetResourceClass sets the value of ResourceClass for the instance
func (instance *MSCluster_ResourceType) SetPropertyResourceClass(value uint32) (err error) {
	return instance.SetProperty("ResourceClass", (value))
}

// GetResourceClass gets the value of ResourceClass for the instance
func (instance *MSCluster_ResourceType) GetPropertyResourceClass() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResourceClass")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetWprProfiles sets the value of WprProfiles for the instance
func (instance *MSCluster_ResourceType) SetPropertyWprProfiles(value []string) (err error) {
	return instance.SetProperty("WprProfiles", (value))
}

// GetWprProfiles gets the value of WprProfiles for the instance
func (instance *MSCluster_ResourceType) GetPropertyWprProfiles() (value []string, err error) {
	retValue, err := instance.GetProperty("WprProfiles")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetWprStartAfter sets the value of WprStartAfter for the instance
func (instance *MSCluster_ResourceType) SetPropertyWprStartAfter(value uint64) (err error) {
	return instance.SetProperty("WprStartAfter", (value))
}

// GetWprStartAfter gets the value of WprStartAfter for the instance
func (instance *MSCluster_ResourceType) GetPropertyWprStartAfter() (value uint64, err error) {
	retValue, err := instance.GetProperty("WprStartAfter")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

//

// <param name="DisplayName" type="string "></param>
// <param name="DLLName" type="string "></param>
// <param name="IsAlivePollInterval" type="uint32 "></param>
// <param name="LooksAlivePollInterval" type="uint32 "></param>
// <param name="Name" type="string "></param>
func (instance *MSCluster_ResourceType) CreateResourceType( /* IN */ Name string,
	/* IN */ DisplayName string,
	/* IN */ DLLName string,
	/* IN */ LooksAlivePollInterval uint32,
	/* IN */ IsAlivePollInterval uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("CreateResourceType", Name, DisplayName, DLLName, LooksAlivePollInterval, IsAlivePollInterval)
	if err != nil {
		return
	}
	return

}

//

// <param name="Reason" type="string "></param>
func (instance *MSCluster_ResourceType) DeleteResourceType( /* OPTIONAL IN */ Reason string) (err error) {
	_, err = instance.InvokeMethodWithReturn("DeleteResourceType", Reason)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeNames" type="string []"></param>
func (instance *MSCluster_ResourceType) GetPossibleOwners( /* OUT */ NodeNames []string) (err error) {
	_, err = instance.InvokeMethod("GetPossibleOwners")
	if err != nil {
		return
	}
	return

}

//

// <param name="ControlCode" type="int32 "></param>
// <param name="InputBuffer" type="uint8 []"></param>
// <param name="Reason" type="string "></param>

// <param name="OutputBuffer" type="uint8 []"></param>
// <param name="OutputBufferSize" type="int32 "></param>
func (instance *MSCluster_ResourceType) ExecuteResourceTypeControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32,
	/* OPTIONAL IN */ Reason string) (err error) {
	_, err = instance.InvokeMethod("ExecuteResourceTypeControl", ControlCode, InputBuffer, Reason)
	if err != nil {
		return
	}
	return

}

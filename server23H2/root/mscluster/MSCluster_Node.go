// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_Node struct
type MSCluster_Node struct {
	*CIM_UnitaryComputerSystem

	//
	BuildNumber uint32

	//
	Characteristics uint32

	//
	CSDVersion string

	//
	DetectedCloudPlatform uint32

	//
	DrainErrorCode uint32

	//
	DynamicWeight uint32

	//
	FaultDomain []string

	//
	FaultDomainId string

	//
	Flags uint32

	//
	Id string

	//
	MajorVersion uint32

	//
	MinorVersion uint32

	//
	NeedsPreventQuorum uint32

	//
	NodeDrainStatus uint32

	//
	NodeDrainTarget string

	//
	NodeFailbackStatus uint32

	//
	NodeHighestVersion uint32

	//
	NodeInstanceID string

	//
	NodeLowestVersion uint32

	//
	NodeWeight uint32

	//
	PrivateProperties MSCluster_Property

	//
	State uint32

	//
	StatusInformation uint32

	//
	UniqueID string
}

func NewMSCluster_NodeEx1(instance *cim.WmiInstance) (newInstance *MSCluster_Node, err error) {
	tmp, err := NewCIM_UnitaryComputerSystemEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Node{
		CIM_UnitaryComputerSystem: tmp,
	}
	return
}

func NewMSCluster_NodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_Node, err error) {
	tmp, err := NewCIM_UnitaryComputerSystemEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Node{
		CIM_UnitaryComputerSystem: tmp,
	}
	return
}

// SetBuildNumber sets the value of BuildNumber for the instance
func (instance *MSCluster_Node) SetPropertyBuildNumber(value uint32) (err error) {
	return instance.SetProperty("BuildNumber", (value))
}

// GetBuildNumber gets the value of BuildNumber for the instance
func (instance *MSCluster_Node) GetPropertyBuildNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("BuildNumber")
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

// SetCharacteristics sets the value of Characteristics for the instance
func (instance *MSCluster_Node) SetPropertyCharacteristics(value uint32) (err error) {
	return instance.SetProperty("Characteristics", (value))
}

// GetCharacteristics gets the value of Characteristics for the instance
func (instance *MSCluster_Node) GetPropertyCharacteristics() (value uint32, err error) {
	retValue, err := instance.GetProperty("Characteristics")
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

// SetCSDVersion sets the value of CSDVersion for the instance
func (instance *MSCluster_Node) SetPropertyCSDVersion(value string) (err error) {
	return instance.SetProperty("CSDVersion", (value))
}

// GetCSDVersion gets the value of CSDVersion for the instance
func (instance *MSCluster_Node) GetPropertyCSDVersion() (value string, err error) {
	retValue, err := instance.GetProperty("CSDVersion")
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

// SetDetectedCloudPlatform sets the value of DetectedCloudPlatform for the instance
func (instance *MSCluster_Node) SetPropertyDetectedCloudPlatform(value uint32) (err error) {
	return instance.SetProperty("DetectedCloudPlatform", (value))
}

// GetDetectedCloudPlatform gets the value of DetectedCloudPlatform for the instance
func (instance *MSCluster_Node) GetPropertyDetectedCloudPlatform() (value uint32, err error) {
	retValue, err := instance.GetProperty("DetectedCloudPlatform")
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

// SetDrainErrorCode sets the value of DrainErrorCode for the instance
func (instance *MSCluster_Node) SetPropertyDrainErrorCode(value uint32) (err error) {
	return instance.SetProperty("DrainErrorCode", (value))
}

// GetDrainErrorCode gets the value of DrainErrorCode for the instance
func (instance *MSCluster_Node) GetPropertyDrainErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DrainErrorCode")
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

// SetDynamicWeight sets the value of DynamicWeight for the instance
func (instance *MSCluster_Node) SetPropertyDynamicWeight(value uint32) (err error) {
	return instance.SetProperty("DynamicWeight", (value))
}

// GetDynamicWeight gets the value of DynamicWeight for the instance
func (instance *MSCluster_Node) GetPropertyDynamicWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicWeight")
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

// SetFaultDomain sets the value of FaultDomain for the instance
func (instance *MSCluster_Node) SetPropertyFaultDomain(value []string) (err error) {
	return instance.SetProperty("FaultDomain", (value))
}

// GetFaultDomain gets the value of FaultDomain for the instance
func (instance *MSCluster_Node) GetPropertyFaultDomain() (value []string, err error) {
	retValue, err := instance.GetProperty("FaultDomain")
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

// SetFaultDomainId sets the value of FaultDomainId for the instance
func (instance *MSCluster_Node) SetPropertyFaultDomainId(value string) (err error) {
	return instance.SetProperty("FaultDomainId", (value))
}

// GetFaultDomainId gets the value of FaultDomainId for the instance
func (instance *MSCluster_Node) GetPropertyFaultDomainId() (value string, err error) {
	retValue, err := instance.GetProperty("FaultDomainId")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSCluster_Node) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSCluster_Node) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetId sets the value of Id for the instance
func (instance *MSCluster_Node) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_Node) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetMajorVersion sets the value of MajorVersion for the instance
func (instance *MSCluster_Node) SetPropertyMajorVersion(value uint32) (err error) {
	return instance.SetProperty("MajorVersion", (value))
}

// GetMajorVersion gets the value of MajorVersion for the instance
func (instance *MSCluster_Node) GetPropertyMajorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("MajorVersion")
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

// SetMinorVersion sets the value of MinorVersion for the instance
func (instance *MSCluster_Node) SetPropertyMinorVersion(value uint32) (err error) {
	return instance.SetProperty("MinorVersion", (value))
}

// GetMinorVersion gets the value of MinorVersion for the instance
func (instance *MSCluster_Node) GetPropertyMinorVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinorVersion")
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

// SetNeedsPreventQuorum sets the value of NeedsPreventQuorum for the instance
func (instance *MSCluster_Node) SetPropertyNeedsPreventQuorum(value uint32) (err error) {
	return instance.SetProperty("NeedsPreventQuorum", (value))
}

// GetNeedsPreventQuorum gets the value of NeedsPreventQuorum for the instance
func (instance *MSCluster_Node) GetPropertyNeedsPreventQuorum() (value uint32, err error) {
	retValue, err := instance.GetProperty("NeedsPreventQuorum")
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

// SetNodeDrainStatus sets the value of NodeDrainStatus for the instance
func (instance *MSCluster_Node) SetPropertyNodeDrainStatus(value uint32) (err error) {
	return instance.SetProperty("NodeDrainStatus", (value))
}

// GetNodeDrainStatus gets the value of NodeDrainStatus for the instance
func (instance *MSCluster_Node) GetPropertyNodeDrainStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeDrainStatus")
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

// SetNodeDrainTarget sets the value of NodeDrainTarget for the instance
func (instance *MSCluster_Node) SetPropertyNodeDrainTarget(value string) (err error) {
	return instance.SetProperty("NodeDrainTarget", (value))
}

// GetNodeDrainTarget gets the value of NodeDrainTarget for the instance
func (instance *MSCluster_Node) GetPropertyNodeDrainTarget() (value string, err error) {
	retValue, err := instance.GetProperty("NodeDrainTarget")
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

// SetNodeFailbackStatus sets the value of NodeFailbackStatus for the instance
func (instance *MSCluster_Node) SetPropertyNodeFailbackStatus(value uint32) (err error) {
	return instance.SetProperty("NodeFailbackStatus", (value))
}

// GetNodeFailbackStatus gets the value of NodeFailbackStatus for the instance
func (instance *MSCluster_Node) GetPropertyNodeFailbackStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeFailbackStatus")
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

// SetNodeHighestVersion sets the value of NodeHighestVersion for the instance
func (instance *MSCluster_Node) SetPropertyNodeHighestVersion(value uint32) (err error) {
	return instance.SetProperty("NodeHighestVersion", (value))
}

// GetNodeHighestVersion gets the value of NodeHighestVersion for the instance
func (instance *MSCluster_Node) GetPropertyNodeHighestVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeHighestVersion")
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

// SetNodeInstanceID sets the value of NodeInstanceID for the instance
func (instance *MSCluster_Node) SetPropertyNodeInstanceID(value string) (err error) {
	return instance.SetProperty("NodeInstanceID", (value))
}

// GetNodeInstanceID gets the value of NodeInstanceID for the instance
func (instance *MSCluster_Node) GetPropertyNodeInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("NodeInstanceID")
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

// SetNodeLowestVersion sets the value of NodeLowestVersion for the instance
func (instance *MSCluster_Node) SetPropertyNodeLowestVersion(value uint32) (err error) {
	return instance.SetProperty("NodeLowestVersion", (value))
}

// GetNodeLowestVersion gets the value of NodeLowestVersion for the instance
func (instance *MSCluster_Node) GetPropertyNodeLowestVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeLowestVersion")
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

// SetNodeWeight sets the value of NodeWeight for the instance
func (instance *MSCluster_Node) SetPropertyNodeWeight(value uint32) (err error) {
	return instance.SetProperty("NodeWeight", (value))
}

// GetNodeWeight gets the value of NodeWeight for the instance
func (instance *MSCluster_Node) GetPropertyNodeWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeWeight")
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
func (instance *MSCluster_Node) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", (value))
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_Node) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
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

// SetState sets the value of State for the instance
func (instance *MSCluster_Node) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSCluster_Node) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
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

// SetStatusInformation sets the value of StatusInformation for the instance
func (instance *MSCluster_Node) SetPropertyStatusInformation(value uint32) (err error) {
	return instance.SetProperty("StatusInformation", (value))
}

// GetStatusInformation gets the value of StatusInformation for the instance
func (instance *MSCluster_Node) GetPropertyStatusInformation() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusInformation")
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

// SetUniqueID sets the value of UniqueID for the instance
func (instance *MSCluster_Node) SetPropertyUniqueID(value string) (err error) {
	return instance.SetProperty("UniqueID", (value))
}

// GetUniqueID gets the value of UniqueID for the instance
func (instance *MSCluster_Node) GetPropertyUniqueID() (value string, err error) {
	retValue, err := instance.GetProperty("UniqueID")
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

//

// <param name="DrainType" type="uint32 "></param>
// <param name="Reason" type="string "></param>
// <param name="TargetNode" type="string "></param>
func (instance *MSCluster_Node) Pause( /* IN */ DrainType uint32,
	/* IN */ TargetNode string,
	/* OPTIONAL IN */ Reason string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Pause", DrainType, TargetNode, Reason)
	if err != nil {
		return
	}
	return

}

//

// <param name="FailbackType" type="uint32 "></param>
// <param name="Reason" type="string "></param>
func (instance *MSCluster_Node) Resume( /* IN */ FailbackType uint32,
	/* OPTIONAL IN */ Reason string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Resume", FailbackType, Reason)
	if err != nil {
		return
	}
	return

}

//

// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *MSCluster_Node) WillEvictLoseQuorum( /* OPTIONAL IN */ Reason string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("WillEvictLoseQuorum", Reason)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="Reason" type="string "></param>

// <param name="ReturnValue" type="bool "></param>
func (instance *MSCluster_Node) WillOfflineLoseQuorum( /* OPTIONAL IN */ Reason string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("WillOfflineLoseQuorum", Reason)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="ControlCode" type="int32 "></param>
// <param name="InputBuffer" type="uint8 []"></param>
// <param name="Reason" type="string "></param>

// <param name="OutputBuffer" type="uint8 []"></param>
// <param name="OutputBufferSize" type="int32 "></param>
func (instance *MSCluster_Node) ExecuteNodeControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32,
	/* OPTIONAL IN */ Reason string) (err error) {
	_, err = instance.InvokeMethod("ExecuteNodeControl", ControlCode, InputBuffer, Reason)
	if err != nil {
		return
	}
	return

}

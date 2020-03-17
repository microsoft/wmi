// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_RegisteredProfile struct
type CIM_RegisteredProfile struct {
	CIM_ManagedElement

	//
	AdvertiseTypeDescriptions []string

	//
	AdvertiseTypes []RegisteredProfile_AdvertiseTypes

	//
	ImplementedFeatures []string

	//
	OtherRegisteredOrganization string

	//
	RegisteredName string

	//
	RegisteredOrganization RegisteredProfile_RegisteredOrganization

	//
	RegisteredVersion string
}

// SetAdvertiseTypeDescriptions sets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredProfile) SetPropertyAdvertiseTypeDescriptions(value []string) (err error) {
	return instance.SetProperty("AdvertiseTypeDescriptions", value)
}

// GetAdvertiseTypeDescriptions gets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredProfile) GetPropertyAdvertiseTypeDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("AdvertiseTypeDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAdvertiseTypes sets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredProfile) SetPropertyAdvertiseTypes(value []RegisteredProfile_AdvertiseTypes) (err error) {
	return instance.SetProperty("AdvertiseTypes", value)
}

// GetAdvertiseTypes gets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredProfile) GetPropertyAdvertiseTypes() (value []RegisteredProfile_AdvertiseTypes, err error) {
	retValue, err := instance.GetProperty("AdvertiseTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]RegisteredProfile_AdvertiseTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetImplementedFeatures sets the value of ImplementedFeatures for the instance
func (instance *CIM_RegisteredProfile) SetPropertyImplementedFeatures(value []string) (err error) {
	return instance.SetProperty("ImplementedFeatures", value)
}

// GetImplementedFeatures gets the value of ImplementedFeatures for the instance
func (instance *CIM_RegisteredProfile) GetPropertyImplementedFeatures() (value []string, err error) {
	retValue, err := instance.GetProperty("ImplementedFeatures")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherRegisteredOrganization sets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) SetPropertyOtherRegisteredOrganization(value string) (err error) {
	return instance.SetProperty("OtherRegisteredOrganization", value)
}

// GetOtherRegisteredOrganization gets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) GetPropertyOtherRegisteredOrganization() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRegisteredOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredName sets the value of RegisteredName for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredName(value string) (err error) {
	return instance.SetProperty("RegisteredName", value)
}

// GetRegisteredName gets the value of RegisteredName for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredName() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredOrganization sets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredOrganization(value RegisteredProfile_RegisteredOrganization) (err error) {
	return instance.SetProperty("RegisteredOrganization", value)
}

// GetRegisteredOrganization gets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredOrganization() (value RegisteredProfile_RegisteredOrganization, err error) {
	retValue, err := instance.GetProperty("RegisteredOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(RegisteredProfile_RegisteredOrganization)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredVersion sets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredVersion(value string) (err error) {
	return instance.SetProperty("RegisteredVersion", value)
}

// GetRegisteredVersion gets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredVersion() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="EnumerationContext" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_RegisteredProfile) CloseConformantInstances( /* IN */ EnumerationContext string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CloseConformantInstances", EnumerationContext)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ContinueOnError" type="bool "></param>
// <param name="IncludedPropertyList" type="string []"></param>
// <param name="MaxObjectCount" type="uint32 "></param>
// <param name="OperationTimeout" type="uint32 "></param>
// <param name="ResultClass" type="string "></param>

// <param name="EndOfSequence" type="bool "></param>
// <param name="EnumerationContext" type="string "></param>
// <param name="InstanceType" type="RegisteredProfile_InstanceType []"></param>
// <param name="InstanceWithPathList" type="CIM_ManagedElement []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_RegisteredProfile) OpenConformantInstances( /* IN */ ResultClass string,
	/* IN */ IncludedPropertyList []string,
	/* IN */ OperationTimeout uint32,
	/* IN */ ContinueOnError bool,
	/* IN */ MaxObjectCount uint32,
	/* OUT */ EnumerationContext string,
	/* OUT */ EndOfSequence bool,
	/* OUT */ InstanceType []RegisteredProfile_InstanceType,
	/* OUT */ InstanceWithPathList []CIM_ManagedElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("OpenConformantInstances", ResultClass, IncludedPropertyList, OperationTimeout, ContinueOnError, MaxObjectCount)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EnumerationContext" type="string "></param>
// <param name="MaxObjectCount" type="uint32 "></param>

// <param name="EndOfSequence" type="bool "></param>
// <param name="EnumerationContext" type="string "></param>
// <param name="InstanceType" type="RegisteredProfile_InstanceType []"></param>
// <param name="InstanceWithPathList" type="CIM_ManagedElement []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_RegisteredProfile) PullConformantInstances( /* IN */ MaxObjectCount uint32,
	/* IN/OUT */ EnumerationContext string,
	/* OUT */ EndOfSequence bool,
	/* OUT */ InstanceType []RegisteredProfile_InstanceType,
	/* OUT */ InstanceWithPathList []CIM_ManagedElement) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("PullConformantInstances", MaxObjectCount)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

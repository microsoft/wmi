// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_ResourceGroup struct
type MSCluster_ResourceGroup struct {
	MSCluster_LogicalElement

	//
	AntiAffinityClassNames []string

	//
	AutoFailbackType uint32

	//
	CCFEpoch uint64

	//
	CCFEpochHigh uint64

	//
	ColdStartSetting uint32

	//
	DefaultOwner uint32

	//
	FailbackWindowEnd int32

	//
	FailbackWindowStart int32

	//
	FailoverPeriod uint32

	//
	FailoverThreshold uint32

	//
	FaultDomain uint32

	//
	GroupType uint32

	//
	Id string

	//
	IsCore bool

	//
	OwnerNode string

	//
	PersistentState bool

	//
	PlacementOptions uint32

	//
	PreferredSite []string

	//
	Priority uint32

	//
	PrivateProperties MSCluster_Property

	//
	ResiliencyPeriod uint32

	//
	State uint32

	//
	StatusInformation uint64

	//
	UpdateDomain uint32
}

// SetAntiAffinityClassNames sets the value of AntiAffinityClassNames for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyAntiAffinityClassNames(value []string) (err error) {
	return instance.SetProperty("AntiAffinityClassNames", value)
}

// GetAntiAffinityClassNames gets the value of AntiAffinityClassNames for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyAntiAffinityClassNames() (value []string, err error) {
	retValue, err := instance.GetProperty("AntiAffinityClassNames")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoFailbackType sets the value of AutoFailbackType for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyAutoFailbackType(value uint32) (err error) {
	return instance.SetProperty("AutoFailbackType", value)
}

// GetAutoFailbackType gets the value of AutoFailbackType for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyAutoFailbackType() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoFailbackType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCCFEpoch sets the value of CCFEpoch for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyCCFEpoch(value uint64) (err error) {
	return instance.SetProperty("CCFEpoch", value)
}

// GetCCFEpoch gets the value of CCFEpoch for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyCCFEpoch() (value uint64, err error) {
	retValue, err := instance.GetProperty("CCFEpoch")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCCFEpochHigh sets the value of CCFEpochHigh for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyCCFEpochHigh(value uint64) (err error) {
	return instance.SetProperty("CCFEpochHigh", value)
}

// GetCCFEpochHigh gets the value of CCFEpochHigh for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyCCFEpochHigh() (value uint64, err error) {
	retValue, err := instance.GetProperty("CCFEpochHigh")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetColdStartSetting sets the value of ColdStartSetting for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyColdStartSetting(value uint32) (err error) {
	return instance.SetProperty("ColdStartSetting", value)
}

// GetColdStartSetting gets the value of ColdStartSetting for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyColdStartSetting() (value uint32, err error) {
	retValue, err := instance.GetProperty("ColdStartSetting")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDefaultOwner sets the value of DefaultOwner for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyDefaultOwner(value uint32) (err error) {
	return instance.SetProperty("DefaultOwner", value)
}

// GetDefaultOwner gets the value of DefaultOwner for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyDefaultOwner() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultOwner")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailbackWindowEnd sets the value of FailbackWindowEnd for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyFailbackWindowEnd(value int32) (err error) {
	return instance.SetProperty("FailbackWindowEnd", value)
}

// GetFailbackWindowEnd gets the value of FailbackWindowEnd for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyFailbackWindowEnd() (value int32, err error) {
	retValue, err := instance.GetProperty("FailbackWindowEnd")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailbackWindowStart sets the value of FailbackWindowStart for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyFailbackWindowStart(value int32) (err error) {
	return instance.SetProperty("FailbackWindowStart", value)
}

// GetFailbackWindowStart gets the value of FailbackWindowStart for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyFailbackWindowStart() (value int32, err error) {
	retValue, err := instance.GetProperty("FailbackWindowStart")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailoverPeriod sets the value of FailoverPeriod for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyFailoverPeriod(value uint32) (err error) {
	return instance.SetProperty("FailoverPeriod", value)
}

// GetFailoverPeriod gets the value of FailoverPeriod for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyFailoverPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailoverPeriod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailoverThreshold sets the value of FailoverThreshold for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyFailoverThreshold(value uint32) (err error) {
	return instance.SetProperty("FailoverThreshold", value)
}

// GetFailoverThreshold gets the value of FailoverThreshold for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyFailoverThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("FailoverThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFaultDomain sets the value of FaultDomain for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyFaultDomain(value uint32) (err error) {
	return instance.SetProperty("FaultDomain", value)
}

// GetFaultDomain gets the value of FaultDomain for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyFaultDomain() (value uint32, err error) {
	retValue, err := instance.GetProperty("FaultDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGroupType sets the value of GroupType for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyGroupType(value uint32) (err error) {
	return instance.SetProperty("GroupType", value)
}

// GetGroupType gets the value of GroupType for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyGroupType() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetId sets the value of Id for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsCore sets the value of IsCore for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyIsCore(value bool) (err error) {
	return instance.SetProperty("IsCore", value)
}

// GetIsCore gets the value of IsCore for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyIsCore() (value bool, err error) {
	retValue, err := instance.GetProperty("IsCore")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwnerNode sets the value of OwnerNode for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyOwnerNode(value string) (err error) {
	return instance.SetProperty("OwnerNode", value)
}

// GetOwnerNode gets the value of OwnerNode for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyOwnerNode() (value string, err error) {
	retValue, err := instance.GetProperty("OwnerNode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPersistentState sets the value of PersistentState for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyPersistentState(value bool) (err error) {
	return instance.SetProperty("PersistentState", value)
}

// GetPersistentState gets the value of PersistentState for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyPersistentState() (value bool, err error) {
	retValue, err := instance.GetProperty("PersistentState")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPlacementOptions sets the value of PlacementOptions for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyPlacementOptions(value uint32) (err error) {
	return instance.SetProperty("PlacementOptions", value)
}

// GetPlacementOptions gets the value of PlacementOptions for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyPlacementOptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("PlacementOptions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreferredSite sets the value of PreferredSite for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyPreferredSite(value []string) (err error) {
	return instance.SetProperty("PreferredSite", value)
}

// GetPreferredSite gets the value of PreferredSite for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyPreferredSite() (value []string, err error) {
	retValue, err := instance.GetProperty("PreferredSite")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriority sets the value of Priority for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyPriority(value uint32) (err error) {
	return instance.SetProperty("Priority", value)
}

// GetPriority gets the value of Priority for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("Priority")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrivateProperties sets the value of PrivateProperties for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", value)
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
	retValue, err := instance.GetProperty("PrivateProperties")
	if err != nil {
		return
	}
	value, ok := retValue.(MSCluster_Property)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResiliencyPeriod sets the value of ResiliencyPeriod for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyResiliencyPeriod(value uint32) (err error) {
	return instance.SetProperty("ResiliencyPeriod", value)
}

// GetResiliencyPeriod gets the value of ResiliencyPeriod for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyResiliencyPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResiliencyPeriod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusInformation sets the value of StatusInformation for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyStatusInformation(value uint64) (err error) {
	return instance.SetProperty("StatusInformation", value)
}

// GetStatusInformation gets the value of StatusInformation for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyStatusInformation() (value uint64, err error) {
	retValue, err := instance.GetProperty("StatusInformation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpdateDomain sets the value of UpdateDomain for the instance
func (instance *MSCluster_ResourceGroup) SetPropertyUpdateDomain(value uint32) (err error) {
	return instance.SetProperty("UpdateDomain", value)
}

// GetUpdateDomain gets the value of UpdateDomain for the instance
func (instance *MSCluster_ResourceGroup) GetPropertyUpdateDomain() (value uint32, err error) {
	retValue, err := instance.GetProperty("UpdateDomain")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Flags" type="uint32 "></param>
// <param name="TimeOut" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) BringOnline( /* IN */ TimeOut uint32,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("BringOnline", TimeOut, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Parameters" type="MSCluster_Property "></param>
// <param name="TimeOut" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) TakeOffline( /* IN */ TimeOut uint32,
	/* IN */ Parameters MSCluster_Property,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("TakeOffline", TimeOut, Parameters, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Parameters" type="uint8 []"></param>
// <param name="TimeOut" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) TakeOfflineParams( /* IN */ TimeOut uint32,
	/* IN */ Parameters []uint8,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("TakeOfflineParams", TimeOut, Parameters, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeName" type="string "></param>
// <param name="TimeOut" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) MoveToNewNode( /* IN */ NodeName string,
	/* IN */ TimeOut uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("MoveToNewNode", NodeName, TimeOut)
	if err != nil {
		return
	}
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="NodeName" type="string "></param>
// <param name="Parameters" type="MSCluster_Property "></param>
func (instance *MSCluster_ResourceGroup) MoveToNewNodeEx( /* IN */ NodeName string,
	/* IN */ Parameters MSCluster_Property,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("MoveToNewNodeEx", NodeName, Parameters, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="NodeName" type="string "></param>
// <param name="Parameters" type="uint8 []"></param>
func (instance *MSCluster_ResourceGroup) MoveToNewNodeParams( /* IN */ NodeName string,
	/* IN */ Parameters []uint8,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("MoveToNewNodeParams", NodeName, Parameters, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="Flags" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) CancelOperation( /* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("CancelOperation", Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="NewName" type="string "></param>
func (instance *MSCluster_ResourceGroup) Rename( /* IN */ NewName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	return

}

//

// <param name="GroupName" type="string "></param>
// <param name="GroupType" type="uint32 "></param>
// <param name="Id" type="string "></param>

// <param name="Id" type="string "></param>
func (instance *MSCluster_ResourceGroup) CreateGroup( /* IN */ GroupName string,
	/* IN */ GroupType uint32,
	/* IN/OUT */ Id string) (err error) {
	_, err = instance.InvokeMethod("CreateGroup", GroupName, GroupType)
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_ResourceGroup) DeleteGroup() (err error) {
	_, err = instance.InvokeMethodWithReturn("DeleteGroup")
	if err != nil {
		return
	}
	return

}

//

// <param name="Options" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) DestroyGroup( /* IN */ Options uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("DestroyGroup", Options)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeNames" type="string []"></param>
func (instance *MSCluster_ResourceGroup) SetPreferredOwners( /* IN */ NodeNames []string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetPreferredOwners", NodeNames)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeNames" type="string []"></param>
func (instance *MSCluster_ResourceGroup) GetPreferredOwners( /* OUT */ NodeNames []string) (err error) {
	_, err = instance.InvokeMethod("GetPreferredOwners")
	if err != nil {
		return
	}
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) GetGroupType() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetGroupType")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="GroupType" type="uint32 "></param>
func (instance *MSCluster_ResourceGroup) SetGroupType( /* IN */ GroupType uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetGroupType", GroupType)
	if err != nil {
		return
	}
	return

}

//

// <param name="ControlCode" type="int32 "></param>
// <param name="InputBuffer" type="uint8 []"></param>

// <param name="OutputBuffer" type="uint8 []"></param>
// <param name="OutputBufferSize" type="int32 "></param>
func (instance *MSCluster_ResourceGroup) ExecuteGroupControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32) (err error) {
	_, err = instance.InvokeMethod("ExecuteGroupControl", ControlCode, InputBuffer)
	if err != nil {
		return
	}
	return

}

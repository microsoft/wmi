// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

// MSCluster_Resource struct
type MSCluster_Resource struct {
	MSCluster_LogicalElement

	//
	CoreResource bool

	//
	CryptoCheckpoints []string

	//
	DeadlockTimeout uint32

	//
	DeleteRequiresAllNodes bool

	//
	EmbeddedFailureAction uint32

	//
	Id string

	//
	IsAlivePollInterval uint32

	//
	IsClusterSharedVolume bool

	//
	LastOperationStatusCode uint64

	//
	LocalQuorumCapable bool

	//
	LooksAlivePollInterval uint32

	//
	MonitorProcessId uint32

	//
	OwnerGroup string

	//
	OwnerNode string

	//
	PendingTimeout uint32

	//
	PersistentState bool

	//
	PrivateProperties MSCluster_Property

	//
	QuorumCapable bool

	//
	RegistryCheckpoints []string

	//
	RequiredDependencyClasses []uint32

	//
	RequiredDependencyTypes []string

	//
	ResourceClass uint32

	//
	ResourceSpecificData1 uint64

	//
	ResourceSpecificData2 uint64

	//
	ResourceSpecificStatus string

	//
	RestartAction uint32

	//
	RestartDelay uint32

	//
	RestartPeriod uint32

	//
	RestartThreshold uint32

	//
	RetryPeriodOnFailure uint32

	//
	SeparateMonitor bool

	//
	State uint32

	//
	StatusInformation uint64

	//
	Subclass uint32

	//
	Type string
}

// SetCoreResource sets the value of CoreResource for the instance
func (instance *MSCluster_Resource) SetPropertyCoreResource(value bool) (err error) {
	return instance.SetProperty("CoreResource", value)
}

// GetCoreResource gets the value of CoreResource for the instance
func (instance *MSCluster_Resource) GetPropertyCoreResource() (value bool, err error) {
	retValue, err := instance.GetProperty("CoreResource")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCryptoCheckpoints sets the value of CryptoCheckpoints for the instance
func (instance *MSCluster_Resource) SetPropertyCryptoCheckpoints(value []string) (err error) {
	return instance.SetProperty("CryptoCheckpoints", value)
}

// GetCryptoCheckpoints gets the value of CryptoCheckpoints for the instance
func (instance *MSCluster_Resource) GetPropertyCryptoCheckpoints() (value []string, err error) {
	retValue, err := instance.GetProperty("CryptoCheckpoints")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeadlockTimeout sets the value of DeadlockTimeout for the instance
func (instance *MSCluster_Resource) SetPropertyDeadlockTimeout(value uint32) (err error) {
	return instance.SetProperty("DeadlockTimeout", value)
}

// GetDeadlockTimeout gets the value of DeadlockTimeout for the instance
func (instance *MSCluster_Resource) GetPropertyDeadlockTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeadlockTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeleteRequiresAllNodes sets the value of DeleteRequiresAllNodes for the instance
func (instance *MSCluster_Resource) SetPropertyDeleteRequiresAllNodes(value bool) (err error) {
	return instance.SetProperty("DeleteRequiresAllNodes", value)
}

// GetDeleteRequiresAllNodes gets the value of DeleteRequiresAllNodes for the instance
func (instance *MSCluster_Resource) GetPropertyDeleteRequiresAllNodes() (value bool, err error) {
	retValue, err := instance.GetProperty("DeleteRequiresAllNodes")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEmbeddedFailureAction sets the value of EmbeddedFailureAction for the instance
func (instance *MSCluster_Resource) SetPropertyEmbeddedFailureAction(value uint32) (err error) {
	return instance.SetProperty("EmbeddedFailureAction", value)
}

// GetEmbeddedFailureAction gets the value of EmbeddedFailureAction for the instance
func (instance *MSCluster_Resource) GetPropertyEmbeddedFailureAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("EmbeddedFailureAction")
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
func (instance *MSCluster_Resource) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", value)
}

// GetId gets the value of Id for the instance
func (instance *MSCluster_Resource) GetPropertyId() (value string, err error) {
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

// SetIsAlivePollInterval sets the value of IsAlivePollInterval for the instance
func (instance *MSCluster_Resource) SetPropertyIsAlivePollInterval(value uint32) (err error) {
	return instance.SetProperty("IsAlivePollInterval", value)
}

// GetIsAlivePollInterval gets the value of IsAlivePollInterval for the instance
func (instance *MSCluster_Resource) GetPropertyIsAlivePollInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("IsAlivePollInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsClusterSharedVolume sets the value of IsClusterSharedVolume for the instance
func (instance *MSCluster_Resource) SetPropertyIsClusterSharedVolume(value bool) (err error) {
	return instance.SetProperty("IsClusterSharedVolume", value)
}

// GetIsClusterSharedVolume gets the value of IsClusterSharedVolume for the instance
func (instance *MSCluster_Resource) GetPropertyIsClusterSharedVolume() (value bool, err error) {
	retValue, err := instance.GetProperty("IsClusterSharedVolume")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLastOperationStatusCode sets the value of LastOperationStatusCode for the instance
func (instance *MSCluster_Resource) SetPropertyLastOperationStatusCode(value uint64) (err error) {
	return instance.SetProperty("LastOperationStatusCode", value)
}

// GetLastOperationStatusCode gets the value of LastOperationStatusCode for the instance
func (instance *MSCluster_Resource) GetPropertyLastOperationStatusCode() (value uint64, err error) {
	retValue, err := instance.GetProperty("LastOperationStatusCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalQuorumCapable sets the value of LocalQuorumCapable for the instance
func (instance *MSCluster_Resource) SetPropertyLocalQuorumCapable(value bool) (err error) {
	return instance.SetProperty("LocalQuorumCapable", value)
}

// GetLocalQuorumCapable gets the value of LocalQuorumCapable for the instance
func (instance *MSCluster_Resource) GetPropertyLocalQuorumCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("LocalQuorumCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLooksAlivePollInterval sets the value of LooksAlivePollInterval for the instance
func (instance *MSCluster_Resource) SetPropertyLooksAlivePollInterval(value uint32) (err error) {
	return instance.SetProperty("LooksAlivePollInterval", value)
}

// GetLooksAlivePollInterval gets the value of LooksAlivePollInterval for the instance
func (instance *MSCluster_Resource) GetPropertyLooksAlivePollInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("LooksAlivePollInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMonitorProcessId sets the value of MonitorProcessId for the instance
func (instance *MSCluster_Resource) SetPropertyMonitorProcessId(value uint32) (err error) {
	return instance.SetProperty("MonitorProcessId", value)
}

// GetMonitorProcessId gets the value of MonitorProcessId for the instance
func (instance *MSCluster_Resource) GetPropertyMonitorProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("MonitorProcessId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwnerGroup sets the value of OwnerGroup for the instance
func (instance *MSCluster_Resource) SetPropertyOwnerGroup(value string) (err error) {
	return instance.SetProperty("OwnerGroup", value)
}

// GetOwnerGroup gets the value of OwnerGroup for the instance
func (instance *MSCluster_Resource) GetPropertyOwnerGroup() (value string, err error) {
	retValue, err := instance.GetProperty("OwnerGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwnerNode sets the value of OwnerNode for the instance
func (instance *MSCluster_Resource) SetPropertyOwnerNode(value string) (err error) {
	return instance.SetProperty("OwnerNode", value)
}

// GetOwnerNode gets the value of OwnerNode for the instance
func (instance *MSCluster_Resource) GetPropertyOwnerNode() (value string, err error) {
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

// SetPendingTimeout sets the value of PendingTimeout for the instance
func (instance *MSCluster_Resource) SetPropertyPendingTimeout(value uint32) (err error) {
	return instance.SetProperty("PendingTimeout", value)
}

// GetPendingTimeout gets the value of PendingTimeout for the instance
func (instance *MSCluster_Resource) GetPropertyPendingTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("PendingTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPersistentState sets the value of PersistentState for the instance
func (instance *MSCluster_Resource) SetPropertyPersistentState(value bool) (err error) {
	return instance.SetProperty("PersistentState", value)
}

// GetPersistentState gets the value of PersistentState for the instance
func (instance *MSCluster_Resource) GetPropertyPersistentState() (value bool, err error) {
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

// SetPrivateProperties sets the value of PrivateProperties for the instance
func (instance *MSCluster_Resource) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", value)
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_Resource) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
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

// SetQuorumCapable sets the value of QuorumCapable for the instance
func (instance *MSCluster_Resource) SetPropertyQuorumCapable(value bool) (err error) {
	return instance.SetProperty("QuorumCapable", value)
}

// GetQuorumCapable gets the value of QuorumCapable for the instance
func (instance *MSCluster_Resource) GetPropertyQuorumCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("QuorumCapable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegistryCheckpoints sets the value of RegistryCheckpoints for the instance
func (instance *MSCluster_Resource) SetPropertyRegistryCheckpoints(value []string) (err error) {
	return instance.SetProperty("RegistryCheckpoints", value)
}

// GetRegistryCheckpoints gets the value of RegistryCheckpoints for the instance
func (instance *MSCluster_Resource) GetPropertyRegistryCheckpoints() (value []string, err error) {
	retValue, err := instance.GetProperty("RegistryCheckpoints")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredDependencyClasses sets the value of RequiredDependencyClasses for the instance
func (instance *MSCluster_Resource) SetPropertyRequiredDependencyClasses(value []uint32) (err error) {
	return instance.SetProperty("RequiredDependencyClasses", value)
}

// GetRequiredDependencyClasses gets the value of RequiredDependencyClasses for the instance
func (instance *MSCluster_Resource) GetPropertyRequiredDependencyClasses() (value []uint32, err error) {
	retValue, err := instance.GetProperty("RequiredDependencyClasses")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiredDependencyTypes sets the value of RequiredDependencyTypes for the instance
func (instance *MSCluster_Resource) SetPropertyRequiredDependencyTypes(value []string) (err error) {
	return instance.SetProperty("RequiredDependencyTypes", value)
}

// GetRequiredDependencyTypes gets the value of RequiredDependencyTypes for the instance
func (instance *MSCluster_Resource) GetPropertyRequiredDependencyTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("RequiredDependencyTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceClass sets the value of ResourceClass for the instance
func (instance *MSCluster_Resource) SetPropertyResourceClass(value uint32) (err error) {
	return instance.SetProperty("ResourceClass", value)
}

// GetResourceClass gets the value of ResourceClass for the instance
func (instance *MSCluster_Resource) GetPropertyResourceClass() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResourceClass")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceSpecificData1 sets the value of ResourceSpecificData1 for the instance
func (instance *MSCluster_Resource) SetPropertyResourceSpecificData1(value uint64) (err error) {
	return instance.SetProperty("ResourceSpecificData1", value)
}

// GetResourceSpecificData1 gets the value of ResourceSpecificData1 for the instance
func (instance *MSCluster_Resource) GetPropertyResourceSpecificData1() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceSpecificData1")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceSpecificData2 sets the value of ResourceSpecificData2 for the instance
func (instance *MSCluster_Resource) SetPropertyResourceSpecificData2(value uint64) (err error) {
	return instance.SetProperty("ResourceSpecificData2", value)
}

// GetResourceSpecificData2 gets the value of ResourceSpecificData2 for the instance
func (instance *MSCluster_Resource) GetPropertyResourceSpecificData2() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResourceSpecificData2")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceSpecificStatus sets the value of ResourceSpecificStatus for the instance
func (instance *MSCluster_Resource) SetPropertyResourceSpecificStatus(value string) (err error) {
	return instance.SetProperty("ResourceSpecificStatus", value)
}

// GetResourceSpecificStatus gets the value of ResourceSpecificStatus for the instance
func (instance *MSCluster_Resource) GetPropertyResourceSpecificStatus() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceSpecificStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartAction sets the value of RestartAction for the instance
func (instance *MSCluster_Resource) SetPropertyRestartAction(value uint32) (err error) {
	return instance.SetProperty("RestartAction", value)
}

// GetRestartAction gets the value of RestartAction for the instance
func (instance *MSCluster_Resource) GetPropertyRestartAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("RestartAction")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartDelay sets the value of RestartDelay for the instance
func (instance *MSCluster_Resource) SetPropertyRestartDelay(value uint32) (err error) {
	return instance.SetProperty("RestartDelay", value)
}

// GetRestartDelay gets the value of RestartDelay for the instance
func (instance *MSCluster_Resource) GetPropertyRestartDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("RestartDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartPeriod sets the value of RestartPeriod for the instance
func (instance *MSCluster_Resource) SetPropertyRestartPeriod(value uint32) (err error) {
	return instance.SetProperty("RestartPeriod", value)
}

// GetRestartPeriod gets the value of RestartPeriod for the instance
func (instance *MSCluster_Resource) GetPropertyRestartPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("RestartPeriod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRestartThreshold sets the value of RestartThreshold for the instance
func (instance *MSCluster_Resource) SetPropertyRestartThreshold(value uint32) (err error) {
	return instance.SetProperty("RestartThreshold", value)
}

// GetRestartThreshold gets the value of RestartThreshold for the instance
func (instance *MSCluster_Resource) GetPropertyRestartThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("RestartThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetryPeriodOnFailure sets the value of RetryPeriodOnFailure for the instance
func (instance *MSCluster_Resource) SetPropertyRetryPeriodOnFailure(value uint32) (err error) {
	return instance.SetProperty("RetryPeriodOnFailure", value)
}

// GetRetryPeriodOnFailure gets the value of RetryPeriodOnFailure for the instance
func (instance *MSCluster_Resource) GetPropertyRetryPeriodOnFailure() (value uint32, err error) {
	retValue, err := instance.GetProperty("RetryPeriodOnFailure")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSeparateMonitor sets the value of SeparateMonitor for the instance
func (instance *MSCluster_Resource) SetPropertySeparateMonitor(value bool) (err error) {
	return instance.SetProperty("SeparateMonitor", value)
}

// GetSeparateMonitor gets the value of SeparateMonitor for the instance
func (instance *MSCluster_Resource) GetPropertySeparateMonitor() (value bool, err error) {
	retValue, err := instance.GetProperty("SeparateMonitor")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSCluster_Resource) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSCluster_Resource) GetPropertyState() (value uint32, err error) {
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
func (instance *MSCluster_Resource) SetPropertyStatusInformation(value uint64) (err error) {
	return instance.SetProperty("StatusInformation", value)
}

// GetStatusInformation gets the value of StatusInformation for the instance
func (instance *MSCluster_Resource) GetPropertyStatusInformation() (value uint64, err error) {
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

// SetSubclass sets the value of Subclass for the instance
func (instance *MSCluster_Resource) SetPropertySubclass(value uint32) (err error) {
	return instance.SetProperty("Subclass", value)
}

// GetSubclass gets the value of Subclass for the instance
func (instance *MSCluster_Resource) GetPropertySubclass() (value uint32, err error) {
	retValue, err := instance.GetProperty("Subclass")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSCluster_Resource) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSCluster_Resource) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
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

// <param name="Group" type="string "></param>
// <param name="Id" type="string "></param>
// <param name="ResourceName" type="string "></param>
// <param name="ResourceType" type="string "></param>
// <param name="SeparateMonitor" type="bool "></param>

// <param name="Id" type="string "></param>
func (instance *MSCluster_Resource) CreateResource( /* IN */ Group string,
	/* IN */ ResourceName string,
	/* IN */ ResourceType string,
	/* IN */ SeparateMonitor bool,
	/* IN/OUT */ Id string) (err error) {
	_, err = instance.InvokeMethod("CreateResource", Group, ResourceName, ResourceType, SeparateMonitor)
	if err != nil {
		return
	}
	return

}

//

// <param name="Options" type="uint32 "></param>
func (instance *MSCluster_Resource) DeleteResource( /* IN */ Options uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("DeleteResource", Options)
	if err != nil {
		return
	}
	return

}

//

// <param name="Group" type="string "></param>
func (instance *MSCluster_Resource) MoveToNewGroup( /* IN */ Group string) (err error) {
	_, err = instance.InvokeMethodWithReturn("MoveToNewGroup", Group)
	if err != nil {
		return
	}
	return

}

//

// <param name="Resource" type="string "></param>
func (instance *MSCluster_Resource) AddDependency( /* IN */ Resource string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddDependency", Resource)
	if err != nil {
		return
	}
	return

}

//

// <param name="Resource" type="string "></param>
func (instance *MSCluster_Resource) RemoveDependency( /* IN */ Resource string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveDependency", Resource)
	if err != nil {
		return
	}
	return

}

//

// <param name="Expression" type="string "></param>
func (instance *MSCluster_Resource) SetDependencies( /* IN */ Expression string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetDependencies", Expression)
	if err != nil {
		return
	}
	return

}

//

// <param name="AsResourceIds" type="bool "></param>

// <param name="Expression" type="string "></param>
func (instance *MSCluster_Resource) GetDependencies( /* OUT */ Expression string,
	/* OPTIONAL IN */ AsResourceIds bool) (err error) {
	_, err = instance.InvokeMethod("GetDependencies", AsResourceIds)
	if err != nil {
		return
	}
	return

}

//

// <param name="TimeOut" type="uint32 "></param>
func (instance *MSCluster_Resource) BringOnline( /* IN */ TimeOut uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("BringOnline", TimeOut)
	if err != nil {
		return
	}
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Parameters" type="MSCluster_Property "></param>
// <param name="TimeOut" type="uint32 "></param>
func (instance *MSCluster_Resource) TakeOffline( /* IN */ TimeOut uint32,
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
func (instance *MSCluster_Resource) TakeOfflineParams( /* IN */ TimeOut uint32,
	/* IN */ Parameters []uint8,
	/* IN */ Flags uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("TakeOfflineParams", TimeOut, Parameters, Flags)
	if err != nil {
		return
	}
	return

}

//

// <param name="newName" type="string "></param>
func (instance *MSCluster_Resource) Rename( /* IN */ newName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Rename", newName)
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_Resource) FailResource() (err error) {
	_, err = instance.InvokeMethodWithReturn("FailResource")
	if err != nil {
		return
	}
	return

}

//

// <param name="CheckpointName" type="string "></param>
func (instance *MSCluster_Resource) AddRegistryCheckpoint( /* IN */ CheckpointName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddRegistryCheckpoint", CheckpointName)
	if err != nil {
		return
	}
	return

}

//

// <param name="CheckpointName" type="string "></param>
func (instance *MSCluster_Resource) RemoveRegistryCheckpoint( /* IN */ CheckpointName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveRegistryCheckpoint", CheckpointName)
	if err != nil {
		return
	}
	return

}

//

// <param name="CheckpointName" type="string "></param>
func (instance *MSCluster_Resource) AddCryptoCheckpoint( /* IN */ CheckpointName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddCryptoCheckpoint", CheckpointName)
	if err != nil {
		return
	}
	return

}

//

// <param name="CheckpointName" type="string "></param>
func (instance *MSCluster_Resource) RemoveCryptoCheckpoint( /* IN */ CheckpointName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveCryptoCheckpoint", CheckpointName)
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_Resource) RenewAddress() (err error) {
	_, err = instance.InvokeMethodWithReturn("RenewAddress")
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_Resource) ReleaseAddress() (err error) {
	_, err = instance.InvokeMethodWithReturn("ReleaseAddress")
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeNames" type="string []"></param>
func (instance *MSCluster_Resource) GetPossibleOwners( /* OUT */ NodeNames []string) (err error) {
	_, err = instance.InvokeMethod("GetPossibleOwners")
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeName" type="string "></param>
func (instance *MSCluster_Resource) AddPossibleOwner( /* IN */ NodeName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddPossibleOwner", NodeName)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeName" type="string "></param>
func (instance *MSCluster_Resource) RemovePossibleOwner( /* IN */ NodeName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemovePossibleOwner", NodeName)
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_Resource) UpdateVirtualMachine() (err error) {
	_, err = instance.InvokeMethodWithReturn("UpdateVirtualMachine")
	if err != nil {
		return
	}
	return

}

//

// <param name="ConfigurationDestinationPath" type="string "></param>
// <param name="DestinationPaths" type="string []"></param>
// <param name="ResourceDestinationPools" type="string []"></param>
// <param name="SnapshotDestinationPath" type="string "></param>
// <param name="SourcePaths" type="string []"></param>
// <param name="SwapFileDestinationPath" type="string "></param>
func (instance *MSCluster_Resource) MigrateVirtualMachine( /* IN */ SnapshotDestinationPath string,
	/* IN */ ConfigurationDestinationPath string,
	/* IN */ SwapFileDestinationPath string,
	/* IN */ SourcePaths []string,
	/* IN */ DestinationPaths []string,
	/* IN */ ResourceDestinationPools []string) (err error) {
	_, err = instance.InvokeMethodWithReturn("MigrateVirtualMachine", SnapshotDestinationPath, ConfigurationDestinationPath, SwapFileDestinationPath, SourcePaths, DestinationPaths, ResourceDestinationPools)
	if err != nil {
		return
	}
	return

}

//

// <param name="StorageDevice" type="MSCluster_AvailableDisk "></param>
func (instance *MSCluster_Resource) AttachStorageDevice( /* IN */ StorageDevice MSCluster_AvailableDisk) (err error) {
	_, err = instance.InvokeMethodWithReturn("AttachStorageDevice", StorageDevice)
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
func (instance *MSCluster_Resource) ExecuteResourceControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32) (err error) {
	_, err = instance.InvokeMethod("ExecuteResourceControl", ControlCode, InputBuffer)
	if err != nil {
		return
	}
	return

}

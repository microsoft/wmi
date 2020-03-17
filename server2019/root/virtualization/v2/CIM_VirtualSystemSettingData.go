// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// CIM_VirtualSystemSettingData struct
type CIM_VirtualSystemSettingData struct {
	CIM_SettingData

	// Action to take for the virtual system when the software executed by the virtual system fails. Failures in this case means a failure that is detectable by the host platform, such as a non-interuptable wait state condition.
	AutomaticRecoveryAction VirtualSystemSettingData_AutomaticRecoveryAction

	// Action to take for the virtual system when the host is shut down.
	AutomaticShutdownAction VirtualSystemSettingData_AutomaticShutdownAction

	// Action to take for the virtual system when the host is started.
	AutomaticStartupAction VirtualSystemSettingData_AutomaticStartupAction

	// Delay applicable to startup action. The value shall be in the interval variant of the datetime datatype.
	AutomaticStartupActionDelay string

	// Number indicating the relative sequence of virtual system activation when the host system is started. A lower number indicates earlier activation. If one or more configurations show the same value, the sequence is implementation dependent. A value of 0 indicates that the sequence is implementation dependent.
	AutomaticStartupActionSequenceNumber uint16

	// Filepath of a directory where information about the virtual system configuration is stored.Format shall be URI based on RFC 2079.
	ConfigurationDataRoot string

	// Filepath of a file where information about the virtual system configuration is stored. A relative path appends to the value of the ConfigurationDataRoot property.Format shall be URI based on RFC 2079.
	ConfigurationFile string

	// Unique id of the virtual system configuration. Note that the ConfigurationID is different from the InstanceID as it is assigned by the implementation to a virtual system or a virtual system configuration. It is not a key, and the same value may occur within more than one instance.
	ConfigurationID string

	// Time when the virtual system configuration was created.
	CreationTime string

	// Filepath of a directory where log information about the virtual system is stored. A relative path appends to the value of the ConfigurationDataRoot property.Format shall be URI based on RFC 2079.
	LogDataRoot string

	// End-user supplied notes that are related to the virtual system.
	Notes []string

	// Filepath of a file where recovery relateded information of the virtual system is stored.Format shall be URI based on RFC 2079.
	RecoveryFile string

	// Filepath of a directory where information about virtual system snapshots is stored. A relative path appends to the value of the ConfigurationDataRoot property.Format shall be URI based on RFC 2079.
	SnapshotDataRoot string

	// Filepath of a directory where suspend related information about the virtual system is stored. A relative path appends to the value of the ConfigurationDataRoot property.Format shall be URI based on RFC 2079.
	SuspendDataRoot string

	// Filepath of a directory where swapfiles of the virtual system are stored. A relative path appends to the value of the ConfigurationDataRoot property.Format shall be URI based on RFC 2079.
	SwapFileDataRoot string

	// VirtualSystemIdentifier shall reflect a unique name for the system as it is used within the virtualization platform. Note that the VirtualSystemIdentifier is not the hostname assigned to the operating system instance running within the virtual system, nor is it an IP address or MAC address assigned to any of its network ports.
	///On create requests VirtualSystemIdentifier may contain implementation specific rules (like simple patterns or regular expresssion) that may be interpreted by the implementation when assigning a VirtualSystemIdentifier.
	VirtualSystemIdentifier string

	// VirtualSystemType shall reflect a particular type of virtual system.
	///The property value shall conform to this format (in ABNF): vs-type = dmtf-value / other-org-value / legacy-value; dmtf-value = "DMTF:" defining-org ":" org-vs-type; other-org-value = defining-org ":" org-vs-type;
	///Where: dmtf-value:
	///is a property value defined by DMTF and is defined in the description of this property. other-org-value:
	///is a property value defined by a business entity other than DMTF and is not defined in the description of this property. legacy-value:
	///is a property value defined by a business entity other than DMTF and is not defined in the description of this property. These values are permitted but recommended to be deprecated over time. defining-org:
	///is an identifier for the business entity that defines the virtual system type. It shall include a copyrighted, trademarked, or otherwise unique name that is owned by that business entity. It shall not be "DMTF" and shall not contain a colon (:). org-vs-type:
	///is an identifier for the virtual system type within the defining business entity. It shall be unique within the defining-org. It may use any character allowed for CIM strings, except for the following: U0000-U001F (Unicode C0 controls) U0020 (space), note that the reason is that OVF allows for multiple space-separated vs-type values in this property. U007F (Unicode C0 controls) U0080-U009F (Unicode C1 controls)
	///If there is a need to structure the value into segments, the segments should be separated with a single colon (:).
	///The values of this property shall be processed case sensitively. They are intended to be processed programmatically (instead of being a display name) and should be short.
	///As stated in the class description, instances of this class may be used for various purposes. A management application intending to use an instance of this class as input parameter to an operation that creates or modifies a virtual system should first determine the set of valid virtual system types that are supported by the virtualization platform hosting the virtual system by inspecting values of array property VirtualSystemTypesSupported of the instance of class CIM_VirtualSystemManagementCapabilities that describes the capabilities of the virtualization platform.
	///The following DMTF values are defined: DMTF:unknown - the virtual system type is unknown or cannot be determined
	VirtualSystemType string
}

// SetAutomaticRecoveryAction sets the value of AutomaticRecoveryAction for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyAutomaticRecoveryAction(value VirtualSystemSettingData_AutomaticRecoveryAction) (err error) {
	return instance.SetProperty("AutomaticRecoveryAction", value)
}

// GetAutomaticRecoveryAction gets the value of AutomaticRecoveryAction for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyAutomaticRecoveryAction() (value VirtualSystemSettingData_AutomaticRecoveryAction, err error) {
	retValue, err := instance.GetProperty("AutomaticRecoveryAction")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualSystemSettingData_AutomaticRecoveryAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutomaticShutdownAction sets the value of AutomaticShutdownAction for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyAutomaticShutdownAction(value VirtualSystemSettingData_AutomaticShutdownAction) (err error) {
	return instance.SetProperty("AutomaticShutdownAction", value)
}

// GetAutomaticShutdownAction gets the value of AutomaticShutdownAction for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyAutomaticShutdownAction() (value VirtualSystemSettingData_AutomaticShutdownAction, err error) {
	retValue, err := instance.GetProperty("AutomaticShutdownAction")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualSystemSettingData_AutomaticShutdownAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutomaticStartupAction sets the value of AutomaticStartupAction for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyAutomaticStartupAction(value VirtualSystemSettingData_AutomaticStartupAction) (err error) {
	return instance.SetProperty("AutomaticStartupAction", value)
}

// GetAutomaticStartupAction gets the value of AutomaticStartupAction for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyAutomaticStartupAction() (value VirtualSystemSettingData_AutomaticStartupAction, err error) {
	retValue, err := instance.GetProperty("AutomaticStartupAction")
	if err != nil {
		return
	}
	value, ok := retValue.(VirtualSystemSettingData_AutomaticStartupAction)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutomaticStartupActionDelay sets the value of AutomaticStartupActionDelay for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyAutomaticStartupActionDelay(value string) (err error) {
	return instance.SetProperty("AutomaticStartupActionDelay", value)
}

// GetAutomaticStartupActionDelay gets the value of AutomaticStartupActionDelay for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyAutomaticStartupActionDelay() (value string, err error) {
	retValue, err := instance.GetProperty("AutomaticStartupActionDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutomaticStartupActionSequenceNumber sets the value of AutomaticStartupActionSequenceNumber for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyAutomaticStartupActionSequenceNumber(value uint16) (err error) {
	return instance.SetProperty("AutomaticStartupActionSequenceNumber", value)
}

// GetAutomaticStartupActionSequenceNumber gets the value of AutomaticStartupActionSequenceNumber for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyAutomaticStartupActionSequenceNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("AutomaticStartupActionSequenceNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigurationDataRoot sets the value of ConfigurationDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyConfigurationDataRoot(value string) (err error) {
	return instance.SetProperty("ConfigurationDataRoot", value)
}

// GetConfigurationDataRoot gets the value of ConfigurationDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyConfigurationDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationDataRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigurationFile sets the value of ConfigurationFile for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyConfigurationFile(value string) (err error) {
	return instance.SetProperty("ConfigurationFile", value)
}

// GetConfigurationFile gets the value of ConfigurationFile for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyConfigurationFile() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationFile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConfigurationID sets the value of ConfigurationID for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyConfigurationID(value string) (err error) {
	return instance.SetProperty("ConfigurationID", value)
}

// GetConfigurationID gets the value of ConfigurationID for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyConfigurationID() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigurationID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLogDataRoot sets the value of LogDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyLogDataRoot(value string) (err error) {
	return instance.SetProperty("LogDataRoot", value)
}

// GetLogDataRoot gets the value of LogDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyLogDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("LogDataRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotes sets the value of Notes for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyNotes(value []string) (err error) {
	return instance.SetProperty("Notes", value)
}

// GetNotes gets the value of Notes for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyNotes() (value []string, err error) {
	retValue, err := instance.GetProperty("Notes")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryFile sets the value of RecoveryFile for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyRecoveryFile(value string) (err error) {
	return instance.SetProperty("RecoveryFile", value)
}

// GetRecoveryFile gets the value of RecoveryFile for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyRecoveryFile() (value string, err error) {
	retValue, err := instance.GetProperty("RecoveryFile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSnapshotDataRoot sets the value of SnapshotDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertySnapshotDataRoot(value string) (err error) {
	return instance.SetProperty("SnapshotDataRoot", value)
}

// GetSnapshotDataRoot gets the value of SnapshotDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertySnapshotDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("SnapshotDataRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuspendDataRoot sets the value of SuspendDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertySuspendDataRoot(value string) (err error) {
	return instance.SetProperty("SuspendDataRoot", value)
}

// GetSuspendDataRoot gets the value of SuspendDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertySuspendDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("SuspendDataRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwapFileDataRoot sets the value of SwapFileDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertySwapFileDataRoot(value string) (err error) {
	return instance.SetProperty("SwapFileDataRoot", value)
}

// GetSwapFileDataRoot gets the value of SwapFileDataRoot for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertySwapFileDataRoot() (value string, err error) {
	retValue, err := instance.GetProperty("SwapFileDataRoot")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemIdentifier sets the value of VirtualSystemIdentifier for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyVirtualSystemIdentifier(value string) (err error) {
	return instance.SetProperty("VirtualSystemIdentifier", value)
}

// GetVirtualSystemIdentifier gets the value of VirtualSystemIdentifier for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyVirtualSystemIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemIdentifier")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualSystemType sets the value of VirtualSystemType for the instance
func (instance *CIM_VirtualSystemSettingData) SetPropertyVirtualSystemType(value string) (err error) {
	return instance.SetProperty("VirtualSystemType", value)
}

// GetVirtualSystemType gets the value of VirtualSystemType for the instance
func (instance *CIM_VirtualSystemSettingData) GetPropertyVirtualSystemType() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualSystemType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

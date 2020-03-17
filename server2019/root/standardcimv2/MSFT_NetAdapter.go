// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetAdapter struct
type MSFT_NetAdapter struct {
	CIM_NetworkPort

	//
	AdminLocked bool

	//
	ComponentID string

	//
	ConnectorPresent bool

	//
	DeviceName string

	//
	DeviceWakeUpEnable bool

	//
	DriverDate string

	//
	DriverDateData uint64

	//
	DriverDescription string

	//
	DriverMajorNdisVersion uint8

	//
	DriverMinorNdisVersion uint8

	//
	DriverName string

	//
	DriverProvider string

	//
	DriverVersionString string

	//
	EndPointInterface bool

	//
	HardwareInterface bool

	//
	Hidden bool

	//
	HigherLayerInterfaceIndices []uint32

	//
	IMFilter bool

	//
	InterfaceAdminStatus uint32

	//
	InterfaceDescription string

	//
	InterfaceGuid string

	//
	InterfaceIndex uint32

	//
	InterfaceName string

	//
	InterfaceOperationalStatus uint32

	//
	InterfaceType uint32

	//
	iSCSIInterface bool

	//
	LowerLayerInterfaceIndices []uint32

	//
	MajorDriverVersion uint16

	//
	MediaConnectState uint32

	//
	MediaDuplexState uint32

	//
	MinorDriverVersion uint16

	//
	MtuSize uint32

	//
	NdisMedium uint32

	//
	NdisPhysicalMedium uint32

	//
	NetLuid uint64

	//
	NetLuidIndex uint32

	//
	NotUserRemovable bool

	//
	OperationalStatusDownDefaultPortNotAuthenticated bool

	//
	OperationalStatusDownInterfacePaused bool

	//
	OperationalStatusDownLowPowerState bool

	//
	OperationalStatusDownMediaDisconnected bool

	//
	PnPDeviceID string

	//
	PromiscuousMode bool

	//
	ReceiveLinkSpeed uint64

	//
	State uint32

	//
	TransmitLinkSpeed uint64

	//
	Virtual bool

	//
	VlanID uint16

	//
	WdmInterface bool
}

// SetAdminLocked sets the value of AdminLocked for the instance
func (instance *MSFT_NetAdapter) SetPropertyAdminLocked(value bool) (err error) {
	return instance.SetProperty("AdminLocked", value)
}

// GetAdminLocked gets the value of AdminLocked for the instance
func (instance *MSFT_NetAdapter) GetPropertyAdminLocked() (value bool, err error) {
	retValue, err := instance.GetProperty("AdminLocked")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComponentID sets the value of ComponentID for the instance
func (instance *MSFT_NetAdapter) SetPropertyComponentID(value string) (err error) {
	return instance.SetProperty("ComponentID", value)
}

// GetComponentID gets the value of ComponentID for the instance
func (instance *MSFT_NetAdapter) GetPropertyComponentID() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectorPresent sets the value of ConnectorPresent for the instance
func (instance *MSFT_NetAdapter) SetPropertyConnectorPresent(value bool) (err error) {
	return instance.SetProperty("ConnectorPresent", value)
}

// GetConnectorPresent gets the value of ConnectorPresent for the instance
func (instance *MSFT_NetAdapter) GetPropertyConnectorPresent() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectorPresent")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceName sets the value of DeviceName for the instance
func (instance *MSFT_NetAdapter) SetPropertyDeviceName(value string) (err error) {
	return instance.SetProperty("DeviceName", value)
}

// GetDeviceName gets the value of DeviceName for the instance
func (instance *MSFT_NetAdapter) GetPropertyDeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("DeviceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDeviceWakeUpEnable sets the value of DeviceWakeUpEnable for the instance
func (instance *MSFT_NetAdapter) SetPropertyDeviceWakeUpEnable(value bool) (err error) {
	return instance.SetProperty("DeviceWakeUpEnable", value)
}

// GetDeviceWakeUpEnable gets the value of DeviceWakeUpEnable for the instance
func (instance *MSFT_NetAdapter) GetPropertyDeviceWakeUpEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("DeviceWakeUpEnable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverDate sets the value of DriverDate for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverDate(value string) (err error) {
	return instance.SetProperty("DriverDate", value)
}

// GetDriverDate gets the value of DriverDate for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverDate() (value string, err error) {
	retValue, err := instance.GetProperty("DriverDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverDateData sets the value of DriverDateData for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverDateData(value uint64) (err error) {
	return instance.SetProperty("DriverDateData", value)
}

// GetDriverDateData gets the value of DriverDateData for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverDateData() (value uint64, err error) {
	retValue, err := instance.GetProperty("DriverDateData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverDescription sets the value of DriverDescription for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverDescription(value string) (err error) {
	return instance.SetProperty("DriverDescription", value)
}

// GetDriverDescription gets the value of DriverDescription for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DriverDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverMajorNdisVersion sets the value of DriverMajorNdisVersion for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverMajorNdisVersion(value uint8) (err error) {
	return instance.SetProperty("DriverMajorNdisVersion", value)
}

// GetDriverMajorNdisVersion gets the value of DriverMajorNdisVersion for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverMajorNdisVersion() (value uint8, err error) {
	retValue, err := instance.GetProperty("DriverMajorNdisVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverMinorNdisVersion sets the value of DriverMinorNdisVersion for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverMinorNdisVersion(value uint8) (err error) {
	return instance.SetProperty("DriverMinorNdisVersion", value)
}

// GetDriverMinorNdisVersion gets the value of DriverMinorNdisVersion for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverMinorNdisVersion() (value uint8, err error) {
	retValue, err := instance.GetProperty("DriverMinorNdisVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverName sets the value of DriverName for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverName(value string) (err error) {
	return instance.SetProperty("DriverName", value)
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverName() (value string, err error) {
	retValue, err := instance.GetProperty("DriverName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverProvider sets the value of DriverProvider for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverProvider(value string) (err error) {
	return instance.SetProperty("DriverProvider", value)
}

// GetDriverProvider gets the value of DriverProvider for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverProvider() (value string, err error) {
	retValue, err := instance.GetProperty("DriverProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverVersionString sets the value of DriverVersionString for the instance
func (instance *MSFT_NetAdapter) SetPropertyDriverVersionString(value string) (err error) {
	return instance.SetProperty("DriverVersionString", value)
}

// GetDriverVersionString gets the value of DriverVersionString for the instance
func (instance *MSFT_NetAdapter) GetPropertyDriverVersionString() (value string, err error) {
	retValue, err := instance.GetProperty("DriverVersionString")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEndPointInterface sets the value of EndPointInterface for the instance
func (instance *MSFT_NetAdapter) SetPropertyEndPointInterface(value bool) (err error) {
	return instance.SetProperty("EndPointInterface", value)
}

// GetEndPointInterface gets the value of EndPointInterface for the instance
func (instance *MSFT_NetAdapter) GetPropertyEndPointInterface() (value bool, err error) {
	retValue, err := instance.GetProperty("EndPointInterface")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardwareInterface sets the value of HardwareInterface for the instance
func (instance *MSFT_NetAdapter) SetPropertyHardwareInterface(value bool) (err error) {
	return instance.SetProperty("HardwareInterface", value)
}

// GetHardwareInterface gets the value of HardwareInterface for the instance
func (instance *MSFT_NetAdapter) GetPropertyHardwareInterface() (value bool, err error) {
	retValue, err := instance.GetProperty("HardwareInterface")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHidden sets the value of Hidden for the instance
func (instance *MSFT_NetAdapter) SetPropertyHidden(value bool) (err error) {
	return instance.SetProperty("Hidden", value)
}

// GetHidden gets the value of Hidden for the instance
func (instance *MSFT_NetAdapter) GetPropertyHidden() (value bool, err error) {
	retValue, err := instance.GetProperty("Hidden")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHigherLayerInterfaceIndices sets the value of HigherLayerInterfaceIndices for the instance
func (instance *MSFT_NetAdapter) SetPropertyHigherLayerInterfaceIndices(value []uint32) (err error) {
	return instance.SetProperty("HigherLayerInterfaceIndices", value)
}

// GetHigherLayerInterfaceIndices gets the value of HigherLayerInterfaceIndices for the instance
func (instance *MSFT_NetAdapter) GetPropertyHigherLayerInterfaceIndices() (value []uint32, err error) {
	retValue, err := instance.GetProperty("HigherLayerInterfaceIndices")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIMFilter sets the value of IMFilter for the instance
func (instance *MSFT_NetAdapter) SetPropertyIMFilter(value bool) (err error) {
	return instance.SetProperty("IMFilter", value)
}

// GetIMFilter gets the value of IMFilter for the instance
func (instance *MSFT_NetAdapter) GetPropertyIMFilter() (value bool, err error) {
	retValue, err := instance.GetProperty("IMFilter")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceAdminStatus sets the value of InterfaceAdminStatus for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceAdminStatus(value uint32) (err error) {
	return instance.SetProperty("InterfaceAdminStatus", value)
}

// GetInterfaceAdminStatus gets the value of InterfaceAdminStatus for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceAdminStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceAdminStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", value)
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceGuid sets the value of InterfaceGuid for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceGuid(value string) (err error) {
	return instance.SetProperty("InterfaceGuid", value)
}

// GetInterfaceGuid gets the value of InterfaceGuid for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", value)
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceName sets the value of InterfaceName for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceName(value string) (err error) {
	return instance.SetProperty("InterfaceName", value)
}

// GetInterfaceName gets the value of InterfaceName for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceName() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceOperationalStatus sets the value of InterfaceOperationalStatus for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceOperationalStatus(value uint32) (err error) {
	return instance.SetProperty("InterfaceOperationalStatus", value)
}

// GetInterfaceOperationalStatus gets the value of InterfaceOperationalStatus for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceOperationalStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceOperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceType sets the value of InterfaceType for the instance
func (instance *MSFT_NetAdapter) SetPropertyInterfaceType(value uint32) (err error) {
	return instance.SetProperty("InterfaceType", value)
}

// GetInterfaceType gets the value of InterfaceType for the instance
func (instance *MSFT_NetAdapter) GetPropertyInterfaceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetiSCSIInterface sets the value of iSCSIInterface for the instance
func (instance *MSFT_NetAdapter) SetPropertyiSCSIInterface(value bool) (err error) {
	return instance.SetProperty("iSCSIInterface", value)
}

// GetiSCSIInterface gets the value of iSCSIInterface for the instance
func (instance *MSFT_NetAdapter) GetPropertyiSCSIInterface() (value bool, err error) {
	retValue, err := instance.GetProperty("iSCSIInterface")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLowerLayerInterfaceIndices sets the value of LowerLayerInterfaceIndices for the instance
func (instance *MSFT_NetAdapter) SetPropertyLowerLayerInterfaceIndices(value []uint32) (err error) {
	return instance.SetProperty("LowerLayerInterfaceIndices", value)
}

// GetLowerLayerInterfaceIndices gets the value of LowerLayerInterfaceIndices for the instance
func (instance *MSFT_NetAdapter) GetPropertyLowerLayerInterfaceIndices() (value []uint32, err error) {
	retValue, err := instance.GetProperty("LowerLayerInterfaceIndices")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMajorDriverVersion sets the value of MajorDriverVersion for the instance
func (instance *MSFT_NetAdapter) SetPropertyMajorDriverVersion(value uint16) (err error) {
	return instance.SetProperty("MajorDriverVersion", value)
}

// GetMajorDriverVersion gets the value of MajorDriverVersion for the instance
func (instance *MSFT_NetAdapter) GetPropertyMajorDriverVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("MajorDriverVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaConnectState sets the value of MediaConnectState for the instance
func (instance *MSFT_NetAdapter) SetPropertyMediaConnectState(value uint32) (err error) {
	return instance.SetProperty("MediaConnectState", value)
}

// GetMediaConnectState gets the value of MediaConnectState for the instance
func (instance *MSFT_NetAdapter) GetPropertyMediaConnectState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaConnectState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMediaDuplexState sets the value of MediaDuplexState for the instance
func (instance *MSFT_NetAdapter) SetPropertyMediaDuplexState(value uint32) (err error) {
	return instance.SetProperty("MediaDuplexState", value)
}

// GetMediaDuplexState gets the value of MediaDuplexState for the instance
func (instance *MSFT_NetAdapter) GetPropertyMediaDuplexState() (value uint32, err error) {
	retValue, err := instance.GetProperty("MediaDuplexState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinorDriverVersion sets the value of MinorDriverVersion for the instance
func (instance *MSFT_NetAdapter) SetPropertyMinorDriverVersion(value uint16) (err error) {
	return instance.SetProperty("MinorDriverVersion", value)
}

// GetMinorDriverVersion gets the value of MinorDriverVersion for the instance
func (instance *MSFT_NetAdapter) GetPropertyMinorDriverVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinorDriverVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMtuSize sets the value of MtuSize for the instance
func (instance *MSFT_NetAdapter) SetPropertyMtuSize(value uint32) (err error) {
	return instance.SetProperty("MtuSize", value)
}

// GetMtuSize gets the value of MtuSize for the instance
func (instance *MSFT_NetAdapter) GetPropertyMtuSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MtuSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisMedium sets the value of NdisMedium for the instance
func (instance *MSFT_NetAdapter) SetPropertyNdisMedium(value uint32) (err error) {
	return instance.SetProperty("NdisMedium", value)
}

// GetNdisMedium gets the value of NdisMedium for the instance
func (instance *MSFT_NetAdapter) GetPropertyNdisMedium() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisMedium")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisPhysicalMedium sets the value of NdisPhysicalMedium for the instance
func (instance *MSFT_NetAdapter) SetPropertyNdisPhysicalMedium(value uint32) (err error) {
	return instance.SetProperty("NdisPhysicalMedium", value)
}

// GetNdisPhysicalMedium gets the value of NdisPhysicalMedium for the instance
func (instance *MSFT_NetAdapter) GetPropertyNdisPhysicalMedium() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisPhysicalMedium")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetLuid sets the value of NetLuid for the instance
func (instance *MSFT_NetAdapter) SetPropertyNetLuid(value uint64) (err error) {
	return instance.SetProperty("NetLuid", value)
}

// GetNetLuid gets the value of NetLuid for the instance
func (instance *MSFT_NetAdapter) GetPropertyNetLuid() (value uint64, err error) {
	retValue, err := instance.GetProperty("NetLuid")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetLuidIndex sets the value of NetLuidIndex for the instance
func (instance *MSFT_NetAdapter) SetPropertyNetLuidIndex(value uint32) (err error) {
	return instance.SetProperty("NetLuidIndex", value)
}

// GetNetLuidIndex gets the value of NetLuidIndex for the instance
func (instance *MSFT_NetAdapter) GetPropertyNetLuidIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetLuidIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNotUserRemovable sets the value of NotUserRemovable for the instance
func (instance *MSFT_NetAdapter) SetPropertyNotUserRemovable(value bool) (err error) {
	return instance.SetProperty("NotUserRemovable", value)
}

// GetNotUserRemovable gets the value of NotUserRemovable for the instance
func (instance *MSFT_NetAdapter) GetPropertyNotUserRemovable() (value bool, err error) {
	retValue, err := instance.GetProperty("NotUserRemovable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatusDownDefaultPortNotAuthenticated sets the value of OperationalStatusDownDefaultPortNotAuthenticated for the instance
func (instance *MSFT_NetAdapter) SetPropertyOperationalStatusDownDefaultPortNotAuthenticated(value bool) (err error) {
	return instance.SetProperty("OperationalStatusDownDefaultPortNotAuthenticated", value)
}

// GetOperationalStatusDownDefaultPortNotAuthenticated gets the value of OperationalStatusDownDefaultPortNotAuthenticated for the instance
func (instance *MSFT_NetAdapter) GetPropertyOperationalStatusDownDefaultPortNotAuthenticated() (value bool, err error) {
	retValue, err := instance.GetProperty("OperationalStatusDownDefaultPortNotAuthenticated")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatusDownInterfacePaused sets the value of OperationalStatusDownInterfacePaused for the instance
func (instance *MSFT_NetAdapter) SetPropertyOperationalStatusDownInterfacePaused(value bool) (err error) {
	return instance.SetProperty("OperationalStatusDownInterfacePaused", value)
}

// GetOperationalStatusDownInterfacePaused gets the value of OperationalStatusDownInterfacePaused for the instance
func (instance *MSFT_NetAdapter) GetPropertyOperationalStatusDownInterfacePaused() (value bool, err error) {
	retValue, err := instance.GetProperty("OperationalStatusDownInterfacePaused")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatusDownLowPowerState sets the value of OperationalStatusDownLowPowerState for the instance
func (instance *MSFT_NetAdapter) SetPropertyOperationalStatusDownLowPowerState(value bool) (err error) {
	return instance.SetProperty("OperationalStatusDownLowPowerState", value)
}

// GetOperationalStatusDownLowPowerState gets the value of OperationalStatusDownLowPowerState for the instance
func (instance *MSFT_NetAdapter) GetPropertyOperationalStatusDownLowPowerState() (value bool, err error) {
	retValue, err := instance.GetProperty("OperationalStatusDownLowPowerState")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationalStatusDownMediaDisconnected sets the value of OperationalStatusDownMediaDisconnected for the instance
func (instance *MSFT_NetAdapter) SetPropertyOperationalStatusDownMediaDisconnected(value bool) (err error) {
	return instance.SetProperty("OperationalStatusDownMediaDisconnected", value)
}

// GetOperationalStatusDownMediaDisconnected gets the value of OperationalStatusDownMediaDisconnected for the instance
func (instance *MSFT_NetAdapter) GetPropertyOperationalStatusDownMediaDisconnected() (value bool, err error) {
	retValue, err := instance.GetProperty("OperationalStatusDownMediaDisconnected")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPnPDeviceID sets the value of PnPDeviceID for the instance
func (instance *MSFT_NetAdapter) SetPropertyPnPDeviceID(value string) (err error) {
	return instance.SetProperty("PnPDeviceID", value)
}

// GetPnPDeviceID gets the value of PnPDeviceID for the instance
func (instance *MSFT_NetAdapter) GetPropertyPnPDeviceID() (value string, err error) {
	retValue, err := instance.GetProperty("PnPDeviceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPromiscuousMode sets the value of PromiscuousMode for the instance
func (instance *MSFT_NetAdapter) SetPropertyPromiscuousMode(value bool) (err error) {
	return instance.SetProperty("PromiscuousMode", value)
}

// GetPromiscuousMode gets the value of PromiscuousMode for the instance
func (instance *MSFT_NetAdapter) GetPropertyPromiscuousMode() (value bool, err error) {
	retValue, err := instance.GetProperty("PromiscuousMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceiveLinkSpeed sets the value of ReceiveLinkSpeed for the instance
func (instance *MSFT_NetAdapter) SetPropertyReceiveLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("ReceiveLinkSpeed", value)
}

// GetReceiveLinkSpeed gets the value of ReceiveLinkSpeed for the instance
func (instance *MSFT_NetAdapter) GetPropertyReceiveLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveLinkSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetState sets the value of State for the instance
func (instance *MSFT_NetAdapter) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", value)
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetAdapter) GetPropertyState() (value uint32, err error) {
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

// SetTransmitLinkSpeed sets the value of TransmitLinkSpeed for the instance
func (instance *MSFT_NetAdapter) SetPropertyTransmitLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("TransmitLinkSpeed", value)
}

// GetTransmitLinkSpeed gets the value of TransmitLinkSpeed for the instance
func (instance *MSFT_NetAdapter) GetPropertyTransmitLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmitLinkSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtual sets the value of Virtual for the instance
func (instance *MSFT_NetAdapter) SetPropertyVirtual(value bool) (err error) {
	return instance.SetProperty("Virtual", value)
}

// GetVirtual gets the value of Virtual for the instance
func (instance *MSFT_NetAdapter) GetPropertyVirtual() (value bool, err error) {
	retValue, err := instance.GetProperty("Virtual")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVlanID sets the value of VlanID for the instance
func (instance *MSFT_NetAdapter) SetPropertyVlanID(value uint16) (err error) {
	return instance.SetProperty("VlanID", value)
}

// GetVlanID gets the value of VlanID for the instance
func (instance *MSFT_NetAdapter) GetPropertyVlanID() (value uint16, err error) {
	retValue, err := instance.GetProperty("VlanID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWdmInterface sets the value of WdmInterface for the instance
func (instance *MSFT_NetAdapter) SetPropertyWdmInterface(value bool) (err error) {
	return instance.SetProperty("WdmInterface", value)
}

// GetWdmInterface gets the value of WdmInterface for the instance
func (instance *MSFT_NetAdapter) GetPropertyWdmInterface() (value bool, err error) {
	retValue, err := instance.GetProperty("WdmInterface")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="CmdletOutput" type="MSFT_NetAdapter "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapter) Enable( /* OUT */ CmdletOutput MSFT_NetAdapter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdletOutput" type="MSFT_NetAdapter "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapter) Disable( /* OUT */ CmdletOutput MSFT_NetAdapter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdletOutput" type="MSFT_NetAdapter "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapter) Restart( /* OUT */ CmdletOutput MSFT_NetAdapter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Restart")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdletOutput" type="MSFT_NetAdapter "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapter) Lock( /* OUT */ CmdletOutput MSFT_NetAdapter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Lock")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdletOutput" type="MSFT_NetAdapter "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapter) Unlock( /* OUT */ CmdletOutput MSFT_NetAdapter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Unlock")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="NewName" type="string "></param>

// <param name="CmdletOutput" type="MSFT_NetAdapter "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapter) Rename( /* IN */ NewName string,
	/* OUT */ CmdletOutput MSFT_NetAdapter) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Rename", NewName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

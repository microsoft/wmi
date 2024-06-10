// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_Connectivity02 struct
type MDM_Policy_Config01_Connectivity02 struct {
	*cim.WmiInstance

	//
	AllowBluetooth int32

	//
	AllowCellularData int32

	//
	AllowCellularDataRoaming int32

	//
	AllowConnectedDevices int32

	//
	AllowPhonePCLinking int32

	//
	AllowVPNOverCellular int32

	//
	AllowVPNRoamingOverCellular int32

	//
	DiablePrintingOverHTTP string

	//
	DisableDownloadingOfPrintDriversOverHTTP string

	//
	DisableInternetDownloadForWebPublishingAndOnlineOrderingWizards string

	//
	DisallowNetworkConnectivityActiveTests int32

	//
	HardenedUNCPaths string

	//
	InstanceID string

	//
	ParentID string

	//
	ProhibitInstallationAndConfigurationOfNetworkBridge string
}

func NewMDM_Policy_Config01_Connectivity02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Connectivity02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Connectivity02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Connectivity02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Connectivity02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Connectivity02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowBluetooth sets the value of AllowBluetooth for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowBluetooth(value int32) (err error) {
	return instance.SetProperty("AllowBluetooth", (value))
}

// GetAllowBluetooth gets the value of AllowBluetooth for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowBluetooth() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowBluetooth")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowCellularData sets the value of AllowCellularData for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowCellularData(value int32) (err error) {
	return instance.SetProperty("AllowCellularData", (value))
}

// GetAllowCellularData gets the value of AllowCellularData for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowCellularData() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCellularData")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowCellularDataRoaming sets the value of AllowCellularDataRoaming for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowCellularDataRoaming(value int32) (err error) {
	return instance.SetProperty("AllowCellularDataRoaming", (value))
}

// GetAllowCellularDataRoaming gets the value of AllowCellularDataRoaming for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowCellularDataRoaming() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCellularDataRoaming")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowConnectedDevices sets the value of AllowConnectedDevices for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowConnectedDevices(value int32) (err error) {
	return instance.SetProperty("AllowConnectedDevices", (value))
}

// GetAllowConnectedDevices gets the value of AllowConnectedDevices for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowConnectedDevices() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowConnectedDevices")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowPhonePCLinking sets the value of AllowPhonePCLinking for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowPhonePCLinking(value int32) (err error) {
	return instance.SetProperty("AllowPhonePCLinking", (value))
}

// GetAllowPhonePCLinking gets the value of AllowPhonePCLinking for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowPhonePCLinking() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPhonePCLinking")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowVPNOverCellular sets the value of AllowVPNOverCellular for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowVPNOverCellular(value int32) (err error) {
	return instance.SetProperty("AllowVPNOverCellular", (value))
}

// GetAllowVPNOverCellular gets the value of AllowVPNOverCellular for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowVPNOverCellular() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowVPNOverCellular")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetAllowVPNRoamingOverCellular sets the value of AllowVPNRoamingOverCellular for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyAllowVPNRoamingOverCellular(value int32) (err error) {
	return instance.SetProperty("AllowVPNRoamingOverCellular", (value))
}

// GetAllowVPNRoamingOverCellular gets the value of AllowVPNRoamingOverCellular for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyAllowVPNRoamingOverCellular() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowVPNRoamingOverCellular")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDiablePrintingOverHTTP sets the value of DiablePrintingOverHTTP for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyDiablePrintingOverHTTP(value string) (err error) {
	return instance.SetProperty("DiablePrintingOverHTTP", (value))
}

// GetDiablePrintingOverHTTP gets the value of DiablePrintingOverHTTP for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyDiablePrintingOverHTTP() (value string, err error) {
	retValue, err := instance.GetProperty("DiablePrintingOverHTTP")
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

// SetDisableDownloadingOfPrintDriversOverHTTP sets the value of DisableDownloadingOfPrintDriversOverHTTP for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyDisableDownloadingOfPrintDriversOverHTTP(value string) (err error) {
	return instance.SetProperty("DisableDownloadingOfPrintDriversOverHTTP", (value))
}

// GetDisableDownloadingOfPrintDriversOverHTTP gets the value of DisableDownloadingOfPrintDriversOverHTTP for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyDisableDownloadingOfPrintDriversOverHTTP() (value string, err error) {
	retValue, err := instance.GetProperty("DisableDownloadingOfPrintDriversOverHTTP")
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

// SetDisableInternetDownloadForWebPublishingAndOnlineOrderingWizards sets the value of DisableInternetDownloadForWebPublishingAndOnlineOrderingWizards for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyDisableInternetDownloadForWebPublishingAndOnlineOrderingWizards(value string) (err error) {
	return instance.SetProperty("DisableInternetDownloadForWebPublishingAndOnlineOrderingWizards", (value))
}

// GetDisableInternetDownloadForWebPublishingAndOnlineOrderingWizards gets the value of DisableInternetDownloadForWebPublishingAndOnlineOrderingWizards for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyDisableInternetDownloadForWebPublishingAndOnlineOrderingWizards() (value string, err error) {
	retValue, err := instance.GetProperty("DisableInternetDownloadForWebPublishingAndOnlineOrderingWizards")
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

// SetDisallowNetworkConnectivityActiveTests sets the value of DisallowNetworkConnectivityActiveTests for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyDisallowNetworkConnectivityActiveTests(value int32) (err error) {
	return instance.SetProperty("DisallowNetworkConnectivityActiveTests", (value))
}

// GetDisallowNetworkConnectivityActiveTests gets the value of DisallowNetworkConnectivityActiveTests for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyDisallowNetworkConnectivityActiveTests() (value int32, err error) {
	retValue, err := instance.GetProperty("DisallowNetworkConnectivityActiveTests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetHardenedUNCPaths sets the value of HardenedUNCPaths for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyHardenedUNCPaths(value string) (err error) {
	return instance.SetProperty("HardenedUNCPaths", (value))
}

// GetHardenedUNCPaths gets the value of HardenedUNCPaths for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyHardenedUNCPaths() (value string, err error) {
	retValue, err := instance.GetProperty("HardenedUNCPaths")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetProhibitInstallationAndConfigurationOfNetworkBridge sets the value of ProhibitInstallationAndConfigurationOfNetworkBridge for the instance
func (instance *MDM_Policy_Config01_Connectivity02) SetPropertyProhibitInstallationAndConfigurationOfNetworkBridge(value string) (err error) {
	return instance.SetProperty("ProhibitInstallationAndConfigurationOfNetworkBridge", (value))
}

// GetProhibitInstallationAndConfigurationOfNetworkBridge gets the value of ProhibitInstallationAndConfigurationOfNetworkBridge for the instance
func (instance *MDM_Policy_Config01_Connectivity02) GetPropertyProhibitInstallationAndConfigurationOfNetworkBridge() (value string, err error) {
	retValue, err := instance.GetProperty("ProhibitInstallationAndConfigurationOfNetworkBridge")
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

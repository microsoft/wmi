// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Config01_AppVirtualization02 struct
type MDM_Policy_Config01_AppVirtualization02 struct {
	*cim.WmiInstance

	//
	AllowAppVClient string

	//
	AllowDynamicVirtualization string

	//
	AllowPackageCleanup string

	//
	AllowPackageScripts string

	//
	AllowPublishingRefreshUX string

	//
	AllowReportingServer string

	//
	AllowRoamingFileExclusions string

	//
	AllowRoamingRegistryExclusions string

	//
	AllowStreamingAutoload string

	//
	ClientCoexistenceAllowMigrationmode string

	//
	InstanceID string

	//
	IntegrationAllowRootGlobal string

	//
	IntegrationAllowRootUser string

	//
	ParentID string

	//
	PublishingAllowServer1 string

	//
	PublishingAllowServer2 string

	//
	PublishingAllowServer3 string

	//
	PublishingAllowServer4 string

	//
	PublishingAllowServer5 string

	//
	StreamingAllowCertificateFilterForClient_SSL string

	//
	StreamingAllowHighCostLaunch string

	//
	StreamingAllowLocationProvider string

	//
	StreamingAllowPackageInstallationRoot string

	//
	StreamingAllowPackageSourceRoot string

	//
	StreamingAllowReestablishmentInterval string

	//
	StreamingAllowReestablishmentRetries string

	//
	StreamingSharedContentStoreMode string

	//
	StreamingSupportBranchCache string

	//
	StreamingVerifyCertificateRevocationList string

	//
	VirtualComponentsAllowList string
}

func NewMDM_Policy_Config01_AppVirtualization02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_AppVirtualization02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_AppVirtualization02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_AppVirtualization02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_AppVirtualization02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_AppVirtualization02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAppVClient sets the value of AllowAppVClient for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowAppVClient(value string) (err error) {
	return instance.SetProperty("AllowAppVClient", (value))
}

// GetAllowAppVClient gets the value of AllowAppVClient for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowAppVClient() (value string, err error) {
	retValue, err := instance.GetProperty("AllowAppVClient")
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

// SetAllowDynamicVirtualization sets the value of AllowDynamicVirtualization for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowDynamicVirtualization(value string) (err error) {
	return instance.SetProperty("AllowDynamicVirtualization", (value))
}

// GetAllowDynamicVirtualization gets the value of AllowDynamicVirtualization for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowDynamicVirtualization() (value string, err error) {
	retValue, err := instance.GetProperty("AllowDynamicVirtualization")
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

// SetAllowPackageCleanup sets the value of AllowPackageCleanup for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowPackageCleanup(value string) (err error) {
	return instance.SetProperty("AllowPackageCleanup", (value))
}

// GetAllowPackageCleanup gets the value of AllowPackageCleanup for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowPackageCleanup() (value string, err error) {
	retValue, err := instance.GetProperty("AllowPackageCleanup")
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

// SetAllowPackageScripts sets the value of AllowPackageScripts for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowPackageScripts(value string) (err error) {
	return instance.SetProperty("AllowPackageScripts", (value))
}

// GetAllowPackageScripts gets the value of AllowPackageScripts for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowPackageScripts() (value string, err error) {
	retValue, err := instance.GetProperty("AllowPackageScripts")
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

// SetAllowPublishingRefreshUX sets the value of AllowPublishingRefreshUX for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowPublishingRefreshUX(value string) (err error) {
	return instance.SetProperty("AllowPublishingRefreshUX", (value))
}

// GetAllowPublishingRefreshUX gets the value of AllowPublishingRefreshUX for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowPublishingRefreshUX() (value string, err error) {
	retValue, err := instance.GetProperty("AllowPublishingRefreshUX")
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

// SetAllowReportingServer sets the value of AllowReportingServer for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowReportingServer(value string) (err error) {
	return instance.SetProperty("AllowReportingServer", (value))
}

// GetAllowReportingServer gets the value of AllowReportingServer for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowReportingServer() (value string, err error) {
	retValue, err := instance.GetProperty("AllowReportingServer")
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

// SetAllowRoamingFileExclusions sets the value of AllowRoamingFileExclusions for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowRoamingFileExclusions(value string) (err error) {
	return instance.SetProperty("AllowRoamingFileExclusions", (value))
}

// GetAllowRoamingFileExclusions gets the value of AllowRoamingFileExclusions for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowRoamingFileExclusions() (value string, err error) {
	retValue, err := instance.GetProperty("AllowRoamingFileExclusions")
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

// SetAllowRoamingRegistryExclusions sets the value of AllowRoamingRegistryExclusions for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowRoamingRegistryExclusions(value string) (err error) {
	return instance.SetProperty("AllowRoamingRegistryExclusions", (value))
}

// GetAllowRoamingRegistryExclusions gets the value of AllowRoamingRegistryExclusions for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowRoamingRegistryExclusions() (value string, err error) {
	retValue, err := instance.GetProperty("AllowRoamingRegistryExclusions")
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

// SetAllowStreamingAutoload sets the value of AllowStreamingAutoload for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyAllowStreamingAutoload(value string) (err error) {
	return instance.SetProperty("AllowStreamingAutoload", (value))
}

// GetAllowStreamingAutoload gets the value of AllowStreamingAutoload for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyAllowStreamingAutoload() (value string, err error) {
	retValue, err := instance.GetProperty("AllowStreamingAutoload")
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

// SetClientCoexistenceAllowMigrationmode sets the value of ClientCoexistenceAllowMigrationmode for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyClientCoexistenceAllowMigrationmode(value string) (err error) {
	return instance.SetProperty("ClientCoexistenceAllowMigrationmode", (value))
}

// GetClientCoexistenceAllowMigrationmode gets the value of ClientCoexistenceAllowMigrationmode for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyClientCoexistenceAllowMigrationmode() (value string, err error) {
	retValue, err := instance.GetProperty("ClientCoexistenceAllowMigrationmode")
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
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyInstanceID() (value string, err error) {
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

// SetIntegrationAllowRootGlobal sets the value of IntegrationAllowRootGlobal for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyIntegrationAllowRootGlobal(value string) (err error) {
	return instance.SetProperty("IntegrationAllowRootGlobal", (value))
}

// GetIntegrationAllowRootGlobal gets the value of IntegrationAllowRootGlobal for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyIntegrationAllowRootGlobal() (value string, err error) {
	retValue, err := instance.GetProperty("IntegrationAllowRootGlobal")
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

// SetIntegrationAllowRootUser sets the value of IntegrationAllowRootUser for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyIntegrationAllowRootUser(value string) (err error) {
	return instance.SetProperty("IntegrationAllowRootUser", (value))
}

// GetIntegrationAllowRootUser gets the value of IntegrationAllowRootUser for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyIntegrationAllowRootUser() (value string, err error) {
	retValue, err := instance.GetProperty("IntegrationAllowRootUser")
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
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyParentID() (value string, err error) {
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

// SetPublishingAllowServer1 sets the value of PublishingAllowServer1 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyPublishingAllowServer1(value string) (err error) {
	return instance.SetProperty("PublishingAllowServer1", (value))
}

// GetPublishingAllowServer1 gets the value of PublishingAllowServer1 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyPublishingAllowServer1() (value string, err error) {
	retValue, err := instance.GetProperty("PublishingAllowServer1")
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

// SetPublishingAllowServer2 sets the value of PublishingAllowServer2 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyPublishingAllowServer2(value string) (err error) {
	return instance.SetProperty("PublishingAllowServer2", (value))
}

// GetPublishingAllowServer2 gets the value of PublishingAllowServer2 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyPublishingAllowServer2() (value string, err error) {
	retValue, err := instance.GetProperty("PublishingAllowServer2")
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

// SetPublishingAllowServer3 sets the value of PublishingAllowServer3 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyPublishingAllowServer3(value string) (err error) {
	return instance.SetProperty("PublishingAllowServer3", (value))
}

// GetPublishingAllowServer3 gets the value of PublishingAllowServer3 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyPublishingAllowServer3() (value string, err error) {
	retValue, err := instance.GetProperty("PublishingAllowServer3")
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

// SetPublishingAllowServer4 sets the value of PublishingAllowServer4 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyPublishingAllowServer4(value string) (err error) {
	return instance.SetProperty("PublishingAllowServer4", (value))
}

// GetPublishingAllowServer4 gets the value of PublishingAllowServer4 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyPublishingAllowServer4() (value string, err error) {
	retValue, err := instance.GetProperty("PublishingAllowServer4")
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

// SetPublishingAllowServer5 sets the value of PublishingAllowServer5 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyPublishingAllowServer5(value string) (err error) {
	return instance.SetProperty("PublishingAllowServer5", (value))
}

// GetPublishingAllowServer5 gets the value of PublishingAllowServer5 for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyPublishingAllowServer5() (value string, err error) {
	retValue, err := instance.GetProperty("PublishingAllowServer5")
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

// SetStreamingAllowCertificateFilterForClient_SSL sets the value of StreamingAllowCertificateFilterForClient_SSL for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowCertificateFilterForClient_SSL(value string) (err error) {
	return instance.SetProperty("StreamingAllowCertificateFilterForClient_SSL", (value))
}

// GetStreamingAllowCertificateFilterForClient_SSL gets the value of StreamingAllowCertificateFilterForClient_SSL for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowCertificateFilterForClient_SSL() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowCertificateFilterForClient_SSL")
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

// SetStreamingAllowHighCostLaunch sets the value of StreamingAllowHighCostLaunch for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowHighCostLaunch(value string) (err error) {
	return instance.SetProperty("StreamingAllowHighCostLaunch", (value))
}

// GetStreamingAllowHighCostLaunch gets the value of StreamingAllowHighCostLaunch for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowHighCostLaunch() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowHighCostLaunch")
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

// SetStreamingAllowLocationProvider sets the value of StreamingAllowLocationProvider for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowLocationProvider(value string) (err error) {
	return instance.SetProperty("StreamingAllowLocationProvider", (value))
}

// GetStreamingAllowLocationProvider gets the value of StreamingAllowLocationProvider for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowLocationProvider() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowLocationProvider")
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

// SetStreamingAllowPackageInstallationRoot sets the value of StreamingAllowPackageInstallationRoot for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowPackageInstallationRoot(value string) (err error) {
	return instance.SetProperty("StreamingAllowPackageInstallationRoot", (value))
}

// GetStreamingAllowPackageInstallationRoot gets the value of StreamingAllowPackageInstallationRoot for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowPackageInstallationRoot() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowPackageInstallationRoot")
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

// SetStreamingAllowPackageSourceRoot sets the value of StreamingAllowPackageSourceRoot for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowPackageSourceRoot(value string) (err error) {
	return instance.SetProperty("StreamingAllowPackageSourceRoot", (value))
}

// GetStreamingAllowPackageSourceRoot gets the value of StreamingAllowPackageSourceRoot for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowPackageSourceRoot() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowPackageSourceRoot")
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

// SetStreamingAllowReestablishmentInterval sets the value of StreamingAllowReestablishmentInterval for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowReestablishmentInterval(value string) (err error) {
	return instance.SetProperty("StreamingAllowReestablishmentInterval", (value))
}

// GetStreamingAllowReestablishmentInterval gets the value of StreamingAllowReestablishmentInterval for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowReestablishmentInterval() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowReestablishmentInterval")
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

// SetStreamingAllowReestablishmentRetries sets the value of StreamingAllowReestablishmentRetries for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingAllowReestablishmentRetries(value string) (err error) {
	return instance.SetProperty("StreamingAllowReestablishmentRetries", (value))
}

// GetStreamingAllowReestablishmentRetries gets the value of StreamingAllowReestablishmentRetries for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingAllowReestablishmentRetries() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingAllowReestablishmentRetries")
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

// SetStreamingSharedContentStoreMode sets the value of StreamingSharedContentStoreMode for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingSharedContentStoreMode(value string) (err error) {
	return instance.SetProperty("StreamingSharedContentStoreMode", (value))
}

// GetStreamingSharedContentStoreMode gets the value of StreamingSharedContentStoreMode for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingSharedContentStoreMode() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingSharedContentStoreMode")
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

// SetStreamingSupportBranchCache sets the value of StreamingSupportBranchCache for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingSupportBranchCache(value string) (err error) {
	return instance.SetProperty("StreamingSupportBranchCache", (value))
}

// GetStreamingSupportBranchCache gets the value of StreamingSupportBranchCache for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingSupportBranchCache() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingSupportBranchCache")
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

// SetStreamingVerifyCertificateRevocationList sets the value of StreamingVerifyCertificateRevocationList for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyStreamingVerifyCertificateRevocationList(value string) (err error) {
	return instance.SetProperty("StreamingVerifyCertificateRevocationList", (value))
}

// GetStreamingVerifyCertificateRevocationList gets the value of StreamingVerifyCertificateRevocationList for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyStreamingVerifyCertificateRevocationList() (value string, err error) {
	retValue, err := instance.GetProperty("StreamingVerifyCertificateRevocationList")
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

// SetVirtualComponentsAllowList sets the value of VirtualComponentsAllowList for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) SetPropertyVirtualComponentsAllowList(value string) (err error) {
	return instance.SetProperty("VirtualComponentsAllowList", (value))
}

// GetVirtualComponentsAllowList gets the value of VirtualComponentsAllowList for the instance
func (instance *MDM_Policy_Config01_AppVirtualization02) GetPropertyVirtualComponentsAllowList() (value string, err error) {
	retValue, err := instance.GetProperty("VirtualComponentsAllowList")
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

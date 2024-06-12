// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DSCLocalConfigurationManager struct
type MSFT_DSCLocalConfigurationManager struct {
	*cim.WmiInstance
}

func NewMSFT_DSCLocalConfigurationManagerEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCLocalConfigurationManager, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCLocalConfigurationManager{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DSCLocalConfigurationManagerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCLocalConfigurationManager, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCLocalConfigurationManager{
		WmiInstance: tmp,
	}
	return
}

//

// <param name="ConfigurationData" type="uint8 []"></param>
// <param name="force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) SendConfiguration( /* IN */ ConfigurationData []uint8,
	/* IN */ force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SendConfiguration", ConfigurationData, force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ConfigurationData" type="uint8 []"></param>
// <param name="force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) SendConfigurationApply( /* IN */ ConfigurationData []uint8,
	/* IN */ force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SendConfigurationApply", ConfigurationData, force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="configurationData" type="uint8 []"></param>

// <param name="configurations" type="OMI_BaseResource []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) GetConfiguration( /* IN */ configurationData []uint8,
	/* OUT */ configurations []OMI_BaseResource) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetConfiguration", configurationData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="configurationData" type="uint8 []"></param>

// <param name="InDesiredState" type="bool "></param>
// <param name="ResourcesInDesiredState" type="MSFT_ResourceInDesiredState []"></param>
// <param name="ResourcesNotInDesiredState" type="MSFT_ResourceNotInDesiredState []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) TestConfiguration( /* IN */ configurationData []uint8,
	/* OUT */ InDesiredState bool,
	/* OUT */ ResourcesInDesiredState []MSFT_ResourceInDesiredState,
	/* OUT */ ResourcesNotInDesiredState []MSFT_ResourceNotInDesiredState) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("TestConfiguration", configurationData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) ApplyConfiguration( /* IN */ force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ApplyConfiguration", force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ConfigurationData" type="uint8 []"></param>
// <param name="force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) SendMetaConfigurationApply( /* IN */ ConfigurationData []uint8,
	/* IN */ force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SendMetaConfigurationApply", ConfigurationData, force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MetaConfiguration" type="MSFT_DSCMetaConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) GetMetaConfiguration( /* OUT */ MetaConfiguration MSFT_DSCMetaConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetMetaConfiguration")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="configurationNumber" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) RollBack( /* IN */ configurationNumber uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RollBack", configurationNumber)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Flags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) PerformRequiredConfigurationChecks( /* IN */ Flags uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("PerformRequiredConfigurationChecks", Flags)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) StopConfiguration( /* IN */ force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopConfiguration", force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="All" type="bool "></param>

// <param name="configurationStatus" type="MSFT_DSCConfigurationStatus []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) GetConfigurationStatus( /* IN */ All bool,
	/* OUT */ configurationStatus []MSFT_DSCConfigurationStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetConfigurationStatus", All)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ConfigurationData" type="uint8 []"></param>
// <param name="force" type="bool "></param>
// <param name="jobId" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) SendConfigurationApplyAsync( /* IN */ ConfigurationData []uint8,
	/* IN */ force bool,
	/* IN */ jobId string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SendConfigurationApplyAsync", ConfigurationData, force, jobId)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="jobId" type="string "></param>
// <param name="resumeOutputBookmark" type="uint8 []"></param>

// <param name="output" type="MSFT_DSCConfigurationOutput []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) GetConfigurationResultOutput( /* IN */ jobId string,
	/* IN */ resumeOutputBookmark []uint8,
	/* OUT */ output []MSFT_DSCConfigurationOutput) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetConfigurationResultOutput", jobId, resumeOutputBookmark)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Stage" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) RemoveConfiguration( /* IN */ Stage uint32,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveConfiguration", Stage, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ModuleName" type="string "></param>
// <param name="resourceProperty" type="uint8 []"></param>
// <param name="ResourceType" type="string "></param>

// <param name="configurations" type="OMI_BaseResource "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) ResourceGet( /* IN */ ResourceType string,
	/* IN */ ModuleName string,
	/* IN */ resourceProperty []uint8,
	/* OUT */ configurations OMI_BaseResource) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ResourceGet", ResourceType, ModuleName, resourceProperty)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ModuleName" type="string "></param>
// <param name="resourceProperty" type="uint8 []"></param>
// <param name="ResourceType" type="string "></param>

// <param name="RebootRequired" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) ResourceSet( /* IN */ ResourceType string,
	/* IN */ ModuleName string,
	/* IN */ resourceProperty []uint8,
	/* OUT */ RebootRequired bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ResourceSet", ResourceType, ModuleName, resourceProperty)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ModuleName" type="string "></param>
// <param name="resourceProperty" type="uint8 []"></param>
// <param name="ResourceType" type="string "></param>

// <param name="InDesiredState" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) ResourceTest( /* IN */ ResourceType string,
	/* IN */ ModuleName string,
	/* IN */ resourceProperty []uint8,
	/* OUT */ InDesiredState bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ResourceTest", ResourceType, ModuleName, resourceProperty)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="BreakAll" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) EnableDebugConfiguration( /* IN */ BreakAll bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableDebugConfiguration", BreakAll)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DSCLocalConfigurationManager) DisableDebugConfiguration() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DisableDebugConfiguration")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

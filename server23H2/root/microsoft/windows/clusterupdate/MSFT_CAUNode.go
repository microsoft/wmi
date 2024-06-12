// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CAUNode struct
type MSFT_CAUNode struct {
	*cim.WmiInstance

	//
	OrchestratorGuid string
}

func NewMSFT_CAUNodeEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAUNode, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CAUNode{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CAUNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CAUNode, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CAUNode{
		WmiInstance: tmp,
	}
	return
}

// SetOrchestratorGuid sets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAUNode) SetPropertyOrchestratorGuid(value string) (err error) {
	return instance.SetProperty("OrchestratorGuid", (value))
}

// GetOrchestratorGuid gets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAUNode) GetPropertyOrchestratorGuid() (value string, err error) {
	retValue, err := instance.GetProperty("OrchestratorGuid")
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

// <param name="ReturnValue" type="bool "></param>
func (instance *MSFT_CAUNode) RebootRequired() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RebootRequired")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

//

// <param name="FeatureUpdatesOnly" type="bool "></param>
// <param name="IncludeRecommendedUpdates" type="bool "></param>
// <param name="QueryString" type="string "></param>

// <param name="Info" type="MSFT_CAU_ScanUpdateInfo []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) ScanUpdates( /* IN */ QueryString string,
	/* IN */ IncludeRecommendedUpdates bool,
	/* IN */ FeatureUpdatesOnly bool,
	/* OUT */ Info []MSFT_CAU_ScanUpdateInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ScanUpdates", QueryString, IncludeRecommendedUpdates, FeatureUpdatesOnly)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FeatureUpdatesOnly" type="bool "></param>
// <param name="IncludeRecommendedUpdates" type="bool "></param>
// <param name="QueryString" type="string "></param>

// <param name="Info" type="MSFT_CAU_DownloadUpdateInfo []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) DownloadUpdates( /* IN */ QueryString string,
	/* IN */ IncludeRecommendedUpdates bool,
	/* IN */ FeatureUpdatesOnly bool,
	/* OUT */ Info []MSFT_CAU_DownloadUpdateInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DownloadUpdates", QueryString, IncludeRecommendedUpdates, FeatureUpdatesOnly)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FeatureUpdatesOnly" type="bool "></param>
// <param name="IncludeRecommendedUpdates" type="bool "></param>
// <param name="QueryString" type="string "></param>

// <param name="Info" type="MSFT_CAU_InstallUpdateInfo []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) InstallUpdates( /* IN */ QueryString string,
	/* IN */ IncludeRecommendedUpdates bool,
	/* IN */ FeatureUpdatesOnly bool,
	/* OUT */ Info []MSFT_CAU_InstallUpdateInfo) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("InstallUpdates", QueryString, IncludeRecommendedUpdates, FeatureUpdatesOnly)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="QueryString" type="string "></param>

// <param name="Result" type="MSFT_CAU_CommitFUpdateWuaResult "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) CommitFUpdatesWua( /* IN */ QueryString string,
	/* OUT */ Result MSFT_CAU_CommitFUpdateWuaResult) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CommitFUpdatesWua", QueryString)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InstallerProgramPath" type="string "></param>
// <param name="Parameters" type="string "></param>
// <param name="RequireSmbEncryption" type="bool "></param>
// <param name="UpdatePath" type="string "></param>

// <param name="Result" type="MSFT_CAU_Update_Installer_Result "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) RunUpdateInstaller( /* IN */ InstallerProgramPath string,
	/* IN */ Parameters string,
	/* IN */ UpdatePath string,
	/* IN */ RequireSmbEncryption bool,
	/* OUT */ Result MSFT_CAU_Update_Installer_Result) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RunUpdateInstaller", InstallerProgramPath, Parameters, UpdatePath, RequireSmbEncryption)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CmdLineArgs" type="string "></param>
// <param name="PathToSetup" type="string "></param>

// <param name="Result" type="MSFT_CAU_Legacy_Upgrade_Install_Result "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) RunUpgradeLegacyMediaInterface( /* IN */ PathToSetup string,
	/* IN */ CmdLineArgs string,
	/* OUT */ Result MSFT_CAU_Legacy_Upgrade_Install_Result) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RunUpgradeLegacyMediaInterface", PathToSetup, CmdLineArgs)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PathToDeploymentCab" type="string "></param>
// <param name="PathToSetup" type="string "></param>

// <param name="Result" type="MSFT_CAU_Audi_Install_Result "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) RunUpgradeInstallAudiInterface( /* IN */ PathToSetup string,
	/* IN */ PathToDeploymentCab string,
	/* OUT */ Result MSFT_CAU_Audi_Install_Result) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RunUpgradeInstallAudiInterface", PathToSetup, PathToDeploymentCab)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PathToDeploymentCab" type="string "></param>
// <param name="PathToSetup" type="string "></param>
// <param name="UpdateId" type="string "></param>

// <param name="Result" type="MSFT_CAU_Audi_Commit_Result "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) RunUpgradeCommitAudiInterface( /* IN */ PathToSetup string,
	/* IN */ UpdateId string,
	/* IN */ PathToDeploymentCab string,
	/* OUT */ Result MSFT_CAU_Audi_Commit_Result) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RunUpgradeCommitAudiInterface", PathToSetup, UpdateId, PathToDeploymentCab)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PathToDeploymentCab" type="string "></param>
// <param name="PathToSetup" type="string "></param>
// <param name="UpdateId" type="string "></param>

// <param name="ErrorCode" type="MSFT_CAU_Audi_GetPostRebootResult_Result "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) GetPostRebootResult( /* IN */ PathToSetup string,
	/* IN */ UpdateId string,
	/* IN */ PathToDeploymentCab string,
	/* OUT */ ErrorCode MSFT_CAU_Audi_GetPostRebootResult_Result) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPostRebootResult", PathToSetup, UpdateId, PathToDeploymentCab)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PathToDeploymentCab" type="string "></param>
// <param name="PathToFiles" type="string "></param>

// <param name="Result" type="MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_CAUNode) GenerateDeviceInfoFilesFU( /* IN */ PathToFiles string,
	/* IN */ PathToDeploymentCab string,
	/* OUT */ Result MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GenerateDeviceInfoFilesFU", PathToFiles, PathToDeploymentCab)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

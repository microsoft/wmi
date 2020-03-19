// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetBranchCacheOrchestrator struct
type MSFT_NetBranchCacheOrchestrator struct {
	*CIM_ManagedElement
}

func NewMSFT_NetBranchCacheOrchestratorEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetBranchCacheOrchestrator, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheOrchestrator{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_NetBranchCacheOrchestratorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetBranchCacheOrchestrator, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetBranchCacheOrchestrator{
		CIM_ManagedElement: tmp,
	}
	return
}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCDistributed( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCDistributed", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="ServerNames" type="string []"></param>
// <param name="UseVersion" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCHostedClientByServerNames( /* IN */ ServerNames []string,
	/* IN */ UseVersion uint32,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCHostedClientByServerNames", ServerNames, UseVersion, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="UseSCP" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCHostedClientByUseSCP( /* IN */ UseSCP bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCHostedClientByUseSCP", UseSCP, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCLocal( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCLocal", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Disable_BC( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable_BC", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="DataCacheExtension" type="MSFT_NetBranchCacheDataCacheExtension []"></param>
// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Remove_BCDataCacheExtensionByDataCacheExtension( /* IN */ DataCacheExtension []MSFT_NetBranchCacheDataCacheExtension,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove_BCDataCacheExtensionByDataCacheExtension", DataCacheExtension, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Path" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Remove_BCDataCacheExtensionByPath( /* IN */ Path string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove_BCDataCacheExtensionByPath", Path, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Cache" type="MSFT_NetBranchCacheCache []"></param>
// <param name="Defragment" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="MoveTo" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Percentage" type="uint32 "></param>
// <param name="SizeBytes" type="uint64 "></param>

// <param name="cmdletOutput" type="MSFT_NetBranchCacheCache []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Set_BCCacheByCache( /* IN */ Cache []MSFT_NetBranchCacheCache,
	/* IN */ MoveTo string,
	/* IN */ Percentage uint32,
	/* IN */ SizeBytes uint64,
	/* IN */ Defragment bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput []MSFT_NetBranchCacheCache) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set_BCCacheByCache", Cache, MoveTo, Percentage, SizeBytes, Defragment, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Defragment" type="bool "></param>
// <param name="Force" type="bool "></param>
// <param name="MoveTo" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Path" type="string "></param>
// <param name="Percentage" type="uint32 "></param>
// <param name="SizeBytes" type="uint64 "></param>

// <param name="cmdletOutput" type="MSFT_NetBranchCacheCache "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Set_BCCacheByPath( /* IN */ Path string,
	/* IN */ MoveTo string,
	/* IN */ Percentage uint32,
	/* IN */ SizeBytes uint64,
	/* IN */ Defragment bool,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput MSFT_NetBranchCacheCache) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Set_BCCacheByPath", Path, MoveTo, Percentage, SizeBytes, Defragment, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Clear_BCCache( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Clear_BCCache", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Filename" type="string "></param>
// <param name="FilePassphrase" type="string "></param>
// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Import_BCSecretKey( /* IN */ Filename string,
	/* IN */ FilePassphrase string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Import_BCSecretKey", Filename, FilePassphrase, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Filename" type="string "></param>
// <param name="FilePassphrase" type="string "></param>
// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Export_BCSecretKey( /* IN */ Filename string,
	/* IN */ FilePassphrase string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Export_BCSecretKey", Filename, FilePassphrase, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Mode" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Set_BCAuthentication( /* IN */ Mode uint32,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set_BCAuthentication", Mode, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Version" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCDowngrading( /* IN */ Version uint32,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCDowngrading", Version, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCServeOnBattery( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCServeOnBattery", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Disable_BCServeOnBattery( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable_BCServeOnBattery", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="LatencyMilliseconds" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Set_BCMinSMBLatency( /* IN */ LatencyMilliseconds uint32,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set_BCMinSMBLatency", LatencyMilliseconds, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Passphrase" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Set_BCSecretKey( /* IN */ Passphrase string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set_BCSecretKey", Passphrase, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Path" type="string "></param>
// <param name="Percentage" type="uint32 "></param>

// <param name="cmdletOutput" type="MSFT_NetBranchCacheDataCacheExtension "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Add_BCDataCacheExtensionByPercentage( /* IN */ Percentage uint32,
	/* IN */ Path string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput MSFT_NetBranchCacheDataCacheExtension) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add_BCDataCacheExtensionByPercentage", Percentage, Path, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="Path" type="string "></param>
// <param name="SizeBytes" type="uint64 "></param>

// <param name="cmdletOutput" type="MSFT_NetBranchCacheDataCacheExtension "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Add_BCDataCacheExtensionBySizeBytes( /* IN */ SizeBytes uint64,
	/* IN */ Path string,
	/* IN */ PassThru bool,
	/* IN */ Force bool,
	/* OUT */ cmdletOutput MSFT_NetBranchCacheDataCacheExtension) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Add_BCDataCacheExtensionBySizeBytes", SizeBytes, Path, PassThru, Force)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Disable_BCDowngrading( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable_BCDowngrading", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="RegisterSCP" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Enable_BCHostedServer( /* IN */ RegisterSCP bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable_BCHostedServer", RegisterSCP, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="ResetFWRulesOnly" type="bool "></param>
// <param name="ResetPerfCountersOnly" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Reset_BC( /* IN */ ResetFWRulesOnly bool,
	/* IN */ ResetPerfCountersOnly bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Reset_BC", ResetFWRulesOnly, ResetPerfCountersOnly, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Path" type="string []"></param>
// <param name="Recurse" type="bool "></param>
// <param name="ReferenceFile" type="string "></param>
// <param name="StageData" type="bool "></param>
// <param name="StagingPath" type="string "></param>
// <param name="UseVersion" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Publish_BCWebHashes( /* IN */ Path []string,
	/* IN */ UseVersion uint32,
	/* IN */ StageData bool,
	/* IN */ StagingPath string,
	/* IN */ ReferenceFile string,
	/* IN */ Recurse bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Publish_BCWebHashes", Path, UseVersion, StageData, StagingPath, ReferenceFile, Recurse, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Path" type="string []"></param>
// <param name="Recurse" type="bool "></param>
// <param name="ReferenceFile" type="string "></param>
// <param name="StageData" type="bool "></param>
// <param name="StagingPath" type="string "></param>
// <param name="UseVersion" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Publish_BCFileHashes( /* IN */ Path []string,
	/* IN */ UseVersion uint32,
	/* IN */ StageData bool,
	/* IN */ StagingPath string,
	/* IN */ ReferenceFile string,
	/* IN */ Recurse bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Publish_BCFileHashes", Path, UseVersion, StageData, StagingPath, ReferenceFile, Recurse, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Destination" type="string "></param>
// <param name="ExportDataCache" type="bool "></param>
// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Export_BCCachePackageByExportDataCache( /* IN */ ExportDataCache bool,
	/* IN */ Destination string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Export_BCCachePackageByExportDataCache", ExportDataCache, Destination, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Destination" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="OutputReferenceFile" type="string "></param>
// <param name="StagingPath" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Export_BCCachePackageByStagingPath( /* IN */ StagingPath string,
	/* IN */ Destination string,
	/* IN */ OutputReferenceFile string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Export_BCCachePackageByStagingPath", StagingPath, Destination, OutputReferenceFile, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="Path" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Import_BCCachePackage( /* IN */ Path string,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Import_BCCachePackage", Path, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Force" type="bool "></param>
// <param name="TimeDays" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetBranchCacheOrchestrator) Set_BCDataCacheEntryMaxAge( /* IN */ TimeDays uint32,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Set_BCDataCacheEntryMaxAge", TimeDays, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

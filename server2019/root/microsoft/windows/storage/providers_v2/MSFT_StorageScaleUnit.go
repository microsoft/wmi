// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

// MSFT_StorageScaleUnit struct
type MSFT_StorageScaleUnit struct {
	MSFT_StorageFaultDomain
}

//

// <param name="EnableMaintenanceMode" type="bool "></param>
// <param name="IgnoreDetachedVirtualDisks" type="bool "></param>
// <param name="Manufacturer" type="string "></param>
// <param name="Model" type="string "></param>
// <param name="Timeout" type="uint32 "></param>
// <param name="ValidateMaintenanceMode" type="bool "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_StorageScaleUnit) Maintenance( /* IN */ ValidateMaintenanceMode bool,
	/* IN */ EnableMaintenanceMode bool,
	/* IN */ Timeout uint32,
	/* IN */ Model string,
	/* IN */ Manufacturer string,
	/* IN */ IgnoreDetachedVirtualDisks bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Maintenance", ValidateMaintenanceMode, EnableMaintenanceMode, Timeout, Model, Manufacturer, IgnoreDetachedVirtualDisks)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

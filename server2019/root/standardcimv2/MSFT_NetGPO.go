// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetGPO struct
type MSFT_NetGPO struct {
	CIM_SettingData
}

//

// <param name="DomainController" type="string "></param>
// <param name="PolicyStore" type="string "></param>

// <param name="GPOSession" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetGPO) Open( /* IN */ PolicyStore string,
	/* IN */ DomainController string,
	/* OUT */ GPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Open", PolicyStore, DomainController)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GPOSession" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetGPO) Save( /* IN */ GPOSession string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Save", GPOSession)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

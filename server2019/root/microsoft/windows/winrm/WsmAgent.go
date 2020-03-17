// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.winrm
//////////////////////////////////////////////
package winrm

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// WsmAgent struct
type WsmAgent struct {
	cim.WmiInstance
}

//

// <param name="detail" type="string "></param>
// <param name="key" type="string "></param>

// <param name="data" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *WsmAgent) GetData( /* IN */ key string,
	/* IN */ detail string,
	/* OUT */ data string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetData", key, detail)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="data" type="string "></param>
// <param name="detail" type="string "></param>
// <param name="key" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WsmAgent) SetData( /* IN */ key string,
	/* IN */ detail string,
	/* IN */ data string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetData", key, detail, data)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="detail" type="string "></param>
// <param name="key" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *WsmAgent) RemoveData( /* IN */ key string,
	/* IN */ detail string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveData", key, detail)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

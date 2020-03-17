// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DtcTransactionsStatisticsTask struct
type MSFT_DtcTransactionsStatisticsTask struct {
	cim.WmiInstance
}

//

// <param name="DtcName" type="string "></param>

// <param name="cmdletOutput" type="DtcTransactionsStatistics "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DtcTransactionsStatisticsTask) Get( /* IN */ DtcName string,
	/* OUT */ cmdletOutput DtcTransactionsStatistics) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Get", DtcName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

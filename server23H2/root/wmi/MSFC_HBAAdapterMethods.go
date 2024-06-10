// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_HBAAdapterMethods struct
type MSFC_HBAAdapterMethods struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewMSFC_HBAAdapterMethodsEx1(instance *cim.WmiInstance) (newInstance *MSFC_HBAAdapterMethods, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAAdapterMethods{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_HBAAdapterMethodsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_HBAAdapterMethods, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_HBAAdapterMethods{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSFC_HBAAdapterMethods) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSFC_HBAAdapterMethods) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSFC_HBAAdapterMethods) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSFC_HBAAdapterMethods) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// <param name="DiscoveredPortIndex" type="uint32 "></param>
// <param name="PortIndex" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PortAttributes" type="MSFC_HBAPortAttributesResults "></param>
func (instance *MSFC_HBAAdapterMethods) GetDiscoveredPortAttributes( /* IN */ PortIndex uint32,
	/* IN */ DiscoveredPortIndex uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PortAttributes MSFC_HBAPortAttributesResults) (err error) {
	_, err = instance.InvokeMethod("GetDiscoveredPortAttributes", PortIndex, DiscoveredPortIndex)
	if err != nil {
		return
	}
	return

}

//

// <param name="wwn" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="PortAttributes" type="MSFC_HBAPortAttributesResults "></param>
func (instance *MSFC_HBAAdapterMethods) GetPortAttributesByWWN( /* IN */ wwn []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ PortAttributes MSFC_HBAPortAttributesResults) (err error) {
	_, err = instance.InvokeMethod("GetPortAttributesByWWN", wwn)
	if err != nil {
		return
	}
	return

}

func (instance *MSFC_HBAAdapterMethods) RefreshInformation() (err error) {
	_, err = instance.InvokeMethodWithReturn("RefreshInformation")
	if err != nil {
		return
	}
	return

}

//

// <param name="PortWWN" type="uint8 []"></param>
// <param name="RequestBuffer" type="uint8 []"></param>
// <param name="RequestBufferCount" type="uint32 "></param>

// <param name="ActualResponseBufferCount" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="TotalResponseBufferCount" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendCTPassThru( /* IN */ PortWWN []uint8,
	/* IN */ RequestBufferCount uint32,
	/* IN */ RequestBuffer []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalResponseBufferCount uint32,
	/* OUT */ ActualResponseBufferCount uint32,
	/* OUT */ ResponseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendCTPassThru", PortWWN, RequestBufferCount, RequestBuffer)
	if err != nil {
		return
	}
	return

}

//

// <param name="wwn" type="uint8 []"></param>
// <param name="wwntype" type="uint32 "></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ResponseBufferCount" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendRNID( /* IN */ wwn []uint8,
	/* IN */ wwntype uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ResponseBufferCount uint32,
	/* OUT */ ResponseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendRNID", wwn, wwntype)
	if err != nil {
		return
	}
	return

}

//

// <param name="DestFCID" type="uint32 "></param>
// <param name="DestWWN" type="uint8 []"></param>
// <param name="NodeIdDataFormat" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="ActualRspBufferSize" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="RspBuffer" type="uint8 []"></param>
// <param name="TotalRspBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendRNIDV2( /* IN */ PortWWN []uint8,
	/* IN */ DestWWN []uint8,
	/* IN */ DestFCID uint32,
	/* IN */ NodeIdDataFormat uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalRspBufferSize uint32,
	/* OUT */ ActualRspBufferSize uint32,
	/* OUT */ RspBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendRNIDV2", PortWWN, DestWWN, DestFCID, NodeIdDataFormat)
	if err != nil {
		return
	}
	return

}

//

// <param name="HBAStatus" type="uint32 "></param>
// <param name="MgmtInfo" type="HBAFC3MgmtInfo "></param>
func (instance *MSFC_HBAAdapterMethods) GetFC3MgmtInfo( /* OUT */ HBAStatus uint32,
	/* OUT */ MgmtInfo HBAFC3MgmtInfo) (err error) {
	_, err = instance.InvokeMethod("GetFC3MgmtInfo")
	if err != nil {
		return
	}
	return

}

//

// <param name="MgmtInfo" type="HBAFC3MgmtInfo "></param>

// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SetFC3MgmtInfo( /* IN */ MgmtInfo HBAFC3MgmtInfo,
	/* OUT */ HBAStatus uint32) (err error) {
	_, err = instance.InvokeMethod("SetFC3MgmtInfo", MgmtInfo)
	if err != nil {
		return
	}
	return

}

//

// <param name="agent_domain" type="uint32 "></param>
// <param name="AgentWWN" type="uint8 []"></param>
// <param name="portIndex" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="ActualRspBufferSize" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="RspBuffer" type="uint8 []"></param>
// <param name="TotalRspBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendRPL( /* IN */ PortWWN []uint8,
	/* IN */ AgentWWN []uint8,
	/* IN */ agent_domain uint32,
	/* IN */ portIndex uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalRspBufferSize uint32,
	/* OUT */ ActualRspBufferSize uint32,
	/* OUT */ RspBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendRPL", PortWWN, AgentWWN, agent_domain, portIndex)
	if err != nil {
		return
	}
	return

}

//

// <param name="AgentDomain" type="uint32 "></param>
// <param name="AgentWWN" type="uint8 []"></param>
// <param name="ObjectPortNumber" type="uint32 "></param>
// <param name="ObjectWWN" type="uint8 []"></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="ActualRspBufferSize" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="RspBuffer" type="uint8 []"></param>
// <param name="TotalRspBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendRPS( /* IN */ PortWWN []uint8,
	/* IN */ AgentWWN []uint8,
	/* IN */ ObjectWWN []uint8,
	/* IN */ AgentDomain uint32,
	/* IN */ ObjectPortNumber uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalRspBufferSize uint32,
	/* OUT */ ActualRspBufferSize uint32,
	/* OUT */ RspBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendRPS", PortWWN, AgentWWN, ObjectWWN, AgentDomain, ObjectPortNumber)
	if err != nil {
		return
	}
	return

}

//

// <param name="Domain" type="uint32 "></param>
// <param name="PortWWN" type="uint8 []"></param>
// <param name="WWN" type="uint8 []"></param>

// <param name="ActualRspBufferSize" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="RspBuffer" type="uint8 []"></param>
// <param name="TotalRspBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendSRL( /* IN */ PortWWN []uint8,
	/* IN */ WWN []uint8,
	/* IN */ Domain uint32,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalRspBufferSize uint32,
	/* OUT */ ActualRspBufferSize uint32,
	/* OUT */ RspBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendSRL", PortWWN, WWN, Domain)
	if err != nil {
		return
	}
	return

}

//

// <param name="DestWWN" type="uint8 []"></param>
// <param name="Function" type="uint8 "></param>
// <param name="SourceWWN" type="uint8 []"></param>
// <param name="Type" type="uint8 "></param>

// <param name="ActualRspBufferSize" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="RspBuffer" type="uint8 []"></param>
// <param name="TotalRspBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendLIRR( /* IN */ SourceWWN []uint8,
	/* IN */ DestWWN []uint8,
	/* IN */ Function uint8,
	/* IN */ Type uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalRspBufferSize uint32,
	/* OUT */ ActualRspBufferSize uint32,
	/* OUT */ RspBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendLIRR", SourceWWN, DestWWN, Function, Type)
	if err != nil {
		return
	}
	return

}

//

// <param name="FC4Type" type="uint8 "></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="FC4Statistics" type="MSFC_FC4STATISTICS "></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) GetFC4Statistics( /* IN */ PortWWN []uint8,
	/* IN */ FC4Type uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ FC4Statistics MSFC_FC4STATISTICS) (err error) {
	_, err = instance.InvokeMethod("GetFC4Statistics", PortWWN, FC4Type)
	if err != nil {
		return
	}
	return

}

//

// <param name="ScsiId" type="HBAScsiID "></param>

// <param name="FC4Statistics" type="MSFC_FC4STATISTICS "></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) GetFCPStatistics( /* IN */ ScsiId HBAScsiID,
	/* OUT */ HBAStatus uint32,
	/* OUT */ FC4Statistics MSFC_FC4STATISTICS) (err error) {
	_, err = instance.InvokeMethod("GetFCPStatistics", ScsiId)
	if err != nil {
		return
	}
	return

}

//

// <param name="Cdb" type="uint8 []"></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="FcLun" type="uint64 "></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ResponseBufferSize" type="uint32 "></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
// <param name="SenseBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) ScsiInquiry( /* IN */ Cdb []uint8,
	/* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ FcLun uint64,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ResponseBufferSize uint32,
	/* OUT */ SenseBufferSize uint32,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ ResponseBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("ScsiInquiry", Cdb, HbaPortWWN, DiscoveredPortWWN, FcLun)
	if err != nil {
		return
	}
	return

}

//

// <param name="Cdb" type="uint8 []"></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="FcLun" type="uint64 "></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ResponseBufferSize" type="uint32 "></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
// <param name="SenseBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) ScsiReadCapacity( /* IN */ Cdb []uint8,
	/* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* IN */ FcLun uint64,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ResponseBufferSize uint32,
	/* OUT */ SenseBufferSize uint32,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ ResponseBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("ScsiReadCapacity", Cdb, HbaPortWWN, DiscoveredPortWWN, FcLun)
	if err != nil {
		return
	}
	return

}

//

// <param name="Cdb" type="uint8 []"></param>
// <param name="DiscoveredPortWWN" type="uint8 []"></param>
// <param name="HbaPortWWN" type="uint8 []"></param>

// <param name="HBAStatus" type="uint32 "></param>
// <param name="ResponseBuffer" type="uint8 []"></param>
// <param name="ResponseBufferSize" type="uint32 "></param>
// <param name="ScsiStatus" type="uint8 "></param>
// <param name="SenseBuffer" type="uint8 []"></param>
// <param name="SenseBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) ScsiReportLuns( /* IN */ Cdb []uint8,
	/* IN */ HbaPortWWN []uint8,
	/* IN */ DiscoveredPortWWN []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ ResponseBufferSize uint32,
	/* OUT */ SenseBufferSize uint32,
	/* OUT */ ScsiStatus uint8,
	/* OUT */ ResponseBuffer []uint8,
	/* OUT */ SenseBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("ScsiReportLuns", Cdb, HbaPortWWN, DiscoveredPortWWN)
	if err != nil {
		return
	}
	return

}

//

// <param name="EventCount" type="uint32 "></param>
// <param name="Events" type="MSFC_EventBuffer []"></param>
// <param name="HBAStatus" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) GetEventBuffer( /* OUT */ HBAStatus uint32,
	/* OUT */ EventCount uint32,
	/* OUT */ Events []MSFC_EventBuffer) (err error) {
	_, err = instance.InvokeMethod("GetEventBuffer")
	if err != nil {
		return
	}
	return

}

//

// <param name="DestWWN" type="uint8 []"></param>
// <param name="PortWWN" type="uint8 []"></param>

// <param name="ActualRspBufferSize" type="uint32 "></param>
// <param name="HBAStatus" type="uint32 "></param>
// <param name="RspBuffer" type="uint8 []"></param>
// <param name="TotalRspBufferSize" type="uint32 "></param>
func (instance *MSFC_HBAAdapterMethods) SendRLS( /* IN */ PortWWN []uint8,
	/* IN */ DestWWN []uint8,
	/* OUT */ HBAStatus uint32,
	/* OUT */ TotalRspBufferSize uint32,
	/* OUT */ ActualRspBufferSize uint32,
	/* OUT */ RspBuffer []uint8) (err error) {
	_, err = instance.InvokeMethod("SendRLS", PortWWN, DestWWN)
	if err != nil {
		return
	}
	return

}

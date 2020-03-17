// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_ClusteringService struct
type CIM_ClusteringService struct {
	CIM_Service
}

//

// <param name="CS" type="CIM_ComputerSystem "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ClusteringService) AddNode( /* IN */ CS CIM_ComputerSystem) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddNode", CS)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CS" type="CIM_ComputerSystem "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_ClusteringService) EvictNode( /* IN */ CS CIM_ComputerSystem) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EvictNode", CS)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

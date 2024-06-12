// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source BusError_RequestType
//////////////////////////////////////////////
package wmi

// BusError_RequestType
type BusError_RequestType int

const (
	// Generic_Error enum
	BusError_RequestType_Generic_Error BusError_RequestType = 0
	// Generic_Read enum
	BusError_RequestType_Generic_Read BusError_RequestType = 1
	// Generic_Write enum
	BusError_RequestType_Generic_Write BusError_RequestType = 2
	// Data_Read enum
	BusError_RequestType_Data_Read BusError_RequestType = 3
	// Data_Write enum
	BusError_RequestType_Data_Write BusError_RequestType = 4
	// Instruction_Fetch enum
	BusError_RequestType_Instruction_Fetch BusError_RequestType = 5
	// Prefetch enum
	BusError_RequestType_Prefetch BusError_RequestType = 6
	// Injection enum
	BusError_RequestType_Injection BusError_RequestType = 7
	// Snoop enum
	BusError_RequestType_Snoop BusError_RequestType = 8
)

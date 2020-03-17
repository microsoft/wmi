// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source AlertIndication_ProbableCause
//////////////////////////////////////////////
package v2

// AlertIndication_ProbableCause
type AlertIndication_ProbableCause int

const (
	// Unknown enum
	AlertIndication_ProbableCause_Unknown AlertIndication_ProbableCause = 0
	// Other enum
	AlertIndication_ProbableCause_Other AlertIndication_ProbableCause = 1
	// Adapter_Card_Error enum
	AlertIndication_ProbableCause_Adapter_Card_Error AlertIndication_ProbableCause = 2
	// Application_Subsystem_Failure enum
	AlertIndication_ProbableCause_Application_Subsystem_Failure AlertIndication_ProbableCause = 3
	// Bandwidth_Reduced enum
	AlertIndication_ProbableCause_Bandwidth_Reduced AlertIndication_ProbableCause = 4
	// Connection_Establishment_Error enum
	AlertIndication_ProbableCause_Connection_Establishment_Error AlertIndication_ProbableCause = 5
	// Communications_Protocol_Error enum
	AlertIndication_ProbableCause_Communications_Protocol_Error AlertIndication_ProbableCause = 6
	// Communications_Subsystem_Failure enum
	AlertIndication_ProbableCause_Communications_Subsystem_Failure AlertIndication_ProbableCause = 7
	// Configuration_Customization_Error enum
	AlertIndication_ProbableCause_Configuration_Customization_Error AlertIndication_ProbableCause = 8
	// Congestion enum
	AlertIndication_ProbableCause_Congestion AlertIndication_ProbableCause = 9
	// Corrupt_Data enum
	AlertIndication_ProbableCause_Corrupt_Data AlertIndication_ProbableCause = 10
	// CPU_Cycles_Limit_Exceeded enum
	AlertIndication_ProbableCause_CPU_Cycles_Limit_Exceeded AlertIndication_ProbableCause = 11
	// Dataset_Modem_Error enum
	AlertIndication_ProbableCause_Dataset_Modem_Error AlertIndication_ProbableCause = 12
	// Degraded_Signal enum
	AlertIndication_ProbableCause_Degraded_Signal AlertIndication_ProbableCause = 13
	// DTE_DCE_Interface_Error enum
	AlertIndication_ProbableCause_DTE_DCE_Interface_Error AlertIndication_ProbableCause = 14
	// Enclosure_Door_Open enum
	AlertIndication_ProbableCause_Enclosure_Door_Open AlertIndication_ProbableCause = 15
	// Equipment_Malfunction enum
	AlertIndication_ProbableCause_Equipment_Malfunction AlertIndication_ProbableCause = 16
	// Excessive_Vibration enum
	AlertIndication_ProbableCause_Excessive_Vibration AlertIndication_ProbableCause = 17
	// File_Format_Error enum
	AlertIndication_ProbableCause_File_Format_Error AlertIndication_ProbableCause = 18
	// Fire_Detected enum
	AlertIndication_ProbableCause_Fire_Detected AlertIndication_ProbableCause = 19
	// Flood_Detected enum
	AlertIndication_ProbableCause_Flood_Detected AlertIndication_ProbableCause = 20
	// Framing_Error enum
	AlertIndication_ProbableCause_Framing_Error AlertIndication_ProbableCause = 21
	// HVAC_Problem enum
	AlertIndication_ProbableCause_HVAC_Problem AlertIndication_ProbableCause = 22
	// Humidity_Unacceptable enum
	AlertIndication_ProbableCause_Humidity_Unacceptable AlertIndication_ProbableCause = 23
	// I_O_Device_Error enum
	AlertIndication_ProbableCause_I_O_Device_Error AlertIndication_ProbableCause = 24
	// Input_Device_Error enum
	AlertIndication_ProbableCause_Input_Device_Error AlertIndication_ProbableCause = 25
	// LAN_Error enum
	AlertIndication_ProbableCause_LAN_Error AlertIndication_ProbableCause = 26
	// Non_Toxic_Leak_Detected enum
	AlertIndication_ProbableCause_Non_Toxic_Leak_Detected AlertIndication_ProbableCause = 27
	// Local_Node_Transmission_Error enum
	AlertIndication_ProbableCause_Local_Node_Transmission_Error AlertIndication_ProbableCause = 28
	// Loss_of_Frame enum
	AlertIndication_ProbableCause_Loss_of_Frame AlertIndication_ProbableCause = 29
	// Loss_of_Signal enum
	AlertIndication_ProbableCause_Loss_of_Signal AlertIndication_ProbableCause = 30
	// Material_Supply_Exhausted enum
	AlertIndication_ProbableCause_Material_Supply_Exhausted AlertIndication_ProbableCause = 31
	// Multiplexer_Problem enum
	AlertIndication_ProbableCause_Multiplexer_Problem AlertIndication_ProbableCause = 32
	// Out_of_Memory enum
	AlertIndication_ProbableCause_Out_of_Memory AlertIndication_ProbableCause = 33
	// Output_Device_Error enum
	AlertIndication_ProbableCause_Output_Device_Error AlertIndication_ProbableCause = 34
	// Performance_Degraded enum
	AlertIndication_ProbableCause_Performance_Degraded AlertIndication_ProbableCause = 35
	// Power_Problem enum
	AlertIndication_ProbableCause_Power_Problem AlertIndication_ProbableCause = 36
	// Pressure_Unacceptable enum
	AlertIndication_ProbableCause_Pressure_Unacceptable AlertIndication_ProbableCause = 37
	// Processor_Problem__Internal_Machine_Error_ enum
	AlertIndication_ProbableCause_Processor_Problem__Internal_Machine_Error_ AlertIndication_ProbableCause = 38
	// Pump_Failure enum
	AlertIndication_ProbableCause_Pump_Failure AlertIndication_ProbableCause = 39
	// Queue_Size_Exceeded enum
	AlertIndication_ProbableCause_Queue_Size_Exceeded AlertIndication_ProbableCause = 40
	// Receive_Failure enum
	AlertIndication_ProbableCause_Receive_Failure AlertIndication_ProbableCause = 41
	// Receiver_Failure enum
	AlertIndication_ProbableCause_Receiver_Failure AlertIndication_ProbableCause = 42
	// Remote_Node_Transmission_Error enum
	AlertIndication_ProbableCause_Remote_Node_Transmission_Error AlertIndication_ProbableCause = 43
	// Resource_at_or_Nearing_Capacity enum
	AlertIndication_ProbableCause_Resource_at_or_Nearing_Capacity AlertIndication_ProbableCause = 44
	// Response_Time_Excessive enum
	AlertIndication_ProbableCause_Response_Time_Excessive AlertIndication_ProbableCause = 45
	// Retransmission_Rate_Excessive enum
	AlertIndication_ProbableCause_Retransmission_Rate_Excessive AlertIndication_ProbableCause = 46
	// Software_Error enum
	AlertIndication_ProbableCause_Software_Error AlertIndication_ProbableCause = 47
	// Software_Program_Abnormally_Terminated enum
	AlertIndication_ProbableCause_Software_Program_Abnormally_Terminated AlertIndication_ProbableCause = 48
	// Software_Program_Error__Incorrect_Results_ enum
	AlertIndication_ProbableCause_Software_Program_Error__Incorrect_Results_ AlertIndication_ProbableCause = 49
	// Storage_Capacity_Problem enum
	AlertIndication_ProbableCause_Storage_Capacity_Problem AlertIndication_ProbableCause = 50
	// Temperature_Unacceptable enum
	AlertIndication_ProbableCause_Temperature_Unacceptable AlertIndication_ProbableCause = 51
	// Threshold_Crossed enum
	AlertIndication_ProbableCause_Threshold_Crossed AlertIndication_ProbableCause = 52
	// Timing_Problem enum
	AlertIndication_ProbableCause_Timing_Problem AlertIndication_ProbableCause = 53
	// Toxic_Leak_Detected enum
	AlertIndication_ProbableCause_Toxic_Leak_Detected AlertIndication_ProbableCause = 54
	// Transmit_Failure enum
	AlertIndication_ProbableCause_Transmit_Failure AlertIndication_ProbableCause = 55
	// Transmitter_Failure enum
	AlertIndication_ProbableCause_Transmitter_Failure AlertIndication_ProbableCause = 56
	// Underlying_Resource_Unavailable enum
	AlertIndication_ProbableCause_Underlying_Resource_Unavailable AlertIndication_ProbableCause = 57
	// Version_MisMatch enum
	AlertIndication_ProbableCause_Version_MisMatch AlertIndication_ProbableCause = 58
	// Previous_Alert_Cleared enum
	AlertIndication_ProbableCause_Previous_Alert_Cleared AlertIndication_ProbableCause = 59
	// Login_Attempts_Failed enum
	AlertIndication_ProbableCause_Login_Attempts_Failed AlertIndication_ProbableCause = 60
	// Software_Virus_Detected enum
	AlertIndication_ProbableCause_Software_Virus_Detected AlertIndication_ProbableCause = 61
	// Hardware_Security_Breached enum
	AlertIndication_ProbableCause_Hardware_Security_Breached AlertIndication_ProbableCause = 62
	// Denial_of_Service_Detected enum
	AlertIndication_ProbableCause_Denial_of_Service_Detected AlertIndication_ProbableCause = 63
	// Security_Credential_MisMatch enum
	AlertIndication_ProbableCause_Security_Credential_MisMatch AlertIndication_ProbableCause = 64
	// Unauthorized_Access enum
	AlertIndication_ProbableCause_Unauthorized_Access AlertIndication_ProbableCause = 65
	// Alarm_Received enum
	AlertIndication_ProbableCause_Alarm_Received AlertIndication_ProbableCause = 66
	// Loss_of_Pointer enum
	AlertIndication_ProbableCause_Loss_of_Pointer AlertIndication_ProbableCause = 67
	// Payload_Mismatch enum
	AlertIndication_ProbableCause_Payload_Mismatch AlertIndication_ProbableCause = 68
	// Transmission_Error enum
	AlertIndication_ProbableCause_Transmission_Error AlertIndication_ProbableCause = 69
	// Excessive_Error_Rate enum
	AlertIndication_ProbableCause_Excessive_Error_Rate AlertIndication_ProbableCause = 70
	// Trace_Problem enum
	AlertIndication_ProbableCause_Trace_Problem AlertIndication_ProbableCause = 71
	// Element_Unavailable enum
	AlertIndication_ProbableCause_Element_Unavailable AlertIndication_ProbableCause = 72
	// Element_Missing enum
	AlertIndication_ProbableCause_Element_Missing AlertIndication_ProbableCause = 73
	// Loss_of_Multi_Frame enum
	AlertIndication_ProbableCause_Loss_of_Multi_Frame AlertIndication_ProbableCause = 74
	// Broadcast_Channel_Failure enum
	AlertIndication_ProbableCause_Broadcast_Channel_Failure AlertIndication_ProbableCause = 75
	// Invalid_Message_Received enum
	AlertIndication_ProbableCause_Invalid_Message_Received AlertIndication_ProbableCause = 76
	// Routing_Failure enum
	AlertIndication_ProbableCause_Routing_Failure AlertIndication_ProbableCause = 77
	// Backplane_Failure enum
	AlertIndication_ProbableCause_Backplane_Failure AlertIndication_ProbableCause = 78
	// Identifier_Duplication enum
	AlertIndication_ProbableCause_Identifier_Duplication AlertIndication_ProbableCause = 79
	// Protection_Path_Failure enum
	AlertIndication_ProbableCause_Protection_Path_Failure AlertIndication_ProbableCause = 80
	// Sync_Loss_or_Mismatch enum
	AlertIndication_ProbableCause_Sync_Loss_or_Mismatch AlertIndication_ProbableCause = 81
	// Terminal_Problem enum
	AlertIndication_ProbableCause_Terminal_Problem AlertIndication_ProbableCause = 82
	// Real_Time_Clock_Failure enum
	AlertIndication_ProbableCause_Real_Time_Clock_Failure AlertIndication_ProbableCause = 83
	// Antenna_Failure enum
	AlertIndication_ProbableCause_Antenna_Failure AlertIndication_ProbableCause = 84
	// Battery_Charging_Failure enum
	AlertIndication_ProbableCause_Battery_Charging_Failure AlertIndication_ProbableCause = 85
	// Disk_Failure enum
	AlertIndication_ProbableCause_Disk_Failure AlertIndication_ProbableCause = 86
	// Frequency_Hopping_Failure enum
	AlertIndication_ProbableCause_Frequency_Hopping_Failure AlertIndication_ProbableCause = 87
	// Loss_of_Redundancy enum
	AlertIndication_ProbableCause_Loss_of_Redundancy AlertIndication_ProbableCause = 88
	// Power_Supply_Failure enum
	AlertIndication_ProbableCause_Power_Supply_Failure AlertIndication_ProbableCause = 89
	// Signal_Quality_Problem enum
	AlertIndication_ProbableCause_Signal_Quality_Problem AlertIndication_ProbableCause = 90
	// Battery_Discharging enum
	AlertIndication_ProbableCause_Battery_Discharging AlertIndication_ProbableCause = 91
	// Battery_Failure enum
	AlertIndication_ProbableCause_Battery_Failure AlertIndication_ProbableCause = 92
	// Commercial_Power_Problem enum
	AlertIndication_ProbableCause_Commercial_Power_Problem AlertIndication_ProbableCause = 93
	// Fan_Failure enum
	AlertIndication_ProbableCause_Fan_Failure AlertIndication_ProbableCause = 94
	// Engine_Failure enum
	AlertIndication_ProbableCause_Engine_Failure AlertIndication_ProbableCause = 95
	// Sensor_Failure enum
	AlertIndication_ProbableCause_Sensor_Failure AlertIndication_ProbableCause = 96
	// Fuse_Failure enum
	AlertIndication_ProbableCause_Fuse_Failure AlertIndication_ProbableCause = 97
	// Generator_Failure enum
	AlertIndication_ProbableCause_Generator_Failure AlertIndication_ProbableCause = 98
	// Low_Battery enum
	AlertIndication_ProbableCause_Low_Battery AlertIndication_ProbableCause = 99
	// Low_Fuel enum
	AlertIndication_ProbableCause_Low_Fuel AlertIndication_ProbableCause = 100
	// Low_Water enum
	AlertIndication_ProbableCause_Low_Water AlertIndication_ProbableCause = 101
	// Explosive_Gas enum
	AlertIndication_ProbableCause_Explosive_Gas AlertIndication_ProbableCause = 102
	// High_Winds enum
	AlertIndication_ProbableCause_High_Winds AlertIndication_ProbableCause = 103
	// Ice_Buildup enum
	AlertIndication_ProbableCause_Ice_Buildup AlertIndication_ProbableCause = 104
	// Smoke enum
	AlertIndication_ProbableCause_Smoke AlertIndication_ProbableCause = 105
	// Memory_Mismatch enum
	AlertIndication_ProbableCause_Memory_Mismatch AlertIndication_ProbableCause = 106
	// Out_of_CPU_Cycles enum
	AlertIndication_ProbableCause_Out_of_CPU_Cycles AlertIndication_ProbableCause = 107
	// Software_Environment_Problem enum
	AlertIndication_ProbableCause_Software_Environment_Problem AlertIndication_ProbableCause = 108
	// Software_Download_Failure enum
	AlertIndication_ProbableCause_Software_Download_Failure AlertIndication_ProbableCause = 109
	// Element_Reinitialized enum
	AlertIndication_ProbableCause_Element_Reinitialized AlertIndication_ProbableCause = 110
	// Timeout enum
	AlertIndication_ProbableCause_Timeout AlertIndication_ProbableCause = 111
	// Logging_Problems enum
	AlertIndication_ProbableCause_Logging_Problems AlertIndication_ProbableCause = 112
	// Leak_Detected enum
	AlertIndication_ProbableCause_Leak_Detected AlertIndication_ProbableCause = 113
	// Protection_Mechanism_Failure enum
	AlertIndication_ProbableCause_Protection_Mechanism_Failure AlertIndication_ProbableCause = 114
	// Protecting_Resource_Failure enum
	AlertIndication_ProbableCause_Protecting_Resource_Failure AlertIndication_ProbableCause = 115
	// Database_Inconsistency enum
	AlertIndication_ProbableCause_Database_Inconsistency AlertIndication_ProbableCause = 116
	// Authentication_Failure enum
	AlertIndication_ProbableCause_Authentication_Failure AlertIndication_ProbableCause = 117
	// Breach_of_Confidentiality enum
	AlertIndication_ProbableCause_Breach_of_Confidentiality AlertIndication_ProbableCause = 118
	// Cable_Tamper enum
	AlertIndication_ProbableCause_Cable_Tamper AlertIndication_ProbableCause = 119
	// Delayed_Information enum
	AlertIndication_ProbableCause_Delayed_Information AlertIndication_ProbableCause = 120
	// Duplicate_Information enum
	AlertIndication_ProbableCause_Duplicate_Information AlertIndication_ProbableCause = 121
	// Information_Missing enum
	AlertIndication_ProbableCause_Information_Missing AlertIndication_ProbableCause = 122
	// Information_Modification enum
	AlertIndication_ProbableCause_Information_Modification AlertIndication_ProbableCause = 123
	// Information_Out_of_Sequence enum
	AlertIndication_ProbableCause_Information_Out_of_Sequence AlertIndication_ProbableCause = 124
	// Key_Expired enum
	AlertIndication_ProbableCause_Key_Expired AlertIndication_ProbableCause = 125
	// Non_Repudiation_Failure enum
	AlertIndication_ProbableCause_Non_Repudiation_Failure AlertIndication_ProbableCause = 126
	// Out_of_Hours_Activity enum
	AlertIndication_ProbableCause_Out_of_Hours_Activity AlertIndication_ProbableCause = 127
	// Out_of_Service enum
	AlertIndication_ProbableCause_Out_of_Service AlertIndication_ProbableCause = 128
	// Procedural_Error enum
	AlertIndication_ProbableCause_Procedural_Error AlertIndication_ProbableCause = 129
	// Unexpected_Information enum
	AlertIndication_ProbableCause_Unexpected_Information AlertIndication_ProbableCause = 130
)

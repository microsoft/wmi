// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source ProtocolEndpoint_ProtocolIFType
//////////////////////////////////////////////
package v2

// ProtocolEndpoint_ProtocolIFType
type ProtocolEndpoint_ProtocolIFType int

const (
	// Unknown enum
	ProtocolEndpoint_ProtocolIFType_Unknown ProtocolEndpoint_ProtocolIFType = 0
	// Other enum
	ProtocolEndpoint_ProtocolIFType_Other ProtocolEndpoint_ProtocolIFType = 1
	// Regular_1822 enum
	ProtocolEndpoint_ProtocolIFType_Regular_1822 ProtocolEndpoint_ProtocolIFType = 2
	// HDH_1822 enum
	ProtocolEndpoint_ProtocolIFType_HDH_1822 ProtocolEndpoint_ProtocolIFType = 3
	// DDN_X_25 enum
	ProtocolEndpoint_ProtocolIFType_DDN_X_25 ProtocolEndpoint_ProtocolIFType = 4
	// RFC877_X_25 enum
	ProtocolEndpoint_ProtocolIFType_RFC877_X_25 ProtocolEndpoint_ProtocolIFType = 5
	// Ethernet_CSMA_CD enum
	ProtocolEndpoint_ProtocolIFType_Ethernet_CSMA_CD ProtocolEndpoint_ProtocolIFType = 6
	// ISO_802_3_CSMA_CD enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_3_CSMA_CD ProtocolEndpoint_ProtocolIFType = 7
	// ISO_802_4_Token_Bus enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_4_Token_Bus ProtocolEndpoint_ProtocolIFType = 8
	// ISO_802_5_Token_Ring enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_5_Token_Ring ProtocolEndpoint_ProtocolIFType = 9
	// ISO_802_6_MAN enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_6_MAN ProtocolEndpoint_ProtocolIFType = 10
	// StarLAN enum
	ProtocolEndpoint_ProtocolIFType_StarLAN ProtocolEndpoint_ProtocolIFType = 11
	// Proteon_10Mbit enum
	ProtocolEndpoint_ProtocolIFType_Proteon_10Mbit ProtocolEndpoint_ProtocolIFType = 12
	// Proteon_80Mbit enum
	ProtocolEndpoint_ProtocolIFType_Proteon_80Mbit ProtocolEndpoint_ProtocolIFType = 13
	// HyperChannel enum
	ProtocolEndpoint_ProtocolIFType_HyperChannel ProtocolEndpoint_ProtocolIFType = 14
	// FDDI enum
	ProtocolEndpoint_ProtocolIFType_FDDI ProtocolEndpoint_ProtocolIFType = 15
	// LAP_B enum
	ProtocolEndpoint_ProtocolIFType_LAP_B ProtocolEndpoint_ProtocolIFType = 16
	// SDLC enum
	ProtocolEndpoint_ProtocolIFType_SDLC ProtocolEndpoint_ProtocolIFType = 17
	// DS1 enum
	ProtocolEndpoint_ProtocolIFType_DS1 ProtocolEndpoint_ProtocolIFType = 18
	// E1 enum
	ProtocolEndpoint_ProtocolIFType_E1 ProtocolEndpoint_ProtocolIFType = 19
	// Basic_ISDN enum
	ProtocolEndpoint_ProtocolIFType_Basic_ISDN ProtocolEndpoint_ProtocolIFType = 20
	// Primary_ISDN enum
	ProtocolEndpoint_ProtocolIFType_Primary_ISDN ProtocolEndpoint_ProtocolIFType = 21
	// Proprietary_Point_to_Point_Serial enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Point_to_Point_Serial ProtocolEndpoint_ProtocolIFType = 22
	// PPP enum
	ProtocolEndpoint_ProtocolIFType_PPP ProtocolEndpoint_ProtocolIFType = 23
	// Software_Loopback enum
	ProtocolEndpoint_ProtocolIFType_Software_Loopback ProtocolEndpoint_ProtocolIFType = 24
	// EON enum
	ProtocolEndpoint_ProtocolIFType_EON ProtocolEndpoint_ProtocolIFType = 25
	// Ethernet_3Mbit enum
	ProtocolEndpoint_ProtocolIFType_Ethernet_3Mbit ProtocolEndpoint_ProtocolIFType = 26
	// NSIP enum
	ProtocolEndpoint_ProtocolIFType_NSIP ProtocolEndpoint_ProtocolIFType = 27
	// SLIP enum
	ProtocolEndpoint_ProtocolIFType_SLIP ProtocolEndpoint_ProtocolIFType = 28
	// Ultra enum
	ProtocolEndpoint_ProtocolIFType_Ultra ProtocolEndpoint_ProtocolIFType = 29
	// DS3 enum
	ProtocolEndpoint_ProtocolIFType_DS3 ProtocolEndpoint_ProtocolIFType = 30
	// SIP enum
	ProtocolEndpoint_ProtocolIFType_SIP ProtocolEndpoint_ProtocolIFType = 31
	// Frame_Relay enum
	ProtocolEndpoint_ProtocolIFType_Frame_Relay ProtocolEndpoint_ProtocolIFType = 32
	// RS_232 enum
	ProtocolEndpoint_ProtocolIFType_RS_232 ProtocolEndpoint_ProtocolIFType = 33
	// Parallel enum
	ProtocolEndpoint_ProtocolIFType_Parallel ProtocolEndpoint_ProtocolIFType = 34
	// ARCNet enum
	ProtocolEndpoint_ProtocolIFType_ARCNet ProtocolEndpoint_ProtocolIFType = 35
	// ARCNet_Plus enum
	ProtocolEndpoint_ProtocolIFType_ARCNet_Plus ProtocolEndpoint_ProtocolIFType = 36
	// ATM enum
	ProtocolEndpoint_ProtocolIFType_ATM ProtocolEndpoint_ProtocolIFType = 37
	// MIO_X_25 enum
	ProtocolEndpoint_ProtocolIFType_MIO_X_25 ProtocolEndpoint_ProtocolIFType = 38
	// SONET enum
	ProtocolEndpoint_ProtocolIFType_SONET ProtocolEndpoint_ProtocolIFType = 39
	// X_25_PLE enum
	ProtocolEndpoint_ProtocolIFType_X_25_PLE ProtocolEndpoint_ProtocolIFType = 40
	// ISO_802_211c enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_211c ProtocolEndpoint_ProtocolIFType = 41
	// LocalTalk enum
	ProtocolEndpoint_ProtocolIFType_LocalTalk ProtocolEndpoint_ProtocolIFType = 42
	// SMDS_DXI enum
	ProtocolEndpoint_ProtocolIFType_SMDS_DXI ProtocolEndpoint_ProtocolIFType = 43
	// Frame_Relay_Service enum
	ProtocolEndpoint_ProtocolIFType_Frame_Relay_Service ProtocolEndpoint_ProtocolIFType = 44
	// V_35 enum
	ProtocolEndpoint_ProtocolIFType_V_35 ProtocolEndpoint_ProtocolIFType = 45
	// HSSI enum
	ProtocolEndpoint_ProtocolIFType_HSSI ProtocolEndpoint_ProtocolIFType = 46
	// HIPPI enum
	ProtocolEndpoint_ProtocolIFType_HIPPI ProtocolEndpoint_ProtocolIFType = 47
	// Modem enum
	ProtocolEndpoint_ProtocolIFType_Modem ProtocolEndpoint_ProtocolIFType = 48
	// AAL5 enum
	ProtocolEndpoint_ProtocolIFType_AAL5 ProtocolEndpoint_ProtocolIFType = 49
	// SONET_Path enum
	ProtocolEndpoint_ProtocolIFType_SONET_Path ProtocolEndpoint_ProtocolIFType = 50
	// SONET_VT enum
	ProtocolEndpoint_ProtocolIFType_SONET_VT ProtocolEndpoint_ProtocolIFType = 51
	// SMDS_ICIP enum
	ProtocolEndpoint_ProtocolIFType_SMDS_ICIP ProtocolEndpoint_ProtocolIFType = 52
	// Proprietary_Virtual_Internal enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Virtual_Internal ProtocolEndpoint_ProtocolIFType = 53
	// Proprietary_Multiplexor enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Multiplexor ProtocolEndpoint_ProtocolIFType = 54
	// IEEE_802_12 enum
	ProtocolEndpoint_ProtocolIFType_IEEE_802_12 ProtocolEndpoint_ProtocolIFType = 55
	// Fibre_Channel enum
	ProtocolEndpoint_ProtocolIFType_Fibre_Channel ProtocolEndpoint_ProtocolIFType = 56
	// HIPPI_Interface enum
	ProtocolEndpoint_ProtocolIFType_HIPPI_Interface ProtocolEndpoint_ProtocolIFType = 57
	// Frame_Relay_Interconnect enum
	ProtocolEndpoint_ProtocolIFType_Frame_Relay_Interconnect ProtocolEndpoint_ProtocolIFType = 58
	// ATM_Emulated_LAN_for_802_3 enum
	ProtocolEndpoint_ProtocolIFType_ATM_Emulated_LAN_for_802_3 ProtocolEndpoint_ProtocolIFType = 59
	// ATM_Emulated_LAN_for_802_5 enum
	ProtocolEndpoint_ProtocolIFType_ATM_Emulated_LAN_for_802_5 ProtocolEndpoint_ProtocolIFType = 60
	// ATM_Emulated_Circuit enum
	ProtocolEndpoint_ProtocolIFType_ATM_Emulated_Circuit ProtocolEndpoint_ProtocolIFType = 61
	// Fast_Ethernet__100BaseT_ enum
	ProtocolEndpoint_ProtocolIFType_Fast_Ethernet__100BaseT_ ProtocolEndpoint_ProtocolIFType = 62
	// ISDN enum
	ProtocolEndpoint_ProtocolIFType_ISDN ProtocolEndpoint_ProtocolIFType = 63
	// V_11 enum
	ProtocolEndpoint_ProtocolIFType_V_11 ProtocolEndpoint_ProtocolIFType = 64
	// V_36 enum
	ProtocolEndpoint_ProtocolIFType_V_36 ProtocolEndpoint_ProtocolIFType = 65
	// G703_at_64K enum
	ProtocolEndpoint_ProtocolIFType_G703_at_64K ProtocolEndpoint_ProtocolIFType = 66
	// G703_at_2Mb enum
	ProtocolEndpoint_ProtocolIFType_G703_at_2Mb ProtocolEndpoint_ProtocolIFType = 67
	// QLLC enum
	ProtocolEndpoint_ProtocolIFType_QLLC ProtocolEndpoint_ProtocolIFType = 68
	// Fast_Ethernet_100BaseFX enum
	ProtocolEndpoint_ProtocolIFType_Fast_Ethernet_100BaseFX ProtocolEndpoint_ProtocolIFType = 69
	// Channel enum
	ProtocolEndpoint_ProtocolIFType_Channel ProtocolEndpoint_ProtocolIFType = 70
	// IEEE_802_11 enum
	ProtocolEndpoint_ProtocolIFType_IEEE_802_11 ProtocolEndpoint_ProtocolIFType = 71
	// IBM_260_370_OEMI_Channel enum
	ProtocolEndpoint_ProtocolIFType_IBM_260_370_OEMI_Channel ProtocolEndpoint_ProtocolIFType = 72
	// ESCON enum
	ProtocolEndpoint_ProtocolIFType_ESCON ProtocolEndpoint_ProtocolIFType = 73
	// Data_Link_Switching enum
	ProtocolEndpoint_ProtocolIFType_Data_Link_Switching ProtocolEndpoint_ProtocolIFType = 74
	// ISDN_S_T_Interface enum
	ProtocolEndpoint_ProtocolIFType_ISDN_S_T_Interface ProtocolEndpoint_ProtocolIFType = 75
	// ISDN_U_Interface enum
	ProtocolEndpoint_ProtocolIFType_ISDN_U_Interface ProtocolEndpoint_ProtocolIFType = 76
	// LAP_D enum
	ProtocolEndpoint_ProtocolIFType_LAP_D ProtocolEndpoint_ProtocolIFType = 77
	// IP_Switch enum
	ProtocolEndpoint_ProtocolIFType_IP_Switch ProtocolEndpoint_ProtocolIFType = 78
	// Remote_Source_Route_Bridging enum
	ProtocolEndpoint_ProtocolIFType_Remote_Source_Route_Bridging ProtocolEndpoint_ProtocolIFType = 79
	// ATM_Logical enum
	ProtocolEndpoint_ProtocolIFType_ATM_Logical ProtocolEndpoint_ProtocolIFType = 80
	// DS0 enum
	ProtocolEndpoint_ProtocolIFType_DS0 ProtocolEndpoint_ProtocolIFType = 81
	// DS0_Bundle enum
	ProtocolEndpoint_ProtocolIFType_DS0_Bundle ProtocolEndpoint_ProtocolIFType = 82
	// BSC enum
	ProtocolEndpoint_ProtocolIFType_BSC ProtocolEndpoint_ProtocolIFType = 83
	// Async enum
	ProtocolEndpoint_ProtocolIFType_Async ProtocolEndpoint_ProtocolIFType = 84
	// Combat_Net_Radio enum
	ProtocolEndpoint_ProtocolIFType_Combat_Net_Radio ProtocolEndpoint_ProtocolIFType = 85
	// ISO_802_5r_DTR enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_5r_DTR ProtocolEndpoint_ProtocolIFType = 86
	// Ext_Pos_Loc_Report_System enum
	ProtocolEndpoint_ProtocolIFType_Ext_Pos_Loc_Report_System ProtocolEndpoint_ProtocolIFType = 87
	// AppleTalk_Remote_Access_Protocol enum
	ProtocolEndpoint_ProtocolIFType_AppleTalk_Remote_Access_Protocol ProtocolEndpoint_ProtocolIFType = 88
	// Proprietary_Connectionless enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Connectionless ProtocolEndpoint_ProtocolIFType = 89
	// ITU_X_29_Host_PAD enum
	ProtocolEndpoint_ProtocolIFType_ITU_X_29_Host_PAD ProtocolEndpoint_ProtocolIFType = 90
	// ITU_X_3_Terminal_PAD enum
	ProtocolEndpoint_ProtocolIFType_ITU_X_3_Terminal_PAD ProtocolEndpoint_ProtocolIFType = 91
	// Frame_Relay_MPI enum
	ProtocolEndpoint_ProtocolIFType_Frame_Relay_MPI ProtocolEndpoint_ProtocolIFType = 92
	// ITU_X_213 enum
	ProtocolEndpoint_ProtocolIFType_ITU_X_213 ProtocolEndpoint_ProtocolIFType = 93
	// ADSL enum
	ProtocolEndpoint_ProtocolIFType_ADSL ProtocolEndpoint_ProtocolIFType = 94
	// RADSL enum
	ProtocolEndpoint_ProtocolIFType_RADSL ProtocolEndpoint_ProtocolIFType = 95
	// SDSL enum
	ProtocolEndpoint_ProtocolIFType_SDSL ProtocolEndpoint_ProtocolIFType = 96
	// VDSL enum
	ProtocolEndpoint_ProtocolIFType_VDSL ProtocolEndpoint_ProtocolIFType = 97
	// ISO_802_5_CRFP enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_5_CRFP ProtocolEndpoint_ProtocolIFType = 98
	// Myrinet enum
	ProtocolEndpoint_ProtocolIFType_Myrinet ProtocolEndpoint_ProtocolIFType = 99
	// Voice_Receive_and_Transmit enum
	ProtocolEndpoint_ProtocolIFType_Voice_Receive_and_Transmit ProtocolEndpoint_ProtocolIFType = 100
	// Voice_Foreign_Exchange_Office enum
	ProtocolEndpoint_ProtocolIFType_Voice_Foreign_Exchange_Office ProtocolEndpoint_ProtocolIFType = 101
	// Voice_Foreign_Exchange_Service enum
	ProtocolEndpoint_ProtocolIFType_Voice_Foreign_Exchange_Service ProtocolEndpoint_ProtocolIFType = 102
	// Voice_Encapsulation enum
	ProtocolEndpoint_ProtocolIFType_Voice_Encapsulation ProtocolEndpoint_ProtocolIFType = 103
	// Voice_over_IP enum
	ProtocolEndpoint_ProtocolIFType_Voice_over_IP ProtocolEndpoint_ProtocolIFType = 104
	// ATM_DXI enum
	ProtocolEndpoint_ProtocolIFType_ATM_DXI ProtocolEndpoint_ProtocolIFType = 105
	// ATM_FUNI enum
	ProtocolEndpoint_ProtocolIFType_ATM_FUNI ProtocolEndpoint_ProtocolIFType = 106
	// ATM_IMA enum
	ProtocolEndpoint_ProtocolIFType_ATM_IMA ProtocolEndpoint_ProtocolIFType = 107
	// PPP_Multilink_Bundle enum
	ProtocolEndpoint_ProtocolIFType_PPP_Multilink_Bundle ProtocolEndpoint_ProtocolIFType = 108
	// IP_over_CDLC enum
	ProtocolEndpoint_ProtocolIFType_IP_over_CDLC ProtocolEndpoint_ProtocolIFType = 109
	// IP_over_CLAW enum
	ProtocolEndpoint_ProtocolIFType_IP_over_CLAW ProtocolEndpoint_ProtocolIFType = 110
	// Stack_to_Stack enum
	ProtocolEndpoint_ProtocolIFType_Stack_to_Stack ProtocolEndpoint_ProtocolIFType = 111
	// Virtual_IP_Address enum
	ProtocolEndpoint_ProtocolIFType_Virtual_IP_Address ProtocolEndpoint_ProtocolIFType = 112
	// MPC enum
	ProtocolEndpoint_ProtocolIFType_MPC ProtocolEndpoint_ProtocolIFType = 113
	// IP_over_ATM enum
	ProtocolEndpoint_ProtocolIFType_IP_over_ATM ProtocolEndpoint_ProtocolIFType = 114
	// ISO_802_5j_Fibre_Token_Ring enum
	ProtocolEndpoint_ProtocolIFType_ISO_802_5j_Fibre_Token_Ring ProtocolEndpoint_ProtocolIFType = 115
	// TDLC enum
	ProtocolEndpoint_ProtocolIFType_TDLC ProtocolEndpoint_ProtocolIFType = 116
	// Gigabit_Ethernet enum
	ProtocolEndpoint_ProtocolIFType_Gigabit_Ethernet ProtocolEndpoint_ProtocolIFType = 117
	// HDLC enum
	ProtocolEndpoint_ProtocolIFType_HDLC ProtocolEndpoint_ProtocolIFType = 118
	// LAP_F enum
	ProtocolEndpoint_ProtocolIFType_LAP_F ProtocolEndpoint_ProtocolIFType = 119
	// V_37 enum
	ProtocolEndpoint_ProtocolIFType_V_37 ProtocolEndpoint_ProtocolIFType = 120
	// X_25_MLP enum
	ProtocolEndpoint_ProtocolIFType_X_25_MLP ProtocolEndpoint_ProtocolIFType = 121
	// X_25_Hunt_Group enum
	ProtocolEndpoint_ProtocolIFType_X_25_Hunt_Group ProtocolEndpoint_ProtocolIFType = 122
	// Transp_HDLC enum
	ProtocolEndpoint_ProtocolIFType_Transp_HDLC ProtocolEndpoint_ProtocolIFType = 123
	// Interleave_Channel enum
	ProtocolEndpoint_ProtocolIFType_Interleave_Channel ProtocolEndpoint_ProtocolIFType = 124
	// FAST_Channel enum
	ProtocolEndpoint_ProtocolIFType_FAST_Channel ProtocolEndpoint_ProtocolIFType = 125
	// IP__for_APPN_HPR_in_IP_Networks_ enum
	ProtocolEndpoint_ProtocolIFType_IP__for_APPN_HPR_in_IP_Networks_ ProtocolEndpoint_ProtocolIFType = 126
	// CATV_MAC_Layer enum
	ProtocolEndpoint_ProtocolIFType_CATV_MAC_Layer ProtocolEndpoint_ProtocolIFType = 127
	// CATV_Downstream enum
	ProtocolEndpoint_ProtocolIFType_CATV_Downstream ProtocolEndpoint_ProtocolIFType = 128
	// CATV_Upstream enum
	ProtocolEndpoint_ProtocolIFType_CATV_Upstream ProtocolEndpoint_ProtocolIFType = 129
	// Avalon_12MPP_Switch enum
	ProtocolEndpoint_ProtocolIFType_Avalon_12MPP_Switch ProtocolEndpoint_ProtocolIFType = 130
	// Tunnel enum
	ProtocolEndpoint_ProtocolIFType_Tunnel ProtocolEndpoint_ProtocolIFType = 131
	// Coffee enum
	ProtocolEndpoint_ProtocolIFType_Coffee ProtocolEndpoint_ProtocolIFType = 132
	// Circuit_Emulation_Service enum
	ProtocolEndpoint_ProtocolIFType_Circuit_Emulation_Service ProtocolEndpoint_ProtocolIFType = 133
	// ATM_SubInterface enum
	ProtocolEndpoint_ProtocolIFType_ATM_SubInterface ProtocolEndpoint_ProtocolIFType = 134
	// Layer_2_VLAN_using_802_1Q enum
	ProtocolEndpoint_ProtocolIFType_Layer_2_VLAN_using_802_1Q ProtocolEndpoint_ProtocolIFType = 135
	// Layer_3_VLAN_using_IP enum
	ProtocolEndpoint_ProtocolIFType_Layer_3_VLAN_using_IP ProtocolEndpoint_ProtocolIFType = 136
	// Layer_3_VLAN_using_IPX enum
	ProtocolEndpoint_ProtocolIFType_Layer_3_VLAN_using_IPX ProtocolEndpoint_ProtocolIFType = 137
	// Digital_Power_Line enum
	ProtocolEndpoint_ProtocolIFType_Digital_Power_Line ProtocolEndpoint_ProtocolIFType = 138
	// Multimedia_Mail_over_IP enum
	ProtocolEndpoint_ProtocolIFType_Multimedia_Mail_over_IP ProtocolEndpoint_ProtocolIFType = 139
	// DTM enum
	ProtocolEndpoint_ProtocolIFType_DTM ProtocolEndpoint_ProtocolIFType = 140
	// DCN enum
	ProtocolEndpoint_ProtocolIFType_DCN ProtocolEndpoint_ProtocolIFType = 141
	// IP_Forwarding enum
	ProtocolEndpoint_ProtocolIFType_IP_Forwarding ProtocolEndpoint_ProtocolIFType = 142
	// MSDSL enum
	ProtocolEndpoint_ProtocolIFType_MSDSL ProtocolEndpoint_ProtocolIFType = 143
	// IEEE_1394 enum
	ProtocolEndpoint_ProtocolIFType_IEEE_1394 ProtocolEndpoint_ProtocolIFType = 144
	// IF_GSN_HIPPI_6400 enum
	ProtocolEndpoint_ProtocolIFType_IF_GSN_HIPPI_6400 ProtocolEndpoint_ProtocolIFType = 145
	// DVB_RCC_MAC_Layer enum
	ProtocolEndpoint_ProtocolIFType_DVB_RCC_MAC_Layer ProtocolEndpoint_ProtocolIFType = 146
	// DVB_RCC_Downstream enum
	ProtocolEndpoint_ProtocolIFType_DVB_RCC_Downstream ProtocolEndpoint_ProtocolIFType = 147
	// DVB_RCC_Upstream enum
	ProtocolEndpoint_ProtocolIFType_DVB_RCC_Upstream ProtocolEndpoint_ProtocolIFType = 148
	// ATM_Virtual enum
	ProtocolEndpoint_ProtocolIFType_ATM_Virtual ProtocolEndpoint_ProtocolIFType = 149
	// MPLS_Tunnel enum
	ProtocolEndpoint_ProtocolIFType_MPLS_Tunnel ProtocolEndpoint_ProtocolIFType = 150
	// SRP enum
	ProtocolEndpoint_ProtocolIFType_SRP ProtocolEndpoint_ProtocolIFType = 151
	// Voice_over_ATM enum
	ProtocolEndpoint_ProtocolIFType_Voice_over_ATM ProtocolEndpoint_ProtocolIFType = 152
	// Voice_over_Frame_Relay enum
	ProtocolEndpoint_ProtocolIFType_Voice_over_Frame_Relay ProtocolEndpoint_ProtocolIFType = 153
	// ISDL enum
	ProtocolEndpoint_ProtocolIFType_ISDL ProtocolEndpoint_ProtocolIFType = 154
	// Composite_Link enum
	ProtocolEndpoint_ProtocolIFType_Composite_Link ProtocolEndpoint_ProtocolIFType = 155
	// SS7_Signaling_Link enum
	ProtocolEndpoint_ProtocolIFType_SS7_Signaling_Link ProtocolEndpoint_ProtocolIFType = 156
	// Proprietary_P2P_Wireless enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_P2P_Wireless ProtocolEndpoint_ProtocolIFType = 157
	// Frame_Forward enum
	ProtocolEndpoint_ProtocolIFType_Frame_Forward ProtocolEndpoint_ProtocolIFType = 158
	// RFC1483_Multiprotocol_over_ATM enum
	ProtocolEndpoint_ProtocolIFType_RFC1483_Multiprotocol_over_ATM ProtocolEndpoint_ProtocolIFType = 159
	// USB enum
	ProtocolEndpoint_ProtocolIFType_USB ProtocolEndpoint_ProtocolIFType = 160
	// IEEE_802_3ad_Link_Aggregate enum
	ProtocolEndpoint_ProtocolIFType_IEEE_802_3ad_Link_Aggregate ProtocolEndpoint_ProtocolIFType = 161
	// BGP_Policy_Accounting enum
	ProtocolEndpoint_ProtocolIFType_BGP_Policy_Accounting ProtocolEndpoint_ProtocolIFType = 162
	// FRF__16_Multilink_FR enum
	ProtocolEndpoint_ProtocolIFType_FRF__16_Multilink_FR ProtocolEndpoint_ProtocolIFType = 163
	// H_323_Gatekeeper enum
	ProtocolEndpoint_ProtocolIFType_H_323_Gatekeeper ProtocolEndpoint_ProtocolIFType = 164
	// H_323_Proxy enum
	ProtocolEndpoint_ProtocolIFType_H_323_Proxy ProtocolEndpoint_ProtocolIFType = 165
	// MPLS enum
	ProtocolEndpoint_ProtocolIFType_MPLS ProtocolEndpoint_ProtocolIFType = 166
	// Multi_Frequency_Signaling_Link enum
	ProtocolEndpoint_ProtocolIFType_Multi_Frequency_Signaling_Link ProtocolEndpoint_ProtocolIFType = 167
	// HDSL_2 enum
	ProtocolEndpoint_ProtocolIFType_HDSL_2 ProtocolEndpoint_ProtocolIFType = 168
	// S_HDSL enum
	ProtocolEndpoint_ProtocolIFType_S_HDSL ProtocolEndpoint_ProtocolIFType = 169
	// DS1_Facility_Data_Link enum
	ProtocolEndpoint_ProtocolIFType_DS1_Facility_Data_Link ProtocolEndpoint_ProtocolIFType = 170
	// Packet_over_SONET_SDH enum
	ProtocolEndpoint_ProtocolIFType_Packet_over_SONET_SDH ProtocolEndpoint_ProtocolIFType = 171
	// DVB_ASI_Input enum
	ProtocolEndpoint_ProtocolIFType_DVB_ASI_Input ProtocolEndpoint_ProtocolIFType = 172
	// DVB_ASI_Output enum
	ProtocolEndpoint_ProtocolIFType_DVB_ASI_Output ProtocolEndpoint_ProtocolIFType = 173
	// Power_Line enum
	ProtocolEndpoint_ProtocolIFType_Power_Line ProtocolEndpoint_ProtocolIFType = 174
	// Non_Facility_Associated_Signaling enum
	ProtocolEndpoint_ProtocolIFType_Non_Facility_Associated_Signaling ProtocolEndpoint_ProtocolIFType = 175
	// TR008 enum
	ProtocolEndpoint_ProtocolIFType_TR008 ProtocolEndpoint_ProtocolIFType = 176
	// GR303_RDT enum
	ProtocolEndpoint_ProtocolIFType_GR303_RDT ProtocolEndpoint_ProtocolIFType = 177
	// GR303_IDT enum
	ProtocolEndpoint_ProtocolIFType_GR303_IDT ProtocolEndpoint_ProtocolIFType = 178
	// ISUP enum
	ProtocolEndpoint_ProtocolIFType_ISUP ProtocolEndpoint_ProtocolIFType = 179
	// Proprietary_Wireless_MAC_Layer enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Wireless_MAC_Layer ProtocolEndpoint_ProtocolIFType = 180
	// Proprietary_Wireless_Downstream enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Wireless_Downstream ProtocolEndpoint_ProtocolIFType = 181
	// Proprietary_Wireless_Upstream enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Wireless_Upstream ProtocolEndpoint_ProtocolIFType = 182
	// HIPERLAN_Type_2 enum
	ProtocolEndpoint_ProtocolIFType_HIPERLAN_Type_2 ProtocolEndpoint_ProtocolIFType = 183
	// Proprietary_Broadband_Wireless_Access_Point_to_Mulipoint enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_Broadband_Wireless_Access_Point_to_Mulipoint ProtocolEndpoint_ProtocolIFType = 184
	// SONET_Overhead_Channel enum
	ProtocolEndpoint_ProtocolIFType_SONET_Overhead_Channel ProtocolEndpoint_ProtocolIFType = 185
	// Digital_Wrapper_Overhead_Channel enum
	ProtocolEndpoint_ProtocolIFType_Digital_Wrapper_Overhead_Channel ProtocolEndpoint_ProtocolIFType = 186
	// ATM_Adaptation_Layer_2 enum
	ProtocolEndpoint_ProtocolIFType_ATM_Adaptation_Layer_2 ProtocolEndpoint_ProtocolIFType = 187
	// Radio_MAC enum
	ProtocolEndpoint_ProtocolIFType_Radio_MAC ProtocolEndpoint_ProtocolIFType = 188
	// ATM_Radio enum
	ProtocolEndpoint_ProtocolIFType_ATM_Radio ProtocolEndpoint_ProtocolIFType = 189
	// Inter_Machine_Trunk enum
	ProtocolEndpoint_ProtocolIFType_Inter_Machine_Trunk ProtocolEndpoint_ProtocolIFType = 190
	// MVL_DSL enum
	ProtocolEndpoint_ProtocolIFType_MVL_DSL ProtocolEndpoint_ProtocolIFType = 191
	// Long_Read_DSL enum
	ProtocolEndpoint_ProtocolIFType_Long_Read_DSL ProtocolEndpoint_ProtocolIFType = 192
	// Frame_Relay_DLCI_Endpoint enum
	ProtocolEndpoint_ProtocolIFType_Frame_Relay_DLCI_Endpoint ProtocolEndpoint_ProtocolIFType = 193
	// ATM_VCI_Endpoint enum
	ProtocolEndpoint_ProtocolIFType_ATM_VCI_Endpoint ProtocolEndpoint_ProtocolIFType = 194
	// Optical_Channel enum
	ProtocolEndpoint_ProtocolIFType_Optical_Channel ProtocolEndpoint_ProtocolIFType = 195
	// Optical_Transport enum
	ProtocolEndpoint_ProtocolIFType_Optical_Transport ProtocolEndpoint_ProtocolIFType = 196
	// Proprietary_ATM enum
	ProtocolEndpoint_ProtocolIFType_Proprietary_ATM ProtocolEndpoint_ProtocolIFType = 197
	// Voice_over_Cable enum
	ProtocolEndpoint_ProtocolIFType_Voice_over_Cable ProtocolEndpoint_ProtocolIFType = 198
	// Infiniband enum
	ProtocolEndpoint_ProtocolIFType_Infiniband ProtocolEndpoint_ProtocolIFType = 199
	// TE_Link enum
	ProtocolEndpoint_ProtocolIFType_TE_Link ProtocolEndpoint_ProtocolIFType = 200
	// Q_2931 enum
	ProtocolEndpoint_ProtocolIFType_Q_2931 ProtocolEndpoint_ProtocolIFType = 201
	// Virtual_Trunk_Group enum
	ProtocolEndpoint_ProtocolIFType_Virtual_Trunk_Group ProtocolEndpoint_ProtocolIFType = 202
	// SIP_Trunk_Group enum
	ProtocolEndpoint_ProtocolIFType_SIP_Trunk_Group ProtocolEndpoint_ProtocolIFType = 203
	// SIP_Signaling enum
	ProtocolEndpoint_ProtocolIFType_SIP_Signaling ProtocolEndpoint_ProtocolIFType = 204
	// CATV_Upstream_Channel enum
	ProtocolEndpoint_ProtocolIFType_CATV_Upstream_Channel ProtocolEndpoint_ProtocolIFType = 205
	// Econet enum
	ProtocolEndpoint_ProtocolIFType_Econet ProtocolEndpoint_ProtocolIFType = 206
	// FSAN_155Mb_PON enum
	ProtocolEndpoint_ProtocolIFType_FSAN_155Mb_PON ProtocolEndpoint_ProtocolIFType = 207
	// FSAN_622Mb_PON enum
	ProtocolEndpoint_ProtocolIFType_FSAN_622Mb_PON ProtocolEndpoint_ProtocolIFType = 208
	// Transparent_Bridge enum
	ProtocolEndpoint_ProtocolIFType_Transparent_Bridge ProtocolEndpoint_ProtocolIFType = 209
	// Line_Group enum
	ProtocolEndpoint_ProtocolIFType_Line_Group ProtocolEndpoint_ProtocolIFType = 210
	// Voice_E_M_Feature_Group enum
	ProtocolEndpoint_ProtocolIFType_Voice_E_M_Feature_Group ProtocolEndpoint_ProtocolIFType = 211
	// Voice_FGD_EANA enum
	ProtocolEndpoint_ProtocolIFType_Voice_FGD_EANA ProtocolEndpoint_ProtocolIFType = 212
	// Voice_DID enum
	ProtocolEndpoint_ProtocolIFType_Voice_DID ProtocolEndpoint_ProtocolIFType = 213
	// MPEG_Transport enum
	ProtocolEndpoint_ProtocolIFType_MPEG_Transport ProtocolEndpoint_ProtocolIFType = 214
	// _6To4 enum
	ProtocolEndpoint_ProtocolIFType__6To4 ProtocolEndpoint_ProtocolIFType = 215
	// GTP enum
	ProtocolEndpoint_ProtocolIFType_GTP ProtocolEndpoint_ProtocolIFType = 216
	// Paradyne_EtherLoop_1 enum
	ProtocolEndpoint_ProtocolIFType_Paradyne_EtherLoop_1 ProtocolEndpoint_ProtocolIFType = 217
	// Paradyne_EtherLoop_2 enum
	ProtocolEndpoint_ProtocolIFType_Paradyne_EtherLoop_2 ProtocolEndpoint_ProtocolIFType = 218
	// Optical_Channel_Group enum
	ProtocolEndpoint_ProtocolIFType_Optical_Channel_Group ProtocolEndpoint_ProtocolIFType = 219
	// HomePNA enum
	ProtocolEndpoint_ProtocolIFType_HomePNA ProtocolEndpoint_ProtocolIFType = 220
	// GFP enum
	ProtocolEndpoint_ProtocolIFType_GFP ProtocolEndpoint_ProtocolIFType = 221
	// ciscoISLvlan enum
	ProtocolEndpoint_ProtocolIFType_ciscoISLvlan ProtocolEndpoint_ProtocolIFType = 222
	// actelisMetaLOOP enum
	ProtocolEndpoint_ProtocolIFType_actelisMetaLOOP ProtocolEndpoint_ProtocolIFType = 223
	// Fcip enum
	ProtocolEndpoint_ProtocolIFType_Fcip ProtocolEndpoint_ProtocolIFType = 224
	// IANA_Reserved enum
	ProtocolEndpoint_ProtocolIFType_IANA_Reserved ProtocolEndpoint_ProtocolIFType = 225
	// IPv4 enum
	ProtocolEndpoint_ProtocolIFType_IPv4 ProtocolEndpoint_ProtocolIFType = 4096
	// IPv6 enum
	ProtocolEndpoint_ProtocolIFType_IPv6 ProtocolEndpoint_ProtocolIFType = 4097
	// IPv4_v6 enum
	ProtocolEndpoint_ProtocolIFType_IPv4_v6 ProtocolEndpoint_ProtocolIFType = 4098
	// IPX enum
	ProtocolEndpoint_ProtocolIFType_IPX ProtocolEndpoint_ProtocolIFType = 4099
	// DECnet enum
	ProtocolEndpoint_ProtocolIFType_DECnet ProtocolEndpoint_ProtocolIFType = 4100
	// SNA enum
	ProtocolEndpoint_ProtocolIFType_SNA ProtocolEndpoint_ProtocolIFType = 4101
	// CONP enum
	ProtocolEndpoint_ProtocolIFType_CONP ProtocolEndpoint_ProtocolIFType = 4102
	// CLNP enum
	ProtocolEndpoint_ProtocolIFType_CLNP ProtocolEndpoint_ProtocolIFType = 4103
	// VINES enum
	ProtocolEndpoint_ProtocolIFType_VINES ProtocolEndpoint_ProtocolIFType = 4104
	// XNS enum
	ProtocolEndpoint_ProtocolIFType_XNS ProtocolEndpoint_ProtocolIFType = 4105
	// ISDN_B_Channel_Endpoint enum
	ProtocolEndpoint_ProtocolIFType_ISDN_B_Channel_Endpoint ProtocolEndpoint_ProtocolIFType = 4106
	// ISDN_D_Channel_Endpoint enum
	ProtocolEndpoint_ProtocolIFType_ISDN_D_Channel_Endpoint ProtocolEndpoint_ProtocolIFType = 4107
	// BGP enum
	ProtocolEndpoint_ProtocolIFType_BGP ProtocolEndpoint_ProtocolIFType = 4108
	// OSPF enum
	ProtocolEndpoint_ProtocolIFType_OSPF ProtocolEndpoint_ProtocolIFType = 4109
	// UDP enum
	ProtocolEndpoint_ProtocolIFType_UDP ProtocolEndpoint_ProtocolIFType = 4110
	// TCP enum
	ProtocolEndpoint_ProtocolIFType_TCP ProtocolEndpoint_ProtocolIFType = 4111
	// _802_11a enum
	ProtocolEndpoint_ProtocolIFType__802_11a ProtocolEndpoint_ProtocolIFType = 4112
	// _802_11b enum
	ProtocolEndpoint_ProtocolIFType__802_11b ProtocolEndpoint_ProtocolIFType = 4113
	// _802_11g enum
	ProtocolEndpoint_ProtocolIFType__802_11g ProtocolEndpoint_ProtocolIFType = 4114
	// _802_11h enum
	ProtocolEndpoint_ProtocolIFType__802_11h ProtocolEndpoint_ProtocolIFType = 4115
	// NFS enum
	ProtocolEndpoint_ProtocolIFType_NFS ProtocolEndpoint_ProtocolIFType = 4200
	// CIFS enum
	ProtocolEndpoint_ProtocolIFType_CIFS ProtocolEndpoint_ProtocolIFType = 4201
	// DAFS enum
	ProtocolEndpoint_ProtocolIFType_DAFS ProtocolEndpoint_ProtocolIFType = 4202
	// WebDAV enum
	ProtocolEndpoint_ProtocolIFType_WebDAV ProtocolEndpoint_ProtocolIFType = 4203
	// HTTP enum
	ProtocolEndpoint_ProtocolIFType_HTTP ProtocolEndpoint_ProtocolIFType = 4204
	// FTP enum
	ProtocolEndpoint_ProtocolIFType_FTP ProtocolEndpoint_ProtocolIFType = 4205
	// NDMP enum
	ProtocolEndpoint_ProtocolIFType_NDMP ProtocolEndpoint_ProtocolIFType = 4300
	// Telnet enum
	ProtocolEndpoint_ProtocolIFType_Telnet ProtocolEndpoint_ProtocolIFType = 4400
	// SSH enum
	ProtocolEndpoint_ProtocolIFType_SSH ProtocolEndpoint_ProtocolIFType = 4401
	// SM_CLP enum
	ProtocolEndpoint_ProtocolIFType_SM_CLP ProtocolEndpoint_ProtocolIFType = 4402
	// SMTP enum
	ProtocolEndpoint_ProtocolIFType_SMTP ProtocolEndpoint_ProtocolIFType = 4403
	// LDAP enum
	ProtocolEndpoint_ProtocolIFType_LDAP ProtocolEndpoint_ProtocolIFType = 4404
	// RDP enum
	ProtocolEndpoint_ProtocolIFType_RDP ProtocolEndpoint_ProtocolIFType = 4405
	// HTTPS enum
	ProtocolEndpoint_ProtocolIFType_HTTPS ProtocolEndpoint_ProtocolIFType = 4406
	// DMTF_Reserved enum
	ProtocolEndpoint_ProtocolIFType_DMTF_Reserved ProtocolEndpoint_ProtocolIFType = 4407
	// Vendor_Reserved enum
	ProtocolEndpoint_ProtocolIFType_Vendor_Reserved ProtocolEndpoint_ProtocolIFType = 4408
)

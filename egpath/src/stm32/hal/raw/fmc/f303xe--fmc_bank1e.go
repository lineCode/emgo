// +build f303xe

// Peripheral: FMC_Bank1E_Periph  Flexible Memory Controller Bank1E.
// Instances:
//  FMC_Bank1E  mmap.FMC_Bank1E_R_BASE
// Registers:
//  0x00 32  BWTR[7] NOR/PSRAM write timing registers.
// Import:
//  stm32/o/f303xe/mmap
package fmc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	EADDSET   BWTR_Bits = 0x0F << 0  //+ ADDSET[3:0] bits (Address setup phase duration).
	EADDSET_0 BWTR_Bits = 0x01 << 0  //  Bit 0.
	EADDSET_1 BWTR_Bits = 0x02 << 0  //  Bit 1.
	EADDSET_2 BWTR_Bits = 0x04 << 0  //  Bit 2.
	EADDSET_3 BWTR_Bits = 0x08 << 0  //  Bit 3.
	EADDHLD   BWTR_Bits = 0x0F << 4  //+ ADDHLD[3:0] bits (Address-hold phase duration).
	EADDHLD_0 BWTR_Bits = 0x01 << 4  //  Bit 0.
	EADDHLD_1 BWTR_Bits = 0x02 << 4  //  Bit 1.
	EADDHLD_2 BWTR_Bits = 0x04 << 4  //  Bit 2.
	EADDHLD_3 BWTR_Bits = 0x08 << 4  //  Bit 3.
	EDATAST   BWTR_Bits = 0xFF << 8  //+ DATAST [3:0] bits (Data-phase duration).
	EDATAST_0 BWTR_Bits = 0x01 << 8  //  Bit 0.
	EDATAST_1 BWTR_Bits = 0x02 << 8  //  Bit 1.
	EDATAST_2 BWTR_Bits = 0x04 << 8  //  Bit 2.
	EDATAST_3 BWTR_Bits = 0x08 << 8  //  Bit 3.
	EDATAST_4 BWTR_Bits = 0x10 << 8  //  Bit 4.
	EDATAST_5 BWTR_Bits = 0x20 << 8  //  Bit 5.
	EDATAST_6 BWTR_Bits = 0x40 << 8  //  Bit 6.
	EDATAST_7 BWTR_Bits = 0x80 << 8  //  Bit 7.
	ECLKDIV   BWTR_Bits = 0x0F << 20 //+ CLKDIV[3:0] bits (Clock divide ratio).
	ECLKDIV_0 BWTR_Bits = 0x01 << 20 //  Bit 0.
	ECLKDIV_1 BWTR_Bits = 0x02 << 20 //  Bit 1.
	ECLKDIV_2 BWTR_Bits = 0x04 << 20 //  Bit 2.
	ECLKDIV_3 BWTR_Bits = 0x08 << 20 //  Bit 3.
	EDATLAT   BWTR_Bits = 0x0F << 24 //+ DATLA[3:0] bits (Data latency).
	EDATLAT_0 BWTR_Bits = 0x01 << 24 //  Bit 0.
	EDATLAT_1 BWTR_Bits = 0x02 << 24 //  Bit 1.
	EDATLAT_2 BWTR_Bits = 0x04 << 24 //  Bit 2.
	EDATLAT_3 BWTR_Bits = 0x08 << 24 //  Bit 3.
	EACCMOD   BWTR_Bits = 0x03 << 28 //+ ACCMOD[1:0] bits (Access mode).
	EACCMOD_0 BWTR_Bits = 0x01 << 28 //  Bit 0.
	EACCMOD_1 BWTR_Bits = 0x02 << 28 //  Bit 1.
)

const (
	EADDSETn = 0
	EADDHLDn = 4
	EDATASTn = 8
	ECLKDIVn = 20
	EDATLATn = 24
	EACCMODn = 28
)

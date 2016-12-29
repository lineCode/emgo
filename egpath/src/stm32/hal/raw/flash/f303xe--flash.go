// +build f303xe

// Peripheral: FLASH_Periph  FLASH Registers.
// Instances:
//  FLASH  mmap.FLASH_R_BASE
// Registers:
//  0x00 32  ACR     Access control register.
//  0x04 32  KEYR    Key register.
//  0x08 32  OPTKEYR Option key register.
//  0x0C 32  SR      Status register.
//  0x10 32  CR      Control register.
//  0x14 32  AR      Address register.
//  0x1C 32  OBR     Option byte register.
//  0x20 32  WRPR    Write register.
// Import:
//  stm32/o/f303xe/mmap
package flash

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	LATENCY   ACR_Bits = 0x03 << 0 //+ LATENCY[2:0] bits (Latency).
	LATENCY_0 ACR_Bits = 0x01 << 0 //  Bit 0.
	LATENCY_1 ACR_Bits = 0x02 << 0 //  Bit 1.
	HLFCYA    ACR_Bits = 0x01 << 3 //+ Flash Half Cycle Access Enable.
	PRFTBE    ACR_Bits = 0x01 << 4 //+ Prefetch Buffer Enable.
	PRFTBS    ACR_Bits = 0x01 << 5 //+
)

const (
	LATENCYn = 0
	HLFCYAn  = 3
	PRFTBEn  = 4
	PRFTBSn  = 5
)

const (
	FKEYR KEYR_Bits = 0xFFFFFFFF << 0 //+ FPEC Key.
)

const (
	FKEYRn = 0
)

const (
	BSY    SR_Bits = 0x01 << 0 //+ Busy.
	PGERR  SR_Bits = 0x01 << 2 //+ Programming Error.
	WRPERR SR_Bits = 0x01 << 4 //+ Write Protection Error.
	EOP    SR_Bits = 0x01 << 5 //+ End of operation.
)

const (
	BSYn    = 0
	PGERRn  = 2
	WRPERRn = 4
	EOPn    = 5
)

const (
	PG         CR_Bits = 0x01 << 0  //+ Programming.
	PER        CR_Bits = 0x01 << 1  //+ Page Erase.
	MER        CR_Bits = 0x01 << 2  //+ Mass Erase.
	OPTPG      CR_Bits = 0x01 << 4  //+ Option Byte Programming.
	OPTER      CR_Bits = 0x01 << 5  //+ Option Byte Erase.
	STRT       CR_Bits = 0x01 << 6  //+ Start.
	LOCK       CR_Bits = 0x01 << 7  //+ Lock.
	OPTWRE     CR_Bits = 0x01 << 9  //+ Option Bytes Write Enable.
	ERRIE      CR_Bits = 0x01 << 10 //+ Error Interrupt Enable.
	EOPIE      CR_Bits = 0x01 << 12 //+ End of operation interrupt enable.
	OBL_LAUNCH CR_Bits = 0x01 << 13 //+ OptionBytes Loader Launch.
)

const (
	PGn         = 0
	PERn        = 1
	MERn        = 2
	OPTPGn      = 4
	OPTERn      = 5
	STRTn       = 6
	LOCKn       = 7
	OPTWREn     = 9
	ERRIEn      = 10
	EOPIEn      = 12
	OBL_LAUNCHn = 13
)

const (
	FAR AR_Bits = 0xFFFFFFFF << 0 //+ Flash Address.
)

const (
	FARn = 0
)

const (
	OPTERR       OBR_Bits = 0x01 << 0 //+ Option Byte Error.
	RDPRT1       OBR_Bits = 0x01 << 1 //+ Read protection Level 1.
	RDPRT2       OBR_Bits = 0x01 << 2 //+ Read protection Level 2.
	USER         OBR_Bits = 0x77 << 8 //+ User Option Bytes.
	IWDG_SW      OBR_Bits = 0x01 << 8 //  IWDG SW.
	nRST_STOP    OBR_Bits = 0x02 << 8 //  nRST_STOP.
	nRST_STDBY   OBR_Bits = 0x04 << 8 //  nRST_STDBY.
	nBOOT1       OBR_Bits = 0x10 << 8 //  nBOOT1.
	VDDA_MONITOR OBR_Bits = 0x20 << 8 //  VDDA_MONITOR.
	SRAM_PE      OBR_Bits = 0x40 << 8 //  SRAM_PE.
)

const (
	OPTERRn = 0
	RDPRT1n = 1
	RDPRT2n = 2
	USERn   = 8
)

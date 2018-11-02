## HW info

* `lshw -short` (with dock/keyboard & a Logitech mouse & micro SD)

```
H/W path       Device  Class          Description
=================================================
                       system         T101HA (ASUS-TabletSKU)
/0                     bus            T101HA
/0/0                   memory         64KiB BIOS
/0/b                   memory         2GiB System Memory
/0/b/0                 memory         2GiB DIMM DDR3 1600 MHz (0.6 ns)
/0/b/1                 memory         DIMMProject-Id-Version: lshwReport-Msgid-Bugs-To: FULL NAME <EMAIL@ADDRESS>PO-Revi
/0/12                  memory         224KiB L1 cache
/0/13                  memory         2MiB L2 cache
/0/14                  processor      Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz
/0/100                 bridge         Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series SoC Transaction Registe
/0/100/2               display        Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCI Configuration Regis
/0/100/3               multimedia     Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Imaging Unit
/0/100/a               generic        Intel Corporation
/0/100/b               generic        Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Power Management Contro
/0/100/14              bus            Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series USB xHCI Controller
/0/100/14/0    usb1    bus            xHCI Host Controller
/0/100/14/0/2          input          USB Receiver
/0/100/14/0/3          communication  Bluetooth wireless interface
/0/100/14/0/4          input          ASUS HID Device
/0/100/14/1    usb2    bus            xHCI Host Controller
/0/100/1a              generic        Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Trusted Execution Engin
/0/100/1c              bridge         Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCI Express Port #1
/0/100/1c/0    wlp1s0  network        QCA9377 802.11ac Wireless Network Adapter
/0/100/1f              bridge         Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCU
```


* `cat /proc/cpuinfo`

```
processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 76
model name	: Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz
stepping	: 4
microcode	: 0x410
cpu MHz		: 1435.764
cache size	: 1024 KB
physical id	: 0
siblings	: 4
core id		: 0
cpu cores	: 4
apicid		: 0
initial apicid	: 0
fpu		: yes
fpu_exception	: yes
cpuid level	: 11
wp		: yes
flags		: fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx rdtscp lm constant_tsc arch_perfmon pebs bts rep_good nopl xtopology tsc_reliable nonstop_tsc cpuid aperfmperf tsc_known_freq pni pclmulqdq dtes64 monitor ds_cpl vmx est tm2 ssse3 cx16 xtpr pdcm sse4_1 sse4_2 movbe popcnt tsc_deadline_timer aes rdrand lahf_lm 3dnowprefetch epb pti ibrs ibpb stibp tpr_shadow vnmi flexpriority ept vpid tsc_adjust smep erms dtherm ida arat
bugs		: cpu_meltdown spectre_v1 spectre_v2
bogomips	: 2880.00
clflush size	: 64
cache_alignment	: 64
address sizes	: 36 bits physical, 48 bits virtual
power management:

... to 4 core
```


* `lspci`

```
00:00.0 Host bridge: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series SoC Transaction Register (rev 36)
00:02.0 VGA compatible controller: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCI Configuration Registers (rev 36)
00:03.0 Multimedia controller: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Imaging Unit (rev 36)
00:0a.0 Non-VGA unclassified device: Intel Corporation Device 22d8 (rev 36)
00:0b.0 Signal processing controller: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Power Management Controller (rev 36)
00:14.0 USB controller: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series USB xHCI Controller (rev 36)
00:1a.0 Encryption controller: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series Trusted Execution Engine (rev 36)
00:1c.0 PCI bridge: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCI Express Port #1 (rev 36)
00:1f.0 ISA bridge: Intel Corporation Atom/Celeron/Pentium Processor x5-E8000/J3xxx/N3xxx Series PCU (rev 36)
01:00.0 Network controller: Qualcomm Atheros QCA9377 802.11ac Wireless Network Adapter (rev 31)
```


* `lsusb` (with dock/keyboard & a Logitech mouse)

```
Bus 002 Device 001: ID 1d6b:0003 Linux Foundation 3.0 root hub
Bus 001 Device 004: ID 0b05:183d ASUSTek Computer, Inc. 
Bus 001 Device 003: ID 13d3:3501 IMC Networks 
Bus 001 Device 002: ID 046d:c534 Logitech, Inc. Unifying Receiver
Bus 001 Device 001: ID 1d6b:0002 Linux Foundation 2.0 root hub
```


* `lsblk`

```
NAME         MAJ:MIN RM   SIZE RO TYPE MOUNTPOINT
mmcblk0      179:0    0  58.2G  0 disk 
├─mmcblk0p1  179:1    0   260M  0 part /boot/efi
└─mmcblk0p2  179:2    0    58G  0 part /
mmcblk0boot0 179:8    0     4M  1 disk 
mmcblk0boot1 179:16   0     4M  1 disk 
mmcblk1      179:24   0  14.9G  0 disk (micro SD)
└─mmcblk1p1  179:25   0  14.9G  0 part 
zram0        252:0    0 608.7M  0 disk [SWAP]
```


* `sudo dmidecode`

```
# dmidecode 3.1
Getting SMBIOS data from sysfs.
SMBIOS 3.0.0 present.
Table at 0x7715F000.

Handle 0x0000, DMI type 0, 24 bytes
BIOS Information
	Vendor: American Megatrends Inc.
	Version: T101HA.305
	Release Date: 01/24/2018
	Address: 0xF0000
	Runtime Size: 64 kB
	ROM Size: 6144 kB
	Characteristics:
		PCI is supported
		BIOS is upgradeable
		BIOS shadowing is allowed
		Boot from CD is supported
		Selectable boot is supported
		BIOS ROM is socketed
		EDD is supported
		5.25"/1.2 MB floppy services are supported (int 13h)
		3.5"/720 kB floppy services are supported (int 13h)
		3.5"/2.88 MB floppy services are supported (int 13h)
		Print screen service is supported (int 5h)
		Serial services are supported (int 14h)
		Printer services are supported (int 17h)
		ACPI is supported
		USB legacy is supported
		Smart battery is supported
		BIOS boot specification is supported
		Targeted content distribution is supported
		UEFI is supported
	BIOS Revision: 5.6

Handle 0x0001, DMI type 1, 27 bytes
System Information
	Manufacturer: ASUSTeK COMPUTER INC.
	Product Name: T101HA
	Version: 1.0       
	Serial Number: (removed)
	UUID: (removed)
	Wake-up Type: Power Switch
	SKU Number: ASUS-TabletSKU
	Family: T

Handle 0x0002, DMI type 2, 15 bytes
Base Board Information
	Manufacturer: ASUSTeK COMPUTER INC.
	Product Name: T101HA
	Version: 1.0       
	Serial Number: (removed)
	Asset Tag: (removed)
	Features:
		Board is a hosting board
		Board is replaceable
	Location In Chassis: MIDDLE              
	Chassis Handle: 0x0003
	Type: Motherboard
	Contained Object Handles: 0

Handle 0x0003, DMI type 3, 22 bytes
Chassis Information
	Manufacturer: ASUSTeK COMPUTER INC.
	Type: Detachable
	Lock: Not Present
	Version: 1.0       
	Serial Number: (removed)     
	Asset Tag: No Asset Tag    
	Boot-up State: Safe
	Power Supply State: Safe
	Thermal State: Safe
	Security Status: None
	OEM Information: 0x00000000
	Height: Unspecified
	Number Of Power Cords: 1
	Contained Elements: 0
	SKU Number: Default string

Handle 0x0008, DMI type 10, 10 bytes
On Board Device 1 Information
	Type: Video
	Status: Enabled
	Description:  VGA
On Board Device 2 Information
	Type: Ethernet
	Status: Enabled
	Description:  GLAN
On Board Device 3 Information
	Type: Ethernet
	Status: Enabled
	Description:  WLAN

Handle 0x0009, DMI type 11, 5 bytes
OEM Strings
	String 1: (removed)
	String 2: (removed)
	String 3: (removed)
	String 4: (removed)
	String 5:  
	String 6:  
	String 7:  
	String 8:  
	String 9:  
	String 10:  

Handle 0x000A, DMI type 12, 5 bytes
System Configuration Options
	Option 1: SMI:00B2CA
	Option 2: DSN:B7CCD53C                        
	Option 3: DSN:FFFFFFFFFFFF                    
	Option 4: DSN:FFFFFFFFFFFF                    

Handle 0x000B, DMI type 16, 23 bytes
Physical Memory Array
	Location: System Board Or Motherboard
	Use: System Memory
	Error Correction Type: Multi-bit ECC
	Maximum Capacity: 8 GB
	Error Information Handle: Not Provided
	Number Of Devices: 2

Handle 0x000C, DMI type 19, 31 bytes
Memory Array Mapped Address
	Starting Address: 0x00000000000
	Ending Address: 0x0007FFFFFFF
	Range Size: 2 GB
	Physical Array Handle: 0x000B
	Partition Width: 2

Handle 0x000D, DMI type 17, 40 bytes
Memory Device
	Array Handle: 0x000B
	Error Information Handle: Not Provided
	Total Width: 8 bits
	Data Width: 64 bits
	Size: 2048 MB
	Form Factor: DIMM
	Set: None
	Locator: A1_DIMM0
	Bank Locator: A1_BANK0
	Type: DDR3
	Type Detail: Unknown
	Speed: 1600 MT/s
	Manufacturer: A1_Manufacturer0
	Serial Number: A1_SerNum0
	Asset Tag: A1_AssetTagNum0
	Part Number: Array1_PartNumber0
	Rank: Unknown
	Configured Clock Speed: 1600 MT/s
	Minimum Voltage: 1.5 V
	Maximum Voltage: 1.5 V
	Configured Voltage: 1.35 V

Handle 0x000E, DMI type 20, 35 bytes
Memory Device Mapped Address
	Starting Address: 0x00000000000
	Ending Address: 0x0007FFFFFFF
	Range Size: 2 GB
	Physical Device Handle: 0x000D
	Memory Array Mapped Address Handle: 0x000C
	Partition Row Position: Unknown
	Interleave Position: 1
	Interleaved Data Depth: 2

Handle 0x000F, DMI type 17, 40 bytes
Memory Device
	Array Handle: 0x000B
	Error Information Handle: Not Provided
	Total Width: Unknown
	Data Width: 64 bits
	Size: No Module Installed
	Form Factor: DIMM
	Set: None
	Locator: A1_DIMM1
	Bank Locator: A1_BANK1
	Type: Unknown
	Type Detail: Unknown
	Speed: Unknown
	Manufacturer: A1_Manufacturer1
	Serial Number: A1_SerNum1
	Asset Tag: A1_AssetTagNum1
	Part Number: Array1_PartNumber1
	Rank: Unknown
	Configured Clock Speed: Unknown
	Minimum Voltage: Unknown
	Maximum Voltage: Unknown
	Configured Voltage: Unknown

Handle 0x0010, DMI type 126, 35 bytes
Inactive

Handle 0x0011, DMI type 32, 20 bytes
System Boot Information
	Status: No errors detected

Handle 0x0012, DMI type 7, 19 bytes
Cache Information
	Socket Designation: CPU Internal L1
	Configuration: Enabled, Not Socketed, Level 1
	Operational Mode: Write Back
	Location: Internal
	Installed Size: 224 kB
	Maximum Size: 224 kB
	Supported SRAM Types:
		Unknown
	Installed SRAM Type: Unknown
	Speed: Unknown
	Error Correction Type: Single-bit ECC
	System Type: Other
	Associativity: Other

Handle 0x0013, DMI type 7, 19 bytes
Cache Information
	Socket Designation: CPU Internal L2
	Configuration: Enabled, Not Socketed, Level 2
	Operational Mode: Write Back
	Location: Internal
	Installed Size: 2048 kB
	Maximum Size: 2048 kB
	Supported SRAM Types:
		Unknown
	Installed SRAM Type: Unknown
	Speed: Unknown
	Error Correction Type: Single-bit ECC
	System Type: Unified
	Associativity: 16-way Set-associative

Handle 0x0014, DMI type 4, 48 bytes
Processor Information
	Socket Designation: SOCKET 0
	Type: Central Processor
	Family: Atom
	Manufacturer: Intel
	ID: C4 06 04 00 FF FB EB BF
	Signature: Type 0, Family 6, Model 76, Stepping 4
	Flags:
		FPU (Floating-point unit on-chip)
		VME (Virtual mode extension)
		DE (Debugging extension)
		PSE (Page size extension)
		TSC (Time stamp counter)
		MSR (Model specific registers)
		PAE (Physical address extension)
		MCE (Machine check exception)
		CX8 (CMPXCHG8 instruction supported)
		APIC (On-chip APIC hardware supported)
		SEP (Fast system call)
		MTRR (Memory type range registers)
		PGE (Page global enable)
		MCA (Machine check architecture)
		CMOV (Conditional move instruction supported)
		PAT (Page attribute table)
		PSE-36 (36-bit page size extension)
		CLFSH (CLFLUSH instruction supported)
		DS (Debug store)
		ACPI (ACPI supported)
		MMX (MMX technology supported)
		FXSR (FXSAVE and FXSTOR instructions supported)
		SSE (Streaming SIMD extensions)
		SSE2 (Streaming SIMD extensions 2)
		SS (Self-snoop)
		HTT (Multi-threading)
		TM (Thermal monitor supported)
		PBE (Pending break enabled)
	Version: Intel(R) Atom(TM) x5-Z8350 CPU @ 1.44GHz
	Voltage: 1.2 V
	External Clock: 80 MHz
	Max Speed: 2400 MHz
	Current Speed: 1440 MHz
	Status: Populated, Enabled
	Upgrade: Socket BGA1155
	L1 Cache Handle: 0x0012
	L2 Cache Handle: 0x0013
	L3 Cache Handle: Not Provided
	Serial Number: Not Specified
	Asset Tag: Fill By OEM
	Part Number: Fill By OEM
	Core Count: 4
	Core Enabled: 4
	Thread Count: 4
	Characteristics:
		64-bit capable

Handle 0x0015, DMI type 248, 9 bytes
OEM-specific Type
	Header and Data:
		F8 09 15 00 00 02 6C 00 01
	Strings:
		BDF_low

Handle 0x0016, DMI type 13, 22 bytes
BIOS Language Information
	Language Description Format: Long
	Installable Languages: 1
		en|US|iso8859-1
	Currently Installed Language: en|US|iso8859-1

Handle 0x0017, DMI type 127, 4 bytes
End Of Table
```


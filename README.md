# Ubuntu 18.04 on ASUS Transformer Book T101HA

official spec:
* cpu: Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz
* RAM: DDR3 1600MHz 2G
* disk: eMMC 64G
* network: wifi 802.11ac, bluetooth 4.1

## Status
### Work out-of-box:
* wifi (seems only 2.4G 802.11b/g)
* bluetooth
* touchpad (2 point scroll works)
* touch screen (single point)
* micro SD card reader

### Need config:
* sound (speaker ok, headphone jack not work)
* screen & touch screen rotation
* screen brightness

### Not work:
* webcam video
* webcam audio (not test)
* mic

### Not test:
* micro HDMI
* micro HDMI audio output

### Hotkeys:
* screen brightness: no
* screen on/off: yes (screen setting popup)
* sound: yes
* touchpad on/off: no
* wifi/bt switch: no
* sleep: no

### Buttons:
* power: yes
* vol up/down: yes

### power:
* suspend to RAM: yes
* suspend to disk: no

## HW info
see `hwinfo.md`

## setup/install
see `setup.md` and `daemon.md`

## bench
* `uname -a`

```
Linux cs8425-T101HA 4.15.0-38-generic #41-Ubuntu SMP Wed Oct 10 10:59:38 UTC 2018 x86_64 x86_64 x86_64 GNU/Linux
```


* `7z b`

```
7-Zip [64] 16.02 : Copyright (c) 1999-2016 Igor Pavlov : 2016-05-21
p7zip Version 16.02 (locale=zh_TW.UTF-8,Utf16=on,HugeFiles=on,64 bits,4 CPUs Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz (406C4),ASM,AES-NI)

      Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz (406C4)
CPU Freq:  1876  1891  1912  1912  1913  1913  1900  1913  1912

RAM size:    1825 MB,  # CPU hardware threads:   4
RAM usage:    882 MB,  # Benchmark threads:      4

                       Compressing  |                  Decompressing
Dict     Speed Usage    R/U Rating  |      Speed Usage    R/U Rating
         KiB/s     %   MIPS   MIPS  |      KiB/s     %   MIPS   MIPS

22:       3717   335   1080   3617  |      67603   391   1473   5768
23:       3685   346   1084   3755  |      61303   365   1455   5304
24:       3586   353   1091   3856  |      65106   392   1457   5715
25:       3557   360   1128   4062  |      63024   393   1428   5609
----------------------------------  | ------------------------------
Avr:             349   1096   3822  |              385   1453   5599
Tot:             367   1274   4711
```


# Ubuntu 18.04 on ASUS Transformer Book T101HA

spec:
* cpu: Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz
* RAM: DDR3 1600MHz 2G
* disk: eMMC 64G
* network: wifi 802.11ac, bluetooth 4.1

## Status
### Work out-of-box:
* wifi (only 2.4G 802.11b/g)
* bluetooth
* touchpad (2 point scroll works)
* touch screen (single point)

### Need config:
sound (speaker ok, headphone jack not work)
screen & touch screen rotation
screen brightness

### Not work:
* webcam video
* webcam audio (not test)
* mic


### Hotkeys:
* screen brightness: no
* screen on/off: yes (screen setting popup)
* sound: yes
* touchpad on/off: no
* wifi/bt switch: no
* sleep: no

### Buttons:
* power: no
* vol up/down: yes

### power:
* suspend to RAM: yes
* suspend to disk: no

## Sound
edit `/etc/modprobe.d/blacklist.conf`
and add the following line :
```
blacklist snd_hdmi_lpe_audio
```
do reboot


## Brightness
edit `/etc/initramfs-tools/modules`
and add the following lines :
```
pwm-lpss
pwm-lpss-platform
```

create/edit `/etc/initramfs-tools/conf.d/noresume.conf`
and add the following lines if only use zram :
```
# Disable resume (this system has no swap/only zram)
RESUME=none
```

run as root: `update-initramfs -u`
do reboot

## Tweaks
### zram

1. install zram-config: `apt install zram-config`
2. edit `/usr/bin/init-zram-swapping`

```
# load dependency modules
#NRDEVICES=$(grep -c ^processor /proc/cpuinfo | sed 's/^0$/1/')
NRDEVICES=1


# Calculate memory to use for zram (1/2 of ram)
totalmem=`LC_ALL=C free | grep -e "^Mem:" | sed -e 's/^Mem: *//' -e 's/  *.*//'`
#mem=$(((totalmem / 2 / ${NRDEVICES}) * 1024))
mem=$(((totalmem / 3 / ${NRDEVICES}) * 1024))
```



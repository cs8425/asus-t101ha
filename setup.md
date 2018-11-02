## Screen rotate
for X user:
```
xinput set-prop 'SIS0457:00 0457:11ED' 'Coordinate Transformation Matrix' 0 1 0 -1 0 1 0 0 1 # touch screen
xrandr --output 'DSI-1' --rotate right # screen
```
do not need reboot

---
for kernel console:
edit `/etc/default/grub`,
add `fbcon=rotate:1` to line:

```
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash"
```

like

```
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash fbcon=rotate:1"
```

run as root: `update-grub`
need reboot

## Sound
edit `/etc/modprobe.d/blacklist.conf`
and add the following line :
```
blacklist snd_hdmi_lpe_audio
```
need reboot

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
need reboot

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

### screen rotate scripts
see `daemon.md`


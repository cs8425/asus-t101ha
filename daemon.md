# daemon for auto dock & screen rotate

## install
1. put `80-dock.rules` to `/etc/udev/rules.d/80-dock.rules`, owner root.root, mode 644
2. put `dock.sh` to `/usr/bin/dock.sh`, owner root.root, mode 755
3. run `udevadm control --reload-rules && udevadm trigger` as root
4. start `t101srv` as user after login


## scripts

* `rotlock.sh` : for rotate lock
* `dock.sh` : set dock or not

## TODO:

* [ ] add rotate setting script/method
* [ ] less memory & cpu usage

## arch/method

* file + libuv fsevent daemon
* file + golang file polling daemon

* unix socket + golang daemon (current)
* unix socket + libuv daemon

* FUSE daemon

### unix socket/files

* `/dev/shm/T101HA.lock` : for daemon lock
* `/dev/shm/T101HA-lock` : for rotate lock, `1` = lock, `0` = unlock, `T` = toggle
* `/dev/shm/T101HA-dock` : for dock, `1` = docked, `0` = undock




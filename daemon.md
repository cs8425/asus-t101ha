# daemon for auto dock & screen rotate

## build

1. run `go build -ldflags="-s -w" t101srv.go`
2. put `t101srv` to anyway you want

## install
1. put `80-dock.rules` to `/etc/udev/rules.d/80-dock.rules`, owner root.root, mode 644
2. put `dock.sh` to `/usr/bin/dock.sh`, owner root.root, mode 755
3. run `udevadm control --reload-rules && udevadm trigger` as root
4. start `t101srv` as user after login


## scripts

* `rotlock.sh` : for rotate lock
* `dock.sh` : set dock or not

## TODO:

* [ ] enable/disable touch screen when docked
* [ ] add rotate setting script/method
* [ ] less memory & cpu usage

## arch/method

* file + libuv fsevent daemon
  * [+] can use `cat` and `echo >` to read/write state
  * [+] wake up on write
  * [-] may need to re-write data (eg: toggle)
  * [-] need c dependency
* file + golang file polling daemon
  * [+] easy to code & port
  * [-] more power usage
  * [-] may need to re-write date (eg: toggle)

* unix socket + golang daemon (current)
  * [-] need `nc -U` to send status
  * [-] no state feedback
  * [-] a little more memory usage (4176 kByte)
  * [+] wake up on write
  * [+] easy to build without dependency
* unix socket + libuv daemon
  * [=] same as golang version
  * [+] maybe less memory usage
  * [-] need c dependency

* FUSE daemon
  * [+] can use `cat` and `echo >` to read/write state
  * [+] wake up on read/write
  * [-] more complex

### unix socket/files

* `/dev/shm/T101HA.lock` : for daemon lock
* `/dev/shm/T101HA-lock` : for rotate lock, `1` = lock, `0` = unlock, `T` = toggle
* `/dev/shm/T101HA-dock` : for dock, `1` = docked, `0` = undock




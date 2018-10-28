#!/bin/bash

SOCK="/dev/shm/T101HA-dock"

case "${ACTION:-$1}" in
	1 | add)
		echo 1 | nc -q 0 -U $SOCK
		;;
	0 | remove)
		echo 0 | nc -q 0 -U $SOCK
		;;

	*)
		echo "Usage: $0 {add|1|remove|0}"
		;;
esac


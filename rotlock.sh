#!/bin/bash

SOCK="/dev/shm/T101HA-lock"

case "$1" in
	1 | lock)
		echo 1 | nc -q 0 -U $SOCK
		;;
	0 | unlock)
		echo 0 | nc -q 0 -U $SOCK
		;;

	T)
		echo T | nc -q 0 -U $SOCK
		;;

	*)
		echo "Usage: $0 {lock|1|unlock|0|T}"
		;;
esac


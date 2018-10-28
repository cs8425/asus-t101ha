package main

import (
	"flag"
	"fmt"
//	"log"
	"runtime"
//	"sync"
	"time"
	"os/exec"

//	"bufio"
//	"os"
	"math"
	"net"
	"path/filepath"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	ACC_X_RAW = "/sys/bus/iio/devices/iio:device0/in_accel_x_raw"
	ACC_SCALE = "/sys/bus/iio/devices/iio:device0/in_accel_scale"

	ACC_DEV_PATH = "/sys/bus/iio/devices/"
	ACC_NAME = "accel_3d"
	ACC_G float64 = 0.7

	LOCK_PATH = "/dev/shm/T101HA-lock"
	DOCK_PATH = "/dev/shm/T101HA-dock"

	SCREEN = "DSI-1"
	TOUCH = "SIS0457:00 0457:11ED"
)

const (
        INIT = iota
        DOCKED
        LOCK
        UNLOCK
)

var ROT = flag.Int("r", 1, "rot")
var T = flag.Int("t", 1, "touch")

var acc *Acc

func main() {
	flag.Parse()

	// runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	fmt.Println("init...")

	acc = NewAcc(ACC_DEV_PATH, ACC_NAME)
	fmt.Println("ACC_DEV = ", acc)

	x, y, z, ok := acc.GetAcc()
	fmt.Println("ACC = ", x, y, z, ok, acc.readRot())

	//doRot(acc.readRot(), true)
	/*if *T == 1 {
		doRot(*ROT, true)
	} else {
		doRot(*ROT, false)
	}*/
	//return

	dockedCh := make(chan bool, 1) // from udev
	rotLockCh := make(chan bool, 1) // from user shortcut
	rotCh := make(chan int, 1) // from sysfs
	exitCh := make(chan bool, 1) // for acc worker exit

	go pollDock(dockedCh)
	go pollLock(rotLockCh)

	// TODO: check and start
	docked := true
	rotLock := false
	rot := 1 // 0 = normal, 1 = right, 2 = inverse, 3 = left

	state := DOCKED
	state0 := DOCKED
	for {
		select {
		case docked = <-dockedCh:
			if docked {
				// docked >> set to right
				state = DOCKED
			} else {
				state = UNLOCK
			}

		case rotLock = <-rotLockCh:
			if docked { // TODO: touch on/off
				continue
			}
			if rotLock {
				state = LOCK
			} else {
				state = UNLOCK
			}

		case rot = <-rotCh:
			fmt.Println("[rot]", rot)
		}

		switch state {
		case DOCKED: // acc off, touch on&off
			if state0 == UNLOCK {
			// stop poll acc
				exitCh <- true
			}
			if state0 == LOCK {
			// force rotate
				doRot(1, true)
			}

		case LOCK: // acc off, touch on
			if state0 == UNLOCK {
			// stop poll acc
				exitCh <- true
			}

		case UNLOCK: // acc on, touch on
			if state0 != UNLOCK {
			// start poll acc
				go pollAccWorker(rotCh, exitCh)
			}
			// do rot
			doRot(rot, true)
		}

		state0 = state
	}
}

func poller(sock string, handler func (c net.Conn)) {
	ln, err := net.Listen("unix", sock)
	if err != nil {
		fmt.Println("[dock]Listen error:", err)
		return
	}

	for {
		fd, err := ln.Accept()
		if err != nil {
			fmt.Println("[dock]accept error:", err)
			return
		}

		handler(fd)
	}
}

func pollDock(dockedCh chan bool) {
	dock0 := true
	buf := make([]byte, 2, 2)
	handler := func (c net.Conn) {
		defer c.Close()
		_, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0]
		fmt.Println("[dock]got:", string(data))

		var dock bool
		switch data {
		case '1':
			dock = true

		case '0':
			dock = false
		}

		// send only change
		if dock != dock0 {
			dockedCh <- dock
			dock0 = dock
		}
	}

	poller(DOCK_PATH, handler)
}

func pollLock(rotLockCh chan bool) {
	lock0 := false
	buf := make([]byte, 2, 2)
	handler := func (c net.Conn) {
		defer c.Close()
		_, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0]
		fmt.Println("[lock]got:", string(data))

		var lock bool
		switch data {
		case 'T':
			lock0 = !lock0
			rotLockCh <- lock0
			return

		case '1':
			lock = true

		case '0':
			lock = false
		}

		// send only change
		if lock != lock0 {
			rotLockCh <- lock
			lock0 = lock
		}
	}

	poller(LOCK_PATH, handler)
}

func pollAccWorker(rotCh chan int, exitCh chan bool) {
	rot0 := acc.readRot()
	for {
		select {
		case <-exitCh:
			return
		default:
		}

		rot := acc.readRot()
		if rot != rot0 && rot != -1 {
			rotCh <- rot
			rot0 = rot
		}

		time.Sleep(1500 * time.Millisecond)
	}
}

func doRot(rot int, touch bool) {
	rot2mat := map[int]string{
		0: "1 0 0 0 1 0 0 0 1",
		1: "0 1 0 -1 0 1 0 0 1",
		2: "-1 0 1 0 -1 1 0 0 1",
		3: "0 -1 1 1 0 0 0 0 1",
	}
	rot2str := map[int]string{
		0: "normal",
		1: "right",
		2: "inverted",
		3: "left",
	}

	// xrandr --output 'DSI-1' --rotate right
	cmd := exec.Command("xrandr", "--output", SCREEN, "--rotate", rot2str[rot])
	err := cmd.Run()
	fmt.Println("[xrandr]", rot2str[rot], err)

	// xinput set-prop 'SIS0457:00 0457:11ED' 'Coordinate Transformation Matrix' 0 1 0 -1 0 1 0 0 1
	args := []string{"set-prop", TOUCH, "Coordinate Transformation Matrix"}
	mat := strings.Split(rot2mat[rot], " ")
	args = append(args, mat...)
	cmd = exec.Command("xinput", args...)
	err = cmd.Run()
	fmt.Println("[xinput]", args, err)


	setTouch(touch)
}

func setTouch(touch bool) {
	t2str := map[bool]string{
		false: "0",
		true: "1",
	}
	// xinput set-prop 'SIS0457:00 0457:11ED' 'Device Enabled' 1
	cmd := exec.Command("xinput", "set-prop", TOUCH, "Device Enabled", t2str[touch])
	err := cmd.Run()
	fmt.Println("[xinput]setTouch()", touch, err)
}

func readFloat(path string) (float64, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return 0.0, err
	}

	n, err := strconv.ParseFloat(string(buf[:len(buf)-1]), 64)
	if err != nil {
		return 0.0, err
	}

	return n, nil
}

func readInt(path string) (int, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	n, err := strconv.ParseInt(string(buf[:len(buf)-1]), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

type Acc struct {
	base   string
	//name   string
	scale  float64
}

func checkAccName(dir string, name string) (bool) {
	buf, err := ioutil.ReadFile(filepath.Join(dir, "name"))
	if err != nil {
		return false
	}
	return string(buf) == name + "\n"
}

func NewAcc(sysfs string, name string) (*Acc) {
	files, err := ioutil.ReadDir(sysfs)
	if err != nil {
		//fmt.Println("[acc][ReadDir]err:", err)
		return nil
	}
	//fmt.Println("[acc][ReadDir]", files)

	var base string
	for _, f := range files {
		base = filepath.Join(sysfs, f.Name())
		//fmt.Println("[acc][try dev]", base)
		if checkAccName(base, "accel_3d") {
			break
		}
	}
	if base == "" {
		return nil
	}

	acc := &Acc{
		base: base,
		//name: name,
	}
	_, ok := acc.getScale()
	if !ok {
		return nil
	}
	return acc
}

func (acc *Acc) getScale() (float64, bool) {
	scale, err := readFloat(filepath.Join(acc.base, "in_accel_scale"))
	if err != nil {
		return 0.0, false
	}

	acc.scale = scale
	return scale, true
}

func (acc *Acc) GetAcc() (fx float64, fy float64, fz float64, ok bool) {
	var err error
	ok = false

	x, err := readInt(filepath.Join(acc.base, "in_accel_x_raw"))
	if err != nil {
		return
	}
	y, err := readInt(filepath.Join(acc.base, "in_accel_y_raw"))
	if err != nil {
		return
	}
	z, err := readInt(filepath.Join(acc.base, "in_accel_z_raw"))
	if err != nil {
		return
	}

	return float64(x) * acc.scale, float64(y) * acc.scale, float64(z) * acc.scale, true
}

func (acc *Acc) readRot() (int) {
	x, y, _, ok := acc.GetAcc()
	if !ok {
		return -1 // error
	}

	if math.Abs(x) > ACC_G {
		if x < 0 {
			return 1 // right
		} else {
			return 3 // left
		}
	}
	if math.Abs(y) > ACC_G {
		if y < 0 {
			return 0 // normal
		} else {
			return 2 // inverse
		}
	}

	return -1 // can not detect
}


package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"upspowershow/models"
	"upspowershow/websocketclient"

	"github.com/getlantern/systray"
)

var secondPerTick = time.Duration(5)
var upsInfo models.UPSInfo
var percent int = 0
var dischargePow int = 0
var chargePow int = 0
var dischargeRemainTime int = 0
var chargeRemainTime int = 0
var charging int = 0

func main() {
	go websocketclient.StartWebSocket()
	systray.Run(onReady, onExit)
}

func loadIcon(path string) []byte {
	iconData, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read icon: %v", err)
	}
	return iconData
}

func traySet() {
	var iconPath string
	if charging > 0 {
		iconPath = fmt.Sprintf("icons/charging/icon_%d.ico", percent)
	} else {
		iconPath = fmt.Sprintf("icons/discharging/icon_%d.ico", percent)
	}
	icon := loadIcon(iconPath)
	systray.SetIcon(icon)
	systray.SetTitle("Greengroup UPS Monitor")
	switch charging {
	case 0:
		systray.SetTooltip(
			"电池电量: " + strconv.Itoa(percent) + "%" + "\n" +
				"当前放电中" + "\n" +
				"放电功率: " + strconv.Itoa(dischargePow) + "W" + "\n" +
				"剩余放电时间: " + strconv.Itoa(dischargeRemainTime) + "min")
	case 1:
		systray.SetTooltip(
			"电池电量: " + strconv.Itoa(percent) + "%" + "\n" +
				"当前充电中" + "\n" +
				"充电功率: " + strconv.Itoa(chargePow) + "W" + "\n" +
				"剩余充电时间: " + strconv.Itoa(chargeRemainTime) + "min")
	case 2:
		systray.SetTooltip(
			"电池电量: " + strconv.Itoa(percent) + "%" + "\n" +
				"当前已充满或未使用")
	}
}

func tickEvent() {
	upsInfo = <-websocketclient.GetUPSInfo()
	log.Printf("upsInfo: %v", upsInfo)
	percent = upsInfo.Percent
	chargePow = upsInfo.ChargePow
	if chargePow > 0 {
		chargeRemainTime = upsInfo.ChargeRemainTime
		charging = 1
	} else {
		chargeRemainTime = 0
	}
	dischargePow = upsInfo.DischargePow
	if dischargePow > 0 {
		dischargeRemainTime = upsInfo.DischargeRemainTime
		charging = 0
	} else {
		dischargeRemainTime = 0
	}
	if chargePow == 0 && dischargePow == 0 {
		charging = 2
	}

	traySet()
}

func onReady() {
	mUpdate := systray.AddMenuItem("Update", "Mannually update the battery percentage")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	tickEvent()
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				websocketclient.Disconnect()
				return
			case <-mUpdate.ClickedCh:
				tickEvent()
			}
		}
	}()

	go func() {
		for {
			tickEvent()
			time.Sleep(secondPerTick * time.Second)
		}
	}()
}

func onExit() {
	websocketclient.Disconnect()
}

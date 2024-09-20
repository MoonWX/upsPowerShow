package models

type UPSInfo struct {
	Percent             int `json:"battery_percentage"`
	ChargePow           int `json:"charge_power_all"`
	ChargeRemainTime    int `json:"charge_remain_time"`
	DischargePow        int `json:"discharge_pow"`
	DischargeRemainTime int `json:"discharge_remain_time"`
}

package switchs

import (
	"errors"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"
)

type (
	ListReq struct {
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}

	ListResp struct {
		Success bool     `json:"success"`
		Error   string   `json:"error"`
		Switchs []Switch `json:"switchs"`
		Total   int      `json:"total"`
	}
	Switch struct {
		Id                    uint   `json:"id"`
		Name                  string `json:"name"`
		Device                string `json:"device"`
		DatacenterId          uint   `json:"datacenter_id"`
		CabinetId             uint   `json:"cabinet_id"`
		Type                  string `json:"type"`
		NetworkLayer          int    `json:"network_layer"`
		SwitchModelId         any    `json:"switch_model_id"`
		Lib                   string `json:"lib"`
		Model                 string `json:"model"`
		SnmpVersion           string `json:"snmp_version"`
		SnmpStatus            int    `json:"snmp_status"`
		Remark                string `json:"remark"`
		Capacity              int    `json:"capacity"`
		IsLock                int    `json:"is_lock"`
		RoutingSupport        int    `json:"routing_support"`
		SnmpPort              int    `json:"snmp_port"`
		PrivilegedPassword    any    `json:"privileged_password"`
		TotalPortsCount       int    `json:"total_ports_count"`
		IdlePortsCount        int    `json:"idle_ports_count"`
		MalfunctionPortsCount int    `json:"malfunction_ports_count"`
		UsingPortsCount       int    `json:"using_ports_count"`
		SwitchModelName       string `json:"switch_model_name"`
		ChassisMac            []any  `json:"chassisMac"`
		IpAddress             struct {
			Id             uint   `json:"id"`
			BlockId        uint   `json:"block_id"`
			Address        string `json:"address"`
			AddressString  string `json:"address_string"`
			Ipv4BlockType  string `json:"ipv4_block_type"`
			Ipv4BlockLabel string `json:"ipv4_block_label"`
		} `json:"ip_address"`
		Datacenter struct {
			Id                 int    `json:"id"`
			Name               string `json:"name"`
			Voltage            int    `json:"voltage"`
			SwitchPorts        []any  `json:"switch_ports"`
			ArpOs              int    `json:"arp_os"`
			ReinstallLimit     int    `json:"reinstall_limit"`
			HardwareCheckLimit int    `json:"hardware_check_limit"`
			SerialNumberRuleId any    `json:"serial_number_rule_id"`
			CreatedAt          int    `json:"created_at"`
			UpdatedAt          int    `json:"updated_at"`
		} `json:"datacenter"`
		Cabinet struct {
			Id     int    `json:"id"`
			Name   string `json:"name"`
			IsLock int    `json:"is_lock"`
		} `json:"cabinet"`
		MainIpv4Address struct {
			Id             int    `json:"id"`
			BlockId        int    `json:"block_id"`
			Address        string `json:"address"`
			AddressString  string `json:"address_string"`
			Ipv4BlockType  string `json:"ipv4_block_type"`
			Ipv4BlockLabel string `json:"ipv4_block_label"`
		} `json:"main_ipv4address"`
	}
)

func (r *ListReq) Url() (url string)           { return "/api/admin/switchs" }
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取交换机列表
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5750682
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request)
}

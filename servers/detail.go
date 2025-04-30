package servers

import (
	"errors"
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	DetailReq  struct{ Id uint }
	DetailResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Server  struct {
			Id                 uint        `json:"id"`
			SystemSerialNumber interface{} `json:"system_serial_number"`
			Serialno           interface{} `json:"serialno"`
			Name               string      `json:"name"`
			Hostname           string      `json:"hostname"`
			Type               string      `json:"type"`
			PowerStatus        string      `json:"power_status"`
			Hardware           struct {
				Id      int    `json:"id"`
				Name    string `json:"name"`
				Size    int    `json:"size"`
				NodeNum int    `json:"node_num"`
				Row     int    `json:"row"`
			} `json:"hardware"`
			Card int `json:"card"`
			// Part struct {
			// 	CPU       []interface{} `json:"CPU"`
			// 	Mem       []interface{} `json:"内存"`
			// 	Disk      []interface{} `json:"硬盘"`
			// 	PCIDevice []interface{} `json:"PCI设备"`
			// } `json:"part"`
			Ipcount         int `json:"ipcount"`
			MainIpv4Address struct {
				Address   string `json:"address"`
				Cidr      string `json:"cidr"`
				Netmask   string `json:"netmask"`
				Gateway   string `json:"gateway"`
				IsBlock   bool   `json:"is_block"`
				Vlan      string `json:"vlan"`
				GroupName string `json:"group_name"`
			} `json:"main_ipv4address"`
			MainIpv6Address []interface{} `json:"main_ipv6address"`
			MainMac         string        `json:"main_mac"`
			IpmiMetadata    struct {
				AddressId       int    `json:"address_id"`
				Address         string `json:"address"`
				Username        string `json:"username"`
				Password        string `json:"password"`
				BlockId         int    `json:"block_id"`
				PublicBlockId   string `json:"public_block_id"`
				PublicAddressId string `json:"public_address_id"`
				PublicAddress   string `json:"public_address"`
				IpmiBlockId     int    `json:"ipmi_block_id"`
				IpmiAddressId   int    `json:"ipmi_address_id"`
				IpmiAddress     string `json:"ipmi_address"`
				Cidr            string `json:"cidr"`
			} `json:"ipmi_metadata"`
			Switch []struct {
				Id             int      `json:"id"`
				Switch         string   `json:"switch"`
				CabinetId      int      `json:"cabinet_id"`
				IsLock         int      `json:"is_lock"`
				PortName       string   `json:"port_name"`
				PortId         int      `json:"port_id"`
				MacAddress     string   `json:"mac_address"`
				Disconnected   int      `json:"disconnected"`
				Disabled       bool     `json:"disabled"`
				Type           string   `json:"type"`
				UpPort         int      `json:"up_port"`
				DownPort       int      `json:"down_port"`
				VlanType       string   `json:"vlan_type"`
				Vlans          []string `json:"vlans"`
				BindMac        int      `json:"bind_mac"`
				BindArp        int      `json:"bind_arp"`
				BindAcl        int      `json:"bind_acl"`
				PortRemark     string   `json:"port_remark"`
				HourInFlow     int      `json:"hour_in_flow"`
				HourOutFlow    int      `json:"hour_out_flow"`
				HourTotalFlow  int      `json:"hour_total_flow"`
				DayInFlow      int      `json:"day_in_flow"`
				DayOutFlow     int      `json:"day_out_flow"`
				DayTotalFlow   int      `json:"day_total_flow"`
				MonthInFlow    int      `json:"month_in_flow"`
				MonthOutFlow   int      `json:"month_out_flow"`
				MonthTotalFlow int      `json:"month_total_flow"`
				CardName       string   `json:"card_name"`
			} `json:"switch"`
			Cabinet struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				IsLock int    `json:"is_lock"`
			} `json:"cabinet"`
			ShelfTime  string `json:"shelf_time"`
			Datacenter struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"datacenter"`
			Username             string        `json:"username"`
			Password             string        `json:"password"`
			Remark               string        `json:"remark"`
			Image                string        `json:"image"`
			User                 []interface{} `json:"user"`
			DiskSize             int           `json:"disk_size"`
			Ipv4Blocks           []interface{} `json:"ipv4_blocks"`
			IsFree               bool          `json:"is_free"`
			Package              []interface{} `json:"package"`
			TaskId               bool          `json:"task_id"`
			RemotePort           int           `json:"remote_port"`
			Issues               []interface{} `json:"issues"`
			Intranets            []interface{} `json:"intranets"`
			IsPreinstall         int           `json:"is_preinstall"`
			IsLock               bool          `json:"is_lock"`
			DhcpEnable           int           `json:"dhcp_enable"`
			IsTraffic            bool          `json:"is_traffic"`
			NetworkCard          []interface{} `json:"network_card"`
			BootMode             string        `json:"boot_mode"`
			QrcodeUrl            string        `json:"qrcode_url"`
			DecommissionSchedule interface{}   `json:"decommission_schedule"`
			KvmType              string        `json:"kvm_type"`
		} `json:"server"`
	}
)

func (r *DetailReq) Url() (url string)           { return fmt.Sprintf("/api/admin/servers/%d", r.Id) }
func (r *DetailReq) Method() (method string)     { return http.MethodGet }
func (r *DetailReq) Values() (values url.Values) { return }
func (r *DetailReq) Body() (body any)            { return }

func (r DetailResp) Ok() (ok bool)    { return r.Success }
func (r DetailResp) Err() (err error) { return errors.New(r.Error) }

// Detail 获取服务器详情
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5768957
func Detail(ctx *dcimsdk.Context, request *DetailReq) (resp DetailResp, err error) {
	return dcimsdk.Execute[*DetailReq, DetailResp](ctx, request, dcimsdk.PowerStatusFixed)
}

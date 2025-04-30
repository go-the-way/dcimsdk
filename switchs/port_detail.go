package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	PortDetailReq struct {
		PortId int `url:"port_id" json:"-"`
	}

	PortDetailResp struct {
		Success  bool   `json:"success"`
		Error    string `json:"error"`
		PortInfo struct {
			Id             uint     `json:"id"`           // 端口id
			SwitchId       uint     `json:"switch_id"`    // 交换机id
			PortIndex      string   `json:"port_index"`   // 端口序号
			PortName       string   `json:"port_name"`    // 端口名
			PortType       string   `json:"port_type"`    // 端口类型
			PortRemark     string   `json:"port_remark"`  // 端口备注
			Type           string   `json:"type"`         // 类型
			Relid          uint     `json:"relid"`        // 关联Id
			Disconnected   uint     `json:"disconnected"` // 断开连接
			Disabled       bool     `json:"disabled"`     // 禁用
			VlanType       string   `json:"vlan_type"`    // vlan类型
			Speed          any      `json:"speed"`        // 速度
			Mtu            any      `json:"mtu"`
			BindMac        int      `json:"bind_mac"` // 绑定mac
			BindArp        int      `json:"bind_arp"` // 绑定arp
			BindAcl        int      `json:"bind_acl"`
			UpPort         int      `json:"up_port"`          // 上行端口
			DownPort       int      `json:"down_port"`        // 下载端口
			MacAddress     string   `json:"mac_address"`      // mac地址
			PhysMacAddress any      `json:"phys_mac_address"` // 物理mac地址
			MacReachTime   any      `json:"mac_reach_time"`   // mac延伸时间
			Remarks        string   `json:"remarks"`          // 备注
			FilePath       string   `json:"file_path"`        // 文件路径
			CreatedAt      any      `json:"created_at"`
			UpdatedAt      any      `json:"updated_at"`
			Vlans          []string `json:"vlans"` // vlan
		} `json:"portinfo"` // 端口信息
	}
)

func (r *PortDetailReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/ports/info/%d", r.PortId)
}
func (r *PortDetailReq) Method() (method string)     { return http.MethodGet }
func (r *PortDetailReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *PortDetailReq) Body() (body any)            { return }

// PortDetail 获取端口详情
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5752250
func PortDetail(ctx *dcimsdk.Context, request *PortDetailReq) (resp PortDetailResp, err error) {
	return dcimsdk.Execute[*PortDetailReq, PortDetailResp](ctx, request)
}

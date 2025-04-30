package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	DetailReq struct {
		Id uint `url:"id" json:"id"`
	}

	DetailResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Switch  struct {
			Id                    int    `json:"id"`
			Name                  string `json:"name"`            // 名称
			Device                string `json:"device"`          // 驱动
			DatacenterId          int    `json:"datacenter_id"`   // 机房id
			CabinetId             int    `json:"cabinet_id"`      // 机柜id
			Type                  string `json:"type"`            // 类型
			NetworkLayer          int    `json:"network_layer"`   // 网络层
			SwitchModelId         any    `json:"switch_model_id"` // 交换机型号id
			Lib                   string `json:"lib"`
			Model                 string `json:"model"`
			SnmpSupport           int    `json:"snmp_support"` // snmp控制
			SnmpVersion           string `json:"snmp_version"` // snmp版本
			SnmpStatus            int    `json:"snmp_status"`  // snmp状态
			Remark                string `json:"remark"`       // 备注
			IsLock                int    `json:"is_lock"`      // 锁定
			Capacity              int    `json:"capacity"`
			TotalPortsCount       int    `json:"total_ports_count"`       // 端口总数
			IdlePortsCount        int    `json:"idle_ports_count"`        // 空闲端口总数
			MalfunctionPortsCount int    `json:"malfunction_ports_count"` // 故障端口数
			IpAddress             struct {
				Id      int    `json:"id"`      // ip地址主键id
				Address string `json:"address"` // ip地址
			} `json:"ip_address"`
			Cabinet struct {
				Id   int    `json:"id"`
				Name string `json:"name"` // 机柜名称
			} `json:"cabinet"` // 机柜信息
			MainIpv4Address struct {
				Id      int    `json:"id"`
				Address string `json:"address"` // ip地址
			} `json:"main_ipv4address"` // 主ip地址

			SnmpCredential struct {
				Community string `json:"community"` // 社区
			} `json:"snmp_credential"` // snmp凭证
			ControlSupport    int    `json:"control_support"`  // 手动控制
			ControlProtocol   string `json:"control_protocol"` // 手动控制协议
			ControlCredential struct {
				Username string `json:"username"` // 用户名
				Password string `json:"password"` // 密码
				Port     string `json:"port"`     // 端口
			} `json:"control_credential"` // 手动控制凭证
			AutoControlProtocol   string `json:"auto_control_protocol"` // 自动控制协议
			AutoControlCredential struct {
				Username string `json:"username"` // 用户名
				Password string `json:"password"` // 密码
				Port     string `json:"port"`     // 端口
			} `json:"auto_control_credential"` // 自动控制凭证
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
			Block     struct {
				Id   int    `json:"id"`   // ip段主键id
				Type string `json:"type"` // 类型
			} `json:"block"` // ip段信息
			Address []struct {
				AddressString string `json:"address_string"` // ip地址
				BlockId       int    `json:"block_id"`       // ip段id
				Type          string `json:"type"`           // 类型
				Id            int    `json:"id"`             // ip地址id TODO address_id
			} `json:"address"` // ip地址信息
			Intranets  []any  `json:"intranets"`  // 内网信息
			Brand      string `json:"brand"`      // 品牌
			ModelValue string `json:"modelValue"` // 型号
			ChassisMac []any  `json:"chassisMac"` // 底盘mac
			// ip地址信息
			SwitchModel any `json:"switch_model"` // 交换机型号
		} `json:"switch"` // 交换机信息
	}
)

func (r *DetailReq) Url() (url string)           { return fmt.Sprintf("/api/admin/switchs/%d", r.Id) }
func (r *DetailReq) Method() (method string)     { return http.MethodGet }
func (r *DetailReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *DetailReq) Body() (body any)            { return }

// Detail 获取交换机详情
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5753939
func Detail(ctx *dcimsdk.Context, request *DetailReq) (resp DetailResp, err error) {
	return dcimsdk.Execute[*DetailReq, DetailResp](ctx, request)
}

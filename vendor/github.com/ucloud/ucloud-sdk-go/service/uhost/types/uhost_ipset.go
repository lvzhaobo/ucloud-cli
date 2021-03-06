package types

// UHostIPSet - DescribeUHostInstance
type UHostIPSet struct {

	// 电信: China-telecom，联通: China-unicom， 国际: Internation，BGP: Bgp，内网: Private 双线: Duplet
	Type string

	// IP资源ID (内网IP无对应的资源ID)
	IPId string

	// IP地址
	IP string

	// IP对应的带宽, 单位: Mb (内网IP不显示带宽信息)
	Bandwidth int

	// 是否默认的弹性网卡的信息。true: 是默认弹性网卡；其他值：不是。
	Default string
}

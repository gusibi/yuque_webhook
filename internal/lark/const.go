package lark

const CustomBotHost = "https://open.feishu.cn/open-apis"

var (
	maxConnsPerhost = 100
	defaultTimeout  = 2
	LarkClient      = NewHttpClient(maxConnsPerhost, defaultTimeout)
)

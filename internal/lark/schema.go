package lark

type TextContent struct {
	Text string `json:"text"`
}

// TextMessage 文本消息
type TextMessage struct {
	MessageType string       `json:"msg_type"`
	Content     *TextContent `json:"content"`
}

type PostContentItem struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
	Href string `json:"href,omitempty"`
}

type PostContentData struct {
	Title   string               `json:"title"`
	Content [][]*PostContentItem `json:"content"`
}

type ZhCnPostContentData struct {
	ZhCn PostContentData `json:"zh_cn"`
}

type PostContent struct {
	Post *ZhCnPostContentData `json:"post"`
}

// PostMessage 富文本消息
type PostMessage struct {
	MessageType string       `json:"msg_type"`
	Content     *PostContent `json:"content"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
	EnableForward  bool `json:"enable_forward"`
}

type CardElementText struct {
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

type CardElementAction struct {
	Tag  string          `json:"tag"`
	Text CardElementText `json:"text"`
	Url  string          `json:"url"`
	Type string          `json:"type"`
}

type CardElement struct {
	Tag    string            `json:"tag"`
	Text   CardElementText   `json:"text"`
	Action CardElementAction `json:"actions"`
}

type CardHeader struct {
	Title CardElementText `json:"title"`
}

type Card struct {
	Config   CardConfig     `json:"config"`
	Elements []*CardElement `json:"elements"`
	Header   CardHeader     `json:"header"`
}

// CardMessage 卡片消息
type CardMessage struct {
	MessageType string `json:"msg_type"`
	Card        Card   `json:"card"`
}

/*
发表/更新 文章
{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "title": {
      "content": "螺旋上升",
      "tag": "plain_text"
    }
  },
  "elements": [
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "[魔魔飞书发表了-----一篇文章] (https://www.feishu.cn)"
      }
    },
    {
      "tag": "div",
      "text": {
        "tag": "lark_md",
        "content": "[飞书](https://www.feishu.cn)整合即时沟通、日历、音视频会议、云文档、云盘、工作台等功能于一体，成就组织和个人，更高效、更愉悦。"
      }
    },
    {
      "tag": "hr"
    },
    {
      "tag": "note",
      "elements": [
        {
          "tag": "img",
          "img_key": "img_e344c476-1e58-4492-b40d-7dcffe9d6dfg",
          "alt": {
            "tag": "plain_text",
            "content": "hover"
          }
        },
        {
          "tag": "plain_text",
          "content": "公号hiiiapril推送"
        }
      ]
    }
  ]
}
*/

type BotResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

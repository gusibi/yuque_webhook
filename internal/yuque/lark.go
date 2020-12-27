package yuque

import (
	"context"
	"fmt"
	"github.com/gusibi/yuque_webhook/internal/lark"
	"time"
)

type MessageType string

const (
	TextMessage MessageType = "text"
	PostMessage MessageType = "post"
	CardMessage MessageType = "card"
)

type LarkWebHook struct {
	MessageType MessageType

	HookId string
	// timeout := time.Duration(2) * time.Second
	DefaultTimeout time.Duration
}

func NewLarkWebHook(hookId string, messageType MessageType) *LarkWebHook {
	return &LarkWebHook{
		MessageType:    messageType,
		HookId:         hookId,
		DefaultTimeout: 2,
	}
}

func (l *LarkWebHook) requestToTextMessage(ctx context.Context, req *WebHookRequest) (*lark.TextMessage, error) {
	message := &lark.TextMessage{
		MessageType: "text",
	}
	var content string
	webhookType := req.Data.WebhookSubjectType
	switch webhookType {
	case CommentCreate, CommentReplyCreate:
		content = fmt.Sprintf("文章 《%s》有一条来自%s的新评论", req.Data.CommentAble.Title, req.Data.User.Name)
	case CommentUpdate, CommentReplyUpdate:
		content = fmt.Sprintf("%s 修改了文章 《%s》的评论", req.Data.User.Name, req.Data.CommentAble.Title)
	case CommentDelete, CommentReplyDelete:
		content = fmt.Sprintf("%s 删除了文章 《%s》的评论", req.Data.User.Name, req.Data.CommentAble.Title)
	case Publish:
		content = fmt.Sprintf("「%s」有一篇新文章", req.Data.Book.Name)
	case Update:
		content = fmt.Sprintf("文章「%s」有更新", req.Data.Title)
	case Delete:
		content = fmt.Sprintf("文章「%s」被删除了", req.Data.Title)
	}
	message.Content = &lark.TextContent{Text: content}
	return message, nil
}

func (l *LarkWebHook) requestToPostMessage(ctx context.Context, req *WebHookRequest) (*lark.PostMessage, error) {
	message := &lark.PostMessage{
		MessageType: "post",
	}
	var title string
	var content []*lark.PostContentItem
	webhookType := req.Data.WebhookSubjectType
	switch webhookType {
	case CommentCreate, CommentReplyCreate:
		title = fmt.Sprintf("文章「%s」有一条新评论", req.Data.CommentAble.Title)
		content = []*lark.PostContentItem{
			{
				Tag:  "text",
				Text: "文章：",
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》", req.Data.CommentAble.Title),
				Href: GetArticleUrl(req.Data.User.Login, req.Data.CommentAble.Book.Slug, req.Data.CommentAble.Slug),
			},
			{
				Tag:  "text",
				Text: fmt.Sprintf("有一条来自「%s」的新评论", req.Data.User.Name),
			},
		}
	case CommentUpdate, CommentReplyUpdate:
		title = fmt.Sprintf("文章「%s」评论更新", req.Data.CommentAble.Title)
		content = []*lark.PostContentItem{
			{
				Tag:  "text",
				Text: fmt.Sprintf("「%s」修改了文章：", req.Data.User.Name),
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》", req.Data.CommentAble.Title),
				Href: GetArticleUrl(req.Data.User.Login, req.Data.CommentAble.Book.Slug, req.Data.CommentAble.Slug),
			},
			{
				Tag:  "text",
				Text: "的评论",
			},
		}
	case CommentDelete, CommentReplyDelete:
		title = fmt.Sprintf("文章「%s」删除了一条评论", req.Data.CommentAble.Title)
		content = []*lark.PostContentItem{
			{
				Tag:  "text",
				Text: fmt.Sprintf("「%s」删除了文章：", req.Data.User.Name),
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》", req.Data.CommentAble.Title),
				Href: GetArticleUrl(req.Data.User.Login, req.Data.CommentAble.Book.Slug, req.Data.CommentAble.Slug),
			},
			{
				Tag:  "text",
				Text: "的评论",
			},
		}
	case Publish:
		title = fmt.Sprintf("「%s」有一篇新文章", req.Data.Book.Name)
		content = []*lark.PostContentItem{
			{
				Tag:  "text",
				Text: fmt.Sprintf("「%s」在知识库", req.Data.User.Name),
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》：", req.Data.Book.Name),
				Href: GetBookeUrl(req.Data.User.Login, req.Data.Book.Slug),
			},
			{
				Tag:  "text",
				Text: "发布了一篇新文章：",
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》", req.Data.Title),
				Href: GetArticleUrl(req.Data.User.Login, req.Data.Book.Slug, req.Data.Slug),
			},
		}
	case Update:
		title = fmt.Sprintf("文章「%s」有更新", req.Data.Title)
		content = []*lark.PostContentItem{
			{
				Tag:  "text",
				Text: fmt.Sprintf("「%s」更新了知识库：", req.Data.User.Name),
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》：", req.Data.Book.Name),
				Href: GetBookeUrl(req.Data.User.Login, req.Data.Book.Slug),
			},
			{
				Tag:  "text",
				Text: "中的文章：",
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》", req.Data.Title),
				Href: GetArticleUrl(req.Data.User.Login, req.Data.Book.Slug, req.Data.Slug),
			},
		}
	case Delete:
		title = fmt.Sprintf("文章「%s」被删除了", req.Data.Title)
		content = []*lark.PostContentItem{
			{
				Tag:  "text",
				Text: fmt.Sprintf("「%s」删除了知识库：", req.Data.User.Name),
			},
			{
				Tag:  "a",
				Text: fmt.Sprintf("《%s》：", req.Data.Book.Name),
				Href: GetBookeUrl(req.Data.User.Login, req.Data.Book.Slug),
			},
			{
				Tag:  "text",
				Text: "中的文章：",
			},
			{
				Tag:  "text",
				Text: fmt.Sprintf("《%s》", req.Data.Title),
			},
		}
	}
	message.Content = &lark.PostContent{
		Post: &lark.ZhCnPostContentData{
			ZhCn: lark.PostContentData{
				Title:   title,
				Content: [][]*lark.PostContentItem{content},
			},
		},
	}
	return message, nil
}

func (l *LarkWebHook) requestToCardMessage(ctx context.Context, req *WebHookRequest) (*lark.CardMessage, error) {
	return nil, nil
}

func (l *LarkWebHook) pushTextMessage(ctx context.Context, req *WebHookRequest) (*lark.TextMessage, error) {
	message, err := l.requestToTextMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	lark.LarkClient.BotHook(ctx, l.HookId, message, nil, nil)
	return nil, nil
}

func (l *LarkWebHook) pushPostMessage(ctx context.Context, req *WebHookRequest) (*lark.PostMessage, error) {
	post, err := l.requestToPostMessage(ctx, req)
	if err != nil {
		return nil, err
	}

	lark.LarkClient.BotHook(ctx, l.HookId, nil, post, nil)
	return nil, nil
}

func (l *LarkWebHook) pushCardMessage(ctx context.Context, req *WebHookRequest) (*lark.CardMessage, error) {
	return nil, nil
}

func (l *LarkWebHook) Push(ctx context.Context, req *WebHookRequest) error {
	if req == nil {
		return fmt.Errorf("req is invalid")
	}
	if l.MessageType == TextMessage {
		l.pushTextMessage(ctx, req)

	} else if l.MessageType == PostMessage {
		l.pushPostMessage(ctx, req)

	} else if l.MessageType == CardMessage {
		l.pushCardMessage(ctx, req)
	}
	return nil
}

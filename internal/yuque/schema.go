package yuque

type WebhookSubjectType string

const (
	CommentReplyCreate WebhookSubjectType = "comment_reply_create"
	CommentReplyUpdate WebhookSubjectType = "comment_reply_update"
	CommentReplyDelete WebhookSubjectType = "comment_reply_delete"
	CommentCreate      WebhookSubjectType = "comment_create"
	CommentUpdate      WebhookSubjectType = "comment_update"
	CommentDelete      WebhookSubjectType = "comment_delete"
	Publish            WebhookSubjectType = "publish"
	Update             WebhookSubjectType = "update"
	Delete             WebhookSubjectType = "delete"
)

type ValidRequest interface {
	Validate() error
}

type YuqueUser struct {
	Id             int    `json:"id"`
	Type           string `json:"type"`
	Login          string `json:""`            // 登录用户名
	Name           string `json:"name"`        // 用户名
	Description    string `json:"description"` // 登录用户描述
	AvatarUrl      string `json:"avatar_url"`  // 头像
	FollowersCount int    `json:"followers_count"`
	FollowingCount int    `json:"following_count"`
	CreatedAt      string `json:"created_at"`  //created_at - 注册时间
	UpdatedAt      string `json:"updated_at"`  //updated_at - 更新时间
	Serializer     string `json:"_serializer"` //
}

type YuqueBook struct {
	Id             int       `json:"id"`
	Type           string    `json:"type"`
	Slug           string    `json:"slug"` //slug - 文档路径
	Name           string    `json:"name"` //title - 标题
	UserId         int       `json:"user_id"`
	User           YuqueUser `json:"user"`
	CreatorId      int       `json:"creator_id"`
	Public         int       `json:"public"`          //public - 公开级别 [0 - 私密, 1 - 公开]
	LikesCount     int       `json:"likes_count"`     //likes_count - 赞数量
	WatchesCount   int       `json:"watches_count"`   //watches_count - 赞数量
	ItemsCount     int       `json:"items_count"`     //items_count - 赞数量
	Description    string    `json:"description"`     // 登录用户描述
	CreatedAt      string    `json:"created_at"`      //created_at - 注册时间
	UpdatedAt      string    `json:"updated_at"`      //updated_at - 更新时间
	ContentUpdated string    `json:"content_updated"` //updated_at - 更新时间

}

type CommentAbleData struct {
	Id         int       `json:"id"`
	Slug       string    `json:"slug"`        //slug - 文档路径
	Title      string    `json:"title"`       //title - 标题
	BookId     int       `json:"book_id"`     //book_id - 仓库编号，就是 repo_id
	Book       YuqueBook `json:"book"`        //book - 仓库信息 <BookSerializer>，就是 repo 信息
	Serializer string    `json:"_serializer"` //
}

type WebhookData struct {
	Id                 int                `json:"id"`                   // id - 文档编号
	Slug               string             `json:"slug"`                 //slug - 文档路径
	Title              string             `json:"title"`                //title - 标题
	BookId             int                `json:"book_id"`              //book_id - 仓库编号，就是 repo_id
	Book               YuqueBook          `json:"book"`                 //book - 仓库信息 <BookSerializer>，就是 repo 信息
	UserId             int                `json:"user_id"`              //user_id - 用户/团队编号
	User               YuqueUser          `json:"user"`                 //user - 用户/团队信息 <UserSerializer>
	Format             string             `json:"format"`               //format - 描述了正文的格式 [lake , markdown]
	Body               string             `json:"body"`                 //body - 正文 Markdown 源代码
	BodyDraft          string             `json:"body_draft"`           //body_draft - 草稿 Markdown 源代码
	BodyHtml           string             `json:"body_html"`            //body_html - 转换过后的正文 HTML
	BodyLake           string             `json:"body_lake"`            //body_lake - 语雀 lake 格式的文档内容
	CreatorId          int                `json:"creator_id"`           //creator_id - 文档创建人 User Id
	Public             int                `json:"public"`               //public - 公开级别 [0 - 私密, 1 - 公开]
	Status             int                `json:"status"`               //status - 状态 [0 - 草稿, 1 - 发布]
	ViewStatus         int                `json:"view_status"`          //view_status - 状态 []
	ReadStatus         int                `json:"read_status"`          //read_status - 状态 []
	LikesCount         int                `json:"likes_count"`          //likes_count - 赞数量
	CommentsCount      int                `json:"comments_count"`       //comments_count - 评论数量
	ContentUpdatedAt   string             `json:"content_updated_at"`   //content_updated_at - 文档内容更新时间
	DeletedAt          string             `json:"deleted_at"`           //deleted_at - 删除时间，未删除为 null
	CreatedAt          string             `json:"created_at"`           //created_at - 创建时间
	UpdatedAt          string             `json:"updated_at"`           //updated_at - 更新时间
	Serializer         string             `json:"_serializer"`          //
	WebhookSubjectType WebhookSubjectType `json:"webhook_subject_type"` //
	CommentAble        CommentAbleData    `json:"commentable"`
}

type WebHookRequest struct {
	Data WebhookData `form:"data" json:"data" binding:"required"`
}

func (req *WebHookRequest) Validate() error {
	//if req.ActionType != nil && *req.ActionType != Publish && *req.ActionType != Update && *req.ActionType != Delete {
	//	return fmt.Errorf("invalid action type:%s", *req.ActionType)
	//}
	return nil
}

func RequestValidate(req interface{}) error {
	if vr, ok := req.(ValidRequest); ok {
		if err := vr.Validate(); err != nil {
			return err
		}
	}
	return nil
}

package yuque

import "fmt"

const Host = "https://www.yuque.com"

func GetArticleUrl(user, bookSlug, articleSlug string) string {
	return fmt.Sprintf("%s/%s/%s/%s", Host, user, bookSlug, articleSlug)
}

func GetBookeUrl(user, bookSlug string) string {
	return fmt.Sprintf("%s/%s/%s", Host, user, bookSlug)
}

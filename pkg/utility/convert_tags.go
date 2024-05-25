package utility

import "github.com/eulbyvan/blog_api/internal/entity"

func ConvertTags(tagLabels []string) []entity.Tag {
	var tags []entity.Tag
	for _, label := range tagLabels {
		tags = append(tags, entity.Tag{Label: label})
	}
	return tags
}

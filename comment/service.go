package comment

import "../entity"

// CommentService specifies customer comment related bservice
type CommentService interface {
	Comments() ([]entity.Comment, error)
	Comment(id uint) (entity.Comment, error)
	UpdateComment(comment entity.Comment) (entity.Comment, error)
	DeleteComment(id uint) (entity.Comment, error)
	StoreComment(comment entity.Comment) (entity.Comment, error)
}

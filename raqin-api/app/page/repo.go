package page

import (
	"context"
	"database/sql"
	"errors"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type PageRepo interface {
	// PAGE
	PageByID(id int) (*repo.Page, error)
	PageByRevisionID(id int) (*repo.Page, error)

	// PAGE_REVISION
	NewRevision(pageRev *repo.PageRevision) (*repo.PageRevision, error)
	UpdateRevision(pageRev *repo.PageRevision) (int64, error)
	DeleteRevision(pageRev *repo.PageRevision) (int64, error)
	RevisionsByReviewerID(id int) (*repo.PageRevisionSlice, error)
	RevisionsByPageID(id int) (*repo.PageRevisionSlice, error)

	RevisionByID(id int) (*repo.PageRevision, error)
	AllRevisions() (*repo.PageRevisionSlice, error)

	// PAGE_REVISION_REACTION
	NewReaction(pReact *repo.PageRevisionReaction) error
	UpdateReaction(pReact *repo.PageRevisionReaction) (int64, error)
	ReactionsByRevisionID(id int) (*repo.PageRevisionReactionSlice, error)

	// PAGE_REVISION_COMMENTS
	NewComment(*repo.PageRevisionComment) error
	UpdateComment(*repo.PageRevisionComment) (int64, error)
	DeleteComment(*repo.PageRevisionComment) (int64, error)
	CommentsByRevisionID(id int) (*repo.PageRevisionCommentSlice, error)

	ApproveRevision(id int, stage string) (*repo.Page, error)
	MostApprovedRevisionByPageID(id int) (*repo.PageRevision, error)
}

type pageRepo struct {
	db *sql.DB
}

func NewPageRepo(db *sql.DB) *pageRepo {
	return &pageRepo{db}
}

func (pr *pageRepo) PageByID(id int) (*repo.Page, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	page, err := repo.Pages(
		qm.Where("id = ?", id)).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return page, nil
}

func (pr *pageRepo) PageByRevisionID(id int) (*repo.Page, error) {
	ctx := context.Background()
	tx, err := pr.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	rev, err := pr.RevisionByID(id)
	if err != nil {
		return nil, err
	}

	page, err := repo.Pages(
		qm.Where("approved_revision = ?", id)).One(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if page.ID != rev.PageID {
		return nil, errors.New("page and revision are not related")
	}

	tx.Commit()
	return page, nil
}
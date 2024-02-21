package fault

import "errors"

var (
	//repo
	RepoNotUpdated = errors.New("Update repo info failed")
	RepoNotFound   = errors.New("Repo not found")
	RepoConflict   = errors.New("Repo already exists")
	RepoInsertFail = errors.New("Add repo failed")

	//bookmark
	BookmarkNotFound = errors.New("Bookmark not found")
	BookmarkFail     = errors.New("Add bookmark failed")
	DelBookmarkFail  = errors.New("Delete Bookmark failed")
	BookmarkConflic  = errors.New("Bookmark already exists")

	//genneral
	ErrorSql = errors.New("SQL Error")
)
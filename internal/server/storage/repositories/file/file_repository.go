package file

import (
	"best-goph-keeper/internal/server/database"
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"database/sql"
	"time"
)

type File struct {
	db *database.DB
}

func New(db *database.DB) *File {
	return &File{
		db: db,
	}
}

func (f *File) UploadFile(binaryRequest *model.FileRequest) (*model.File, error) {
	binary := &model.File{}
	if err := f.db.Pool.QueryRow(
		"INSERT INTO file (user_id, name, created_at) VALUES ($1, $2, $3) "+
			"RETURNING file_id, name",
		binaryRequest.UserID,
		binaryRequest.Name,
		time.Now(),
	).Scan(&binary.ID, &binary.Name); err != nil {
		return nil, err
	}
	return binary, nil
}

func (f *File) GetListFile(userId int64) ([]model.File, error) {
	listFile := []model.File{}

	rows, err := f.db.Pool.Query("SELECT file_id, user_id, name, created_at FROM file "+
		"where user_id = $1 and deleted_at IS NULL", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		binary := model.File{}
		err = rows.Scan(&binary.ID, &binary.UserID, &binary.Name, &binary.CreatedAt)
		if err != nil {
			return nil, err
		}
		listFile = append(listFile, binary)
	}
	return listFile, nil
}

func (f *File) FileExists(binaryRequest *model.FileRequest) (bool, error) {
	var exists bool
	row := f.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM file "+
		"where file.user_id = $1 and file.name = $2 and file.deleted_at IS NULL)",
		binaryRequest.UserID,
		binaryRequest.Name)
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (f *File) DeleteFile(binaryRequest *model.FileRequest) (int64, error) {
	var id int64
	if err := f.db.Pool.QueryRow("UPDATE file SET deleted_at = $1 "+
		"where file.user_id = $2 and file.name = $3 RETURNING file_id",
		time.Now(),
		binaryRequest.UserID,
		binaryRequest.Name,
	).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

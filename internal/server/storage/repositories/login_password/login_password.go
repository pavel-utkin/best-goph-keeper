package login_password

import (
	"best-goph-keeper/internal/database"
	"best-goph-keeper/internal/server/model"
	"best-goph-keeper/internal/server/storage/errors"
	"best-goph-keeper/internal/server/storage/vars"
	"database/sql"
	"time"
)

type LoginPassword struct {
	db *database.DB
}

func New(db *database.DB) *LoginPassword {
	return &LoginPassword{
		db: db,
	}
}

func (lp *LoginPassword) CreateLoginPassword(loginPasswordRequest *model.CreateLoginPasswordRequest) (*model.LoginPassword, error) {
	loginPassword := &model.LoginPassword{}
	if err := lp.db.Pool.QueryRow(
		"INSERT INTO login_password (user_id, data, created_at, updated_at) VALUES ($1, $2, $3, $4) "+
			"RETURNING login_password_id, data",
		loginPasswordRequest.UserID,
		loginPasswordRequest.Data,
		time.Now(),
		time.Now(),
	).Scan(&loginPassword.ID, &loginPassword.Data); err != nil {
		return nil, err
	}
	return loginPassword, nil
}

func (lp *LoginPassword) KeyExists(loginPasswordRequest *model.CreateLoginPasswordRequest) (bool, error) {
	var exists bool
	row := lp.db.Pool.QueryRow("SELECT EXISTS(SELECT 1 FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4)",
		string(vars.Name), loginPasswordRequest.Name, loginPasswordRequest.UserID, string(vars.LoginPassword))
	if err := row.Scan(&exists); err != nil {
		return exists, err
	}
	return exists, nil
}

func (lp *LoginPassword) GetNodeLoginPassword(loginPasswordRequest *model.GetNodeLoginPasswordRequest) (*model.LoginPassword, error) {
	loginPassword := &model.LoginPassword{}
	err := lp.db.Pool.QueryRow("SELECT login_password.data FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4",
		string(vars.Name), loginPasswordRequest.Value, loginPasswordRequest.UserID, string(vars.LoginPassword)).
		Scan(&loginPassword.Data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	return loginPassword, nil
}

func (lp *LoginPassword) GetListLoginPassword(userId int64) ([]model.LoginPassword, error) {
	listLoginPassword := []model.LoginPassword{}
	rows, err := lp.db.Pool.Query("SELECT metadata.entity_id, metadata.key, login_password.data, metadata.value, login_password.created_at, "+
		"login_password.updated_at FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where users.user_id = $1 and metadata.type = $2 and login_password.deleted_at IS NULL",
		userId, string(vars.LoginPassword))

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		} else {
			return nil, err
		}
	}
	defer rows.Close()
	for rows.Next() {
		loginPassword := model.LoginPassword{}
		err = rows.Scan(&loginPassword.ID, &loginPassword.Key, &loginPassword.Data, &loginPassword.Value, &loginPassword.CreatedAt, &loginPassword.UpdatedAt)
		if err != nil {
			return nil, err
		}
		listLoginPassword = append(listLoginPassword, loginPassword)
	}
	return listLoginPassword, nil
}

func (lp *LoginPassword) GetIdLoginPassword(value string, userID int64) (int64, error) {
	var loginPasswordID int64
	err := lp.db.Pool.QueryRow("SELECT login_password.login_password_id FROM metadata "+
		"inner join login_password on metadata.entity_id = login_password.login_password_id "+
		"inner join users on login_password.user_id  = users.user_id "+
		"where metadata.key = $1 and metadata.value = $2 and users.user_id = $3 and metadata.type = $4",
		string(vars.Name), value, userID, string(vars.LoginPassword)).
		Scan(&loginPasswordID)
	if err != nil {
		if err == sql.ErrNoRows {
			return loginPasswordID, errors.ErrRecordNotFound
		} else {
			return loginPasswordID, err
		}
	}
	return loginPasswordID, nil
}

func (lp *LoginPassword) DeleteLoginPassword(entityId int64) error {
	var id int64
	layout := "01/02/2006 15:04:05"
	if err := lp.db.Pool.QueryRow("UPDATE login_password SET deleted_at = $1 WHERE login_password_id = $2 RETURNING login_password_id",
		time.Now().Format(layout),
		entityId,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}

func (lp *LoginPassword) UpdateLoginPassword(textID int64, data []byte) error {
	var id int64
	layout := "01/02/2006 15:04:05"
	if err := lp.db.Pool.QueryRow("UPDATE login_password SET data = $1, updated_at = $2 WHERE login_password_id = $3 RETURNING login_password_id",
		data,
		time.Now().Format(layout),
		textID,
	).Scan(&id); err != nil {
		return err
	}
	return nil
}

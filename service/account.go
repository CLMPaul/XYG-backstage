package service

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack/v5"
	"os"
	"strconv"
	"time"
	"xueyigou_demo/constants/redis_keys"
	"xueyigou_demo/models"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"xueyigou_demo/cache"
	"xueyigou_demo/constants/errcodes"
	"xueyigou_demo/db"
	"xueyigou_demo/internal/utils"
)

type accountService struct {
	ctx context.Context
}

type AccountForSelect struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountForAuthenticate struct {
	Name            string `json:"name"`
	ProjectID       string `json:"project_id"`
	Username        string `json:"user_name"`
	IsSuperuser     bool   `json:"is_superuser"`
	IsAdministrator bool   `json:"is_administrator"`
	Locked          bool   `json:"locked"`
	UpdateTime      *time.Time
}

type AccountSession struct {
	AccountID  string
	UpdateTime *time.Time
	ClientIP   string
}

type LoginAccount struct {
	UID             string
	Name            string
	ProjectID       string
	Username        string
	IsSuperuser     bool
	IsAdministrator bool
	IsAPIKey        bool
}

func NewAPIAccount() *LoginAccount {
	return &LoginAccount{
		IsSuperuser:     true,
		IsAdministrator: true,
		IsAPIKey:        true,
	}
}

func (s *accountService) WithContext(ctx context.Context) *accountService {
	return &accountService{ctx: ctx}
}

func (s *accountService) context() context.Context {
	if s.ctx != nil {
		return s.ctx
	}
	return context.Background()
}

func (s *accountService) GetForAuthenticate(userID string) (*AccountForAuthenticate, error) {
	var user AccountForAuthenticate
	if cached, err := cache.GetClient().Get(s.context(), redis_keys.AccountForAuthenticate(userID)).Bytes(); err == nil {
		if err := json.Unmarshal(cached, &user); err != nil {
			logrus.Debugf("cannot parse cached AccountForAuthenticate as json: %v", err)
		} else {
			return &user, nil
		}
	}
	result := db.DB.
		Model(&models.SuperUser{}).
		Select([]string{"project_id", "username", "name", "locked", "is_superuser", "is_administrator", "update_time"}).
		Where(clause.Eq{Column: clause.Column{Name: "id"}, Value: userID}).
		Limit(1).
		Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	if cached, err := json.Marshal(&user); err == nil {
		if err := cache.GetClient().Set(context.Background(), redis_keys.AccountForAuthenticate(userID), cached, 24*time.Hour).Err(); err != nil {
			logrus.Debugf("failed setting cached AccountForAuthenticate: %v", err)
		}
	}
	return &user, nil
}

func (s *accountService) Authenticate(token string) (loginAccount *LoginAccount, err error) {
	sessionKey := redis_keys.Session(token)

	var sessionBytes []byte
	if RedisSupportsGetEx {
		sessionBytes, err = cache.GetClient().GetEx(s.context(), sessionKey, SessionExpiration).Bytes()
	} else {
		sessionBytes, err = cache.GetClient().Get(s.context(), sessionKey).Bytes()
	}

	if err != nil {
		if !errors.Is(err, redis.Nil) {
			err = errors.Wrap(err, "failed reading session")
		} else {
			err = nil
		}
		return
	}

	var session AccountSession
	if err = msgpack.Unmarshal(sessionBytes, &session); err != nil {
		err = errors.Wrap(err, "failed unmarshaling session")
		return
	}

	user, err := s.GetForAuthenticate(session.AccountID)
	if err != nil {
		err = errors.Wrap(err, "failed verifying user from token")
		return
	}
	if user == nil {
		err = errcodes.SessionExpired
		return
	}
	if user.Locked {
		err = errcodes.SessionExpired
		return
	}
	if user.UpdateTime != nil {
		// 如果用户的 UpdateTime 比会话中的新，强制用户下线
		if session.UpdateTime == nil || session.UpdateTime.Before(*user.UpdateTime) {
			err = errcodes.SessionExpired
			return
		}
	}

	// 保活
	if !RedisSupportsGetEx {
		// _ = cache.GetClient().Set(context.Background(), sessionKey, sessionBytes, SessionExpiration).Err()
		cache.GetClient().Expire(context.Background(), sessionKey, SessionExpiration)
	}
	loginAccount = &LoginAccount{
		UID:             session.AccountID,
		Name:            user.Name,
		ProjectID:       user.ProjectID,
		Username:        user.Username,
		IsSuperuser:     user.IsSuperuser,
		IsAdministrator: user.IsSuperuser || user.IsAdministrator,
	}
	return
}

func (s *accountService) NewSession(session AccountSession) (string, error) {
	token := utils.UUIDWithoutDash()
	sessionKey := redis_keys.Session(token)
	sessionBytes, err := msgpack.Marshal(session)
	if err != nil {
		return "", err
	}
	err = cache.GetClient().Set(s.context(), sessionKey, sessionBytes, SessionExpiration).Err()
	if err != nil {
		return "", err
	}
	return token, err
}

func (s *accountService) KeepSession(token string) {
	err := cache.GetClient().Expire(context.Background(), redis_keys.Session(token), SessionExpiration).Err()
	if err != nil {
		logrus.WithError(err).Error("keepsession error")
	}
}

var checkIPChange = true

func init() {
	if val, ok := os.LookupEnv("PLATFORM_CHECK_IP_CHANGE"); ok {
		checkIPChange, _ = strconv.ParseBool(val)
	}
}

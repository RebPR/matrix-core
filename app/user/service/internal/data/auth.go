package data

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"github.com/the-zion/matrix-core/app/user/service/internal/biz"
	"github.com/the-zion/matrix-core/app/user/service/internal/pkg/util"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var _ biz.AuthRepo = (*authRepo)(nil)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "user/data/auth")),
	}
}

func (r *authRepo) FindUserByPhone(ctx context.Context, phone string) (*biz.User, error) {
	user := &User{}
	err := r.data.db.WithContext(ctx).Where("phone = ?", phone).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kerrors.NotFound("phone not found from db", fmt.Sprintf("phone(%s)", phone))
	}
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to find user by phone: phone(%s)", phone))
	}
	return &biz.User{
		Uuid:     user.Uuid,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
		Wechat:   user.Wechat,
		Github:   user.Github,
	}, nil
}

func (r *authRepo) FindUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	user := &User{}
	err := r.data.db.WithContext(ctx).Where("email = ?", email).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kerrors.NotFound("email not found from db", fmt.Sprintf("email(%s)", email))
	}
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to find user by email: email(%s)", email))
	}
	return &biz.User{
		Uuid:     user.Uuid,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
		Wechat:   user.Wechat,
		Github:   user.Github,
	}, nil
}

func (r *authRepo) FindUserByGithub(ctx context.Context, github int32) (*biz.User, error) {
	user := &User{}
	err := r.data.db.WithContext(ctx).Where("github = ?", github).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, kerrors.NotFound("github not found from db", fmt.Sprintf("github(%v)", github))
	}
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to find user by github: github(%v)", github))
	}
	return &biz.User{
		Uuid:     user.Uuid,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
		Wechat:   user.Wechat,
		Github:   user.Github,
	}, nil
}

func (r *authRepo) CreateUserWithPhone(ctx context.Context, phone string) (*biz.User, error) {
	uuid, err := util.UUIdV4()
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to create uuid: uuid(%s)", uuid))
	}

	user := &User{
		Uuid:  uuid,
		Phone: phone,
	}
	err = r.data.DB(ctx).Select("Phone", "Uuid").Create(user).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to create a user: phone(%s)", phone))
	}

	return &biz.User{
		Uuid:  uuid,
		Phone: phone,
	}, nil
}

func (r *authRepo) CreateUserWithEmail(ctx context.Context, email, password string) (*biz.User, error) {
	uuid, err := util.UUIdV4()
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to create uuid: uuid(%s)", uuid))
	}

	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to hash password: password(%s)", password))
	}

	user := &User{
		Uuid:     uuid,
		Email:    email,
		Password: hashPassword,
	}
	err = r.data.DB(ctx).Select("Email", "Uuid", "Password").Create(user).Error
	if err != nil {
		e := err.Error()
		if strings.Contains(e, "Duplicate") {
			return nil, kerrors.Conflict("email conflict", fmt.Sprintf("email(%s)", email))
		} else {
			return nil, errors.Wrapf(err, fmt.Sprintf("fail to create a user: email(%s)", email))
		}
	}

	return &biz.User{
		Uuid:     uuid,
		Email:    email,
		Password: hashPassword,
	}, nil
}

func (r *authRepo) CreateUserWithGithub(ctx context.Context, github int32) (*biz.User, error) {
	uuid, err := util.UUIdV4()
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to create uuid: uuid(%s)", uuid))
	}

	user := &User{
		Uuid:   uuid,
		Github: github,
	}
	err = r.data.DB(ctx).Select("Github", "Uuid").Create(user).Error
	if err != nil {
		e := err.Error()
		if strings.Contains(e, "Duplicate") {
			return nil, kerrors.Conflict("github conflict", fmt.Sprintf("github(%v)", github))
		} else {
			return nil, errors.Wrapf(err, fmt.Sprintf("fail to create a user: github(%v)", github))
		}
	}

	return &biz.User{
		Uuid:   uuid,
		Github: github,
	}, nil
}

func (r *authRepo) CreateUserProfile(ctx context.Context, account, uuid string) error {
	p := &Profile{
		Uuid:     uuid,
		Username: account,
		Updated:  time.Now().Unix(),
	}
	err := r.data.DB(ctx).Select("Uuid", "Username", "Updated").Create(p).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to register a profile: uuid(%s)", uuid))
	}
	return nil
}

func (r *authRepo) CreateUserProfileWithGithub(ctx context.Context, account, github, uuid string) error {
	p := &Profile{
		Uuid:     uuid,
		Username: account,
		Github:   github,
		Updated:  time.Now().Unix(),
	}
	err := r.data.DB(ctx).Select("Uuid", "Username", "Github", "Updated").Create(p).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to register a profile: uuid(%s)", uuid))
	}
	return nil
}

func (r *authRepo) CreateUserProfileUpdate(ctx context.Context, account, uuid string) error {
	pu := &ProfileUpdate{}
	pu.Uuid = uuid
	pu.Username = account
	pu.Updated = time.Now().Unix()
	err := r.data.DB(ctx).Select("Uuid", "Username", "Updated").Create(pu).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to create table: profile_update, uuid(%s)", uuid))
	}
	return nil
}

func (r *authRepo) CreateUserSearch(ctx context.Context, account, uuid string) error {
	user := &biz.UserSearchMap{
		Username:  account,
		Introduce: "",
	}
	body, err := user.MarshalJSON()
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("error marshaling document: account(%s), uuid(%s)", account, uuid))
	}

	req := esapi.IndexRequest{
		Index:      "user",
		DocumentID: uuid,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, r.data.elasticSearch.es)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("error getting user search create response: account(%s), uuid(%s)", account, uuid))
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return errors.Wrapf(err, fmt.Sprintf("error parsing the response body: account(%s), uuid(%s)", account, uuid))
		} else {
			return errors.Errorf(fmt.Sprintf("error indexing document to es: reason(%v), account(%s), uuid(%s)", e, account, uuid))
		}
	}
	return nil
}

func (r *authRepo) SetUserPhone(ctx context.Context, uuid, phone string) error {
	err := r.data.db.WithContext(ctx).Model(&User{}).Where("uuid = ?", uuid).Update("phone", phone).Error
	if err != nil {
		e := err.Error()
		if strings.Contains(e, "Duplicate") {
			return kerrors.Conflict("phone conflict", fmt.Sprintf("uuid(%s), phone(%s)", uuid, phone))
		} else {
			return errors.Wrapf(err, fmt.Sprintf("fail to set user uphone: uuid(%s), phone(%s)", uuid, phone))
		}
	}
	return nil
}

func (r *authRepo) SetUserEmail(ctx context.Context, uuid, email string) error {
	err := r.data.db.WithContext(ctx).Model(&User{}).Where("uuid = ?", uuid).Update("email", email).Error
	if err != nil {
		e := err.Error()
		if strings.Contains(e, "Duplicate") {
			return kerrors.Conflict("email conflict", fmt.Sprintf("uuid(%s), email(%s)", uuid, email))
		} else {
			return errors.Wrapf(err, fmt.Sprintf("fail to set user email: uuid(%s), email(%s)", uuid, email))
		}
	}
	return nil
}

func (r *authRepo) SetUserPassword(ctx context.Context, uuid, password string) error {
	p, err := util.HashPassword(password)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to hash password: %s", password))
	}

	err = r.data.db.WithContext(ctx).Model(&User{}).Where("uuid = ?", uuid).Update("password", p).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to set user password: uuid(%s), password(%s)", uuid, password))
	}
	return nil
}

func (r *authRepo) SendPhoneCode(ctx context.Context, template, phone string) error {
	code := util.RandomNumber()
	err := r.setCodeToCache(ctx, "phone_"+phone, code)
	if err != nil {
		return err
	}

	message := strings.Join([]string{phone, code, template, "phone"}, ";")
	msg := &primitive.Message{
		Topic: "code",
		Body:  []byte(message),
	}
	msg.WithTag("phone")
	err = r.data.codeMqPro.producer.SendOneWay(ctx, msg)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to send code to producer: %s", message))
	}

	return nil
}

func (r *authRepo) SendEmailCode(ctx context.Context, template, email string) error {
	code := util.RandomNumber()
	err := r.setCodeToCache(ctx, "email_"+email, code)
	if err != nil {
		return err
	}

	message := strings.Join([]string{email, code, template, "email"}, ";")
	msg := &primitive.Message{
		Topic: "code",
		Body:  []byte(message),
	}
	msg.WithTag("email")
	err = r.data.codeMqPro.producer.SendOneWay(ctx, msg)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to send code to producer: %s", message))
	}

	return nil
}

func (r *authRepo) VerifyPhoneCode(ctx context.Context, phone, code string) error {
	key := "phone_" + phone
	return r.verifyCode(ctx, key, code)
}

func (r *authRepo) VerifyEmailCode(ctx context.Context, email, code string) error {
	key := "email_" + email
	return r.verifyCode(ctx, key, code)
}

func (r *authRepo) verifyCode(ctx context.Context, key, code string) error {
	codeInCache, err := r.getCodeFromCache(ctx, key)
	if err != nil {
		return err
	}
	if code != codeInCache {
		return errors.Errorf("code error: code(%s)", code)
	}
	r.removeCodeFromCache(ctx, key)
	return nil
}

func (r *authRepo) VerifyPassword(ctx context.Context, account, password, mode string) (*biz.User, error) {
	var err error
	var user *biz.User
	if mode == "phone" {
		user, err = r.FindUserByPhone(ctx, account)
	} else {
		user, err = r.FindUserByEmail(ctx, account)
	}
	if err != nil {
		return nil, err
	}

	pass := util.CheckPasswordHash(password, user.Password)
	if !pass {
		return nil, errors.Errorf("password error: password(%s)", password)
	}
	return user, nil
}

func (r *authRepo) PasswordResetByPhone(ctx context.Context, phone, password string) error {
	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to hash password: password(%s)", password))
	}

	err = r.data.db.WithContext(ctx).Model(&User{}).Where("phone = ?", phone).Update("password", hashPassword).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to reset password: password(%s)", password))
	}
	return nil
}

func (r *authRepo) PasswordResetByEmail(ctx context.Context, email, password string) error {
	hashPassword, err := util.HashPassword(password)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to hash password: password(%s)", password))
	}

	err = r.data.db.WithContext(ctx).Model(&User{}).Where("email = ?", email).Update("password", hashPassword).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to reset password: password(%s)", password))
	}
	return nil
}

func (r *authRepo) GetCosSessionKey(_ context.Context, uuid string) (*biz.Credentials, error) {
	c := r.data.cos.client
	opt := r.data.cos.opt

	optNew := &sts.CredentialOptions{}
	var buf bytes.Buffer

	err := gob.NewEncoder(&buf).Encode(opt)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to encoder cos credential options")
	}

	err = gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(optNew)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to decode cos credential options")
	}

	for _, statement := range optNew.Policy.Statement {
		for index, _ := range statement.Resource {
			statement.Resource[index] += uuid + "/*"
		}
	}
	res, err := c.GetCredential(optNew)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get cos session key")
	}
	return &biz.Credentials{
		TmpSecretKey: res.Credentials.TmpSecretKey,
		TmpSecretID:  res.Credentials.TmpSecretID,
		SessionToken: res.Credentials.SessionToken,
		StartTime:    int64(res.StartTime),
		ExpiredTime:  int64(res.ExpiredTime),
	}, nil
}

func (r *authRepo) UnbindUserPhone(ctx context.Context, uuid string) error {
	err := r.data.db.WithContext(ctx).Model(&User{}).Where("uuid = ?", uuid).Update("phone", nil).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to unbind user phone: uuid(%s)", uuid))
	}
	return nil
}

func (r *authRepo) UnbindUserEmail(ctx context.Context, uuid string) error {
	err := r.data.db.WithContext(ctx).Model(&User{}).Where("uuid = ?", uuid).Update("email", nil).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to unbind user email: uuid(%s)", uuid))
	}
	return nil
}

func (r *authRepo) GetGithubAccessToken(ctx context.Context, code string) (string, error) {
	url := "https://github.com/login/oauth/access_token?client_id=" + r.data.github.clientId + "&client_secret=" + r.data.github.clientSecret + "&code=" + code
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return "", errors.Wrapf(err, fmt.Sprintf("fail to new github access token request: code(%s)", code))
	}

	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, fmt.Sprintf("fail to get github access token: code(%s)", code))
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	data := map[string]string{}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", errors.Wrapf(err, fmt.Sprintf("fail to unmarshal responce data: code(%s)", code))
	}
	token := data["access_token"]
	return token, nil
}

func (r *authRepo) GetGithubUserInfo(ctx context.Context, token string) (map[string]interface{}, error) {
	url := "https://api.github.com/user"
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to new request: token(%s)", token))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "token "+token)
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get github user info: token(%s)", token))
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d", res.StatusCode)
	}

	data := map[string]interface{}{}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to unmarshal responce data: token(%s)", token))
	}
	return data, nil
}

func (r *authRepo) setCodeToCache(ctx context.Context, key, code string) error {
	err := r.data.redisCli.Set(ctx, key, code, time.Minute*2).Err()
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to set code to cache: redis.Set(%v), code(%s)", key, code))
	}
	return nil
}

func (r *authRepo) getCodeFromCache(ctx context.Context, key string) (string, error) {
	code, err := r.data.redisCli.Get(ctx, key).Result()
	if err != nil {
		return "", errors.Wrapf(err, fmt.Sprintf("fail to get code from cache: redis.Get(%v)", key))
	}
	return code, nil
}

func (r *authRepo) removeCodeFromCache(ctx context.Context, key string) {
	_, err := r.data.redisCli.Del(ctx, key).Result()
	if err != nil {
		r.log.Errorf("fail to delete code from cache: redis.Del(key, %v), error(%v)", key, err)
	}
}

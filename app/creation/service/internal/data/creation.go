package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/the-zion/matrix-core/app/creation/service/internal/biz"
	"golang.org/x/sync/errgroup"
	"strconv"
	"strings"
	"time"
)

var _ biz.CreationRepo = (*creationRepo)(nil)

type creationRepo struct {
	data *Data
	log  *log.Helper
}

func NewCreationRepo(data *Data, logger log.Logger) biz.CreationRepo {
	return &creationRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "creation/data/creation")),
	}
}

func (r *creationRepo) GetLeaderBoard(ctx context.Context) ([]*biz.LeaderBoard, error) {
	return r.getLeaderBoardFromCache(ctx)
}

func (r *creationRepo) getLeaderBoardFromCache(ctx context.Context) ([]*biz.LeaderBoard, error) {
	list, err := r.data.redisCli.ZRevRange(ctx, "leaderboard", 0, 9).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "fail to get leader board from cache")
	}

	board := make([]*biz.LeaderBoard, 0)
	for _, item := range list {
		member := strings.Split(item, "%")
		id, err := strconv.ParseInt(member[0], 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("fail to covert string to int64: id(%s)", member[0]))
		}
		board = append(board, &biz.LeaderBoard{
			Id:   int32(id),
			Uuid: member[1],
			Mode: member[2],
		})
	}
	return board, nil
}

func (r *creationRepo) GetCollectArticle(ctx context.Context, id, page int32) ([]*biz.Article, error) {
	if page < 1 {
		page = 1
	}
	index := int(page - 1)
	list := make([]*Collect, 0)
	err := r.data.db.WithContext(ctx).Where("collections_id = ? and mode = ? and status = ?", id, 1, 1).Order("updated_at desc").Offset(index * 10).Limit(10).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get collect article from db: id(%v), page(%v)", id, page))
	}

	article := make([]*biz.Article, 0)
	for _, item := range list {
		article = append(article, &biz.Article{
			ArticleId: item.CreationsId,
			Uuid:      item.Uuid,
		})
	}
	return article, nil
}

func (r *creationRepo) GetCollectArticleCount(ctx context.Context, id int32) (int32, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Model(&Collect{}).Where("collections_id = ? and mode = ? and status = ?", id, 1, 1).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to get collect article count from db: id(%v)", id))
	}
	return int32(count), nil
}

func (r *creationRepo) GetCollectTalk(ctx context.Context, id, page int32) ([]*biz.Talk, error) {
	if page < 1 {
		page = 1
	}
	index := int(page - 1)
	list := make([]*Collect, 0)
	err := r.data.db.WithContext(ctx).Where("collections_id = ? and mode = ? and status = ?", id, 2, 1).Order("updated_at desc").Offset(index * 10).Limit(10).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get collect talk from db: id(%v), page(%v)", id, page))
	}

	talk := make([]*biz.Talk, 0)
	for _, item := range list {
		talk = append(talk, &biz.Talk{
			TalkId: item.CreationsId,
			Uuid:   item.Uuid,
		})
	}
	return talk, nil
}

func (r *creationRepo) GetCollectTalkCount(ctx context.Context, id int32) (int32, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Model(&Collect{}).Where("collections_id = ? and mode = ? and status = ?", id, 2, 1).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to get collect talk count from db: id(%v)", id))
	}
	return int32(count), nil
}

func (r *creationRepo) GetCollectColumn(ctx context.Context, id, page int32) ([]*biz.Column, error) {
	if page < 1 {
		page = 1
	}
	index := int(page - 1)
	list := make([]*Collect, 0)
	err := r.data.db.WithContext(ctx).Where("collections_id = ? and mode = ? and status = ?", id, 3, 1).Order("updated_at desc").Offset(index * 10).Limit(10).Find(&list).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get collect column from db: id(%v), page(%v)", id, page))
	}

	column := make([]*biz.Column, 0)
	for _, item := range list {
		column = append(column, &biz.Column{
			ColumnId: item.CreationsId,
			Uuid:     item.Uuid,
		})
	}
	return column, nil
}

func (r *creationRepo) GetCollectColumnCount(ctx context.Context, id int32) (int32, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Model(&Collect{}).Where("collections_id = ? and mode = ? and status = ?", id, 3, 1).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to get collect column count from db: id(%v)", id))
	}
	return int32(count), nil
}

func (r *creationRepo) GetCollection(ctx context.Context, id int32, uuid string) (*biz.Collections, error) {
	collections := &Collections{}
	err := r.data.db.WithContext(ctx).Where("id = ?", id).First(collections).Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get collection from db: id(%v)", id))
	}
	if collections.Auth == 2 && collections.Uuid != uuid {
		return nil, errors.Errorf("fail to get collection: no auth")
	}
	return &biz.Collections{
		Uuid:      collections.Uuid,
		Name:      collections.Name,
		Introduce: collections.Introduce,
		Auth:      collections.Auth,
	}, nil
}

func (r *creationRepo) GetCollectionListInfo(ctx context.Context, ids []int32) ([]*biz.Collections, error) {
	collectionsListInfo := make([]*biz.Collections, 0)
	exists, unExists, err := r.collectionsListInfoExist(ctx, ids)
	if err != nil {
		return nil, err
	}

	g, _ := errgroup.WithContext(ctx)
	g.Go(r.data.GroupRecover(ctx, func(ctx context.Context) error {
		if len(exists) == 0 {
			return nil
		}
		return r.getCollectionsListInfoFromCache(ctx, exists, &collectionsListInfo)
	}))
	g.Go(r.data.GroupRecover(ctx, func(ctx context.Context) error {
		if len(unExists) == 0 {
			return nil
		}
		return r.getCollectionsListInfoFromDb(ctx, unExists, &collectionsListInfo)
	}))

	err = g.Wait()
	if err != nil {
		return nil, err
	}

	return collectionsListInfo, nil
}

func (r *creationRepo) collectionsListInfoExist(ctx context.Context, ids []int32) ([]int32, []int32, error) {
	exists := make([]int32, 0)
	unExists := make([]int32, 0)
	cmd, err := r.data.redisCli.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, item := range ids {
			pipe.Exists(ctx, "collection_"+strconv.Itoa(int(item)))
		}
		return nil
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, fmt.Sprintf("fail to check if collection exist from cache: ids(%v)", ids))
	}

	for index, item := range cmd {
		exist := item.(*redis.IntCmd).Val()
		if exist == 1 {
			exists = append(exists, ids[index])
		} else {
			unExists = append(unExists, ids[index])
		}
	}
	return exists, unExists, nil
}

func (r *creationRepo) getCollectionsListInfoFromCache(ctx context.Context, exists []int32, collectionsListInfo *[]*biz.Collections) error {
	cmd, err := r.data.redisCli.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, id := range exists {
			pipe.HMGet(ctx, "collection_"+strconv.Itoa(int(id)), "name", "introduce")
		}
		return nil
	})
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to get collections list info from cache: ids(%v)", exists))
	}

	for index, item := range cmd {
		val := []string{"", ""}
		for _index, info := range item.(*redis.SliceCmd).Val() {
			if info == nil {
				break
			}
			val[_index] = info.(string)
		}
		*collectionsListInfo = append(*collectionsListInfo, &biz.Collections{
			Id:        exists[index],
			Name:      val[0],
			Introduce: val[1],
		})
	}
	return nil
}

func (r *creationRepo) getCollectionsListInfoFromDb(ctx context.Context, unExists []int32, collectionsListInfo *[]*biz.Collections) error {
	list := make([]*Collections, 0)
	err := r.data.db.WithContext(ctx).Where("id IN ?", unExists).Find(&list).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to get collections list info from db: ids(%v)", unExists))
	}

	for _, item := range list {
		*collectionsListInfo = append(*collectionsListInfo, &biz.Collections{
			Id:        int32(item.ID),
			Uuid:      item.Uuid,
			Auth:      item.Auth,
			Name:      item.Name,
			Introduce: item.Introduce,
		})
	}

	if len(list) != 0 {
		go r.setCollectionsListInfoToCache(list)
	}
	return nil
}

func (r *creationRepo) setCollectionsListInfoToCache(collectionsList []*Collections) {
	_, err := r.data.redisCli.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		for _, item := range collectionsList {
			key := "collection_" + strconv.Itoa(int(item.ID))
			pipe.HSetNX(context.Background(), key, "uuid", item.Uuid)
			pipe.HSetNX(context.Background(), key, "auth", item.Auth)
			pipe.HSetNX(context.Background(), key, "name", item.Name)
			pipe.HSetNX(context.Background(), key, "introduce", item.Introduce)
			pipe.Expire(context.Background(), key, time.Hour*8)
		}
		return nil
	})

	if err != nil {
		r.log.Errorf("fail to set collections info to cache, err(%s)", err.Error())
	}
}

func (r *creationRepo) GetCollectCount(ctx context.Context, id int32) (int64, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Model(&Collect{}).Where("collections_id = ? and status = ?", id, 1).Limit(1).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to get collect count from db: collections_id(%v)", id))
	}
	return count, nil
}

func (r *creationRepo) GetCollections(ctx context.Context, uuid string, page int32) ([]*biz.Collections, error) {
	collections, err := r.getUserCollectionsListFromCache(ctx, page, uuid)
	if err != nil {
		return nil, err
	}

	size := len(collections)
	if size != 0 {
		return collections, nil
	}

	collections, err = r.getUserCollectionsListFromDB(ctx, page, uuid)
	if err != nil {
		return nil, err
	}

	size = len(collections)
	if size != 0 {
		go r.setUserCollectionsListToCache("user_collections_list_"+uuid, collections)
	}
	return collections, nil
}

func (r *creationRepo) getUserCollectionsListFromCache(ctx context.Context, page int32, uuid string) ([]*biz.Collections, error) {
	if page < 1 {
		page = 1
	}
	index := int64(page - 1)
	list, err := r.data.redisCli.ZRevRange(ctx, "user_collections_list_"+uuid, index*10, index*10+9).Result()
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get user collections list visitor from cache: key(%s), page(%v)", "user_collections_list_", page))
	}

	collections := make([]*biz.Collections, 0)
	for _, item := range list {
		id, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("fail to covert string to int64: id(%s)", item))
		}
		collections = append(collections, &biz.Collections{
			Id: int32(id),
		})
	}
	return collections, nil
}

func (r *creationRepo) getUserCollectionsListFromDB(ctx context.Context, page int32, uuid string) ([]*biz.Collections, error) {
	if page < 1 {
		page = 1
	}
	index := int(page - 1)
	list := make([]*Collections, 0)
	handle := r.data.db.WithContext(ctx).Where("uuid = ?", uuid).Order("id desc").Offset(index * 10).Limit(10).Find(&list)
	err := handle.Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get collections from db: uuid(%s)", uuid))
	}
	collections := make([]*biz.Collections, 0)
	for _, item := range list {
		collections = append(collections, &biz.Collections{
			Id: int32(item.ID),
		})
	}
	return collections, nil
}

func (r *creationRepo) GetCollectionsByVisitor(ctx context.Context, uuid string, page int32) ([]*biz.Collections, error) {
	collections, err := r.getUserCollectionsListVisitorFromCache(ctx, page, uuid)
	if err != nil {
		return nil, err
	}

	size := len(collections)
	if size != 0 {
		return collections, nil
	}

	collections, err = r.getUserCollectionsListVisitorFromDB(ctx, page, uuid)
	if err != nil {
		return nil, err
	}

	size = len(collections)
	if size != 0 {
		go r.setUserCollectionsListToCache("user_collections_list_visitor_"+uuid, collections)
	}
	return collections, nil
}

func (r *creationRepo) getUserCollectionsListVisitorFromCache(ctx context.Context, page int32, uuid string) ([]*biz.Collections, error) {
	if page < 1 {
		page = 1
	}
	index := int64(page - 1)
	list, err := r.data.redisCli.ZRevRange(ctx, "user_collections_list_visitor_"+uuid, index*10, index*10+9).Result()
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get user collections list visitor from cache: key(%s), page(%v)", "user_collections_list_visitor_", page))
	}

	collections := make([]*biz.Collections, 0)
	for _, item := range list {
		id, err := strconv.ParseInt(item, 10, 32)
		if err != nil {
			return nil, errors.Wrapf(err, fmt.Sprintf("fail to covert string to int64: id(%s)", item))
		}
		collections = append(collections, &biz.Collections{
			Id: int32(id),
		})
	}
	return collections, nil
}

func (r *creationRepo) getUserCollectionsListVisitorFromDB(ctx context.Context, page int32, uuid string) ([]*biz.Collections, error) {
	if page < 1 {
		page = 1
	}
	index := int(page - 1)
	list := make([]*Collections, 0)
	handle := r.data.db.WithContext(ctx).Where("uuid = ? and auth = ?", uuid, 1).Order("id desc").Offset(index * 10).Limit(10).Find(&list)
	err := handle.Error
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("fail to get collections visitor from db: uuid(%s)", uuid))
	}
	collections := make([]*biz.Collections, 0)
	for _, item := range list {
		collections = append(collections, &biz.Collections{
			Id: int32(item.ID),
		})
	}
	return collections, nil
}

func (r *creationRepo) setUserCollectionsListToCache(key string, collections []*biz.Collections) {
	_, err := r.data.redisCli.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
		z := make([]*redis.Z, 0)
		for _, item := range collections {
			z = append(z, &redis.Z{
				Score:  float64(item.Id),
				Member: strconv.Itoa(int(item.Id)),
			})
		}
		pipe.ZAddNX(context.Background(), key, z...)
		pipe.Expire(context.Background(), key, time.Hour*8)
		return nil
	})
	if err != nil {
		r.log.Errorf("fail to set talk to cache: collections(%v)", collections)
	}
}

func (r *creationRepo) GetCollectionsCount(ctx context.Context, uuid string) (int32, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Model(&Collections{}).Where("uuid = ?", uuid).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to get collections count from db: uuid(%s)", uuid))
	}
	return int32(count), nil
}

func (r *creationRepo) GetCollectionsVisitorCount(ctx context.Context, uuid string) (int32, error) {
	var count int64
	err := r.data.db.WithContext(ctx).Model(&Collections{}).Where("uuid = ? and auth = ?", uuid, 1).Count(&count).Error
	if err != nil {
		return 0, errors.Wrapf(err, fmt.Sprintf("fail to get collections visitor count from db: uuid(%s)", uuid))
	}
	return int32(count), nil
}

func (r *creationRepo) CreateCollections(ctx context.Context, uuid, name, introduce string, auth int32) error {
	collect := &Collections{
		Uuid:      uuid,
		Name:      name,
		Introduce: introduce,
		Auth:      auth,
	}
	err := r.data.DB(ctx).Create(collect).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to create an collection from db: uuid(%v)", uuid))
	}
	return nil
}

func (r *creationRepo) EditCollections(ctx context.Context, id int32, uuid, name, introduce string, auth int32) error {
	collect := &Collections{
		Name:      name,
		Introduce: introduce,
		Auth:      auth,
	}
	err := r.data.db.WithContext(ctx).Model(&Collections{}).Where("id = ? and uuid = ?", id, uuid).Updates(collect).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to edit an collection from db: id(%v), uuid(%v)", id, uuid))
	}
	return nil
}

func (r *creationRepo) DeleteCollections(ctx context.Context, id int32, uuid string) error {
	collect := &Collections{}
	err := r.data.DB(ctx).Where("id = ? and uuid = ?", id, uuid).Delete(collect).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to delete an collection from db: id(%v), uuid(%v)", id, uuid))
	}
	return nil
}

func (r *creationRepo) SetRecord(ctx context.Context, id, mode int32, uuid, operation, ip string) error {
	record := &Record{
		CreationId: id,
		Mode:       mode,
		Uuid:       uuid,
		Operation:  operation,
		Ip:         ip,
	}
	err := r.data.DB(ctx).Create(record).Error
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("fail to add an record to db: id(%v), uuid(%s), ip(%s)", id, uuid, ip))
	}
	return nil
}

func (r *creationRepo) SetLeaderBoardToCache(ctx context.Context, boardList []*biz.LeaderBoard) {
	_, err := r.data.redisCli.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		z := make([]*redis.Z, 0)
		for _, item := range boardList {
			z = append(z, &redis.Z{
				Score:  float64(item.Agree),
				Member: strconv.Itoa(int(item.Id)) + "%" + item.Uuid + "%" + item.Mode,
			})
		}
		pipe.ZAddNX(context.Background(), "leaderboard", z...)
		return nil
	})
	if err != nil {
		r.log.Errorf("fail to set LeaderBoard to cache: LeaderBoard(%v)", boardList)
	}
}

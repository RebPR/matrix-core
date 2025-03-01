package biz

import "C"
import (
	"context"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/the-zion/matrix-core/api/creation/service/v1"
	"golang.org/x/sync/errgroup"
)

type ColumnRepo interface {
	CreateColumnDraft(ctx context.Context, uuid string) (int32, error)
	CreateColumnFolder(ctx context.Context, id int32, uuid string) error
	CreateColumn(ctx context.Context, id, auth int32, uuid string) error
	CreateColumnStatistic(ctx context.Context, id, auth int32, uuid string) error
	CreateColumnCache(ctx context.Context, id, auth int32, uuid, mode string) error
	CreateColumnSearch(ctx context.Context, id int32, uuid string) error

	AddColumnIncludes(ctx context.Context, id, articleId int32, uuid string) error
	AddColumnIncludesToCache(ctx context.Context, id, articleId int32, uuid string) error
	AddCreationUserColumn(ctx context.Context, uuid string, auth int32) error
	AddUserCreationSubscribe(ctx context.Context, uuid string) error
	ReduceCreationUserColumn(ctx context.Context, auth int32, uuid string) error

	UpdateColumnCache(ctx context.Context, id, auth int32, uuid, mode string) error
	EditColumnCos(ctx context.Context, id int32, uuid string) error
	EditColumnSearch(ctx context.Context, id int32, uuid string) error

	DeleteColumnDraft(ctx context.Context, id int32, uuid string) error
	DeleteColumnStatistic(ctx context.Context, id int32, uuid string) error
	DeleteColumn(ctx context.Context, id int32, uuid string) error
	DeleteColumnCache(ctx context.Context, id, auth int32, uuid string) error
	DeleteColumnSearch(ctx context.Context, id int32, uuid string) error
	DeleteColumnIncludes(ctx context.Context, id, articleId int32, uuid string) error
	DeleteColumnIncludesFromCache(ctx context.Context, id, articleId int32, uuid string) error

	GetColumn(ctx context.Context, id int32) (*Column, error)
	GetLastColumnDraft(ctx context.Context, uuid string) (*ColumnDraft, error)
	GetColumnList(ctx context.Context, page int32) ([]*Column, error)
	GetColumnListHot(ctx context.Context, page int32) ([]*ColumnStatistic, error)
	GetColumnHotFromDB(ctx context.Context, page int32) ([]*ColumnStatistic, error)
	GetUserColumnList(ctx context.Context, page int32, uuid string) ([]*Column, error)
	GetUserColumnListVisitor(ctx context.Context, page int32, uuid string) ([]*Column, error)
	GetColumnCount(ctx context.Context, uuid string) (int32, error)
	GetColumnCountVisitor(ctx context.Context, uuid string) (int32, error)
	GetColumnStatistic(ctx context.Context, id int32, uuid string) (*ColumnStatistic, error)
	GetColumnListStatistic(ctx context.Context, ids []int32) ([]*ColumnStatistic, error)
	GetColumnAgreeJudge(ctx context.Context, id int32, uuid string) (bool, error)
	GetColumnCollectJudge(ctx context.Context, id int32, uuid string) (bool, error)
	GetSubscribeList(ctx context.Context, page int32, uuid string) ([]*Subscribe, error)
	GetSubscribeListCount(ctx context.Context, uuid string) (int32, error)
	GetColumnSubscribes(ctx context.Context, uuid string, ids []int32) ([]*Subscribe, error)
	GetColumnSearch(ctx context.Context, page int32, search, time string) ([]*ColumnSearch, int32, error)
	GetColumnAuth(ctx context.Context, id int32) (int32, error)
	GetCollectionsIdFromColumnCollect(ctx context.Context, id int32) (int32, error)
	GetUserColumnAgree(ctx context.Context, uuid string) (map[int32]bool, error)
	GetUserColumnCollect(ctx context.Context, uuid string) (map[int32]bool, error)
	GetUserSubscribeColumn(ctx context.Context, uuid string) (map[int32]bool, error)
	GetAuthorFromSubscribe(ctx context.Context, id int32) (string, error)
	GetColumnImageReview(ctx context.Context, page int32, uuid string) ([]*ImageReview, error)
	GetColumnContentReview(ctx context.Context, page int32, uuid string) ([]*TextReview, error)

	SendColumn(ctx context.Context, id int32, uuid string) (*ColumnDraft, error)
	SendColumnToMq(ctx context.Context, column *Column, mode string) error
	SendReviewToMq(ctx context.Context, review *ColumnReview) error
	SendScoreToMq(ctx context.Context, score int32, uuid, mode string) error
	SendStatisticToMq(ctx context.Context, id, collectionsId int32, uuid, userUuid, mode string) error
	SendColumnIncludesToMq(ctx context.Context, id, articleId int32, uuid, mode string) error
	SendColumnSubscribeToMq(ctx context.Context, id int32, uuid, mode string) error
	SendColumnImageIrregularToMq(ctx context.Context, review *ImageReview) error
	SendColumnContentIrregularToMq(ctx context.Context, review *TextReview) error

	FreezeColumnCos(ctx context.Context, id int32, uuid string) error

	SetColumnAgree(ctx context.Context, id int32, uuid string) error
	SetUserColumnAgree(ctx context.Context, id int32, userUuid string) error
	SetColumnAgreeToCache(ctx context.Context, id int32, uuid, userUuid string) error
	SendColumnStatisticToMq(ctx context.Context, uuid, userUuid, mode string) error
	SetColumnView(ctx context.Context, id int32, uuid string) error
	SetColumnViewToCache(ctx context.Context, id int32, uuid string) error
	SetColumnUserCollect(ctx context.Context, id, collectionsId int32, userUuid string) error
	SetColumnCollect(ctx context.Context, id int32, uuid string) error
	SetColumnCollectToCache(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error
	SetUserColumnAgreeToCache(ctx context.Context, id int32, userUuid string) error
	SetUserColumnCollectToCache(ctx context.Context, id int32, userUuid string) error
	SetCollectionsColumnCollect(ctx context.Context, id, collectionsId int32, userUuid string) error
	SetCollectionColumn(ctx context.Context, collectionsId int32, userUuid string) error
	SetUserColumnCollect(ctx context.Context, id int32, userUuid string) error
	SetCreationUserCollect(ctx context.Context, userUuid string) error
	SetUserColumnSubscribeToCache(ctx context.Context, id int32, uuid string) error
	SetColumnSubscribeToCache(ctx context.Context, id int32, author, uuid string) error
	SetColumnImageIrregular(ctx context.Context, review *ImageReview) (*ImageReview, error)
	SetColumnImageIrregularToCache(ctx context.Context, review *ImageReview) error
	SetColumnContentIrregular(ctx context.Context, review *TextReview) (*TextReview, error)
	SetColumnContentIrregularToCache(ctx context.Context, review *TextReview) error

	CancelColumnAgree(ctx context.Context, id int32, uuid string) error
	CancelColumnAgreeFromCache(ctx context.Context, id int32, uuid, userUuid string) error
	CancelColumnCollect(ctx context.Context, id int32, uuid string) error
	CancelColumnCollectFromCache(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error
	CancelColumnUserCollect(ctx context.Context, id int32, userUuid string) error
	CancelSubscribeColumn(ctx context.Context, id int32, uuid string) error
	CancelUserColumnAgreeFromCache(ctx context.Context, id int32, userUuid string) error
	CancelUserColumnAgree(ctx context.Context, id int32, userUuid string) error
	CancelUserColumnCollectFromCache(ctx context.Context, id int32, userUuid string) error
	CancelCollectionsColumnCollect(ctx context.Context, id int32, userUuid string) error
	CancelUserColumnCollect(ctx context.Context, id int32, userUuid string) error
	CancelCollectionColumn(ctx context.Context, collectionsId int32, userUuid string) error
	CancelUserColumnSubscribeFromCache(ctx context.Context, id int32, uuid string) error
	CancelColumnSubscribeFromCache(ctx context.Context, id int32, author, uuid string) error

	ReduceCreationUserCollect(ctx context.Context, userUuid string) error
	ReduceUserCreationSubscribe(ctx context.Context, uuid string) error

	SubscribeColumn(ctx context.Context, id int32, author, uuid string) error
	SubscribeJudge(ctx context.Context, id int32, uuid string) (bool, error)
}

type ColumnUseCase struct {
	repo         ColumnRepo
	creationRepo CreationRepo
	tm           Transaction
	re           Recovery
	log          *log.Helper
}

func NewColumnUseCase(repo ColumnRepo, re Recovery, creationRepo CreationRepo, tm Transaction, logger log.Logger) *ColumnUseCase {
	return &ColumnUseCase{
		repo:         repo,
		creationRepo: creationRepo,
		tm:           tm,
		re:           re,
		log:          log.NewHelper(log.With(logger, "module", "creation/biz/columnUseCase")),
	}
}

func (r *ColumnUseCase) GetLastColumnDraft(ctx context.Context, uuid string) (*ColumnDraft, error) {
	draft, err := r.repo.GetLastColumnDraft(ctx, uuid)
	if kerrors.IsNotFound(err) {
		return nil, v1.ErrorRecordNotFound("last draft not found: %s", err.Error())
	}
	if err != nil {
		return nil, v1.ErrorGetColumnDraftFailed("get last draft failed: %s", err.Error())
	}
	return draft, nil
}

func (r *ColumnUseCase) ColumnImageIrregular(ctx context.Context, review *ImageReview) error {
	err := r.repo.SendColumnImageIrregularToMq(ctx, review)
	if err != nil {
		return v1.ErrorSetImageIrregularFailed("set column image irregular to mq failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) ColumnContentIrregular(ctx context.Context, review *TextReview) error {
	err := r.repo.SendColumnContentIrregularToMq(ctx, review)
	if err != nil {
		return v1.ErrorSetContentIrregularFailed("set column content irregular to mq failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) AddColumnImageReviewDbAndCache(ctx context.Context, review *ImageReview) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		review, err := r.repo.SetColumnImageIrregular(ctx, review)
		if err != nil {
			return v1.ErrorSetImageIrregularFailed("set column image irregular failed: %s", err.Error())
		}

		err = r.repo.SetColumnImageIrregularToCache(ctx, review)
		if err != nil {
			return v1.ErrorSetImageIrregularFailed("set column image irregular to cache failed: %s", err.Error())
		}

		return nil
	})
}

func (r *ColumnUseCase) AddColumnContentReviewDbAndCache(ctx context.Context, review *TextReview) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		review, err := r.repo.SetColumnContentIrregular(ctx, review)
		if err != nil {
			return v1.ErrorSetContentIrregularFailed("set column content irregular failed: %s", err.Error())
		}

		err = r.repo.SetColumnContentIrregularToCache(ctx, review)
		if err != nil {
			return v1.ErrorSetContentIrregularFailed("set column content irregular to cache failed: %s", err.Error())
		}

		return nil
	})
}

func (r *ColumnUseCase) GetColumnSearch(ctx context.Context, page int32, search, time string) ([]*ColumnSearch, int32, error) {
	columnList, total, err := r.repo.GetColumnSearch(ctx, page, search, time)
	if err != nil {
		return nil, 0, v1.ErrorGetColumnSearchFailed("get column search failed: %s", err.Error())
	}
	return columnList, total, nil
}

func (r *ColumnUseCase) GetUserColumnAgree(ctx context.Context, uuid string) (map[int32]bool, error) {
	agreeMap, err := r.repo.GetUserColumnAgree(ctx, uuid)
	if err != nil {
		return nil, v1.ErrorGetColumnAgreeFailed("get user column agree failed: %s", err.Error())
	}
	return agreeMap, nil
}

func (r *ColumnUseCase) GetUserColumnCollect(ctx context.Context, uuid string) (map[int32]bool, error) {
	agreeMap, err := r.repo.GetUserColumnCollect(ctx, uuid)
	if err != nil {
		return nil, v1.ErrorGetColumnCollectFailed("get user column collect failed: %s", err.Error())
	}
	return agreeMap, nil
}

func (r *ColumnUseCase) GetUserSubscribeColumn(ctx context.Context, uuid string) (map[int32]bool, error) {
	agreeMap, err := r.repo.GetUserSubscribeColumn(ctx, uuid)
	if err != nil {
		return nil, v1.ErrorGetSubscribeColumnFailed("get user column subscribe failed: %s", err.Error())
	}
	return agreeMap, nil
}

func (r *ColumnUseCase) GetColumnImageReview(ctx context.Context, page int32, uuid string) ([]*ImageReview, error) {
	reviewList, err := r.repo.GetColumnImageReview(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetImageReviewFailed("get column image review failed: %s", err.Error())
	}
	return reviewList, nil
}

func (r *ColumnUseCase) GetColumnContentReview(ctx context.Context, page int32, uuid string) ([]*TextReview, error) {
	reviewList, err := r.repo.GetColumnContentReview(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetContentReviewFailed("get column content review failed: %s", err.Error())
	}
	return reviewList, nil
}

func (r *ColumnUseCase) CreateColumnDraft(ctx context.Context, uuid string) (int32, error) {
	var id int32
	err := r.tm.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		id, err = r.repo.CreateColumnDraft(ctx, uuid)
		if err != nil {
			return v1.ErrorCreateDraftFailed("create column draft failed: %s", err.Error())
		}

		err = r.repo.CreateColumnFolder(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateDraftFailed("create column folder failed: %s", err.Error())
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ColumnUseCase) SubscribeColumn(ctx context.Context, id int32, uuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SetUserColumnSubscribeToCache(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("set user column subscribe to cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendColumnSubscribeToMq(ctx, id, uuid, "set_column_subscribe_db_and_cache")
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("set column subscribe to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) SetColumnSubscribeDbAndCache(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		author, err := r.repo.GetAuthorFromSubscribe(ctx, id)
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("get column author failed: %s", err.Error())
		}
		if author == uuid {
			return v1.ErrorSubscribeColumnFailed("author and uuid are the same")
		}
		err = r.repo.SubscribeColumn(ctx, id, author, uuid)
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("subscribe column failed: %s", err.Error())
		}
		err = r.repo.AddUserCreationSubscribe(ctx, uuid)
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("subscribe column failed: %s", err.Error())
		}
		err = r.repo.SetColumnSubscribeToCache(ctx, id, author, uuid)
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("set user column subscribe to cache failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) CancelSubscribeColumn(ctx context.Context, id int32, uuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.CancelUserColumnSubscribeFromCache(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCancelSubscribeColumnFailed("cancel user column subscribe from cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendColumnSubscribeToMq(ctx, id, uuid, "cancel_column_subscribe_db_and_cache")
		if err != nil {
			return v1.ErrorCancelSubscribeColumnFailed("cancel column subscribe to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) CancelColumnSubscribeDbAndCache(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		author, err := r.repo.GetAuthorFromSubscribe(ctx, id)
		if err != nil {
			return v1.ErrorCancelSubscribeColumnFailed("get column author failed: %s", err.Error())
		}
		err = r.repo.CancelSubscribeColumn(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCancelSubscribeColumnFailed("cancel subscribe column failed: %s", err.Error())
		}
		err = r.repo.ReduceUserCreationSubscribe(ctx, uuid)
		if err != nil {
			return v1.ErrorSubscribeColumnFailed("reduce user creation subscribe failed: %s", err.Error())
		}
		err = r.repo.CancelColumnSubscribeFromCache(ctx, id, author, uuid)
		if err != nil {
			return v1.ErrorCancelSubscribeColumnFailed("cancel user column subscribe from cache failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) SubscribeJudge(ctx context.Context, id int32, uuid string) (bool, error) {
	subscribe, err := r.repo.SubscribeJudge(ctx, id, uuid)
	if err != nil {
		return false, v1.ErrorSubscribeColumnJudgeFailed("get subscribe column judge failed: %s", err.Error())
	}
	return subscribe, nil
}

func (r *ColumnUseCase) SendColumn(ctx context.Context, id int32, uuid, ip string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		draft, err := r.repo.SendColumn(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateColumnFailed("send column failed: %s", err.Error())
		}

		err = r.creationRepo.SetRecord(ctx, id, 2, uuid, "create", ip)
		if err != nil {
			return v1.ErrorSetRecordFailed("set record failed: %s", err.Error())
		}

		err = r.repo.SendReviewToMq(ctx, &ColumnReview{
			Uuid: draft.Uuid,
			Id:   draft.Id,
			Mode: "create",
		})
		if err != nil {
			return v1.ErrorCreateColumnFailed("send create review to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) SendColumnEdit(ctx context.Context, id int32, uuid, ip string) error {
	column, err := r.repo.GetColumn(ctx, id)
	if err != nil {
		return v1.ErrorGetColumnFailed("get column failed: %s", err.Error())
	}

	if column.Uuid != uuid {
		return v1.ErrorGetColumnFailed("send column edit failed: no auth")
	}

	err = r.creationRepo.SetRecord(ctx, id, 2, uuid, "edit", ip)
	if err != nil {
		return v1.ErrorSetRecordFailed("set record failed: %s", err.Error())
	}

	err = r.repo.SendReviewToMq(ctx, &ColumnReview{
		Uuid: column.Uuid,
		Id:   column.ColumnId,
		Mode: "edit",
	})
	if err != nil {
		return v1.ErrorGetColumnFailed("send edit review to mq failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) CreateColumn(ctx context.Context, id, auth int32, uuid string) error {
	err := r.repo.SendColumnToMq(ctx, &Column{
		ColumnId: id,
		Uuid:     uuid,
		Auth:     auth,
	}, "create_column_db_cache_and_search")
	if err != nil {
		return v1.ErrorCreateColumnFailed("create column to mq failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) EditColumn(ctx context.Context, id, auth int32, uuid string) error {
	err := r.repo.SendColumnToMq(ctx, &Column{
		ColumnId: id,
		Auth:     auth,
		Uuid:     uuid,
	}, "edit_column_cos_and_search")
	if err != nil {
		return v1.ErrorEditColumnFailed("edit column to mq failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) DeleteColumn(ctx context.Context, id int32, uuid string) error {
	err := r.repo.SendColumnToMq(ctx, &Column{
		ColumnId: id,
		Uuid:     uuid,
	}, "delete_column_cache_and_search")
	if err != nil {
		return v1.ErrorDeleteColumnFailed("delete column to mq failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) CreateColumnDbCacheAndSearch(ctx context.Context, id, auth int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		err = r.repo.DeleteColumnDraft(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateColumnFailed("delete column draft failed: %s", err.Error())
		}

		err = r.repo.CreateColumn(ctx, id, auth, uuid)
		if err != nil {
			return v1.ErrorCreateColumnFailed("create column failed: %s", err.Error())
		}

		timelineId, err := r.creationRepo.CreateTimeLine(ctx, id, auth, 2, uuid)
		if err != nil {
			return v1.ErrorCreateTimelineFailed("create column timeline failed: %s", err.Error())
		}

		err = r.repo.CreateColumnStatistic(ctx, id, auth, uuid)
		if err != nil {
			return v1.ErrorCreateColumnFailed("create column statistic failed: %s", err.Error())
		}

		err = r.repo.AddCreationUserColumn(ctx, uuid, auth)
		if err != nil {
			return v1.ErrorCreateColumnFailed("add creation column failed: %s", err.Error())
		}

		err = r.repo.CreateColumnCache(ctx, id, auth, uuid, "create")
		if err != nil {
			return v1.ErrorCreateColumnFailed("create column cache failed: %s", err.Error())
		}

		if auth == 2 {
			return nil
		}

		err = r.creationRepo.CreateTimeLineCache(ctx, timelineId, id, 2, uuid)
		if err != nil {
			return v1.ErrorCreateArticleFailed("create column timeline cache failed: %s", err.Error())
		}

		err = r.repo.CreateColumnSearch(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateColumnFailed("create column search failed: %s", err.Error())
		}

		err = r.repo.SendScoreToMq(ctx, 20, uuid, "add_score")
		if err != nil {
			return v1.ErrorCreateColumnFailed("send 20 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) EditColumnCosAndSearch(ctx context.Context, id, auth int32, uuid string) error {
	err := r.repo.UpdateColumnCache(ctx, id, auth, uuid, "edit")
	if err != nil {
		return v1.ErrorEditColumnFailed("edit column cache failed: %s", err.Error())
	}

	err = r.repo.EditColumnCos(ctx, id, uuid)
	if err != nil {
		return v1.ErrorEditColumnFailed("edit column cache failed: %s", err.Error())
	}

	if auth == 2 {
		return nil
	}

	err = r.repo.EditColumnSearch(ctx, id, uuid)
	if err != nil {
		return v1.ErrorEditColumnFailed("edit column search failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) DeleteColumnCacheAndSearch(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		auth, err := r.repo.GetColumnAuth(ctx, id)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("get column auth failed: %s", err.Error())
		}

		timelineId, err := r.creationRepo.GetUserTimeLine(ctx, id, 2)
		if err != nil {
			return v1.ErrorDeleteArticleFailed("get user column timeline failed: %s", err.Error())
		}

		err = r.repo.DeleteColumn(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("delete column failed: %s", err.Error())
		}

		err = r.creationRepo.DeleteTimeLine(ctx, timelineId)
		if err != nil {
			return v1.ErrorDeleteArticleFailed("delete column timeline failed: %s", err.Error())
		}

		err = r.repo.DeleteColumnStatistic(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("delete column statistic failed: %s", err.Error())
		}

		err = r.repo.ReduceCreationUserColumn(ctx, auth, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("delete creation user column failed: %s", err.Error())
		}

		err = r.repo.DeleteColumnCache(ctx, id, auth, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("delete column cache failed: %s", err.Error())
		}

		err = r.repo.FreezeColumnCos(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("freeze column cos failed: %s", err.Error())
		}

		if auth == 2 {
			return nil
		}

		err = r.creationRepo.DeleteTimeLineCache(ctx, timelineId, id, 2, uuid)
		if err != nil {
			return v1.ErrorDeleteArticleFailed("delete column timeline cache failed: %s", err.Error())
		}

		err = r.repo.DeleteColumnSearch(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnFailed("delete column search failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) GetColumnList(ctx context.Context, page int32) ([]*Column, error) {
	columnList, err := r.repo.GetColumnList(ctx, page)
	if err != nil {
		return nil, v1.ErrorGetColumnListFailed("get column list failed: %s", err.Error())
	}
	return columnList, nil
}

func (r *ColumnUseCase) GetColumnListHot(ctx context.Context, page int32) ([]*ColumnStatistic, error) {
	columnList, err := r.repo.GetColumnListHot(ctx, page)
	if err != nil {
		return nil, v1.ErrorGetColumnListFailed("get column hot list failed: %s", err.Error())
	}
	return columnList, nil
}

func (r *ColumnUseCase) GetUserColumnList(ctx context.Context, page int32, uuid string) ([]*Column, error) {
	columnList, err := r.repo.GetUserColumnList(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetColumnListFailed("get user column list failed: %s", err.Error())
	}
	return columnList, nil
}

func (r *ColumnUseCase) GetUserColumnListVisitor(ctx context.Context, page int32, uuid string) ([]*Column, error) {
	columnList, err := r.repo.GetUserColumnListVisitor(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetColumnListFailed("get user column list visitor failed: %s", err.Error())
	}
	return columnList, nil
}

func (r *ColumnUseCase) GetColumnCount(ctx context.Context, uuid string) (int32, error) {
	count, err := r.repo.GetColumnCount(ctx, uuid)
	if err != nil {
		return 0, v1.ErrorGetCountFailed("get column count failed: %s", err.Error())
	}
	return count, nil
}

func (r *ColumnUseCase) GetColumnCountVisitor(ctx context.Context, uuid string) (int32, error) {
	count, err := r.repo.GetColumnCountVisitor(ctx, uuid)
	if err != nil {
		return 0, v1.ErrorGetCountFailed("get column count visitor failed: %s", err.Error())
	}
	return count, nil
}

func (r *ColumnUseCase) GetColumnListStatistic(ctx context.Context, ids []int32) ([]*ColumnStatistic, error) {
	statisticList, err := r.repo.GetColumnListStatistic(ctx, ids)
	if err != nil {
		return nil, v1.ErrorGetStatisticFailed("get column list statistic failed: %s", err.Error())
	}
	return statisticList, nil
}

func (r *ColumnUseCase) GetColumnStatistic(ctx context.Context, id int32, uuid string) (*ColumnStatistic, error) {
	statistic, err := r.repo.GetColumnStatistic(ctx, id, uuid)
	if err != nil {
		return nil, v1.ErrorGetStatisticFailed("get column statistic failed: %s", err.Error())
	}
	return statistic, nil
}

func (r *ColumnUseCase) ColumnStatisticJudge(ctx context.Context, id int32, uuid string) (*ColumnStatisticJudge, error) {
	agree, err := r.repo.GetColumnAgreeJudge(ctx, id, uuid)
	if err != nil {
		return nil, v1.ErrorGetStatisticJudgeFailed("get column statistic judge failed: %s", err.Error())
	}
	collect, err := r.repo.GetColumnCollectJudge(ctx, id, uuid)
	if err != nil {
		return nil, v1.ErrorGetStatisticJudgeFailed("get column statistic judge failed: %s", err.Error())
	}
	return &ColumnStatisticJudge{
		Agree:   agree,
		Collect: collect,
	}, nil
}

func (r *ColumnUseCase) GetSubscribeList(ctx context.Context, page int32, uuid string) ([]*Subscribe, error) {
	subscribe, err := r.repo.GetSubscribeList(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetSubscribeColumnListFailed("get subscribe column list failed: %s", err.Error())
	}
	return subscribe, nil
}

func (r *ColumnUseCase) GetSubscribeListCount(ctx context.Context, uuid string) (int32, error) {
	count, err := r.repo.GetSubscribeListCount(ctx, uuid)
	if err != nil {
		return 0, v1.ErrorGetCountFailed("get subscribe column list count failed: %s", err.Error())
	}
	return count, nil
}

func (r *ColumnUseCase) GetColumnSubscribes(ctx context.Context, uuid string, ids []int32) ([]*Subscribe, error) {
	subscribe, err := r.repo.GetColumnSubscribes(ctx, uuid, ids)
	if err != nil {
		return nil, v1.ErrorGetColumnSubscribesFailed("get column subscribes failed: %s", err.Error())
	}
	return subscribe, nil
}

func (r *ColumnUseCase) SetColumnAgree(ctx context.Context, id int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SetUserColumnAgreeToCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set user column agree to cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, userUuid, "set_column_agree_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("set column agree to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) SetColumnAgreeDbAndCache(ctx context.Context, id int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.SetColumnAgree(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set column agree failed: %s", err.Error())
		}
		err = r.repo.SetUserColumnAgree(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set column agree failed: %s", err.Error())
		}
		err = r.repo.SetColumnAgreeToCache(ctx, id, uuid, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set column agree to cache failed: %s", err.Error())
		}
		err = r.repo.SendColumnStatisticToMq(ctx, uuid, userUuid, "agree")
		if err != nil {
			return v1.ErrorSetAgreeFailed("set column agree to mq failed: %s", err.Error())
		}
		err = r.repo.SendScoreToMq(ctx, 2, uuid, "add_score")
		if err != nil {
			return v1.ErrorSetAgreeFailed("send 2 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) SetColumnView(ctx context.Context, id int32, uuid string) error {
	err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, "", "set_column_view_db_and_cache")
	if err != nil {
		return v1.ErrorSetViewFailed("set column view failed: %s", err.Error())
	}
	return nil
}

func (r *ColumnUseCase) SetColumnViewDbAndCache(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.SetColumnView(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetViewFailed("set column view failed: %s", err.Error())
		}
		err = r.repo.SetColumnViewToCache(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetViewFailed("set column view to cache failed: %s", err.Error())
		}
		err = r.repo.SendColumnStatisticToMq(ctx, uuid, "", "view")
		if err != nil {
			return v1.ErrorSetViewFailed("set column view to mq failed: %s", err.Error())
		}
		err = r.repo.SendScoreToMq(ctx, 1, uuid, "add_score")
		if err != nil {
			return v1.ErrorSetViewFailed("send 1 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) CancelColumnAgree(ctx context.Context, id int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.CancelUserColumnAgreeFromCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel user column agree from cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, userUuid, "cancel_column_agree_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel column agree to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) CancelColumnAgreeDbAndCache(ctx context.Context, id int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.CancelColumnAgree(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel column agree failed: %s", err.Error())
		}
		err = r.repo.CancelUserColumnAgree(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel column agree failed: %s", err.Error())
		}
		err = r.repo.CancelColumnAgreeFromCache(ctx, id, uuid, userUuid)
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel column agree from cache failed: %s", err.Error())
		}
		err = r.repo.SendColumnStatisticToMq(ctx, uuid, userUuid, "agree_cancel")
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel column agree to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) SetColumnCollect(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SetUserColumnCollectToCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set user column collect to cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, collectionsId, uuid, userUuid, "set_column_collect_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("set column collect to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) SetColumnCollectDbAndCache(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.SetCollectionsColumnCollect(ctx, id, collectionsId, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set column user collect failed: %s", err.Error())
		}
		err = r.repo.SetCollectionColumn(ctx, collectionsId, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set column collection collect failed: %s", err.Error())
		}
		err = r.repo.SetUserColumnCollect(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set user column collect failed: %s", err.Error())
		}
		err = r.repo.SetCreationUserCollect(ctx, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set creation user collect failed: %s", err.Error())
		}
		err = r.repo.SetColumnCollect(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set column collect failed: %s", err.Error())
		}
		err = r.repo.SetColumnCollectToCache(ctx, id, collectionsId, uuid, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set column collect to cache failed: %s", err.Error())
		}
		err = r.repo.SendColumnStatisticToMq(ctx, uuid, "", "collect")
		if err != nil {
			return v1.ErrorSetCollectFailed("set column collect to mq failed: %s", err.Error())
		}
		err = r.repo.SendScoreToMq(ctx, 2, uuid, "add_score")
		if err != nil {
			return v1.ErrorSetViewFailed("send 1 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) CancelColumnCollect(ctx context.Context, id int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.CancelUserColumnCollectFromCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel user column collect from cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, userUuid, "cancel_column_collect_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel column collect to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) CancelColumnCollectDbAndCache(ctx context.Context, id int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		collectionsId, err := r.repo.GetCollectionsIdFromColumnCollect(ctx, id)
		if err != nil {
			return v1.ErrorSetCollectFailed("get collections Id from collect failed: %s", err.Error())
		}
		err = r.repo.CancelCollectionsColumnCollect(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel column user collect failed: %s", err.Error())
		}
		err = r.repo.CancelUserColumnCollect(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel user column collect failed: %s", err.Error())
		}
		err = r.repo.CancelCollectionColumn(ctx, collectionsId, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel column collection collect failed: %s", err.Error())
		}
		err = r.repo.ReduceCreationUserCollect(ctx, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("reduce creation user collect failed: %s", err.Error())
		}
		err = r.repo.CancelColumnCollect(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel column collect failed: %s", err.Error())
		}
		err = r.repo.CancelColumnCollectFromCache(ctx, id, collectionsId, uuid, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel column collect from cache failed: %s", err.Error())
		}
		err = r.repo.SendColumnStatisticToMq(ctx, uuid, "", "collect_cancel")
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel column collect to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) AddColumnIncludes(ctx context.Context, id, articleId int32, uuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.AddColumnIncludesToCache(ctx, id, articleId, uuid)
		if err != nil {
			return v1.ErrorAddColumnIncludesFailed("add column includes failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendColumnIncludesToMq(ctx, id, articleId, uuid, "add_column_includes_db_and_cache")
		if err != nil {
			return v1.ErrorAddColumnIncludesFailed("add column includes to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) AddColumnIncludesDbAndCache(ctx context.Context, id, articleId int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.AddColumnIncludes(ctx, id, articleId, uuid)
		if err != nil {
			return v1.ErrorAddColumnIncludesFailed("add column includes failed: %s", err.Error())
		}

		err = r.repo.AddColumnIncludesToCache(ctx, id, articleId, uuid)
		if err != nil {
			return v1.ErrorAddColumnIncludesFailed("add column includes to cache failed: %s", err.Error())
		}
		return nil
	})
}

func (r *ColumnUseCase) DeleteColumnIncludes(ctx context.Context, id, articleId int32, uuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.DeleteColumnIncludesFromCache(ctx, id, articleId, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnIncludesFailed("delete column includes from cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendColumnIncludesToMq(ctx, id, articleId, uuid, "delete_column_includes_db_and_cache")
		if err != nil {
			return v1.ErrorDeleteColumnIncludesFailed("delete column includes to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *ColumnUseCase) DeleteColumnIncludesDbAndCache(ctx context.Context, id, articleId int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.DeleteColumnIncludes(ctx, id, articleId, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnIncludesFailed("delete column includes failed: %s", err.Error())
		}

		err = r.repo.DeleteColumnIncludesFromCache(ctx, id, articleId, uuid)
		if err != nil {
			return v1.ErrorDeleteColumnIncludesFailed("delete column includes from cache failed: %s", err.Error())
		}
		return nil
	})
}

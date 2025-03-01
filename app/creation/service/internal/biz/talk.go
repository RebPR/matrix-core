package biz

import (
	"context"
	kerrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/the-zion/matrix-core/api/creation/service/v1"
	"golang.org/x/sync/errgroup"
)

type TalkRepo interface {
	GetTalk(ctx context.Context, id int32) (*Talk, error)
	GetTalkList(ctx context.Context, page int32) ([]*Talk, error)
	GetTalkListHot(ctx context.Context, page int32) ([]*TalkStatistic, error)
	GetTalkHotFromDB(ctx context.Context, page int32) ([]*TalkStatistic, error)
	GetUserTalkList(ctx context.Context, page int32, uuid string) ([]*Talk, error)
	GetUserTalkListVisitor(ctx context.Context, page int32, uuid string) ([]*Talk, error)
	GetTalkCount(ctx context.Context, uuid string) (int32, error)
	GetTalkCountVisitor(ctx context.Context, uuid string) (int32, error)
	GetTalkListStatistic(ctx context.Context, ids []int32) ([]*TalkStatistic, error)
	GetTalkStatistic(ctx context.Context, id int32, uuid string) (*TalkStatistic, error)
	GetLastTalkDraft(ctx context.Context, uuid string) (*TalkDraft, error)
	GetTalkAgreeJudge(ctx context.Context, id int32, uuid string) (bool, error)
	GetTalkCollectJudge(ctx context.Context, id int32, uuid string) (bool, error)
	GetTalkSearch(ctx context.Context, page int32, search, time string) ([]*TalkSearch, int32, error)
	GetUserTalkAgree(ctx context.Context, uuid string) (map[int32]bool, error)
	GetUserTalkCollect(ctx context.Context, uuid string) (map[int32]bool, error)
	GetTalkAuth(ctx context.Context, id int32) (int32, error)
	GetCollectionsIdFromTalkCollect(ctx context.Context, id int32) (int32, error)
	GetTalkImageReview(ctx context.Context, page int32, uuid string) ([]*ImageReview, error)
	GetTalkContentReview(ctx context.Context, page int32, uuid string) ([]*TextReview, error)

	CreateTalkDraft(ctx context.Context, uuid string) (int32, error)
	CreateTalkFolder(ctx context.Context, id int32, uuid string) error
	CreateTalk(ctx context.Context, id, auth int32, uuid string) error
	CreateTalkStatistic(ctx context.Context, id, auth int32, uuid string) error
	CreateTalkCache(ctx context.Context, id, auth int32, uuid, mode string) error
	CreateTalkSearch(ctx context.Context, id int32, uuid string) error
	AddCreationUserTalk(ctx context.Context, uuid string, auth int32) error
	AddTalkComment(ctx context.Context, id int32) error
	AddTalkCommentToCache(ctx context.Context, id int32, uuid string) error
	ReduceTalkComment(ctx context.Context, id int32) error
	ReduceTalkCommentToCache(ctx context.Context, id int32, uuid string) error
	ReduceCreationUserTalk(ctx context.Context, auth int32, uuid string) error

	SetTalkAgree(ctx context.Context, id int32, uuid string) error
	SetUserTalkAgree(ctx context.Context, id int32, userUuid string) error
	SetTalkView(ctx context.Context, id int32, uuid string) error
	SetTalkViewToCache(ctx context.Context, id int32, uuid string) error
	SetTalkAgreeToCache(ctx context.Context, id int32, uuid, userUuid string) error
	SetTalkCollect(ctx context.Context, id int32, uuid string) error
	SetTalkUserCollect(ctx context.Context, id, collectionsId int32, userUuid string) error
	SetTalkCollectToCache(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error
	SetUserTalkAgreeToCache(ctx context.Context, id int32, userUuid string) error
	SetUserTalkCollectToCache(ctx context.Context, id int32, userUuid string) error
	SetCollectionsTalkCollect(ctx context.Context, id, collectionsId int32, userUuid string) error
	SetCollectionTalk(ctx context.Context, collectionsId int32, userUuid string) error
	SetUserTalkCollect(ctx context.Context, id int32, userUuid string) error
	SetCreationUserCollect(ctx context.Context, userUuid string) error
	SetTalkImageIrregular(ctx context.Context, review *ImageReview) (*ImageReview, error)
	SetTalkImageIrregularToCache(ctx context.Context, review *ImageReview) error
	SetTalkContentIrregular(ctx context.Context, review *TextReview) (*TextReview, error)
	SetTalkContentIrregularToCache(ctx context.Context, review *TextReview) error

	CancelTalkAgree(ctx context.Context, id int32, uuid string) error
	CancelUserTalkAgree(ctx context.Context, id int32, userUuid string) error
	CancelTalkAgreeFromCache(ctx context.Context, id int32, uuid, userUuid string) error
	CancelTalkUserCollect(ctx context.Context, id int32, userUuid string) error
	CancelTalkCollect(ctx context.Context, id int32, uuid string) error
	CancelTalkCollectFromCache(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error
	CancelUserTalkAgreeFromCache(ctx context.Context, id int32, userUuid string) error
	CancelUserTalkCollectFromCache(ctx context.Context, id int32, userUuid string) error
	CancelCollectionsTalkCollect(ctx context.Context, id int32, userUuid string) error
	CancelUserTalkCollect(ctx context.Context, id int32, userUuid string) error
	CancelCollectionTalk(ctx context.Context, collectionsId int32, userUuid string) error
	ReduceCreationUserCollect(ctx context.Context, userUuid string) error

	SendTalk(ctx context.Context, id int32, uuid string) (*TalkDraft, error)
	SendReviewToMq(ctx context.Context, review *TalkReview) error
	SendTalkToMq(ctx context.Context, talk *Talk, mode string) error
	SendTalkStatisticToMq(ctx context.Context, uuid, userUuid, mode string) error
	SendScoreToMq(ctx context.Context, score int32, uuid, mode string) error
	SendStatisticToMq(ctx context.Context, id, collectionsId int32, uuid, userUuid, mode string) error
	SendTalkImageIrregularToMq(ctx context.Context, review *ImageReview) error
	SendTalkContentIrregularToMq(ctx context.Context, review *TextReview) error

	DeleteTalk(ctx context.Context, id int32, uuid string) error
	DeleteTalkDraft(ctx context.Context, id int32, uuid string) error
	DeleteTalkStatistic(ctx context.Context, id int32, uuid string) error
	DeleteTalkCache(ctx context.Context, id, auth int32, uuid string) error
	DeleteTalkSearch(ctx context.Context, id int32, uuid string) error

	FreezeTalkCos(ctx context.Context, id int32, uuid string) error

	UpdateTalkCache(ctx context.Context, id, auth int32, uuid, mode string) error
	EditTalkCos(ctx context.Context, id int32, uuid string) error
	EditTalkSearch(ctx context.Context, id int32, uuid string) error
}

type TalkUseCase struct {
	repo         TalkRepo
	creationRepo CreationRepo
	tm           Transaction
	re           Recovery
	log          *log.Helper
}

func NewTalkUseCase(repo TalkRepo, re Recovery, creationRepo CreationRepo, tm Transaction, logger log.Logger) *TalkUseCase {
	return &TalkUseCase{
		repo:         repo,
		creationRepo: creationRepo,
		tm:           tm,
		re:           re,
		log:          log.NewHelper(log.With(logger, "module", "creation/biz/talkUseCase")),
	}
}

func (r *TalkUseCase) GetLastTalkDraft(ctx context.Context, uuid string) (*TalkDraft, error) {
	draft, err := r.repo.GetLastTalkDraft(ctx, uuid)
	if kerrors.IsNotFound(err) {
		return nil, v1.ErrorRecordNotFound("last draft not found: %s", err.Error())
	}
	if err != nil {
		return nil, v1.ErrorGetTalkDraftFailed("get last draft failed: %s", err.Error())
	}
	return draft, nil
}

func (r *TalkUseCase) GetTalkList(ctx context.Context, page int32) ([]*Talk, error) {
	talkList, err := r.repo.GetTalkList(ctx, page)
	if err != nil {
		return nil, v1.ErrorGetTalkListFailed("get talk list failed: %s", err.Error())
	}
	return talkList, nil
}

func (r *TalkUseCase) GetTalkListHot(ctx context.Context, page int32) ([]*TalkStatistic, error) {
	talkList, err := r.repo.GetTalkListHot(ctx, page)
	if err != nil {
		return nil, v1.ErrorGetTalkListFailed("get talk hot list failed: %s", err.Error())
	}
	return talkList, nil
}

func (r *TalkUseCase) GetUserTalkList(ctx context.Context, page int32, uuid string) ([]*Talk, error) {
	talkList, err := r.repo.GetUserTalkList(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetTalkListFailed("get user talk list failed: %s", err.Error())
	}
	return talkList, nil
}

func (r *TalkUseCase) GetUserTalkListVisitor(ctx context.Context, page int32, uuid string) ([]*Talk, error) {
	talkList, err := r.repo.GetUserTalkListVisitor(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetTalkListFailed("get user talk list visitor failed: %s", err.Error())
	}
	return talkList, nil
}

func (r *TalkUseCase) GetTalkCount(ctx context.Context, uuid string) (int32, error) {
	count, err := r.repo.GetTalkCount(ctx, uuid)
	if err != nil {
		return 0, v1.ErrorGetCountFailed("get talk count failed: %s", err.Error())
	}
	return count, nil
}

func (r *TalkUseCase) GetTalkCountVisitor(ctx context.Context, uuid string) (int32, error) {
	count, err := r.repo.GetTalkCountVisitor(ctx, uuid)
	if err != nil {
		return 0, v1.ErrorGetCountFailed("get talk count visitor failed: %s", err.Error())
	}
	return count, nil
}

func (r *TalkUseCase) GetTalkListStatistic(ctx context.Context, ids []int32) ([]*TalkStatistic, error) {
	statisticList, err := r.repo.GetTalkListStatistic(ctx, ids)
	if err != nil {
		return nil, v1.ErrorGetStatisticFailed("get talk list statistic failed: %s", err.Error())
	}
	return statisticList, nil
}

func (r *TalkUseCase) GetTalkStatistic(ctx context.Context, id int32, uuid string) (*TalkStatistic, error) {
	statistic, err := r.repo.GetTalkStatistic(ctx, id, uuid)
	if err != nil {
		return nil, v1.ErrorGetStatisticFailed("get talk statistic failed: %s", err.Error())
	}
	return statistic, nil
}

func (r *TalkUseCase) GetTalkSearch(ctx context.Context, page int32, search, time string) ([]*TalkSearch, int32, error) {
	talkList, total, err := r.repo.GetTalkSearch(ctx, page, search, time)
	if err != nil {
		return nil, 0, v1.ErrorGetTalkSearchFailed("get talk search failed: %s", err.Error())
	}
	return talkList, total, nil
}

func (r *TalkUseCase) GetUserTalkAgree(ctx context.Context, uuid string) (map[int32]bool, error) {
	agreeMap, err := r.repo.GetUserTalkAgree(ctx, uuid)
	if err != nil {
		return nil, v1.ErrorGetTalkAgreeFailed("get user talk agree failed: %s", err.Error())
	}
	return agreeMap, nil
}

func (r *TalkUseCase) GetUserTalkCollect(ctx context.Context, uuid string) (map[int32]bool, error) {
	collectMap, err := r.repo.GetUserTalkCollect(ctx, uuid)
	if err != nil {
		return nil, v1.ErrorGetTalkAgreeFailed("get user talk collect failed: %s", err.Error())
	}
	return collectMap, nil
}

func (r *TalkUseCase) GetTalkImageReview(ctx context.Context, page int32, uuid string) ([]*ImageReview, error) {
	reviewList, err := r.repo.GetTalkImageReview(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetImageReviewFailed("get talk image review failed: %s", err.Error())
	}
	return reviewList, nil
}

func (r *TalkUseCase) GetTalkContentReview(ctx context.Context, page int32, uuid string) ([]*TextReview, error) {
	reviewList, err := r.repo.GetTalkContentReview(ctx, page, uuid)
	if err != nil {
		return nil, v1.ErrorGetContentReviewFailed("get talk content review failed: %s", err.Error())
	}
	return reviewList, nil
}

func (r *TalkUseCase) CreateTalkDraft(ctx context.Context, uuid string) (int32, error) {
	var id int32
	err := r.tm.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		id, err = r.repo.CreateTalkDraft(ctx, uuid)
		if err != nil {
			return v1.ErrorCreateDraftFailed("create talk draft failed: %s", err.Error())
		}

		err = r.repo.CreateTalkFolder(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateDraftFailed("create talk folder failed: %s", err.Error())
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TalkUseCase) SendTalk(ctx context.Context, id int32, uuid, ip string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		draft, err := r.repo.SendTalk(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateTalkFailed("send talk failed: %s", err.Error())
		}

		err = r.creationRepo.SetRecord(ctx, id, 3, uuid, "create", ip)
		if err != nil {
			return v1.ErrorSetRecordFailed("set record failed: %s", err.Error())
		}

		err = r.repo.SendReviewToMq(ctx, &TalkReview{
			Uuid: draft.Uuid,
			Id:   draft.Id,
			Mode: "create",
		})
		if err != nil {
			return v1.ErrorCreateTalkFailed("send create review to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) SendTalkEdit(ctx context.Context, id int32, uuid, ip string) error {
	talk, err := r.repo.GetTalk(ctx, id)
	if err != nil {
		return v1.ErrorGetTalkFailed("get talk failed: %s", err.Error())
	}

	if talk.Uuid != uuid {
		return v1.ErrorEditTalkFailed("send talk edit failed: no auth")
	}

	err = r.creationRepo.SetRecord(ctx, id, 3, uuid, "edit", ip)
	if err != nil {
		return v1.ErrorSetRecordFailed("set record failed: %s", err.Error())
	}

	err = r.repo.SendReviewToMq(ctx, &TalkReview{
		Uuid: talk.Uuid,
		Id:   talk.TalkId,
		Mode: "edit",
	})
	if err != nil {
		return v1.ErrorEditTalkFailed("send edit review to mq failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) CreateTalk(ctx context.Context, id, auth int32, uuid string) error {
	err := r.repo.SendTalkToMq(ctx, &Talk{
		TalkId: id,
		Uuid:   uuid,
		Auth:   auth,
	}, "create_talk_db_cache_and_search")
	if err != nil {
		return v1.ErrorCreateTalkFailed("create talk to mq failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) EditTalk(ctx context.Context, id, auth int32, uuid string) error {
	err := r.repo.SendTalkToMq(ctx, &Talk{
		TalkId: id,
		Auth:   auth,
		Uuid:   uuid,
	}, "edit_talk_cos_and_search")
	if err != nil {
		return v1.ErrorEditTalkFailed("edit talk to mq failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) DeleteTalk(ctx context.Context, id int32, uuid string) error {
	err := r.repo.SendTalkToMq(ctx, &Talk{
		TalkId: id,
		Uuid:   uuid,
	}, "delete_talk_cache_and_search")
	if err != nil {
		return v1.ErrorDeleteTalkFailed("delete talk to mq failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) CreateTalkDbCacheAndSearch(ctx context.Context, id, auth int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		err = r.repo.DeleteTalkDraft(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateTalkFailed("delete talk draft failed: %s", err.Error())
		}

		err = r.repo.CreateTalk(ctx, id, auth, uuid)
		if err != nil {
			return v1.ErrorCreateTalkFailed("create talk failed: %s", err.Error())
		}

		timelineId, err := r.creationRepo.CreateTimeLine(ctx, id, auth, 3, uuid)
		if err != nil {
			return v1.ErrorCreateTimelineFailed("create talk timeline failed: %s", err.Error())
		}

		err = r.repo.CreateTalkStatistic(ctx, id, auth, uuid)
		if err != nil {
			return v1.ErrorCreateTalkFailed("create talk statistic failed: %s", err.Error())
		}

		err = r.repo.AddCreationUserTalk(ctx, uuid, auth)
		if err != nil {
			return v1.ErrorCreateTalkFailed("add creation talk failed: %s", err.Error())
		}

		err = r.repo.CreateTalkCache(ctx, id, auth, uuid, "create")
		if err != nil {
			return v1.ErrorCreateTalkFailed("create talk cache failed: %s", err.Error())
		}

		if auth == 2 {
			return nil
		}

		err = r.creationRepo.CreateTimeLineCache(ctx, timelineId, id, 3, uuid)
		if err != nil {
			return v1.ErrorCreateArticleFailed("create talk timeline cache failed: %s", err.Error())
		}

		err = r.repo.CreateTalkSearch(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCreateTalkFailed("create talk search failed: %s", err.Error())
		}

		err = r.repo.SendScoreToMq(ctx, 50, uuid, "add_score")
		if err != nil {
			return v1.ErrorCreateTalkFailed("send 50 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) EditTalkCosAndSearch(ctx context.Context, id, auth int32, uuid string) error {
	err := r.repo.UpdateTalkCache(ctx, id, auth, uuid, "edit")
	if err != nil {
		return v1.ErrorEditTalkFailed("edit talk cache failed: %s", err.Error())
	}

	err = r.repo.EditTalkCos(ctx, id, uuid)
	if err != nil {
		return v1.ErrorEditTalkFailed("edit talk cache failed: %s", err.Error())
	}

	if auth == 2 {
		return nil
	}

	err = r.repo.EditTalkSearch(ctx, id, uuid)
	if err != nil {
		return v1.ErrorEditTalkFailed("edit talk search failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) DeleteTalkCacheAndSearch(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		var err error
		auth, err := r.repo.GetTalkAuth(ctx, id)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("get talk auth failed: %s", err.Error())
		}

		timelineId, err := r.creationRepo.GetUserTimeLine(ctx, id, 3)
		if err != nil {
			return v1.ErrorDeleteArticleFailed("get user talk timeline failed: %s", err.Error())
		}

		err = r.repo.DeleteTalk(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("delete talk failed: %s", err.Error())
		}

		err = r.creationRepo.DeleteTimeLine(ctx, timelineId)
		if err != nil {
			return v1.ErrorDeleteArticleFailed("delete talk timeline failed: %s", err.Error())
		}

		err = r.repo.DeleteTalkStatistic(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("delete talk statistic failed: %s", err.Error())
		}

		err = r.repo.ReduceCreationUserTalk(ctx, auth, uuid)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("delete creation user talk failed: %s", err.Error())
		}

		err = r.repo.DeleteTalkCache(ctx, id, auth, uuid)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("delete talk cache failed: %s", err.Error())
		}

		err = r.repo.FreezeTalkCos(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("freeze talk cos failed: %s", err.Error())
		}

		if auth == 2 {
			return nil
		}

		err = r.creationRepo.DeleteTimeLineCache(ctx, timelineId, id, 3, uuid)
		if err != nil {
			return v1.ErrorDeleteArticleFailed("delete talk timeline cache failed: %s", err.Error())
		}

		err = r.repo.DeleteTalkSearch(ctx, id, uuid)
		if err != nil {
			return v1.ErrorDeleteTalkFailed("delete talk search failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) SetTalkAgree(ctx context.Context, id int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SetUserTalkAgreeToCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set user talk agree to cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, userUuid, "set_talk_agree_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("set talk agree to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *TalkUseCase) SetTalkAgreeDbAndCache(ctx context.Context, id int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.SetTalkAgree(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set talk agree failed: %s", err.Error())
		}
		err = r.repo.SetUserTalkAgree(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set talk agree failed: %s", err.Error())
		}
		err = r.repo.SetTalkAgreeToCache(ctx, id, uuid, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set talk agree to cache failed: %s", err.Error())
		}
		err = r.repo.SendTalkStatisticToMq(ctx, uuid, userUuid, "agree")
		if err != nil {
			return v1.ErrorSetAgreeFailed("set talk agree to mq failed: %s", err.Error())
		}
		err = r.repo.SendScoreToMq(ctx, 2, uuid, "add_score")
		if err != nil {
			return v1.ErrorSetAgreeFailed("send 2 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) CancelTalkAgree(ctx context.Context, id int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.CancelUserTalkAgreeFromCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel user talk agree from cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, userUuid, "cancel_talk_agree_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel talk agree to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *TalkUseCase) CancelTalkAgreeDbAndCache(ctx context.Context, id int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.CancelTalkAgree(ctx, id, uuid)
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel talk agree failed: %s", err.Error())
		}
		err = r.repo.CancelUserTalkAgree(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel talk agree failed: %s", err.Error())
		}
		err = r.repo.CancelTalkAgreeFromCache(ctx, id, uuid, userUuid)
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel talk agree from cache failed: %s", err.Error())
		}
		err = r.repo.SendTalkStatisticToMq(ctx, uuid, userUuid, "agree_cancel")
		if err != nil {
			return v1.ErrorCancelAgreeFailed("cancel talk agree to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) SetTalkView(ctx context.Context, id int32, uuid string) error {
	err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, "", "set_talk_view_db_and_cache")
	if err != nil {
		return v1.ErrorSetViewFailed("set talk view failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) SetTalkViewDbAndCache(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.SetTalkView(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetViewFailed("set talk view failed: %s", err.Error())
		}
		err = r.repo.SetTalkViewToCache(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetViewFailed("set talk view to cache failed: %s", err.Error())
		}
		err = r.repo.SendTalkStatisticToMq(ctx, uuid, "", "view")
		if err != nil {
			return v1.ErrorSetViewFailed("set talk view to mq failed: %s", err.Error())
		}
		err = r.repo.SendScoreToMq(ctx, 1, uuid, "add_score")
		if err != nil {
			return v1.ErrorSetViewFailed("send 1 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) SetTalkCollect(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SetUserTalkCollectToCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("set user talk collect to cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, collectionsId, uuid, userUuid, "set_talk_collect_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("set talk collect to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *TalkUseCase) SetTalkCollectDbAndCache(ctx context.Context, id, collectionsId int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.SetCollectionsTalkCollect(ctx, id, collectionsId, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set talk user collect failed: %s", err.Error())
		}
		err = r.repo.SetCollectionTalk(ctx, collectionsId, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set talk collection collect failed: %s", err.Error())
		}
		err = r.repo.SetUserTalkCollect(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set user talk collect failed: %s", err.Error())
		}
		err = r.repo.SetCreationUserCollect(ctx, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set creation user collect failed: %s", err.Error())
		}
		err = r.repo.SetTalkCollect(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set talk collect failed: %s", err.Error())
		}
		err = r.repo.SetTalkCollectToCache(ctx, id, collectionsId, uuid, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("set talk collect to cache failed: %s", err.Error())
		}
		err = r.repo.SendTalkStatisticToMq(ctx, uuid, "", "collect")
		if err != nil {
			return v1.ErrorSetCollectFailed("set talk collect to mq failed: %s", err.Error())
		}
		err = r.repo.SendScoreToMq(ctx, 2, uuid, "add_score")
		if err != nil {
			return v1.ErrorSetViewFailed("send 1 score to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) CancelTalkCollect(ctx context.Context, id int32, uuid, userUuid string) error {
	g, _ := errgroup.WithContext(ctx)
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.CancelUserTalkCollectFromCache(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel user talk collect from cache failed: %s", err.Error())
		}
		return nil
	}))
	g.Go(r.re.GroupRecover(ctx, func(ctx context.Context) error {
		err := r.repo.SendStatisticToMq(ctx, id, 0, uuid, userUuid, "cancel_talk_collect_db_and_cache")
		if err != nil {
			return v1.ErrorSetAgreeFailed("cancel talk collect to mq failed: %s", err.Error())
		}
		return nil
	}))
	return g.Wait()
}

func (r *TalkUseCase) CancelTalkCollectDbAndCache(ctx context.Context, id int32, uuid, userUuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		collectionsId, err := r.repo.GetCollectionsIdFromTalkCollect(ctx, id)
		if err != nil {
			return v1.ErrorSetCollectFailed("get collections Id from collect failed: %s", err.Error())
		}
		err = r.repo.CancelCollectionsTalkCollect(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel talk user collect failed: %s", err.Error())
		}
		err = r.repo.CancelUserTalkCollect(ctx, id, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel user talk collect failed: %s", err.Error())
		}
		err = r.repo.CancelCollectionTalk(ctx, collectionsId, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel talk collection collect failed: %s", err.Error())
		}
		err = r.repo.ReduceCreationUserCollect(ctx, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("reduce creation user collect failed: %s", err.Error())
		}
		err = r.repo.CancelTalkCollect(ctx, id, uuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel talk collect failed: %s", err.Error())
		}
		err = r.repo.CancelTalkCollectFromCache(ctx, id, collectionsId, uuid, userUuid)
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel talk collect from cache failed: %s", err.Error())
		}
		err = r.repo.SendTalkStatisticToMq(ctx, uuid, "", "collect_cancel")
		if err != nil {
			return v1.ErrorSetCollectFailed("cancel talk collect to mq failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) TalkStatisticJudge(ctx context.Context, id int32, uuid string) (*TalkStatisticJudge, error) {
	agree, err := r.repo.GetTalkAgreeJudge(ctx, id, uuid)
	if err != nil {
		return nil, v1.ErrorGetStatisticJudgeFailed("get talk statistic judge failed: %s", err.Error())
	}
	collect, err := r.repo.GetTalkCollectJudge(ctx, id, uuid)
	if err != nil {
		return nil, v1.ErrorGetStatisticJudgeFailed("get talk statistic judge failed: %s", err.Error())
	}
	return &TalkStatisticJudge{
		Agree:   agree,
		Collect: collect,
	}, nil
}

func (r *TalkUseCase) AddTalkComment(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.AddTalkComment(ctx, id)
		if err != nil {
			return v1.ErrorAddCommentFailed("add talk comment failed: %s", err.Error())
		}

		err = r.repo.AddTalkCommentToCache(ctx, id, uuid)
		if err != nil {
			return v1.ErrorAddCommentFailed("add talk comment failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) ReduceTalkComment(ctx context.Context, id int32, uuid string) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		err := r.repo.ReduceTalkComment(ctx, id)
		if err != nil {
			return v1.ErrorReduceCommentFailed("reduce talk comment failed: %s", err.Error())
		}

		err = r.repo.ReduceTalkCommentToCache(ctx, id, uuid)
		if err != nil {
			return v1.ErrorReduceCommentFailed("reduce talk comment failed: %s", err.Error())
		}
		return nil
	})
}

func (r *TalkUseCase) TalkImageIrregular(ctx context.Context, review *ImageReview) error {
	err := r.repo.SendTalkImageIrregularToMq(ctx, review)
	if err != nil {
		return v1.ErrorSetImageIrregularFailed("set talk image irregular to mq failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) TalkContentIrregular(ctx context.Context, review *TextReview) error {
	err := r.repo.SendTalkContentIrregularToMq(ctx, review)
	if err != nil {
		return v1.ErrorSetContentIrregularFailed("set talk content irregular to mq failed: %s", err.Error())
	}
	return nil
}

func (r *TalkUseCase) AddTalkImageReviewDbAndCache(ctx context.Context, review *ImageReview) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		review, err := r.repo.SetTalkImageIrregular(ctx, review)
		if err != nil {
			return v1.ErrorSetImageIrregularFailed("set talk image irregular failed: %s", err.Error())
		}

		err = r.repo.SetTalkImageIrregularToCache(ctx, review)
		if err != nil {
			return v1.ErrorSetImageIrregularFailed("set talk image irregular to cache failed: %s", err.Error())
		}

		return nil
	})
}

func (r *TalkUseCase) AddTalkContentReviewDbAndCache(ctx context.Context, review *TextReview) error {
	return r.tm.ExecTx(ctx, func(ctx context.Context) error {
		review, err := r.repo.SetTalkContentIrregular(ctx, review)
		if err != nil {
			return v1.ErrorSetContentIrregularFailed("set talk content irregular failed: %s", err.Error())
		}

		err = r.repo.SetTalkContentIrregularToCache(ctx, review)
		if err != nil {
			return v1.ErrorSetContentIrregularFailed("set talk content irregular to cache failed: %s", err.Error())
		}

		return nil
	})
}

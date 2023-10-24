package biz

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"kubecit-service/internal/pkg/common"
	"strconv"
	"time"

	pb "kubecit-service/api/helloworld/v1"
	"kubecit-service/ent"
	"kubecit-service/internal/pkg/jwt"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	AccountMethodUsername = "username"
	AccountMethodWeChat   = "wechat"
)
const (
	UserRoleInvalid uint8 = iota
	UserRoleGuest
	UserRoleRegisterUser
	UserRoleLecturer
	UserRoleSuperAdmin
)
const (
	ServiceUserErrorCode = 100000
)
const (
	TypeSystemErrorCode = 1000 + iota*1000
	TypeDatabaseErrorCode
	TypeUserParamErrorCode
)

// ServiceUserErrorCode+TypeSystemErrorCode

const (
	_ = ServiceUserErrorCode + TypeSystemErrorCode + iota
)

// ServiceUserErrorCode+TypeDatabaseErrorCode

const (
	_ = ServiceUserErrorCode + TypeDatabaseErrorCode + iota
	UserSaveDatabaseErrorCode
	UserFindDatabaseErrorCode
)

// ServiceUserErrorCode+TypeUserParamErrorCode

const (
	_ = ServiceUserErrorCode + TypeUserParamErrorCode + iota
	UserUsernameIsExists
	UserUsernameNotExists
	UserPasswordNotMatch
)

type AccountPO struct {
	Id       uint64
	UserId   uint64
	Openid   string
	Password string
	Method   string
}

type UserPO struct {
	Id        uint64
	Username  string
	Channel   string
	RoleId    uint8
	UserId    int
	TeacherId int
}

type BecomeOrderInfo struct {
	Id        int
	UserId    int
	PayType   int
	VipType   int
	PayStatus int
	BizId     int64
	Price     int
	StartAt   time.Time
	ExpireAt  time.Time
}

type VipInfo struct {
	Id         int
	VipType    int
	VipOrderId int64
	StartAt    time.Time
	ExpireAt   time.Time
	UserId     int
}

type AccountRepo interface {
	FindByOpenidAndMethod(ctx context.Context, openid string, method string) (po *AccountPO, err error)
	Save(ctx context.Context, accountPO *AccountPO) error
}

type UserRepo interface {
	FindById(ctx context.Context, id uint64) (po *UserPO, err error)
	Save(ctx context.Context, po *UserPO) error
	SaveAccountAndUserTx(ctx context.Context, accountPO *AccountPO, userPO *UserPO) error
	Become(ctx context.Context, req *BecomeOrderInfo) error
	Callback(ctx context.Context, req *VipInfo, payStatus int) (*VipInfo, error)
	GetOrderInfo(ctx context.Context, bizId int64) (*BecomeOrderInfo, error)
	GetVipInfo(ctx context.Context, userId int) (*VipInfo, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	accountRepo AccountRepo
	userRepo    UserRepo
	log         *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(accountRepo AccountRepo, userRepo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		accountRepo: accountRepo,
		userRepo:    userRepo,
		log:         log.NewHelper(logger)}
}

func (usecase *UserUsecase) LoginByJson(ctx context.Context, request *pb.LoginByJsonRequest) (*pb.LoginByJsonReply, error) {

	accountPO, err := usecase.accountRepo.FindByOpenidAndMethod(ctx, request.Username, AccountMethodUsername)
	if accountPO == nil {
		if _, isEmpty := err.(*ent.NotFoundError); isEmpty {
			return &pb.LoginByJsonReply{
				Meta: usecase.errorMeta("用户名不存在", UserUsernameNotExists),
			}, nil
		} else {
			return &pb.LoginByJsonReply{
				Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
			}, nil
		}
	}
	if accountPO.Password != usecase.md5(request.Password) {
		return &pb.LoginByJsonReply{
			Meta: usecase.errorMeta("密码错误", UserPasswordNotMatch),
		}, nil

	}

	userPO, err := usecase.userRepo.FindById(ctx, accountPO.UserId)
	if err != nil {
		return &pb.LoginByJsonReply{
			Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
		}, nil
	}

	token, _ := jwt.GenerateToken(userPO.Id, userPO.RoleId)
	return &pb.LoginByJsonReply{
		Meta: &pb.Metadata{
			Code:    "0",
			Success: true,
		},
		Data: &pb.LoginByJsonReplyData{AccessToken: token},
	}, nil

}
func (usecase *UserUsecase) RegisterUsername(ctx context.Context, request *pb.RegisterUsernameRequest) (*pb.RegisterUsernameReply, error) {

	accountPO, err := usecase.accountRepo.FindByOpenidAndMethod(ctx, request.Username, AccountMethodUsername)

	if accountPO != nil {
		return &pb.RegisterUsernameReply{
			Meta: usecase.errorMeta("用户名已存在", UserUsernameIsExists),
		}, nil
	}

	if err != nil {
		if _, isEmpty := err.(*ent.NotFoundError); !isEmpty {

			usecase.log.Errorf("register username err: %v", err.Error())
			return &pb.RegisterUsernameReply{
				Meta: usecase.errorMeta("数据库发生错误", UserFindDatabaseErrorCode),
			}, nil

		}
	}

	userPO := &UserPO{
		Username: request.Username,
		Channel:  "",
		RoleId:   UserRoleRegisterUser,
	}
	accountPO = &AccountPO{

		Openid:   request.Username,
		Password: usecase.md5(request.Password),
		Method:   AccountMethodUsername,
	}
	err = usecase.userRepo.SaveAccountAndUserTx(ctx, accountPO, userPO)
	if err != nil {
		usecase.log.Errorf("register username err: %v", err.Error())
		return &pb.RegisterUsernameReply{
			Meta: usecase.errorMeta("数据库发生错误", UserSaveDatabaseErrorCode),
		}, nil
	}
	token, _ := jwt.GenerateToken(userPO.Id, userPO.RoleId)
	return &pb.RegisterUsernameReply{
		Meta: &pb.Metadata{
			Code:    "0",
			Success: true,
		},
		Data: &pb.LoginByJsonReplyData{AccessToken: token},
	}, nil
}
func (usecase *UserUsecase) md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str

}
func (usecase *UserUsecase) errorMeta(msg string, code int) *pb.Metadata {
	return &pb.Metadata{
		Msg:       msg,
		Code:      strconv.FormatInt(int64(code), 10),
		Success:   false,
		Version:   "",
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
	}

}

func (usecase *UserUsecase) CurrentUserInfo(ctx context.Context) (*pb.UserInfoReply, error) {
	UserId, err := common.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	userPO, err := usecase.userRepo.FindById(ctx, UserId)
	if err != nil {
		return &pb.UserInfoReply{}, err
	}
	vipInfo, err := usecase.userRepo.GetVipInfo(ctx, int(UserId))
	if err == nil {
		if vipInfo.ExpireAt.After(time.Now()) {
			return &pb.UserInfoReply{
				Username:  userPO.Username,
				Channel:   userPO.Channel,
				RoleId:    uint32(int32(userPO.RoleId)),
				UserId:    int32(userPO.UserId),
				TeacherId: int32(userPO.TeacherId),
				VipStatus: pb.VipStatus_VIP_STATUS_ACTIVE,
			}, nil
		}
	}
	return &pb.UserInfoReply{
		Username:  userPO.Username,
		Channel:   userPO.Channel,
		RoleId:    uint32(int32(userPO.RoleId)),
		UserId:    int32(userPO.UserId),
		TeacherId: int32(userPO.TeacherId),
		VipStatus: pb.VipStatus_VIP_STATUS_INACTIVE,
	}, nil
}

func (u *UserUsecase) Become(ctx context.Context, req *pb.BecomeVipRequest) (*pb.BecomeVipReply, error) {
	UserId, err := common.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	BizId, err := common.NewIdGenerator(time.Now(), 1)
	if err != nil {
		return &pb.BecomeVipReply{}, err
	}
	becomeReq := &BecomeOrderInfo{
		UserId:  int(UserId),
		PayType: int(req.GetPayType()),
		VipType: int(req.GetVipType()),
		BizId:   BizId.GenID(),
		Price:   int(req.Price),
	}
	err = u.userRepo.Become(ctx, becomeReq)
	if err != nil {
		return &pb.BecomeVipReply{}, err
	}
	return &pb.BecomeVipReply{
		PayLink: "https://mock-alipay.com/hgbonasdg146",
		BizId:   becomeReq.BizId,
	}, nil
}

func (u *UserUsecase) Callback(ctx context.Context, req *pb.TradeCallbackRequest) (*pb.TradeCallbackReply, error) {
	OrderRecord, err := u.userRepo.GetOrderInfo(ctx, req.BizNo)
	if err != nil {
		u.log.Errorf("")
		return nil, err
	}
	var expireAt time.Time
	switch OrderRecord.VipType {
	case int(pb.VipType_VIP_TYPE_MONTHLY):
		expireAt = time.Now().AddDate(0, 1, 0)
	case int(pb.VipType_VIP_TYPE_QUARTERLY):
		expireAt = time.Now().AddDate(0, 3, 0)
	case int(pb.VipType_VIP_TYPE_YEARLY):
		expireAt = time.Now().AddDate(1, 0, 0)
	case int(pb.VipType_VIP_TYPE_FOREVER):
		expireAt = time.Now().AddDate(100, 0, 0)
	default:
		return &pb.TradeCallbackReply{Message: "时间计算异常"}, errors.New("time error")
	}

	_, err = u.userRepo.Callback(ctx, &VipInfo{
		VipType:    OrderRecord.VipType,
		VipOrderId: OrderRecord.BizId,
		ExpireAt:   expireAt,
		UserId:     OrderRecord.UserId,
	}, int(req.PayStatus))
	if err != nil {
		return nil, err
	}
	return &pb.TradeCallbackReply{
		Message: "成功",
	}, nil
}

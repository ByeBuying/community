package protocol

type ResultCode int

const (
	Success                      ResultCode = 0
	Failed                                  = 1   // 요청이 실패하였습니다.
	IpInvalid                               = 2   // 요청한 아이피가 정상적이지 않습니다.
	UserIDNotFound                          = 13  // 유저 아이디가 존재하지 않습니다
	AccessTokenInvalid                      = 101 // 접속 토큰이 유효하지 않음
	UserNotFound                            = 102 // 유저 정보를 찾을 수 없음
	UserExistMnemonic                       = 103 // 유저의 니모닉이 이미 있습니다.
	OTPEncryptionFail                       = 104 // OTP를 암호화 하는 도중 실패하였습니다.
	CIOverlapError                          = 105 // 유저의 CI값이 이미 있습니다.(중복됩니다.)
	UserAddressNotFound                     = 106 // 유저 어드레스(체인)이 아무것도 없을 경우 입니다.
	UserNotSleep                            = 107 // 유저가 휴면 상태가 아닙니다.
	UserNotWithdrawalError                  = 108 // 유저가 탈퇴할 수 없는 상태입니다.
	OTPDeleteFail                           = 109 // OTP를 초기화하는데 실패하였습니다.
	OTPInvalid                              = 110 // OTP가 유효하지 않습니다.
	UserKlayAddressNotFound                 = 300 // 클레이튼 어드레스가 없습니다.
	KlayTxFail                              = 301 // 클레이튼 트렌젝션이 실패했습니다.
	UserEtherAddressNotFound                = 302 // 이더리움 어드레스가 없습니다.
	EtherTxFail                             = 303 // 이더리움 트렌젝션이 실패했습니다.
	KlayTxPending                           = 304 // 클레이튼 트렌젝션이 펜딩상태입니다.
	EtherTxPending                          = 305 // 이더리움 트렌젝션이 펜딩상태입니다.
	NptTxPending                            = 306 // Npt 트렌젝션이 펜딩상태입니다.
	UserNptAddressNotFound                  = 307 // Npt 어드레스가 없습니다.
	NptTxFail                               = 308 // Npt 트렌젝션이 실패했습니다.
	AlreadyTransaction                      = 309 // 이미 진해중인 트랜잭션이 존재한다.
	KlaytkTxPending                         = 310 // Klaytk 트렌젝션이 펜딩상태입니다.
	UserKlaytkAddressNotFound               = 311 // Klaytk 어드레스가 없습니다.
	KlaytkTxFail                            = 312 // Klaytk 트렌젝션이 실패했습니다.
	BlockedUser                             = 400 // Npt 트렌젝션이 블럭되었습니다.
	TronNetworkInvalid                      = 401 // Tron network 가 정상적으로 동작하지 않습니다.
	KlaytnNetworkInvalid                    = 402 // Klaytn network 가 정상적으로 동작하지 않습니다.
	EtherNetworkInvalid                     = 403 // Ether network 가 정상적으로 동작하지 않습니다.
	Maintenance                             = 500 // 전체점검 중
	KisaApiError                            = 600 // KISA API에서 에러가 난 경우.
	KisaApiBodyError                        = 601 // KISA API response body에서 에러가 난 경우
	NiceIdCommonError                       = 602 // NICE ID 공통 모듈 에러.
	StakingOrderStatInvalid                 = 700 // 스테이킹을 할 수 있는 상태가 아니다.
	TxInvalid                               = 701 // 트랜잭션이 올바르지 않습니다. (검증 실패)
	MinStkValError                          = 702 // 스테이킹 최소수량보다 적을경우 에러
	MinClaimValError                        = 703 // 클레임 최소수량보다 적을경우 에러
	InprogressClaim                         = 704 // 클레임이 진행중이다.
	InvalidStkValError                      = 705 // 본인이 가지고 있는 스테이킹 수량보다 더 요청했을 경우 에러
	InvalidStkTimeError                     = 706 // 스테이킹 할 수 있는 시간 (2일)전에 요청했을 경우 에러
	SoldOutAFO                              = 800 // AFO 다팔렸거나, 종료되었거나
	InvalidServiceTxId                      = 801 // ServiceTxId 가 올바르지 않다.
	InvalidMinVal                           = 802 // 최소 수량보다 작은 경우 발생한다.
	InvalidMaxVal                           = 803 // 최대 수량보다 큰 경우 발생한다.
	ExceedAmountPerson                      = 804 // 사용자가 예치할 수 있는 한도를 초과한 경우 발생한다.
	ExceedAmountPool                        = 805 // 풀에서 예치할 수 잇는 한도를 초과한 경우 발생한다.
	ExceedSlippage                          = 806 // 슬리피지 범위를 벗어난 경우 발생
	NotKYCApprovedFriendCode                = 900
	NotExistFriendCode                      = 901  // 존재하지 않는 초대코드
	InvalidEventFriendUser                  = 902  // 이벤트 기간 전 이미 KYC 완료한 가입자가 초대코드 입력시 이벤트 참여 대상자 X
	AlreadyRegisterEventFriend              = 903  // 이미 초대코드를 통해 가입 완료
	InvalidEventFriendPeriod                = 904  // 친구초대 이벤트 기간이 아니거나 이벤트를 중지한 상태
	LimitRegisterEventFriend                = 905  // 친구 초대 참여할 수 있는 사용자가 넘음
	SelfRegisterFriendCode                  = 906  // 본인의 친구코드를 넣었을 경우
	LockTxChainErr                          = 1000 // 체인별로 트랜잭션 요청을 막기위한 코드
	SwapV2Error                             = 1100 //SwapV2 Error
	SwapV2ApproveError                      = 1101 //SwapV2 Error
	InvalidAccessToken                      = 4010 // access_token 이 유효하지 않다.
	ExpireAccessToken                       = 4011 // access_token 이 만료되었다.
	NotLastRegisteredAccessToken            = 4012 // 서버에 마지막에 등록된 access_token 이 아니다(중복 로그인)
	PanicError                              = 5000 // 서버 패닉시 발생하는 공통 코드
)

func (r ResultCode) toString() string {
	switch r {
	case Success:
		return "Success"
	case Failed:
		return "Failed"
	case IpInvalid:
		return "IpInvalid"
	case UserIDNotFound:
		return "UserIDNotFound"
	case AccessTokenInvalid:
		return "AccessTokenInvalid"
	case UserNotFound:
		return "UserNotFound"
	case UserExistMnemonic:
		return "UserExistMnemonic"
	case OTPEncryptionFail:
		return "OTPEncryptionFail"
	case CIOverlapError:
		return "CIOverlapError"
	case UserAddressNotFound:
		return "UserAddressNotFound"
	case UserNotSleep:
		return "UserNotSleep"
	case UserNotWithdrawalError:
		return "UserNotWithdrawalError"
	case OTPDeleteFail:
		return "OTPDeleteFail"
	case OTPInvalid:
		return "OTPInvalid"
	case UserKlayAddressNotFound:
		return "UserKlayAddressNotFound"
	case KlayTxFail:
		return "KlayTxFail"
	case UserEtherAddressNotFound:
		return "UserEtherAddressNotFound"
	case EtherTxFail:
		return "EtherTxFail"
	case KlayTxPending:
		return "KlayTxPending"
	case EtherTxPending:
		return "EtherTxPending"
	case NptTxPending:
		return "NptTxPending"
	case UserNptAddressNotFound:
		return "UserNptAddressNotFound"
	case NptTxFail:
		return "NptTxFail"
	case AlreadyTransaction:
		return "AlreadyTransaction"
	case KlaytkTxPending:
		return "KlaytkTxPending"
	case UserKlaytkAddressNotFound:
		return "UserKlaytkAddressNotFound"
	case KlaytkTxFail:
		return "KlaytkTxFail"
	case BlockedUser:
		return "BlockedUser"
	case TronNetworkInvalid:
		return "TronNetworkInvalid"
	case KlaytnNetworkInvalid:
		return "KlaytnNetworkInvalid"
	case EtherNetworkInvalid:
		return "EtherNetworkInvalid"
	case KisaApiError:
		return "KisaApiError"
	case KisaApiBodyError:
		return "KisaApiBodyError"
	case NiceIdCommonError:
		return "NiceIdCommonError"
	case StakingOrderStatInvalid:
		return "StakingOrderStatInvalid"
	case TxInvalid:
		return "TxInvalid"
	case MinStkValError:
		return "MinStkValError"
	case InvalidStkValError:
		return "InvalidStkValError"
	case InvalidStkTimeError:
		return "InvalidStkTimeError"
	case MinClaimValError:
		return "MinClaimValError"
	case InprogressClaim:
		return "InprogressClaim"
	case SoldOutAFO:
		return "SoldOutAFO"
	case InvalidServiceTxId:
		return "InvalidServiceTxId"
	case InvalidMinVal:
		return "InvalidMinVal"
	case InvalidMaxVal:
		return "InvalidMaxVal"
	case ExceedAmountPerson:
		return "ExceedAmountPerson"
	case ExceedAmountPool:
		return "ExceedAmountPool"
	case ExceedSlippage:
		return "ExceedSlippage"
	case NotKYCApprovedFriendCode:
		return "NotKYCApprovedFriendCode"
	case NotExistFriendCode:
		return "NotExistFriendCode"
	case InvalidEventFriendUser:
		return "InvalidEventFriendUser"
	case AlreadyRegisterEventFriend:
		return "AlreadyRegisterEventFriend"
	case SelfRegisterFriendCode:
		return "SelfRegisterFriendCode"
	case LockTxChainErr:
		return "LockTxChainErr"
	case SwapV2Error:
		return "SwapV2Error"
	case InvalidAccessToken:
		return "InvalidAccessToken"
	case ExpireAccessToken:
		return "ExpireAccessToken"
	case NotLastRegisteredAccessToken:
		return "NotLastRegisteredAccessToken"
	case PanicError:
		return "InternalServerError"
	}
	return ""
}

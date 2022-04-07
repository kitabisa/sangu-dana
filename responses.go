package dana

type ResponseBody struct {
	Response  Response `json:"response" valid:"required"`
	Signature string   `json:"signature" valid:"required"`
}

type Response struct {
	Head ResponseHeader `json:"head" valid:"required"`
	Body interface{}    `json:"body" valid:"required"`
}

type ResponseHeader struct {
	Function  string `json:"function" valid:"required"`
	ClientID  string `json:"clientId" valid:"required"`
	Version   string `json:"version" valid:"required"`
	RespTime  string `json:"respTime" valid:"required"`
	RespMsgID string `json:"reqMsgId" valid:"required"`
}

type OrderResponseData struct {
	MerchantTransID string     `json:"merchantTransId,omitempty" valid:"optional"`
	AcquirementID   string     `json:"acquirementId,omitempty" valid:"optional"`
	CheckoutURL     string     `json:"checkoutUrl,omitempty" valid:"optional"`
	ResultInfo      ResultInfo `json:"resultInfo" valid:"required"`
}

type OrderDetailData struct {
	ResultInfo      ResultInfo     `json:"resultInfo" valid:"required"`
	AcquirementID   string         `json:"acquirementId" valid:"optional"`
	MerchantTransID string         `json:"merchantTransId" valid:"optional"`
	Buyer           InputUserInfo  `json:"buyer" valid:"optional"`
	Seller          InputUserInfo  `json:"seller" valid:"optional"`
	OrderTitle      string         `json:"orderTitle" valid:"optional"`
	ExtendedInfo    string         `json:"extendedInfo" valid:"optional"`
	AmountDetail    AmountDetail   `json:"amountDetail" valid:"optional"`
	TimeDetail      TimeDetail     `json:"timeDetail" valid:"optional"`
	StatusDetail    StatusDetail   `json:"statusDetail" valid:"optional"`
	Goods           []Good         `json:"goods" valid:"optional"`
	ShippingInfo    []ShippingInfo `json:"shippingInfo" valid:"optional"`
	OrderMemo       string         `json:"orderMemo" valid:"optional"`
	PaymentViews    []PaymentView  `json:"paymentViews" valid:"optional"`
}

type RefundResponseData struct {
	ResultInfo ResultInfo `json:"resultInfo" valid:"required"`
	RequestID  string     `json:"requestId,omitempty" valid:"optional"`
	RefundID   string     `json:"refundId,omitempty" valid:"optional"`
}

type ResultInfo struct {
	ResultStatus  string `json:"resultStatus" valid:"optional"`
	ResultCodeID  string `json:"resultCodeId" valid:"optional"`
	ResultMsg     string `json:"resultMsg" valid:"optional"`
	ResultCode    string `json:"resultCode" valid:"optional"`
	ResultMessage string `json:"resultMessage" valid:"optional"`
}

type PayFinishResponse struct {
	Response  ResponsePayFinish `json:"response"`
	Signature string            `json:"signature"`
}

type ResponsePayFinish struct {
	Head ResponseHeader        `json:"head"`
	Body ResponseBodyPayFinish `json:"body"`
}

type ResponseBodyPayFinish struct {
	ResultInfo ResultInfo `json:"resultInfo"`
}

type InputUserInfo struct {
	UserID           string `json:"userId" valid:"optional"`
	ExternalUserID   string `json:"externalUserId" valid:"optional"`
	ExternalUserType string `json:"externalUserType" valid:"optional"`
	Nickname         string `json:"nickname" valid:"optional"`
}

type AmountDetail struct {
	OrderAmount      Amount `json:"orderAmount" valid:"required"`
	PayAmount        Amount `json:"payAmount" valid:"optional"`
	VoidAmount       Amount `json:"voidAmount" valid:"optional"`
	ConfirmAmount    Amount `json:"confirmAmount" valid:"optional"`
	RefundAmount     Amount `json:"refundAmount" valid:"optional"`
	ChargebackAmount Amount `json:"chargebackAmount" valid:"optional"`
	ChargeAmount     Amount `json:"chargeAmount" valid:"optional"`
}

type TimeDetail struct {
	CreatedTime    string   `json:"createdTime" valid:"required"`
	ExpiryTime     string   `json:"expiryTime" valid:"required"`
	PaidTimes      []string `json:"paidTimes" valid:"optional"`
	ConfirmedTimes []string `json:"confirmedTimes" valid:"optional"`
	CancelledTime  string   `json:"cancelledTime" valid:"optional"`
}

type StatusDetail struct {
	AcquirementStatus string `json:"acquirementStatus" valid:"required"`
	Frozen            bool   `json:"frozen" valid:"required"`
}

type PaymentView struct {
	CashierRequestID     string          `json:"cashierRequestId" valid:"required"`
	PaidTime             string          `json:"paidTime" valid:"required"`
	PayOptionInfos       []PayOptionInfo `json:"payOptionInfos" valid:"required"`
	PayRequestExtendInfo string          `json:"payRequestExtendInfo" valid:"optional"`
	ExtendInfo           string          `json:"extendInfo" valid:"optional"`
}

type PayOptionInfo struct {
	PayMethod               string `json:"payMethod" valid:"required"`
	PayAmount               Amount `json:"payAmount" valid:"required"`
	TransAmount             Amount `json:"transAmount" valid:"optional"`
	ChargeAmount            Amount `json:"chargeAmount" valid:"optional"`
	ExtendInfo              string `json:"extendInfo" valid:"optional"`
	PayOptionBillExtendInfo string `json:"payOptionBillExtendInfo" valid:"optional"`
}

type AccessTokenInfo struct {
	AccessToken  string `json:"accessToken"`
	ExpiresIn    string `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
	ReExpiresIn  string `json:"reExpiresIn"`
	TokenStatus  string `json:"tokenStatus"`
}

type UserInfo struct {
	PublicUserID string `json:"publicUserId"`
}

type ApplyAccessToken struct {
	ResultInfo      ResultInfo      `json:"resultInfo"`
	AccessTokenInfo AccessTokenInfo `json:"accessTokenInfo"`
}

type UserResourceInfos struct {
	ResourceType string      `json:"resourceType"`
	Value        interface{} `json:"value"`
}

type UserProfileResponseData struct {
	ResultInfo        ResultInfo          `json:"resultInfo" valid:"required"`
	UserResourceInfos []UserResourceInfos `json:"userResourceInfos" valid:"required"`
}

type InquiryUserInfoResponse struct {
	Result   ResultInfo     `json:"result" valid:"required"`
	UserInfo ResultUserInfo `json:"userInfo"  valid:"required"`
}

type ResultUserInfo struct {
	UserContactInfoEmail string        `json:"USER_CONTACTINFO_EMAIL" valid:"optional"`
	UserName             string        `json:"USER_NAME" valid:"optional"`
	UserAddress          []UserAddress `json:"USER_ADDRESS" valid:"optional"`
	UserContactInfo      string        `json:"USER_CONTACTINFO" valid:"optional"`
}

type UserAddress struct {
	Area     string `json:"area" valid:"optional"`
	Province string `json:"province" valid:"optional"`
	City     string `json:"city" valid:"optional"`
	Address1 string `json:"address1" valid:"optional"`
}

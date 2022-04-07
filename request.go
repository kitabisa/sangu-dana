package dana

import "time"

type RequestBody struct {
	Request   Request `json:"request" valid:"required"`
	Signature string  `json:"signature" valid:"required"`
}

type Request struct {
	Head RequestHeader `json:"head" valid:"required"`
	Body interface{}   `json:"body" valid:"required"`
}

type RequestHeader struct {
	Version      string `json:"version" valid:"required"`
	Function     string `json:"function" valid:"required"`
	ClientID     string `json:"clientId" valid:"required"`
	ReqTime      string `json:"reqTime" valid:"required"`
	ReqMsgID     string `json:"reqMsgId" valid:"required"`
	ClientSecret string `json:"clientSecret" valid:"required"`
	AccessToken  string `json:"accessToken,omitempty" valid:"optional"`
	Reserve      string `json:"reserve,omitempty" valid:"optional"`
}

type OrderRequestData struct {
	Order             Order              `json:"order" valid:"required"`
	MerchantID        string             `json:"merchantId" valid:"required"`
	Mcc               string             `json:"mcc,omitempty" valid:"optional"`
	ProductCode       string             `json:"productCode" valid:"required"`
	EnvInfo           EnvInfo            `json:"envInfo" valid:"required"`
	NotificationUrls  *[]NotificationUrl `json:"notificationUrls,omitempty" valid:"optional"`
	ExtendInfo        string             `json:"extendInfo,omitempty" valid:"optional"`
	PaymentPreference *PaymentPreference `json:"paymentPreference,omitempty" valid:"optional"`
}

type OrderDetailRequestData struct {
	MerchantID      string `json:"merchantId" valid:"required"`
	AcquirementID   string `json:"acquirementId" valid:"optional"`
	MerchantTransID string `json:"merchantTransId" valid:"optional"`
}

type RefundRequestData struct {
	RequestID           string       `json:"requestId" valid:"required"`
	MerchantID          string       `json:"merchantId" valid:"required"`
	AcquirementID       string       `json:"acquirementId,omitempty" valid:"optional"`
	RefundAmount        Amount       `json:"refundAmount,omitempty" valid:"required"`
	RefundAppliedTime   time.Time    `json:"refundAppliedTime,omitempty" valid:"optional"`
	ActorType           string       `json:"actorType,omitempty" valid:"optional"`
	RefundReason        string       `json:"refundReason,omitempty" valid:"optional"`
	ReturnChargeToPayer bool         `json:"returnChargeToPayer,omitempty" valid:"optional"`
	Destination         string       `json:"destination,omitempty" valid:"optional"`
	ExtendInfo          string       `json:"extendInfo,omitempty" valid:"optional"`
	EnvInfo             EnvInfo      `json:"envInfo,omitempty" valid:"optional"`
	AuditInfo           AuditInfo    `json:"auditInfo,omitempty" valid:"optional"`
	ActorContext        ActorContext `json:"actorContext,omitempty" valid:"optional"`
}

type Order struct {
	OrderTitle        string         `json:"orderTitle"`
	OrderAmount       Amount         `json:"orderAmount"`
	MerchantTransID   string         `json:"merchantTransId"`
	MerchantTransType string         `json:"merchantTransType,omitempty"`
	OrderMemo         string         `json:"orderMemo,omitempty"`
	CreatedTime       string         `json:"createdTime,omitempty"`
	ExpiryTime        string         `json:"expiryTime,omitempty"`
	Goods             []Good         `json:"goods,omitempty"`
	ShippingInfo      []ShippingInfo `json:"shippingInfo,omitempty"`
}

type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

type Good struct {
	MerchantGoodsID    string `json:"merchantGoodsId,omitempty"`
	Description        string `json:"description"`
	Category           string `json:"category,omitempty"`
	Price              Amount `json:"price"`
	Unit               string `json:"unit,omitempty"`
	Quantity           string `json:"quantity,omitempty"`
	MerchantShippingID string `json:"merchantShippingId,omitempty"`
	SnapshotURL        string `json:"snapshotUrl,omitempty"`
	ExtendInfo         string `json:"extendInfo,omitempty"`
}

type ShippingInfo struct {
	MerchantShippingID string `json:"merchantShippingId"`
	TrackingNo         string `json:"trackingNo,omitempty"`
	Carrier            string `json:"carrier,omitempty"`
	ChargeAmount       Amount `json:"chargeAmount,omitempty"`
	CountryName        string `json:"countryName"`
	StateName          string `json:"stateName"`
	CityName           string `json:"cityName"`
	AreaName           string `json:"areaName,omitempty"`
	Address1           string `json:"address1"`
	Address2           string `json:"address2,omitempty"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	MobileNo           string `json:"mobileNo,omitempty"`
	PhoneNo            string `json:"phoneNo,omitempty"`
	ZipCode            string `json:"zipCode"`
	Email              string `json:"email,omitempty"`
	FaxNo              string `json:"faxNo,omitempty"`
}

type EnvInfo struct {
	SessionID          string `json:"sessionId,omitempty"`
	TokenID            string `json:"tokenId,omitempty"`
	WebsiteLanguage    string `json:"websiteLanguage,omitempty"`
	ClientIP           string `json:"clientIp,omitempty"`
	OsType             string `json:"osType,omitempty"`
	AppVersion         string `json:"appVersion,omitempty"`
	SdkVersion         string `json:"sdkVersion,omitempty"`
	SourcePlatform     string `json:"sourcePlatform"`
	TerminalType       string `json:"terminalType"`
	ClientKey          string `json:"clientKey,omitempty"`
	OrderTerminalType  string `json:"orderTerminalType"`
	OrderOsType        string `json:"orderOsType,omitempty"`
	MerchantAppVersion string `json:"merchantAppVersion,omitempty"`
	ExtendInfo         string `json:"extendInfo,omitempty"`
}

type AuditInfo struct {
	ActionReason  string `json:"actionReason" valid:"optional"`
	ThirdClientID string `json:"thirdClientId" valid:"optional"`
}

type ActorContext struct {
	ActorID   string `json:"actorId" valid:"required"`
	ActorType string `json:"actorType" valid:"required"`
}

type NotificationUrl struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

type PaymentPreference struct {
	DisabledPayMethods string `json:"disabledPayMethods"`
}

type PayFinishRequest struct {
	Request   RequestPayFinish `json:"request"`
	Signature string           `json:"signature"`
}
type RequestPayFinish struct {
	Head RequestHeader        `json:"head"`
	Body RequestBodyPayFinish `json:"body"`
}

type RequestBodyPayFinish struct {
	AcquirementID     string `json:"acquirementId"`
	MerchantTransID   string `json:"merchantTransId"`
	FinishedTime      string `json:"finishedTime"`
	CreatedTime       string `json:"createdTime"`
	MerchantID        string `json:"merchantId"`
	OrderAmount       Amount `json:"orderAmount"`
	AcquirementStatus string `json:"acquirementStatus"`
	ExtendInfo        string `json:"extendInfo"`
}

type RequestApplyAccessToken struct {
	GrantType    string `json:"grantType"`
	AuthCode     string `json:"authCode"`
	RefreshToken string `json:"refreshToken"`
}

type UserProfileRequestData struct {
	UserResources []string `json:"userResources" valid:"required"`
}

type InquiryUserInfoRequest struct {
	AccessToken string `json:"accessToken" valid:"required"`
	ExtendInfo  string `json:"extendInfo,omitempty" valid:"optional"`
}

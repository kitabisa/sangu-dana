package dana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/tidwall/gjson"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	ORDER_PATH              = "alipayplus/acquiring/order/createOrder.htm"
	QUERY_PATH              = "alipayplus/acquiring/order/query.htm"
	REFUND_PATH             = "alipayplus/acquiring/refund/refund.htm"
	APPLY_ACCESS_TOKEN_PATH = "dana/oauth/auth/applyToken.htm"
	USER_PROFILE_PATH       = "alipayplus/member/query/queryUserProfile.htm"
	DANA_TIME_LAYOUT        = "2006-01-02T15:04:05-07:00"
	CURRENCY_IDR            = "IDR"
	INQUIRY_USER_INFO_PATH  = "v1/customers/user/inquiryUserInfoByAccessToken.htm"

	FUNCTION_CREATE_ORDER       = "dana.acquiring.order.createOrder"
	FUNCTION_QUERY_ORDER        = "dana.acquiring.order.query"
	FUNCTION_REFUND             = "dana.acquiring.refund.refund"
	FUNCTION_APPLY_ACCESS_TOKEN = "dana.oauth.auth.applyToken"
	FUNCTION_USER_PROFILE       = "dana.member.query.queryUserProfile"
	FUNCTION_INQUIRY_USER_INFO  = "customers.openapi.user.inquiryUserInfoByAccessToken"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) Order(reqBody *OrderRequestData, accessToken string) (res ResponseBody, err error) {
	reqBody.Order.OrderAmount.Value = fmt.Sprintf("%v00", reqBody.Order.OrderAmount.Value)

	res, err = gateway.requestToDana(reqBody, accessToken, FUNCTION_CREATE_ORDER, ORDER_PATH)
	if err != nil {
		return
	}

	var orderResponseData OrderResponseData
	err = mapstructure.Decode(res.Response.Body, &orderResponseData)
	if err != nil {
		return
	}

	res.Response.Body = orderResponseData

	return
}

func (gateway *CoreGateway) OrderDetail(reqBody *OrderDetailRequestData, accessToken string) (res ResponseBody, err error) {
	res, err = gateway.requestToDana(reqBody, accessToken, FUNCTION_QUERY_ORDER, QUERY_PATH)
	if err != nil {
		return
	}

	var orderDetailData OrderDetailData
	err = mapstructure.Decode(res.Response.Body, &orderDetailData)
	if err != nil {
		return
	}

	res.Response.Body = orderDetailData

	return
}

func (gateway *CoreGateway) ApplyAccessToken(reqBody *RequestApplyAccessToken) (res ResponseBody, err error) {
	res, err = gateway.requestToDana(reqBody, "", FUNCTION_APPLY_ACCESS_TOKEN, APPLY_ACCESS_TOKEN_PATH)
	if err != nil {
		return
	}

	var applyAccessToken ApplyAccessToken
	err = mapstructure.Decode(res.Response.Body, &applyAccessToken)
	if err != nil {
		return
	}

	res.Response.Body = applyAccessToken

	return
}

func (gateway *CoreGateway) Refund(reqBody *RefundRequestData, accessToken string) (res ResponseBody, err error) {
	reqBody.RefundAmount.Value = fmt.Sprintf("%v00", reqBody.RefundAmount.Value)

	res, err = gateway.requestToDana(reqBody, accessToken, FUNCTION_REFUND, REFUND_PATH)
	if err != nil {
		return
	}

	var RefundResponseData RefundResponseData
	err = mapstructure.Decode(res.Response.Body, &RefundResponseData)
	if err != nil {
		return
	}

	res.Response.Body = RefundResponseData

	return
}

func (gateway *CoreGateway) GenerateSignature(req interface{}) (signature string, err error) {
	signature, err = generateSignature(req, gateway.Client.PrivateKey)
	if err != nil {
		err = fmt.Errorf("failed to generate signature: %v", err)
		return
	}

	return
}

func (gateway *CoreGateway) VerifySignature(res []byte, signature string) (err error) {
	response := gjson.Get(string(res), "request")
	err = verifySignature(response.String(), signature, gateway.Client.PublicKey)
	if err != nil {
		err = fmt.Errorf("could not verify request: %v", err)
	}
	return
}

func (gateway *CoreGateway) UserProfile(reqBody *UserProfileRequestData, accessToken string) (res ResponseBody, err error) {
	res, err = gateway.requestToDana(reqBody, accessToken, FUNCTION_USER_PROFILE, USER_PROFILE_PATH)
	if err != nil {
		return
	}

	var userProfileResponseData UserProfileResponseData
	err = mapstructure.Decode(res.Response.Body, &userProfileResponseData)
	if err != nil {
		return
	}

	res.Response.Body = userProfileResponseData

	return
}

func (gateway *CoreGateway) InquiryUserInfo(reqBody *InquiryUserInfoRequest, accessToken string) (res InquiryUserInfoResponse, err error) {
	var response interface{}
	response, err = gateway.requestToDanaV1(reqBody, accessToken, FUNCTION_INQUIRY_USER_INFO, INQUIRY_USER_INFO_PATH)
	if err != nil {
		return
	}

	var responseByte []byte
	responseByte, err = json.Marshal(response)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseByte, &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) requestToDana(reqBody interface{}, accessToken string, headerFunction string, path string) (res ResponseBody, err error) {
	now := time.Now()

	head := RequestHeader{}
	head.Version = gateway.Client.Version
	head.Function = headerFunction
	head.ClientID = gateway.Client.ClientId
	head.ReqTime = now.Format(DANA_TIME_LAYOUT)
	head.ClientSecret = gateway.Client.ClientSecret

	if accessToken != "" {
		head.AccessToken = accessToken
	}

	var id uuid.UUID
	id, err = uuid.NewUUID()
	if err != nil {
		return res, err
	}

	head.ReqMsgID = id.String()

	req := Request{
		Head: head,
		Body: reqBody,
	}

	sig, err := generateSignature(req, gateway.Client.PrivateKey)
	if err != nil {
		err = fmt.Errorf("failed to generate signature: %v", err)
		return
	}
	orderDetailReq := RequestBody{
		Request:   req,
		Signature: sig,
	}

	reqJson, err := json.Marshal(orderDetailReq)
	if err != nil {
		return
	}

	log.Println("Dana request: ", string(reqJson))
	requestReader := bytes.NewBuffer(reqJson)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	err = gateway.Call("POST", path, headers, requestReader, &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) requestToDanaV1(reqBody interface{}, accessToken string, headerFunction string, path string) (res interface{}, err error) {
	now := time.Now()

	head := RequestHeader{}
	head.Version = gateway.Client.Version
	head.Function = headerFunction
	head.ClientID = gateway.Client.ClientId
	head.ReqTime = now.Format(DANA_TIME_LAYOUT)
	head.ClientSecret = gateway.Client.ClientSecret

	if accessToken != "" {
		head.AccessToken = accessToken
	}

	var id uuid.UUID
	id, err = uuid.NewUUID()
	if err != nil {
		return
	}

	head.ReqMsgID = id.String()

	req := Request{
		Head: head,
		Body: reqBody,
	}

	sig, err := generateSignature(req, gateway.Client.PrivateKey)
	if err != nil {
		err = fmt.Errorf("failed to generate signature: %v", err)
		log.Println("generateSignature Failed: ", err)
		return
	}

	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	log.Println("Dana request: ", string(reqJson))

	headers := map[string]string{
		"Content-Type": "application/json",
		"Client-Id":    gateway.Client.ClientId,
		"Request-Time": now.Format(DANA_TIME_LAYOUT),
		"Signature":    sig,
	}

	bodyReq := bytes.NewBuffer(reqJson)
	err = gateway.Call("POST", path, headers, bodyReq, &res)
	if err != nil {
		log.Println("Failed call dana endpoint: ", err)
		return
	}

	return
}

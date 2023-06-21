package models

const (
	DataFoundStatus          = "200"
	DataWriteStatus          = "201"
	InvalidRequestStatus     = "400"
	DataNotFoundStatus       = "404"
	FailParsingStatus        = "406"
	DataWriteFailStatus      = "500"
	HealthCheckStatus        = "200"
	PermissionDeniedStatus   = "403"
	InvalidDataStatus        = "400"
	FailSettingSessionStatus = "406"
)

const (
	DataWriteMsg            = "Record has been successfully inserted!"
	DataUpdateMsg           = "Record has been successfully updated!"
	DataUpdateFailMsg       = "Record update failed!"
	DataWriteFailMsg        = "Record insertion failed!"
	DataFoundMsg            = "Records have been successfully retrieved!"
	DataNotFoundMsg         = "Records Not Found"
	InvalidRequestMsg       = "Invalid Request. Try Again Using Valid Request"
	FailParsingMsg          = "Something Went Wrong Parsing Data"
	HealthCheckMsg          = "Health Check API"
	PermissionDeniedMsg     = "permission Denied"
	UserNotFoundMsg         = "user Not Found"
	InvalidAmountMsg        = "Entered Amount is Invalid"
	InvalidCardMsg          = "Entered Card Details are Invalid"
	Success                 = "Successful"
	Unsuccessful            = "Unsuccessful"
	UnsuccessfulTransaction = "Unsuccessful Transactions didn't added"
	InsufficientFunds       = "Unsuccessful Insufficient Funds in your account"
	InvalidCredentials      = "Unsuccessful Invalid Credentials"
	NegativeAmount          = "Unsuccessful Negative Amount Entered"
	ExceededLimit           = "Unsuccessful You Exceeded amount limit"
	FailSettingSessionMsg   = "Unable to set Session"
)

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func DataResponse(msg, status string, data interface{}) (res Response) {
	res.Message = msg
	res.Status = status
	res.Data = data
	return
}

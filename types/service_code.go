package types

const (
	// Generic Code

	CodeOK int = iota
	// CodeBindError indicates a parameter missing error
	CodeBindError
	// CodeUnserializeDataError indicates a parsing data error
	CodeUnserializeDataError
	// CodeInvalidParameters tells some wrong data was in the request
	CodeInvalidParameters
	// GetRawDataError tells some wrong data was in the request
	CodeGetRawDataError
)

const (
	// Generic Code -- Database
	// CodeInsertError occurs when insert object into database
	CodeInsertError int = iota + 100
	// CodeSelectError occurs when select object from database
	CodeSelectError
	// CodeNotFound occurs when object with specific condition is not in the
	// database
	CodeNotFound
	// CodeDeleteNoEffect occurs when deleting object has no effect
	CodeDeleteNoEffect
	// CodeDuplicatePrimaryKey occurs when the object's primary key conflicts
	// with something that was already in the database
	CodeDuplicatePrimaryKey
	// CodeUpdateError occurs when update object to database
	CodeUpdateError // 105
	// CodeDeleteError occurs when delete object from database
	CodeDeleteError

	// CodeDeleteError occurs when begin a transaction
	CodeBeginTransactionError

	// CodeDeleteError occurs when commit a transaction
	CodeCommitTransactionError
)

const (
	// Generic Code -- Authentication
	// CodeAuthGenerateTokenError occurs when insert object into database
	CodeAuthGenerateTokenError int = iota + 1000
	CodeAuthenticatePasswordError
	CodeAuthenticatePolicyError

	CodeChangeOwnerError
	CodeGroupCreateError
	CodeAddReadPrivilegeError
	CodeAddWritePrivilegeError
)

const (
	CodeUserIDMissing int = iota + 10000
	CodeUserWrongPassword
	CodeWeakPassword
	CodeInvalidCityCode
)

const (
	CodeSubmissionUploaded int = iota + 11000
	CodeFSExecError
	CodeUploadFileError
	CodeConfigModifyError
	CodeStatError
)

const (
	CodeGoodsStatusUnknown int = iota + 12000
	CodeGoodsStatusFinished
	CodeGoodsStatusCancelled
	CodeGoodsLifeTimeout
	CodeGoodsStatusNotBeUnfinished
	CodeGoodsStatusNotBePending
	CodeGoodsBuyTypeInvalid
	CodeGoodsInsufficientValue
	CodeGoodsOverflowValue
)

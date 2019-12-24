//go:generate stringer -type=CodeType
package types

type CodeType int

const (
	// Generic Code

	CodeOK CodeType = iota
	// CodeBindError indicates a parameter missing error
	CodeBindError
	// CodeUnserializeDataError indicates a parsing data error
	CodeUnserializeDataError
	// CodeInvalidParameters tells some wrong data was in the request
	CodeInvalidParameters
	// GetRawDataError tells some wrong data was in the request
	CodeGetRawDataError

	CodeGenericErrorR
	CodeGenericErrorL = CodeOK
)

const (
	// Generic Code -- Database
	// CodeInsertError occurs when insert object into database
	CodeInsertError CodeType = iota + 100
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

	CodeDatabaseErrorR
	CodeDatabaseErrorL = CodeInsertError
)

const (
	// Generic Code -- Authentication
	// CodeAuthGenerateTokenError occurs when insert object into database
	CodeAuthGenerateTokenError CodeType = iota + 1000
	CodeAuthenticatePasswordError
	CodeAuthenticatePolicyError

	CodeChangeOwnerError
	CodeGroupCreateError
	CodeAddReadPrivilegeError
	CodeAddWritePrivilegeError

	CodeGrantNoEffect
	CodeGrantError

	CodeAuthenticationErrorR
	CodeAuthenticationErrorL = CodeAuthGenerateTokenError
)

const (
	CodeUserIDMissing CodeType = iota + 10000
	CodeUserWrongPassword
	CodeWeakPassword
	CodeInvalidCityCode
	CodeBadPhone

	CodeUserServiceErrorR
	CodeUserServiceErrorL = CodeUserIDMissing
)

const (
	CodeSubmissionUploaded CodeType = iota + 11000
	CodeFSExecError
	CodeUploadFileError
	CodeConfigModifyError
	CodeStatError

	CodeFileSystemErrorR
	CodeFileSystemErrorL = CodeSubmissionUploaded
)

const (
	CodeGoodsStatusUnknown CodeType = iota + 12000
	CodeGoodsStatusFinished
	CodeGoodsStatusCancelled
	CodeGoodsLifeTimeout
	CodeGoodsStatusNotBeUnfinished
	CodeGoodsStatusNotBePending
	CodeGoodsStatusNotBeUnfinishedOrPending
	CodeGoodsBuyTypeInvalid
	CodeGoodsInsufficientValue
	CodeGoodsOverflowValue

	CodeGoodsServiceErrorR
	CodeGoodsServiceErrorL = CodeGoodsStatusUnknown
)

var CodeDesc map[CodeType]string

func init() {
	CodeDesc = make(map[CodeType]string)
	for _, groupCode := range []struct {
		L CodeType
		R CodeType
	}{
		{CodeGenericErrorL, CodeGenericErrorR},
		{CodeDatabaseErrorL, CodeDatabaseErrorR},
		{CodeAuthenticationErrorL, CodeAuthenticationErrorR},
		{CodeFileSystemErrorL, CodeFileSystemErrorR},
		{CodeUserServiceErrorL, CodeUserServiceErrorR},
		{CodeGoodsServiceErrorL, CodeGoodsServiceErrorR},
	} {
		for i := groupCode.L; i < groupCode.R; i++ {
			CodeDesc[i] = i.String()
		}
	}
}

package api

type EnumShareLinkType int

const (
	EnumShareByKey            EnumShareLinkType = 0
	EnumRestricted            EnumShareLinkType = 1
	EnumAuthenticated         EnumShareLinkType = 2
	EnumShareByKeyAndPassword EnumShareLinkType = 3
)

type AccessType int

const (
	NONE       AccessType = 0
	READ       AccessType = 1
	READ_WRITE AccessType = 2
)

type EnumStatus int

const (
	COMMAND_STATUS_DONE        EnumStatus = 0
	COMMAND_STATUS_QUEUED      EnumStatus = 1
	COMMAND_STATUS_IN_PROGRESS EnumStatus = 2
	COMMAND_STATUS_ERROR       EnumStatus = 3
	COMMAND_STATUS_CANCELLED   EnumStatus = 4
)

type EnumCommandErrorCode int

const (
	ERROR_NO_ERROR             EnumCommandErrorCode = 0
	ERROR_MISSING_COMMAND_BODY EnumCommandErrorCode = 1
	ERROR_MISSING_PARAMETERS   EnumCommandErrorCode = 2
	ERROR_INVALID_PARAMETERS   EnumCommandErrorCode = 3
	ERROR_NOT_ALLOWED          EnumCommandErrorCode = 4
	ERROR_INVALID_PATH         EnumCommandErrorCode = 5
	ERROR_FILE_SYSTEM          EnumCommandErrorCode = 6
	ERROR_SAVING               EnumCommandErrorCode = 7
	ERROR_UNKNOWN              EnumCommandErrorCode = 8
)

package settings

type IStatus int

type IResponseType string
type IResponseArgs *[]string

type IResponseBody struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var ResponseTypes = struct {
	ServerInternalError        IResponseType
	ForbiddenError             IResponseType
	CommonError                IResponseType
	UserNameLengthError        IResponseType
	UserNameFormatError        IResponseType
	UserNameExistsError        IResponseType
	UserPasswdLengthError      IResponseType
	UserPasswdFormatError      IResponseType
	UserNameNullError          IResponseType
	UserPasswdNullError        IResponseType
	UserNameOrPasswdError      IResponseType
	ForbiddenOnlyUser          IResponseType
	ForbiddenOnlyVisitor       IResponseType
	ForbiddenOnlyUserOperation IResponseType
	RadioNameLengthError       IResponseType
	RadioNameFormatError       IResponseType
	RadioDescLengthError       IResponseType
	RadioNameExistsError       IResponseType
	AudioExtError              IResponseType

	RadioAuthorError IResponseType
	RadioIIDError    IResponseType

	FMOfflineError IResponseType
}{
	ServerInternalError:        "ServerInternalError",
	ForbiddenError:             "ForbiddenError",
	ForbiddenOnlyUser:          "ForbiddenOnlyUser",
	ForbiddenOnlyVisitor:       "ForbiddenOnlyVisitor",
	ForbiddenOnlyUserOperation: "ForbiddenOnlyUserOperation",
	CommonError:                "CommonError",
	UserNameLengthError:        "UserNameLengthError",
	UserNameFormatError:        "UserNameFormatError",
	UserNameExistsError:        "UserNameExistsError",
	UserPasswdLengthError:      "UserPasswdLengthError",
	UserPasswdFormatError:      "UserPasswdFormatError",
	UserNameNullError:          "UserNameNullError",
	UserPasswdNullError:        "UserPasswdNullError",
	UserNameOrPasswdError:      "UserNameOrPasswdError",
	RadioNameLengthError:       "RadioNameLengthError",
	RadioNameFormatError:       "RadioNameFormatError",
	RadioDescLengthError:       "RadioDescLengthError",
	RadioNameExistsError:       "RadioNameExistsError",
	AudioExtError:              "AudioExtError",
	RadioAuthorError:           "RadioAuthorError",
	RadioIIDError:              "RadioIIDError",
	FMOfflineError:             "FMOfflineError",
}

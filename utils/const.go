package utils

const (
	ErrDataNotFound          = "mongo: no documents in result"
	ErrIDCannotBeEmpty       = "ID cannot be empty"
	ErrUsernameAlreadyExists = "username already exists"
	ErrPhoneAlreadyExists    = "phone already exists"
	ErrEmailAlreadyExists    = "email already exists"
	ErrPasswordNotMatch      = "password not match"
	ErrInvalidPassword       = "invalid password"
	ErrTokenDoesNotExists    = "token does not exists"
	ErrHeaderDoesNotExists   = "header does not exists"
	ErrCompanyDoesNotExists  = "company does not exists"
	ErrEmailAlreadyTaken     = "email already taken"
	ErrAccountAlreadyActive  = "account already active"
	ErrCompanyAlreadyExists  = "company already exists"
)

const (
	DateLayout             = "2006-01-02"
	DatetimeLayout         = "2006-01-02 15:04:05"
	DatetimeTimezoneLayout = "2006-01-02T15:04:05"
	DatetimeRFC3339        = "2006-01-02T15:04:05Z07:00"
)

const (
	ON_PROGRESS_USER_REQUEST_MESSAGE = "User Request: Waiting for approver"
	ON_PROGRESS_USER_REQUEST_CODE    = "on_progress_user_request"

	ON_PROGRESS_ADMIN_APPROVE_MESSAGE = "Admin Approve: Waiting to deliver"
	ON_PROGRESS_ADMIN_APPROVE_CODE    = "on_progress_admin_approve"

	DELIVER_ADMIN_MESSAGE = "Admin Delivering: On progress deliver to user"
	DELIVER_ADMIN_CODE    = "deliver_admin"

	ON_PROGRESS_USER_REJECT_MESSAGE = "User Reject: Item rejected by user Waiting to deliver back"
	ON_PROGRESS_USER_REJECT_CODE    = "on_progress_user_reject"

	DELIVER_USER_MESSAGE = "User Delivering: On progress deliver to admin"
	DELIVER_USER_CODE    = "deliver_user"

	COMPLETE_REJECTED_MESSAGE = "Admin Reject: Item request rejected by admin"
	COMPLETE_REJECTED_CODE    = "complete_rejected"

	COMPLETE_APPROVED_MESSAGE = "User Approve: Item received and approve by user"
	COMPLETE_APPROVED_CODE    = "complete_approved"
)

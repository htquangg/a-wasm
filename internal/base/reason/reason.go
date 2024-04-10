package reason

const (
	Success                   = "base.success"
	UnknownError              = "base.unknown"
	RequestFormatError        = "base.request_format_error"
	UnauthorizedError         = "base.unauthorized_error"
	DatabaseError             = "base.database_error"
	ForbiddenError            = "base.forbidden_error"
	DuplicateRequestError     = "base.duplicate_request_error"
	TooManyWrongAttemptsError = "base.too_many_wrong_attemps_error"
)

const (
	EmailOrPasswordWrong        = "error.object.email_or_password_incorrect"
	InvalidTokenError           = "error.common.invalid_token"
	InvalidScopeError           = "error.common.invalid_scope"
	EndpointNotFound            = "error.endpoint.not_found"
	EndpointHasNotPublished     = "error.endpoint.has_not_published"
	DeploymentNotFound          = "error.deployment.not_found"
	DeploymentAlreadyActivated  = "error.deployment.already_activated"
	EmailDuplicate              = "error.email.duplicate"
	EmailNotFound               = "error.email.not_found"
	RequiredSession             = "error.access_token.session_required"
	SessionNotFound             = "error.session.not_found"
	UserNotFound                = "error.user.not_found"
	SRPNotFound                 = "error.srp.not_found"
	SRPAlreadyVerified          = "error.srp.already_verified"
	SRPChallengeNotFound        = "error.srp_challenge.not_found"
	SRPChallengeAlreadyVerified = "error.srp_challenge.already_verified"
	KeyAttributeNotFound        = "error.key_attribute.not_found"
)
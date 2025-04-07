package reason

const (
	// 成功
	Success = "base.success"

	// 未知错误
	UnknownError = "base.unknown"

	// 参数错误
	RequestFormatError = "base.request_format_error"

	// 未登录
	UnauthorizedError = "base.unauthorized_error"

	// 数据库错误
	DatabaseError = "base.database_error"

	// 禁止访问
	ForbiddenError = "base.forbidden_error"

	// 重复请求
	DuplicateRequestError = "base.duplicate_request_error"
)

var (
	// 密码长度必须在 6-20 之间
	PasswordLengthError = "error.password.invalid_length"

	// 密码不能包含特殊字符
	PasswordSpecialCharacterError = "error.password.special_character"

	// 生成密码哈希失败
	GeneratePasswordHashFailed = "error.password.generate_hash_failed"

	// 邮箱格式错误
	EmailInvalid = "error.email.invalid"

	// 邮箱已存在
	EmailExists = "error.email.exists"

	// 用户不存在
	UserNotFound = "error.user.not_found"

	// 验证码错误
	InvalidVerificationCode = "error.verification_code.invalid"
)

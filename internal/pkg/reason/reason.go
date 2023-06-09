package reason

var (
	CategoryNotFound        = "category not found"
	CategoryCannotCreate    = "cannot Create Category"
	CategoryCannotBrowse    = "cannot Browse Category"
	CategoryCannotUpdate    = "cannot Update Category"
	CategoryCannotDelete    = "cannot Delete Category"
	CategoryCannotGetDetail = "cannot get detail category"
	InternalServerError     = "internal server error"
	RequestFormError        = "request format is not valid"
)
var (
	CurrencyNotFound        = "currency not found"
	CurrencyCannotCreate    = "cannot Create currency"
	CurrencyCannotBrowse    = "cannot Browse currency"
	CurrencyCannotUpdate    = "cannot Update currency"
	CurrencyCannotDelete    = "cannot Delete currency"
	CurrencyCannotGetDetail = "cannot get detail currency"
)

var (
	TransactionCannotCreate    = "cannot Create transaction"
	TransactionCannotBrowse    = "cannot Browse transaction"
	TransactionCannotGetDetail = "cannot get detail transaction"
	TransactionCannotDelete    = "cannot Delete transaction"
	TransactionCannotUpdate    = "cannot Update transaction"
	TransactionNotFound        = "transaction not found"
)
var (
	UserAlreadyExist    = "user already exist"
	RegisterFailed      = "cannot register user"
	UserNotFound        = "user not found"
	UserCannotDelete    = "user cannot delete"
	LoginFailed         = "login failed, please check your email or password"
	SaveToken           = "cannot save refresh token"
	UserSignOut         = "user has sign out"
	UserNotLogin        = "user has not logged in yet"
	NotAuthorized       = "You are not authorized to access this resource"
	ErrAuthorize        = "error occurred when authorizing user"
	UserNotAuthenticate = "user does not have an authentication"
)

var (
	ErrInvalidToken         = "token is invalid"
	ErrNoToken              = "request does not contain an access token"
	InvalidRefreshToken     = "invalid refresh token"
	CannotCreateAccessToken = "cannot create access token"
)

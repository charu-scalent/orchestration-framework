package errors

// ----------------------------------------------------------[ AUTH MICRO-SERVICE ]----------------------------------------------------------
const (
	INVALID_TOKEN            = "Invalid token"
	COULD_NOT_GET_SESSION    = "Could not get session"
	SERVER_DOWN              = "Server down"
	COULD_NOT_CREATE_SESSION = "could not create the session"
	COULD_NOT_DELETE_SESSION = "could not delete the session"
	INVALID_CREDENTIALS      = "invalid credentials"
	INVALID_ORG              = "provide valid organization"
	NO_ORG_REGISTERED        = "no organization is registered"
	INVALID_PASSWORD         = "invalid password"
	AUTH_NOT_RESPONDING      = "auth not responding"
	INVALID_USER             = "invalid user"
)

// ----------------------------------------------------------[ HTTP ERRORS ]-------------------------------------------------------------------

const (
	BAD_REQUEST_ERROR         = "Bad Request"           // HTTP code 400
	UNAUTHORIZED_ERROR        = "Unauthorized"          // HTTP code 401
	FORBIDDEN_ERROR           = "Forbidden"             // HTTP code 403
	NOT_FOUND_ERROR           = "Not Found"             // HTTP code 404
	METHOD_NOT_ALLOWED_ERROR  = "Method Not Allowed"    // HTTP code 405
	REQUEST_TIMEOUT_ERROR     = "Request Timeout"       // HTTP code 408
	CONFLICT_ERROR            = "Conflict"              // HTTP code 409
	EXPECTATION_FAILED_ERROR  = "Expectation Failed"    // HTTP code 417
	INTERNAL_SERVER_ERROR     = "Internal server error" // HTTP code 500
	BAD_GATEWAY_ERROR         = "Bad Gateway"           // HTTP code 502
	SERVICE_UNAVAILABLE_ERROR = "Service Unavailable"   // HTTP code 503
	GATEWAY_TIMEOUT_ERROR     = "Gateway Timeout"       // HTTP code 504
	NO_CONTENT_ERROR          = "No Content"            // HTTP code 204
)

// ----------------------------------------------------------[ USER MICRO-SERVICE ]----------------------------------------------------------
const (
	// Random Error
	INVALID_PAGE_NO         = "Invalid page no"
	MARSHALING_DATA         = "Error while marshalling data: "
	INVALID_REQUEST         = "Invalid request"
	PARSHING_TEMPLATE_ERROR = "Error parsing template"
	MARSHALING_ERROR        = "Error while marshalling"
	SOMETHING_WENT_WRONG    = "OOPS! there is something went wrong"
	EMAIL_SENDING_FAILED    = "Email sending failed"
	OTP_NOT_MATCHED         = "OTP did not match"
	OTP_EXPIRED             = "OTP Expired"
	EMAIL_NOT_REGISTERED    = "Email is not registered with us, please register first"
)

// --------------------------------------------------------[ATS MICRO-SERVICE]------------------------------------------------------------------------------
const (
	STEP_1_INCOMPLETE_ERROR        = "step 1 is not completed, please complete it first"
	STEP_2_INCOMPLETE_ERROR        = "step 2 is not completed, please complete it first"
	STEP_3_INCOMPLETE_ERROR        = "step 3 is not completed, please complete it first"
	STEP_1_ALREADY_COMPLETE_ERROR  = "step 1 is already completed, please update it"
	STEP_2_ALREADY_COMPLETE_ERROR  = "step 2 is already completed, please update it"
	STEP_3_ALREADY_COMPLETE_ERROR  = "step 3 is already completed, please update it"
	ADDRESSES_LIMIT_EXCEEDED       = "maximum number of addresses has been reached"
	LOCATION_ERROR                 = "add valid organziation location"
	ERROR_IN_PARSING_EMPTY_PAGE_NO = "strconv.Atoi: parsing \"\": invalid syntax"
	UNAUTHORIZED_USER              = "unauthorized user"
	JOB_NOT_FOUND                  = "job does not exists for given organization"
	CITY_ID_DOES_NOT_EXIST         = "cityID does not exits"
	STATE_ID_DOES_NOT_EXIST        = "stateID does not exist"
	CITY_NOT_ASSOCIATED_WITH_STATE = "city is not associated with"
	PLEASE_PROVIDE_CITY            = "please provide city"
	PLEASE_PROVIDE_STATE           = "please provide state"
	JOB_DOES_NOT_EXIST             = "job doesn't exist"
	DELETE_LOCATION_RESTRICTION    = "Org must have atleast one location"
	INVALID_PARAMETER              = "Please provide either self or all parameter"
	DEPARTMENT_NOT_EXISTS          = "department does not exists for this organization"
	INVALID_LOCATION               = "Invalid Location"
	PLEASE_CHECK_IN_FIRST          = "Please check_in first"
	ALREADY_CHECKOUT               = "you have already checkout, please update your attendence"
	INVALID_EMAIL                  = "Email Not Registered"
	ATTENDANCE_DOES_NOT_EXIST      = "your attendance does not exist"
)

var DYNAMIC_ERROR string

var ConstantMap = map[string]bool{
	INVALID_PAGE_NO:                true,
	INTERNAL_SERVER_ERROR:          true,
	MARSHALING_DATA:                true,
	INVALID_REQUEST:                true,
	PARSHING_TEMPLATE_ERROR:        true,
	MARSHALING_ERROR:               true,
	INVALID_TOKEN:                  true,
	COULD_NOT_GET_SESSION:          true,
	SERVER_DOWN:                    true,
	COULD_NOT_CREATE_SESSION:       true,
	COULD_NOT_DELETE_SESSION:       true,
	LOCATION_ERROR:                 true,
	STEP_2_INCOMPLETE_ERROR:        true,
	ERROR_IN_PARSING_EMPTY_PAGE_NO: true,
	UNAUTHORIZED_USER:              true,
	JOB_NOT_FOUND:                  true,
	CITY_ID_DOES_NOT_EXIST:         true,
	STATE_ID_DOES_NOT_EXIST:        true,
	CITY_NOT_ASSOCIATED_WITH_STATE: true,
	PLEASE_PROVIDE_CITY:            true,
	PLEASE_PROVIDE_STATE:           true,
	JOB_DOES_NOT_EXIST:             true,
	DELETE_LOCATION_RESTRICTION:    true,
	INVALID_PARAMETER:              true,
	DEPARTMENT_NOT_EXISTS:          true,
	INVALID_LOCATION:               true,
	INVALID_USER:                   true,
	PLEASE_CHECK_IN_FIRST:          true,
	ALREADY_CHECKOUT:               true,
	INVALID_EMAIL:                  true,
	ATTENDANCE_DOES_NOT_EXIST:      true,
}

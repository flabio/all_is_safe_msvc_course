package utils

const (
	LiMIT       int    = 5
	PAGE        string = "page"
	TOTAL_COUNT string = "totalCount"
	PAGE_COUNT  string = "pageCount"
	BEGIN       string = "begin"
)
const (
	ID                    = "id"
	MESSAGE               = "message"
	STATUS                = "status"
	DATA                  = "data"
	CREATED               = "was successfully created"
	UPDATED               = "was updated successfully"
	REMOVED               = "was successfully removed"
	ERROR_QUERY           = "error query, please try again later"
	ERROR_CREATE          = "error creating"
	ERROR_PARSING_BODY    = "error parsing body"
	ERROR_UPDATE          = "error updating"
	ERROR_DELETE          = "error deleting"
	ERROR_REQUIRED_FIELDS = " is required."
	EMPTY                 = ""
)
const (
	DB_DIFF_ID        = "id <>?"
	DB_EQUAL_ID       = "id=?"
	DB_EQUAL_NAME     = "name =?"
	DB_ORDER_DESC     = "id desc"
	DB_EQUAL_EMAIL_ID = "email=?"
)

const (
	ID_NO_EXIST         = "The id not exists"
	NAME_ALREADY_EXIST  = "Name already exists"
	EMAIL_ALREADY_EXIST = "Email already exists"
)
const (
	NAME      = "name"
	SCHOOL_ID = "school_id"
	ACTIVE    = "active"
)
const (
	NAME_IS_FIELD_REQUIRED      = "The field name is required."
	SCHOOL_ID_IS_FIELD_REQUIRED = "The field school_id is required."
)

const (
	NAME_IS_REQUIRED      = "The name is required."
	SCHOOL_ID_IS_REQUIRED = "The id school is required."
)

const (
	AUTHORIZATION string = "Authorization"
	BEARER        string = "Bearer "
	TOKEN_INVALID string = "Token not provided or invalid"
)

const USER string = "user"
const USER_NOT_FOUND string = "username o password incorrect."
const RECOVER_PANIC string = "Se recuperó de un panic:"
const ROL_NOT_FOUND string = "Rol not found"
const STATE_NOT_FOUND string = "State not found"
const FIRST_NAME string = "first_name"
const FIRST_SUR_NAME string = "first_sur_name"
const SECOND_SUR_NAME string = "second_sur_name"
const ADDRESS string = "address"
const PHONE string = "phone"
const ZIP_CODE string = "zip_code"
const PROVIDER_NUMBER string = "provider_number"
const STATE_ID string = "state_id"
const ROL_ID string = "rol_id"
const EMAIL string = "email"
const PASSWORD string = "password"
const TITLE string = "title"
const TIME_HOURS string = "time_hours"
const COURSE_ID string = "course_id"

const COURSE_ID_IS_FIELD_REQUIRED string = "The field course_id is required."
const TITLE_FIELD_IS_REQUIRED string = "The field title is required."
const TIME_HOURS_IS_FIELD_REQUIRED string = "The field time_hours is required."

const FIRST_NAME_IS_FIELD_REQUIRED string = "The field first name is required."
const FIRST_SUR_NAME_IS_FIELD_REQUIRED string = "The field first sur name is required."
const SECOND_SUR_NAME_IS_FIELD_REQUIRED string = "The field second sur name is required."
const ADDRESS_IS_FIELD_REQUIRED string = "The field address is required."
const PHONE_IS_FIELD_REQUIRED string = "The field phone is required."
const ZIP_CODE_IS_FIELD_REQUIRED string = "The field zip code is required."
const PROVIDER_NUMBER_IS_FIELD_REQUIRED string = "The field provider number is required."
const STATE_ID_IS_FIELD_REQUIRED string = "The field state id is required."
const ROL_ID_IS_FIELD_REQUIRED string = "The field rol id is required."
const EMAIL_IS_FIELD_REQUIRED string = "The field email is required."
const PASSWORD_IS_FIELD_REQUIRED string = "The field password is required."
const NAME_FIELD_IS_REQUIRED string = "The name field is required"
const ACTIVE_FIELD_IS_REQUIRED string = "The active field is required"
const CITY_ID_FIELD_IS_REQUIRED string = "The city id field is required"
const FIRST_NAME_IS_MIN_LENGTH_REQUIRED string = "The field first name must have a minimum length of 3 characters."
const FIRST_SUR_NAME_IS_MIN_LENGTH_REQUIRED string = "The field first sur name must have a minimum length of 3 characters."
const SECOND_SUR_NAME_IS_MIN_LENGTH_REQUIRED string = "The field second sur name must have a minimum length of 3 characters."
const ADDRESS_IS_MIN_LENGTH_REQUIRED string = "The field address must have a minimum length of 10 characters."
const PHONE_IS_MIN_LENGTH_REQUIRED string = "The field phone must have a minimum length of 10 characters."
const ZIP_CODE_IS_MIN_LENGTH_REQUIRED string = "The field zip code must have a minimum length of 5 characters."
const STATE_ID_IS_MIN_LENGTH_REQUIRED string = "The field state id "
const FIRST_NAME_IS_REQUIRED string = "The first name is required."
const FIRST_SUR_NAME_IS_REQUIRED string = "The first sur name is required."
const SECOND_SUR_NAME_IS_REQUIRED string = "The second sur name is required."
const ADDRESS_IS_REQUIRED string = "The address is required."
const PHONE_IS_REQUIRED string = "The phone is required."
const ZIP_CODE_IS_REQUIRED string = "The zip code is required."
const PROVIDER_NUMBER_IS_REQUIRED string = "The provider number is required."
const STATE_ID_IS_REQUIRED string = "The state id is required."
const ROL_ID_IS_REQUIRED string = "The rol id is required."
const EMAIL_IS_REQUIRED string = "The email is required."
const PASSWORD_IS_REQUIRED string = "The password is required."
const EMAIL_IS_INVALID string = "The email is invalid."
const CITY_ID_IS_REQUIRED string = "The city id is required."

const TITLE_IS_REQUIRED string = "The title is required."
const COURSE_ID_IS_REQUIRED string = "The course id is required."
const TIME_HOURS_IS_REQUIRED string = "The time hours is required."
const PASSWORD_IS_INVALID_LENGTH string = "The password must have a minimum length of 8 characters."
const (
	MSVC_ROL_URL               = "http://3.90.42.175:3001/api/rol"
	MSVC_PARENTESCO_URL        = "http://3.85.174.77:3002/api/parentesco"
	MSVC_EMERGENCY_CONTACT_URL = "http://3.81.78.83:3003/api/emergency"
	MSVC_CITY_URL              = "http://54.224.197.15:3014/api/cities"
	MSVC_STATES_URL            = "http://54.224.197.15:3014/api/states"
	MSVC_STATES_BY_CITY_URL    = "http://54.224.197.15:3014/api/states/city"
	MSVC_USER_URL              = "http://34.227.88.155:3005/api/user"
	MSVC_AUTH_URL              = "http://34.227.88.155:3005/api/auth"
	MSVC_SCHOOL_URL            = "http://54.91.68.197:3006/api/school"
	MSVC_COURSE_URL            = "http://52.205.201.42:3007/api/course"
)
const GET string = "GET"
const POST string = "POST"
const PUT string = "PUT"
const DELETE string = "DELETE"

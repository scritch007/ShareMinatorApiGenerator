package api

var RequestGetUsersUrl = "config/get_users"

type RequestGetUsersOutput struct {
	Users []Account `json:"users"`
}

var RequestDummyAuthUrl = "auths/dummy.auth"

type RequestDummyAuthInput struct {
	Login         string `json:"login"`
	ChallengeHash string `json:"challenge_hash"`
	Ref           string `json:"ref"`
}
type RequestDummyAuthOutput struct {
	MySelf               Account `json:"my_self"`
	AuthenticationHeader string  `json:"authentication_header"`
}

var RequestDummyGetChallengeUrl = "auths/dummy.get_challenge"

type RequestDummyGetChallengeOutput struct {
	Challenge string `json:"challenge"`
	Ref       string `json:"ref"`
}

var RequestDummyCreateUrl = "auths/dummy.create"

type RequestDummyCreateInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
	IsAdmin  *bool  `json:"is_admin,omitempty"`
}

var RequestLogoutUrl = "auths/logout"
var RequestListUrl = "auths/list"

type RequestListOutput struct {
	Auths []string `json:"auths"`
}

var RequestGetInfoUrl = "config/get_info"

type RequestGetInfoOutput struct {
	ShareLink         bool       `json:"share_link"`
	PasswordProtected bool       `json:"password_protected"`
	NbDownloads       *int       `json:"nb_downloads,omitempty"` // Number of downloads left for this particular sharing
	Access            AccessType `json:"access"`                 // Kind of access user has on this repository
	ShareAccess       AccessType `json:"share_access"`           // Kind of share access user has on this repository
}

package api

type CommandGetUsersOutput struct {
	Users []Account `json:"users"`
}
type CommandDummyAuthInput struct {
	Login        string `json:"login"`
	ChallengHash string `json:"challeng_hash"`
	Ref          string `json:"ref"`
}
type CommandDummyAuthOutput struct {
	MySelf               Account `json:"my_self"`
	AuthenticationHeader string  `json:"authentication_header"`
}
type CommandDummyGetChallengeOutput struct {
	Challenge string `json:"challenge"`
	Ref       string `json:"ref"`
}
type CommandGetInfoOutput struct {
	ShareLink         bool       `json:"share_link"`
	PasswordProtected bool       `json:"password_protected"`
	NbDownloads       *int       `json:"nb_downloads,omitempty"` // Number of downloads left for this particular sharing
	Access            AccessType `json:"access"`                 // Kind of access user has on this repository
	ShareAccess       AccessType `json:"share_access"`           // Kind of share access user has on this repository
}

package api

type StorageItem struct {
	Name        string     `json:"name"`
	IsDir       bool       `json:"isDir"`
	MDate       int64      `json:"mDate"`
	Size        int64      `json:"size"`
	Kind        string     `json:"kind"` // this is the extension of the file. value will be folder for a folder
	Access      AccessType `json:"access"`
	ShareAccess AccessType `json:"share_access"`
}
type CommandStatus struct {
	Status    EnumStatus           `json:"status"`
	Progress  int                  `json:"progress"`
	ErrorCode EnumCommandErrorCode `json:"error_code"`
}
type CommandsSearchParameters struct {
	Status *EnumStatus `json:"status,omitempty"`
}
type ShareLink struct {
	Name     *string           `json:"name,omitempty"`      // Name used for displaying the share link if multiple share links available
	Path     *string           `json:"path,omitempty"`      // Can be empty only if ShareLinkKey is provided
	Key      *string           `json:"key,omitempty"`       // Can be empty only for a creation or on a Get
	User     string            `json:"user"`                // Can be empty only for a creation or on a Get
	UserList *[]string         `json:"user_list,omitempty"` // This is only available for EnumRestricted mode
	Type     EnumShareLinkType `json:"type"`
	Access   *AccessType       `json:"access,omitempty"` // What access would people coming with this link have
}

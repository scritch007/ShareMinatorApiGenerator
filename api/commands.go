package api

type CommandBrowserListInput struct {
	Path            string `json:"path"`
	ShowHiddenFiles *bool  `json:"show_hidden_files,omitempty"` // By default hidden files are hidden
}
type CommandBrowserListOutput struct {
	CurrentItem StorageItem   `json:"current_item"`
	Children    []StorageItem `json:"children"`
}
type CommandBrowserList struct {
	Input  CommandBrowserListInput  `json:"input" bson:"input"`
	Output CommandBrowserListOutput `json:"output" bson:"output"`
}
type CommandBrowserCreateFolderInput struct {
	Path string `json:"path"`
}
type CommandBrowserCreateFolderOutput struct {
	Result StorageItem `json:"result"`
}
type CommandBrowserCreateFolder struct {
	Input  CommandBrowserCreateFolderInput  `json:"input" bson:"input"`
	Output CommandBrowserCreateFolderOutput `json:"output" bson:"output"`
}
type CommandBrowserDeleteInput struct {
	Path string `json:"path"`
}
type CommandBrowserDelete struct {
	Input CommandBrowserDeleteInput `json:"input" bson:"input"`
}
type CommandBrowserDownloadLinkInput struct {
	Path string `json:"path"`
}
type CommandBrowserDownloadLinkOutput struct {
	DownloadLink string `json:"download_link"`
}
type CommandBrowserDownloadLink struct {
	Input  CommandBrowserDownloadLinkInput  `json:"input" bson:"input"`
	Output CommandBrowserDownloadLinkOutput `json:"output" bson:"output"`
}
type CommandBrowserUploadFileInput struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
}
type CommandBrowserUploadFile struct {
	Input CommandBrowserUploadFileInput `json:"input" bson:"input"`
}
type CommandBrowserThumbnailInput struct {
	Path string `json:"path"`
	Size *int   `json:"size,omitempty"` // size of desired picture
}
type CommandBrowserThumbnailOutput struct {
	Content string `json:"content"` // base64 of the image
}
type CommandBrowserThumbnail struct {
	Input  CommandBrowserThumbnailInput  `json:"input" bson:"input"`
	Output CommandBrowserThumbnailOutput `json:"output" bson:"output"`
}
type BrowserCommand struct {
	List         *CommandBrowserList         `json:"list,omitempty" bson:"list,omitempty"`
	CreateFolder *CommandBrowserCreateFolder `json:"create_folder,omitempty" bson:"create_folder,omitempty"`
	Delete       *CommandBrowserDelete       `json:"delete,omitempty" bson:"delete,omitempty"`
	DownloadLink *CommandBrowserDownloadLink `json:"download_link,omitempty" bson:"download_link,omitempty"`
	UploadFile   *CommandBrowserUploadFile   `json:"upload_file,omitempty" bson:"upload_file,omitempty"`
	Thumbnail    *CommandBrowserThumbnail    `json:"thumbnail,omitempty" bson:"thumbnail,omitempty"`
}
type CommandShareLinkListInput struct {
	Path string `json:"path"`
}
type CommandShareLinkListOutput struct {
	ShareLinks []ShareLink `json:"share_links"`
}
type CommandShareLinkList struct {
	Input  CommandShareLinkListInput  `json:"input" bson:"input"`
	Output CommandShareLinkListOutput `json:"output" bson:"output"`
}
type CommandShareLinkCreateInput struct {
	ShareLink ShareLink `json:"share_link"`
}
type CommandShareLinkCreateOutput struct {
	ShareLink ShareLink `json:"share_link"`
}
type CommandShareLinkCreate struct {
	Input  CommandShareLinkCreateInput  `json:"input" bson:"input"`
	Output CommandShareLinkCreateOutput `json:"output" bson:"output"`
}
type CommandShareLinkUpdateInput struct {
	ShareLink ShareLink `json:"share_link"`
}
type CommandShareLinkUpdateOutput struct {
	ShareLink ShareLink `json:"share_link"`
}
type CommandShareLinkUpdate struct {
	Input  CommandShareLinkUpdateInput  `json:"input" bson:"input"`
	Output CommandShareLinkUpdateOutput `json:"output" bson:"output"`
}
type CommandShareLinkDeleteInput struct {
	Key string `json:"key"`
}
type CommandShareLinkDelete struct {
	Input CommandShareLinkDeleteInput `json:"input" bson:"input"`
}
type ShareLinkCommand struct {
	List   *CommandShareLinkList   `json:"list,omitempty" bson:"list,omitempty"`
	Create *CommandShareLinkCreate `json:"create,omitempty" bson:"create,omitempty"`
	Update *CommandShareLinkUpdate `json:"update,omitempty" bson:"update,omitempty"`
	Delete *CommandShareLinkDelete `json:"delete,omitempty" bson:"delete,omitempty"`
}
type EnumAction string

const (
	EnumBrowserList         EnumAction = "browser.list"
	EnumBrowserCreateFolder EnumAction = "browser.create_folder"
	EnumBrowserDelete       EnumAction = "browser.delete"
	EnumBrowserDownloadLink EnumAction = "browser.download_link"
	EnumBrowserUploadFile   EnumAction = "browser.upload_file"
	EnumBrowserThumbnail    EnumAction = "browser.thumbnail"
	EnumShareLinkList       EnumAction = "share_link.list"
	EnumShareLinkCreate     EnumAction = "share_link.create"
	EnumShareLinkUpdate     EnumAction = "share_link.update"
	EnumShareLinkDelete     EnumAction = "share_link.delete"
)

type Command struct {
	Name      EnumAction        `json:"name"`
	CommandId string            `json:"command_id"`
	Timeout   int64             `json:"timeout,omitempty"`
	AuthKey   *string           `json:"auth_key,omitempty"` //Used when calling commands on behalf of a sharedlink
	Password  *string           `json:"password"`           //Used when a share_link requires a password This should be the hash of AuthKey + Password
	State     CommandStatus     `json:"state"`
	Browser   *BrowserCommand   `json:"browser,omitempty"`
	ShareLink *ShareLinkCommand `json:"share_link,omitempty"`
}

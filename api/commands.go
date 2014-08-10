package api
type CommandBrowserListInput struct{
	Path string `json:"path"` 
}
type CommandBrowserListOutput struct{
	CurrentItem StorageItem `json:"current_item"` 
	Children []StorageItem `json:"children"` 
}
type CommandBrowserList struct{
	Input CommandBrowserListInput `json:"input"`
	Output CommandBrowserListOutput `json:"output"`
}
type CommandBrowserCreateFolderInput struct{
	Path string `json:"path"` 
}
type CommandBrowserCreateFolderOutput struct{
	Result StorageItem `json:"result"` 
}
type CommandBrowserCreateFolder struct{
	Input CommandBrowserCreateFolderInput `json:"input"`
	Output CommandBrowserCreateFolderOutput `json:"output"`
}
type CommandBrowserDeleteInput struct{
	Path string `json:"path"` 
}
type CommandBrowserDelete struct{
	Input CommandBrowserDeleteInput `json:"input"`
}
type CommandBrowserDownloadLinkInput struct{
	Path string `json:"path"` 
}
type CommandBrowserDownloadLinkOutput struct{
	DownloadLink string `json:"download_link"` 
}
type CommandBrowserDownloadLink struct{
	Input CommandBrowserDownloadLinkInput `json:"input"`
	Output CommandBrowserDownloadLinkOutput `json:"output"`
}
type CommandBrowserUploadFileInput struct{
	Path string `json:"path"` 
	Size int64 `json:"size"` 
}
type CommandBrowserUploadFile struct{
	Input CommandBrowserUploadFileInput `json:"input"`
}
type BrowserCommand struct{
	List *CommandBrowserList `json:"list,omitempty"`
	CreateFolder *CommandBrowserCreateFolder `json:"create_folder,omitempty"`
	Delete *CommandBrowserDelete `json:"delete,omitempty"`
	DownloadLink *CommandBrowserDownloadLink `json:"download_link,omitempty"`
	UploadFile *CommandBrowserUploadFile `json:"upload_file,omitempty"`
}
type CommandShareLinkListInput struct{
	Path string `json:"path"` 
}
type CommandShareLinkListOutput struct{
	ShareLinks []ShareLink `json:"share_links"` 
}
type CommandShareLinkList struct{
	Input CommandShareLinkListInput `json:"input"`
	Output CommandShareLinkListOutput `json:"output"`
}
type CommandShareLinkCreateInput struct{
	ShareLink ShareLink `json:"share_link"` 
}
type CommandShareLinkCreateOutput struct{
	ShareLink ShareLink `json:"share_link"` 
}
type CommandShareLinkCreate struct{
	Input CommandShareLinkCreateInput `json:"input"`
	Output CommandShareLinkCreateOutput `json:"output"`
}
type CommandShareLinkUpdateInput struct{
	ShareLink ShareLink `json:"share_link"` 
}
type CommandShareLinkUpdateOutput struct{
	ShareLink ShareLink `json:"share_link"` 
}
type CommandShareLinkUpdate struct{
	Input CommandShareLinkUpdateInput `json:"input"`
	Output CommandShareLinkUpdateOutput `json:"output"`
}
type CommandShareLinkDeleteInput struct{
	Key string `json:"key"` 
}
type CommandShareLinkDelete struct{
	Input CommandShareLinkDeleteInput `json:"input"`
}
type ShareLinkCommand struct{
	List *CommandShareLinkList `json:"list,omitempty"`
	Create *CommandShareLinkCreate `json:"create,omitempty"`
	Update *CommandShareLinkUpdate `json:"update,omitempty"`
	Delete *CommandShareLinkDelete `json:"delete,omitempty"`
}
type EnumAction string
const (
	EnumBrowserList EnumAction = "browser.list"
	EnumBrowserCreateFolder EnumAction = "browser.create_folder"
	EnumBrowserDelete EnumAction = "browser.delete"
	EnumBrowserDownloadLink EnumAction = "browser.download_link"
	EnumBrowserUploadFile EnumAction = "browser.upload_file"
	EnumShareLinkList EnumAction = "share_link.list"
	EnumShareLinkCreate EnumAction = "share_link.create"
	EnumShareLinkUpdate EnumAction = "share_link.update"
	EnumShareLinkDelete EnumAction = "share_link.delete"
)
type Command struct{
	Name EnumAction `json:"name"`
	CommandId string `json:"command_id"`
	Timeout int64 `json:"timeout,omitempty"`
	AuthKey *string `json:"auth_key,omitempty"` //Used when calling commands on behalf of a sharedlink
	State CommandStatus `json:"state"`
	Browser *BrowserCommand `json:"browser,omitempty"`
	ShareLink *ShareLinkCommand `json:"share_link,omitempty"`
}

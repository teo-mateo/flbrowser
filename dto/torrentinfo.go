package dto

// FLTorrentInfo - torrent information
type FLTorrentInfo struct {
	ID int `json:"id"`
	Name  string `json:"name"`
	DlURL string `json:"dlurl"`
	DateAded string `json:"dateadded"`
	Size string `json:"size"`
	TimesDownloaded string `json:"timesdownloaded"`
	Seeders string `json:"seeders"`
	Leechers string `json:"leechers"`
}

type RTRTorrentInfo struct{
	ID string `json:"id"`
	Name string `json:"id"`
	CreationDate int `json:"id"`
	IsOpen bool `json:"isopen"`
	IsActive bool `json:"isactive"`
	IsHashChecked bool `json:"ishashchecked"`
	IsHashChecking bool `json:"ishashchecking"`
	IsMultiFile bool `json:"ismultifile"`
	DownTotal int `json:"downtotal"`
	UpTotal int `json:"uptotal"`
	Directory string `json:"directory"`
	CompletedBytes int `json:"completedbytes"`
	LeftBytes int `json:"leftbytes"`
	SizeFiles int `json:"sizefiles"`
	SizeBytes int `json:"sizebytes"`
}
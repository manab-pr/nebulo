package entities

type StorageSummary struct {
	TotalDevices     int   `json:"total_devices"`
	OnlineDevices    int   `json:"online_devices"`
	OfflineDevices   int   `json:"offline_devices"`
	TotalStorage     int64 `json:"total_storage"`
	UsedStorage      int64 `json:"used_storage"`
	AvailableStorage int64 `json:"available_storage"`
	TotalFiles       int   `json:"total_files"`
}

type DeviceStorageInfo struct {
	DeviceID         string `json:"device_id"`
	DeviceName       string `json:"device_name"`
	TotalStorage     int64  `json:"total_storage"`
	UsedStorage      int64  `json:"used_storage"`
	AvailableStorage int64  `json:"available_storage"`
	FileCount        int    `json:"file_count"`
	Status           string `json:"status"`
}

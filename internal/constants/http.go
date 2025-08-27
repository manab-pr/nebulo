package constants

const (
	// Auth routes
	AuthBaseRoute                   = "/auth"
	RegisterRoute                   = "/register"
	LoginRoute                      = "/login"
	VerifyOTPRoute                  = "/verify-otp"

	// User profile routes
	ProfileRoute                    = "/profile"
)

const (
	TransferBaseRoute               = "/transfers"
	GetPendingTransfersRoute        = "/pending/:deviceId"
	CompleteTransferRoute           = "/complete"
	CancelTransferRoute             = "/:id"
)

const (
	FileBaseRoute                   = "/files"
	StoreFileRoute                  = "/store"
	GetFileRoute                    = "/:fileId"
	GetAllFilesRoute                = ""
	DeleteFileRoute                 = "/:fileId"
)

const (
	DeviceBaseRoute                 = "/devices"
	RegisterDeviceRoute             = "/register"
	HeartbeatRoute                  = "/heartbeat"
	GetDevicesRoute                 = ""
	DeleteDeviceRoute               = "/:id"
)

const (
	StorageBaseRoute                = "/storage"
	GetStorageSummaryRoute          = "/summary"
	GetDeviceStorageRoute           = "/device/:deviceId"
)

const (
	SearchBaseRoute                 = "/files"
	SearchFilesRoute                = "/search"
	GetFileLocationRoute            = "/location/:fileId"
)


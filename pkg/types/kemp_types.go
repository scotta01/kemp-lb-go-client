package types

// VirtualServices is a list Kemp Virtual Service
type VirtualServices struct {
	Code   int              `json:"code"`
	VS     []VirtualService `json:"VS"`
	Status string           `json:"status"`
}

// VirtualService is a Kemp Virtual Service
type VirtualService struct {
	Status                  string       `json:"Status"`
	Index                   int          `json:"Index"`
	VSAddress               string       `json:"VSAddress"`
	VSPort                  string       `json:"VSPort"`
	Layer                   int          `json:"Layer"`
	NickName                string       `json:"NickName"`
	Enable                  bool         `json:"Enable"`
	SSLReverse              bool         `json:"SSLReverse"`
	SSLReencrypt            bool         `json:"SSLReencrypt"`
	InterceptMode           int          `json:"InterceptMode"`
	Intercept               bool         `json:"Intercept"`
	InterceptOpts           []string     `json:"InterceptOpts"`
	AlertThreshold          int          `json:"AlertThreshold"`
	OwaspOpts               []string     `json:"OwaspOpts"`
	BlockingParanoia        int          `json:"BlockingParanoia"`
	IPReputationBlocking    bool         `json:"IPReputationBlocking"`
	ExecutingParanoia       int          `json:"ExecutingParanoia"`
	AnomalyScoringThreshold int          `json:"AnomalyScoringThreshold"`
	PCRELimit               int          `json:"PCRELimit"`
	JSONDLimit              int          `json:"JSONDLimit"`
	BodyLimit               int          `json:"BodyLimit"`
	Transactionlimit        int          `json:"Transactionlimit"`
	Transparent             bool         `json:"Transparent"`
	SubnetOriginating       bool         `json:"SubnetOriginating"`
	ServerInit              int          `json:"ServerInit"`
	StartTLSMode            int          `json:"StartTLSMode"`
	Idletime                int          `json:"Idletime"`
	Cache                   bool         `json:"Cache"`
	Compress                bool         `json:"Compress"`
	Verify                  int          `json:"Verify"`
	UseforSnat              bool         `json:"UseforSnat"`
	ForceL4                 bool         `json:"ForceL4"`
	ForceL7                 bool         `json:"ForceL7"`
	MultiConnect            bool         `json:"MultiConnect"`
	ClientCert              int          `json:"ClientCert"`
	SecurityHeaderOptions   int          `json:"SecurityHeaderOptions"`
	SameSite                int          `json:"SameSite"`
	VerifyBearer            bool         `json:"VerifyBearer"`
	ErrorCode               string       `json:"ErrorCode"`
	CheckUse11              bool         `json:"CheckUse1.1"`
	MatchLen                int          `json:"MatchLen"`
	CheckUseGet             int          `json:"CheckUseGet"`
	SSLRewrite              string       `json:"SSLRewrite"`
	VStype                  string       `json:"VStype"`
	FollowVSID              int          `json:"FollowVSID"`
	Protocol                string       `json:"Protocol"`
	Schedule                string       `json:"Schedule"`
	CheckType               string       `json:"CheckType"`
	PersistTimeout          string       `json:"PersistTimeout"`
	CheckPort               string       `json:"CheckPort"`
	HTTPReschedule          bool         `json:"HTTPReschedule"`
	NRules                  int          `json:"NRules"`
	NRequestRules           int          `json:"NRequestRules"`
	NResponseRules          int          `json:"NResponseRules"`
	NMatchBodyRules         int          `json:"NMatchBodyRules"`
	NPreProcessRules        int          `json:"NPreProcessRules"`
	EspEnabled              bool         `json:"EspEnabled"`
	InputAuthMode           int          `json:"InputAuthMode"`
	OutputAuthMode          int          `json:"OutputAuthMode"`
	MasterVS                int          `json:"MasterVS"`
	MasterVSID              int          `json:"MasterVSID"`
	IsTransparent           int          `json:"IsTransparent"`
	AddVia                  int          `json:"AddVia"`
	QoS                     int          `json:"QoS"`
	TlsType                 string       `json:"TlsType"`
	NeedHostName            bool         `json:"NeedHostName"`
	OCSPVerify              bool         `json:"OCSPVerify"`
	AllowHTTP2              bool         `json:"AllowHTTP2"`
	PassCipher              bool         `json:"PassCipher"`
	PassSni                 bool         `json:"PassSni"`
	ChkInterval             int          `json:"ChkInterval"`
	ChkTimeout              int          `json:"ChkTimeout"`
	ChkRetryCount           int          `json:"ChkRetryCount"`
	Bandwidth               int          `json:"Bandwidth"`
	ConnsPerSecLimit        int          `json:"ConnsPerSecLimit"`
	RequestsPerSecLimit     int          `json:"RequestsPerSecLimit"`
	MaxConnsLimit           int          `json:"MaxConnsLimit"`
	RefreshPersist          bool         `json:"RefreshPersist"`
	ResponseStatusRemap     bool         `json:"ResponseStatusRemap"`
	ResponseRemapMsgFormat  int          `json:"ResponseRemapMsgFormat"`
	EnhancedHealthChecks    bool         `json:"EnhancedHealthChecks"`
	RsMinimum               int          `json:"RsMinimum"`
	NumberOfRSs             int          `json:"NumberOfRSs"`
	Rs                      []RealServer `json:"Rs,omitempty"`
}

// SubVS is a Kemp SubVS
type SubVS struct {
	SubVS []struct {
		Status    string `json:"Status"`
		VSIndex   int    `json:"VSIndex"`
		RsIndex   int    `json:"RsIndex"`
		Name      string `json:"Name"`
		Forward   string `json:"Forward"`
		Weight    int    `json:"Weight"`
		Limit     int    `json:"Limit"`
		RateLimit int    `json:"RateLimit"`
		Follow    int    `json:"Follow"`
		Enable    bool   `json:"Enable"`
		Critical  bool   `json:"Critical"`
	} `json:"SubVS,omitempty"`
}

// RealServers is a list of Kemp Real (Backend) Servers
type RealServers struct {
	Code   int          `json:"code"`
	Rs     []RealServer `json:"Rs"`
	Status string       `json:"status"`
}

// RealServer is a Kemp Real (Backend) Server
type RealServer struct {
	Status    string `json:"Status"`
	VSIndex   int    `json:"VSIndex"`
	RsIndex   int    `json:"RsIndex"`
	Addr      string `json:"Addr"`
	Port      int    `json:"Port"`
	DnsName   string `json:"DnsName"`
	Forward   string `json:"Forward"`
	Weight    int    `json:"Weight"`
	Limit     int    `json:"Limit"`
	RateLimit int    `json:"RateLimit"`
	Follow    int    `json:"Follow"`
	Enable    bool   `json:"Enable"`
	Critical  bool   `json:"Critical"`
	Nrules    int    `json:"Nrules"`
}

// Interfaces is a list of Kemp Ethernet Interfaces
type Interfaces struct {
	Code      int         `json:"code"`
	Interface []Interface `json:"Interface"`
	Status    string      `json:"status"`
}

// Interface is a Kemp Ethernet Interface
type Interface struct {
	Id               int    `json:"Id"`
	IPAddress        string `json:"IPAddress"`
	Mtu              string `json:"Mtu"`
	InterfaceType    string `json:"InterfaceType"`
	GeoTrafficEnable bool   `json:"GeoTrafficEnable"`
	DefaultInterface bool   `json:"DefaultInterface"`
}

// Stats is a Kemp Stats including a list of Ethernet Interface IDs
type Stats struct {
	Code int `json:"code"`
	CPU  struct {
		Total struct {
			User      int `json:"User"`
			System    int `json:"System"`
			Idle      int `json:"Idle"`
			IOWaiting int `json:"IOWaiting"`
		} `json:"total"`
		Cpu []struct {
			Cpu          int `json:"cpu"`
			User         int `json:"User"`
			System       int `json:"System"`
			HWInterrupts int `json:"HWInterrupts"`
			SWInterrupts int `json:"SWInterrupts"`
			Idle         int `json:"Idle"`
			IOWaiting    int `json:"IOWaiting"`
		} `json:"cpu"`
	} `json:"CPU"`
	Memory struct {
		MBtotal        int `json:"MBtotal"`
		Memused        int `json:"memused"`
		MBused         int `json:"MBused"`
		Percentmemused int `json:"percentmemused"`
		Memfree        int `json:"memfree"`
		MBfree         int `json:"MBfree"`
		Percentmemfree int `json:"percentmemfree"`
	} `json:"Memory"`
	Network struct {
		Interface []struct {
			Name          string  `json:"Name"`
			IfaceID       int     `json:"ifaceID"`
			Speed         int     `json:"speed"`
			In            float32 `json:"in"`
			Inbytes       int     `json:"inbytes"`
			InbytesTotal  int     `json:"inbytesTotal"`
			Out           float32 `json:"out"`
			Outbytes      int     `json:"outbytes"`
			OutbytesTotal int     `json:"outbytesTotal"`
		} `json:"Interface"`
	} `json:"Network"`
	DiskUsage struct {
		Partition []struct {
			Name        string  `json:"name"`
			GBtotal     float64 `json:"GBtotal"`
			GBused      float64 `json:"GBused"`
			Percentused int     `json:"percentused"`
			GBfree      float64 `json:"GBfree"`
			Percentfree int     `json:"percentfree"`
		} `json:"partition"`
	} `json:"DiskUsage"`
	ClientLimits struct {
		Totals struct {
			CurrentConnections            int `json:"CurrentConnections"`
			CurrentConnectionsBlocked     int `json:"CurrentConnectionsBlocked"`
			SuccessfulConnectionAttempts  int `json:"SuccessfulConnectionAttempts"`
			SuccessfulRequestAttempts     int `json:"SuccessfulRequestAttempts"`
			SuccessfulMatchedRuleAttempts int `json:"SuccessfulMatchedRuleAttempts"`
			ConnectionAttemptsBlocked     int `json:"ConnectionAttemptsBlocked"`
			RequestAttemptsBlocked        int `json:"RequestAttemptsBlocked"`
			MatchedRulesBlocked           int `json:"MatchedRulesBlocked"`
		} `json:"Totals"`
	} `json:"ClientLimits"`
	VStotals struct {
		ConnsPerSec  int `json:"ConnsPerSec"`
		TotalConns   int `json:"TotalConns"`
		BitsPerSec   int `json:"BitsPerSec"`
		TotalBits    int `json:"TotalBits"`
		BytesPerSec  int `json:"BytesPerSec"`
		TotalBytes   int `json:"TotalBytes"`
		PktsPerSec   int `json:"PktsPerSec"`
		TotalPackets int `json:"TotalPackets"`
	} `json:"VStotals"`
	Vs []struct {
		VSAddress           string `json:"VSAddress"`
		VSPort              int    `json:"VSPort"`
		VSProt              string `json:"VSProt"`
		Index               int    `json:"Index"`
		Status              string `json:"Status"`
		ErrorCode           int    `json:"ErrorCode"`
		Enable              int    `json:"Enable"`
		TotalConns          int    `json:"TotalConns"`
		TotalPkts           int    `json:"TotalPkts"`
		TotalBytes          int    `json:"TotalBytes"`
		TotalBits           int    `json:"TotalBits"`
		ActiveConns         int    `json:"ActiveConns"`
		BytesRead           int    `json:"BytesRead"`
		BytesWritten        int    `json:"BytesWritten"`
		ConnsPerSec         int    `json:"ConnsPerSec"`
		ConnsRateBlocked    int    `json:"ConnsRateBlocked"`
		RequestsRateBlocked int    `json:"RequestsRateBlocked"`
		MaxConnsBlocked     int    `json:"MaxConnsBlocked"`
		WafEnable           int    `json:"WafEnable"`
		InterceptMode       int    `json:"InterceptMode"`
	} `json:"Vs"`
	TPS struct {
		Total int `json:"Total"`
		SSL   int `json:"SSL"`
	} `json:"TPS"`
	Timestamp struct {
		Sec    int `json:"Sec"`
		Usec   int `json:"Usec"`
		Period int `json:"Period"`
	} `json:"Timestamp"`
	ChangeTime int    `json:"ChangeTime"`
	Status     string `json:"status"`
}

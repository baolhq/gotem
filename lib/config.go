package lib

type Config struct {
	Backup   bool               `json:"backup"`
	Create   bool               `json:"create"`
	Dotpath  string             `json:"dotpath"`
	Uselink  bool               `json:"uselink"`
	Keepdot  bool               `json:"keepdot"`
	Profiles map[string]Profile `json:"profiles"`
}

type Profile struct {
	Backup      *bool              `json:"backup,omitempty"`
	Create      *bool              `json:"create,omitempty"`
	Dotpath     *string            `json:"dotpath,omitempty"`
	Uselink     *bool              `json:"uselink,omitempty"`
	Keepdot     *bool              `json:"keepdot,omitempty"`
	Directories map[string]Entry   `json:"directories,omitempty"`
	Files       map[string]Entry   `json:"files,omitempty"`
}

type Entry struct {
	Dst string `json:"dst"`
	Src string `json:"src"`
}

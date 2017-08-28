package pcmiler

type Error struct {
	Type            string `json:"string"`
	Code            int    `json:"code"`
	LegacyErrorCode int    `json:"LegacyErrorCode,omitempty"`
	Description     string `json:"description"`
}

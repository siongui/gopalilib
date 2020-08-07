package lib

// PaliSetting is common structure for the setting of Pāli Dictionary and
// Tipiṭaka
type PaliSetting struct {
	IsShowWordPreview bool   `json:"isPreview"`
	P2en              bool   `json:"p2en"`
	P2ja              bool   `json:"p2ja"`
	P2zh              bool   `json:"p2zh"`
	P2vi              bool   `json:"p2vi"`
	P2my              bool   `json:"p2my"`
	DicLangOrder      string `json:"dicLangOrder"`
}

func GetDefaultPaliSetting() PaliSetting {
	return PaliSetting{
		IsShowWordPreview: false,
		P2en:              true,
		P2ja:              true,
		P2zh:              true,
		P2vi:              true,
		P2my:              true,
		DicLangOrder:      "hdr",
	}
}

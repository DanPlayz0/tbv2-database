package database

type Language string

const (
	English   Language = "en"
	French    Language = "fr"
	Spanish   Language = "es"
	German    Language = "de"
	Dutch     Language = "nl"
	Polish    Language = "pl"
	Norwegian Language = "no"
	Turkish   Language = "tr"
	Swedish   Language = "sv"
	Arabic    Language = "ar"
	Hungarian Language = "hu"
)

var Flags = map[Language]string{
	English:   "🇬🇧",
	French:    "🇫🇷",
	Spanish:   "🇪🇸",
	German:    "🇩🇪",
	Dutch:     "🇳🇱",
	Polish:    "🇵🇱",
	Norwegian: "🇳🇴",
	Turkish:   "🇹🇷",
	Swedish:   "🇸🇪",
	Arabic:    "🇸🇦",
	Hungarian: "🇭🇺",
}

// https://discord.com/developers/docs/dispatch/field-values
var Locales = map[string]Language{
	"en-US": English,
	"en-GB": English,
	"zh-CN": English, // Chinese (China)
	"zh-TW": English, // Chinese (Taiwan)
	"cs":    English, // Czech
	"da":    English, // Danish
	"nl":    Dutch,
	"fr":    French,
	"de":    German,
	"hu":    Hungarian,
	"it":    English,   // Italian
	"ja":    English,   // Japanese
	"ko":    English,   // Korean
	"no":    Norwegian, // Norwegian
	"pl":    Polish,
	"pt-BR": English, // Portuguese (Brazil)
	"ru":    English, // Russian
	"es-ES": Spanish,
	"sv-SE": Swedish,
	"tr":    Turkish,
	"bg":    English, // Bulgarian
	"uk":    English, // Ukrainian
	"fi":    English, // Finnish
	"hr":    English, // Croatian
	"ro":    English, // Romanian
	"lt":    English, // Lithuanian
}

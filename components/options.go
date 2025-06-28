package rephtml

type Options struct {
	AllowAudio           bool         // allow audio
	AllowEmbeddedContent bool         // allow embedded content
	AllowImages          bool         // allow images
	AllowInteractive     bool         // allow interactive elements
	AllowMath            bool         // allow math
	AllowScripts         bool         // allow scripting
	AllowSvg             bool         // allow svg
	AllowVideo           bool         // allow video
	CheckIds             bool         // id validation strictness
	Constitution         Constitution // element checks strictness
}

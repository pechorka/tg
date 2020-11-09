package tg

type ParseMod int

const (
	ParseModDefault ParseMod = ""
	ParseModMDV2    ParseMod = "MarkdownV2"
	ParseModMD      ParseMod = "Markdown"
	ParseModHTML    ParseMod = "HTML"
)

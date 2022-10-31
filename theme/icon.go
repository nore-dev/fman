package theme

import (
	"github.com/nore-dev/fman/args"
)

type iconSet struct {
	LeftArrowIcon  rune
	RightArrowIcon rune
	UpArrowIcon    rune

	BreadcrumbArrowIcon rune

	GopherIcon rune

	FileIcon    rune
	FolderIcon  rune
	SymlinkIcon rune

	TimeIcon rune
	SizeIcon rune
	NameIcon rune
}

type iconSets map[string]iconSet

var nerdFont = iconSet{
	LeftArrowIcon:       '\uf060',
	RightArrowIcon:      '\uf061',
	UpArrowIcon:         '\uf062',
	BreadcrumbArrowIcon: '\uf0a4',
	GopherIcon:          '\ue627',
	FileIcon:            '\uf15c',
	FolderIcon:          '\uf07b',
	SymlinkIcon:         '\uf838',
	TimeIcon:            '\uf017',
	SizeIcon:            '\uf200',
	NameIcon:            '\ue612',
}

var emoji = iconSet{
	LeftArrowIcon:       '◀',
	RightArrowIcon:      '▶',
	UpArrowIcon:         '🔼',
	BreadcrumbArrowIcon: '👉',
	GopherIcon:          '🐻',
	FileIcon:            '📄',
	FolderIcon:          '📁',
	SymlinkIcon:         '🔗',
	TimeIcon:            '⏰',
	SizeIcon:            '📊',
	NameIcon:            '🏷',
}

var noIcons = iconSet{
	LeftArrowIcon:       '<',
	RightArrowIcon:      '>',
	UpArrowIcon:         '^',
	BreadcrumbArrowIcon: '>',
}

var iconProviders = iconSets{
	"emoji":    emoji,
	"nerdfont": nerdFont,
	"none":     noIcons,
}

func GetActiveIconTheme() iconSet {
	return iconProviders[args.CommandLine.Icons]
}

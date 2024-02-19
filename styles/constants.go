package styles

const (
	// Layout Properties
	Display             = "display"
	Position            = "position"
	Top                 = "top"
	Right               = "right"
	Bottom              = "bottom"
	Left                = "left"
	Overflow            = "overflow"
	OverflowX           = "overflow-x"
	OverflowY           = "overflow-y"
	ZIndex              = "z-index"
	Flex                = "flex"
	FlexDirection       = "flex-direction"
	JustifyContent      = "justify-content"
	AlignItems          = "align-items"
	AlignContent        = "align-content"
	AlignSelf           = "align-self"
	FlexWrap            = "flex-wrap"
	FlexGrow            = "flex-grow"
	FlexShrink          = "flex-shrink"
	FlexBasis           = "flex-basis"
	Grid                = "grid"
	GridArea            = "grid-area"
	GridAutoColumns     = "grid-auto-columns"
	GridAutoFlow        = "grid-auto-flow"
	GridAutoRows        = "grid-auto-rows"
	GridColumn          = "grid-column"
	GridColumnEnd       = "grid-column-end"
	ColumnGap           = "column-gap"
	GridColumnStart     = "grid-column-start"
	GridRow             = "grid-row"
	Gap                 = "gap"
	GridRowEnd          = "grid-row-end"
	RowGap              = "row-gap"
	GridRowStart        = "grid-row-start"
	GridTemplate        = "grid-template"
	GridTemplateAreas   = "grid-template-areas"
	GridTemplateColumns = "grid-template-columns"
	GridTemplateRows    = "grid-template-rows"

	// Box Model Properties
	Width             = "width"
	MinWidth          = "min-width"
	MaxWidth          = "max-width"
	Height            = "height"
	MinHeight         = "min-height"
	MaxHeight         = "max-height"
	Padding           = "padding"
	PaddingTop        = "padding-top"
	PaddingRight      = "padding-right"
	PaddingBottom     = "padding-bottom"
	PaddingLeft       = "padding-left"
	Margin            = "margin"
	MarginTop         = "margin-top"
	MarginRight       = "margin-right"
	MarginBottom      = "margin-bottom"
	MarginLeft        = "margin-left"
	Border            = "border"
	BorderTop         = "border-top"
	BorderRight       = "border-right"
	BorderBottom      = "border-bottom"
	BorderLeft        = "border-left"
	BorderColor       = "border-color"
	BorderTopColor    = "border-top-color"
	BorderRightColor  = "border-right-color"
	BorderBottomColor = "border-bottom-color"
	BorderLeftColor   = "border-left-color"
	BorderStyle       = "border-style"
	BorderTopStyle    = "border-top-style"
	BorderRightStyle  = "border-right-style"
	BorderBottomStyle = "border-bottom-style"
	BorderLeftStyle   = "border-left-style"
	BorderWidth       = "border-width"
	BorderTopWidth    = "border-top-width"
	BorderRightWidth  = "border-right-width"
	BorderBottomWidth = "border-bottom-width"
	BorderLeftWidth   = "border-left-width"
	BorderRadius      = "border-radius"
	BoxShadow         = "box-shadow"
	Outline           = "outline"
	OutlineStyle      = "outline-style"
	OutlineColor      = "outline-color"
	OutlineWidth      = "outline-width"
	OutlineOffset     = "outline-offset"

	// Fonts & Text Properties
	Color          = "color"
	Font           = "font"
	FontFamily     = "font-family"
	FontSize       = "font-size"
	FontWeight     = "font-weight"
	LineHeight     = "line-height"
	TextAlign      = "text-align"
	TextDecoration = "text-decoration"
	TextTransform  = "text-transform"
	LetterSpacing  = "letter-spacing"
	WhiteSpace     = "white-space"
	TextOverflow   = "text-overflow"
	FontStyle      = "font-style"
	TextShadow     = "text-shadow"
	VerticalAlign  = "vertical-align"
	WordSpacing    = "word-spacing"
	WordBreak      = "word-break"
	TextIndent     = "text-indent"

	// Visual Properties
	BackgroundColor      = "background-color"
	Background           = "background"
	BackgroundImage      = "background-image"
	BackgroundRepeat     = "background-repeat"
	BackgroundSize       = "background-size"
	ObjectFit            = "object-fit"
	Opacity              = "opacity"
	BoxSizing            = "box-sizing"
	Cursor               = "cursor"
	Transition           = "transition"
	Transform            = "transform"
	BackgroundPosition   = "background-position"
	BackgroundAttachment = "background-attachment"
	BackgroundBlendMode  = "background-blend-mode"
	BackfaceVisibility   = "backface-visibility"
	BackdropFilter       = "backdrop-filter"
	Perspective          = "perspective"
	TransformOrigin      = "transform-origin"

	// List properties
	ListStyle     = "list-style"
	ListStyleType = "list-style-type"

	// Table Properties
	BorderCollapse = "border-collapse"
	BorderSpacing  = "border-spacing"
	TableLayout    = "table-layout"
	CaptionSide    = "caption-side"

	// Resource Attributes
	Integrity   = "integrity"
	Crossorigin = "crossorigin"

	// Other Properties
	Visibility    = "visibility"
	Clip          = "clip"
	Content       = "content"
	Filter        = "filter"
	PointerEvents = "pointer-events"
	Resize        = "resize"
	UserSelect    = "user-select"

	// Interaction pseudo-classes
	PseudoHover        = ":hover"
	PseudoActive       = ":active"
	PseudoFocus        = ":focus"
	PseudoFocusVisible = ":focus-visible"
	PseudoFocusWithin  = ":focus-within"

	// Location pseudo-classes
	PseudoLink    = ":link"
	PseudoVisited = ":visited"
	PseudoTarget  = ":target"

	// Input pseudo-classes
	PseudoEnabled          = ":enabled"
	PseudoDisabled         = ":disabled"
	PseudoChecked          = ":checked"
	PseudoIndeterminate    = ":indeterminate"
	PseudoDefault          = ":default"
	PseudoValid            = ":valid"
	PseudoInvalid          = ":invalid"
	PseudoPlaceholderShown = ":placeholder-shown"
	PseudoReadOnly         = ":read-only"
	PseudoReadWrite        = ":read-write"
	PseudoRequired         = ":required"
	PseudoOptional         = ":optional"

	// Tree-structural pseudo-classes
	PseudoFirstChild    = ":first-child"
	PseudoLastChild     = ":last-child"
	PseudoOnlyChild     = ":only-child"
	PseudoFirstOfType   = ":first-of-type"
	PseudoLastOfType    = ":last-of-type"
	PseudoOnlyOfType    = ":only-of-type"
	PseudoEmpty         = ":empty"
	PseudoNthChild      = ":nth-child"
	PseudoNthLastChild  = ":nth-last-child"
	PseudoNthOfType     = ":nth-of-type"
	PseudoNthLastOfType = ":nth-last-of-type"
	PseudoNot           = ":not()"
	PseudoRoot          = ":root"

	// Linguistic pseudo-classes
	PseudoLang = ":lang"

	// Fullscreen pseudo-class
	PseudoFullScreen = ":fullscreen"

	// Animations pseudo-classes
	PseudoPlaying = ":playing"
	PseudoPaused  = ":paused"
)

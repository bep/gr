//go:generate go run generate.go

// Package attr defines markup to create HTML attributes supported by Facebook React.
//
// Created from "HTML Attributes" as defined by Facebook in
// - https://facebook.github.io/react/docs/tags-and-attributes.html
// - http://facebook.github.io/react/docs/special-non-dom-attributes.html
package attr

import "github.com/bep/gr"

// About creates an HTML attribute for 'about'.
func About(v interface{}) gr.Modifier {
	return gr.Prop("about", v)
}

// Accept creates an HTML attribute for 'accept'.
func Accept(v interface{}) gr.Modifier {
	return gr.Prop("accept", v)
}

// AcceptCharset creates an HTML attribute for 'acceptCharset'.
func AcceptCharset(v interface{}) gr.Modifier {
	return gr.Prop("acceptCharset", v)
}

// AccessKey creates an HTML attribute for 'accessKey'.
func AccessKey(v interface{}) gr.Modifier {
	return gr.Prop("accessKey", v)
}

// Action creates an HTML attribute for 'action'.
func Action(v interface{}) gr.Modifier {
	return gr.Prop("action", v)
}

// AllowFullScreen creates an HTML attribute for 'allowFullScreen'.
func AllowFullScreen(v interface{}) gr.Modifier {
	return gr.Prop("allowFullScreen", v)
}

// AllowTransparency creates an HTML attribute for 'allowTransparency'.
func AllowTransparency(v interface{}) gr.Modifier {
	return gr.Prop("allowTransparency", v)
}

// Alt creates an HTML attribute for 'alt'.
func Alt(v interface{}) gr.Modifier {
	return gr.Prop("alt", v)
}

// Async creates an HTML attribute for 'async'.
func Async(v interface{}) gr.Modifier {
	return gr.Prop("async", v)
}

// AutoCapitalize creates an HTML attribute for 'autoCapitalize'.
func AutoCapitalize(v interface{}) gr.Modifier {
	return gr.Prop("autoCapitalize", v)
}

// AutoComplete creates an HTML attribute for 'autoComplete'.
func AutoComplete(v interface{}) gr.Modifier {
	return gr.Prop("autoComplete", v)
}

// AutoCorrect creates an HTML attribute for 'autoCorrect'.
func AutoCorrect(v interface{}) gr.Modifier {
	return gr.Prop("autoCorrect", v)
}

// AutoFocus creates an HTML attribute for 'autoFocus'.
func AutoFocus(v interface{}) gr.Modifier {
	return gr.Prop("autoFocus", v)
}

// AutoPlay creates an HTML attribute for 'autoPlay'.
func AutoPlay(v interface{}) gr.Modifier {
	return gr.Prop("autoPlay", v)
}

// AutoSave creates an HTML attribute for 'autoSave'.
func AutoSave(v interface{}) gr.Modifier {
	return gr.Prop("autoSave", v)
}

// Capture creates an HTML attribute for 'capture'.
func Capture(v interface{}) gr.Modifier {
	return gr.Prop("capture", v)
}

// CellPadding creates an HTML attribute for 'cellPadding'.
func CellPadding(v interface{}) gr.Modifier {
	return gr.Prop("cellPadding", v)
}

// CellSpacing creates an HTML attribute for 'cellSpacing'.
func CellSpacing(v interface{}) gr.Modifier {
	return gr.Prop("cellSpacing", v)
}

// Challenge creates an HTML attribute for 'challenge'.
func Challenge(v interface{}) gr.Modifier {
	return gr.Prop("challenge", v)
}

// CharSet creates an HTML attribute for 'charSet'.
func CharSet(v interface{}) gr.Modifier {
	return gr.Prop("charSet", v)
}

// Checked creates an HTML attribute for 'checked'.
func Checked(v interface{}) gr.Modifier {
	return gr.Prop("checked", v)
}

// Cite creates an HTML attribute for 'cite'.
func Cite(v interface{}) gr.Modifier {
	return gr.Prop("cite", v)
}

// ClassID creates an HTML attribute for 'classID'.
func ClassID(v interface{}) gr.Modifier {
	return gr.Prop("classID", v)
}

// ClassName creates an HTML attribute for 'className'.
func ClassName(v interface{}) gr.Modifier {
	return gr.Prop("className", v)
}

// ColSpan creates an HTML attribute for 'colSpan'.
func ColSpan(v interface{}) gr.Modifier {
	return gr.Prop("colSpan", v)
}

// Color creates an HTML attribute for 'color'.
func Color(v interface{}) gr.Modifier {
	return gr.Prop("color", v)
}

// Cols creates an HTML attribute for 'cols'.
func Cols(v interface{}) gr.Modifier {
	return gr.Prop("cols", v)
}

// Content creates an HTML attribute for 'content'.
func Content(v interface{}) gr.Modifier {
	return gr.Prop("content", v)
}

// ContentEditable creates an HTML attribute for 'contentEditable'.
func ContentEditable(v interface{}) gr.Modifier {
	return gr.Prop("contentEditable", v)
}

// ContextMenu creates an HTML attribute for 'contextMenu'.
func ContextMenu(v interface{}) gr.Modifier {
	return gr.Prop("contextMenu", v)
}

// Controls creates an HTML attribute for 'controls'.
func Controls(v interface{}) gr.Modifier {
	return gr.Prop("controls", v)
}

// Coords creates an HTML attribute for 'coords'.
func Coords(v interface{}) gr.Modifier {
	return gr.Prop("coords", v)
}

// CrossOrigin creates an HTML attribute for 'crossOrigin'.
func CrossOrigin(v interface{}) gr.Modifier {
	return gr.Prop("crossOrigin", v)
}

// DangerouslySetInnerHTML Provides the ability to insert raw HTML, 
// mainly for cooperating with DOM string manipulation libraries.
func DangerouslySetInnerHTML(v interface{}) gr.Modifier {
	return gr.Prop("dangerouslySetInnerHTML", v)
}

// Data creates an HTML attribute for 'data'.
func Data(v interface{}) gr.Modifier {
	return gr.Prop("data", v)
}

// Datatype creates an HTML attribute for 'datatype'.
func Datatype(v interface{}) gr.Modifier {
	return gr.Prop("datatype", v)
}

// DateTime creates an HTML attribute for 'dateTime'.
func DateTime(v interface{}) gr.Modifier {
	return gr.Prop("dateTime", v)
}

// Default creates an HTML attribute for 'default'.
func Default(v interface{}) gr.Modifier {
	return gr.Prop("default", v)
}

// Defer creates an HTML attribute for 'defer'.
func Defer(v interface{}) gr.Modifier {
	return gr.Prop("defer", v)
}

// Dir creates an HTML attribute for 'dir'.
func Dir(v interface{}) gr.Modifier {
	return gr.Prop("dir", v)
}

// Disabled creates an HTML attribute for 'disabled'.
func Disabled(v interface{}) gr.Modifier {
	return gr.Prop("disabled", v)
}

// Download creates an HTML attribute for 'download'.
func Download(v interface{}) gr.Modifier {
	return gr.Prop("download", v)
}

// Draggable creates an HTML attribute for 'draggable'.
func Draggable(v interface{}) gr.Modifier {
	return gr.Prop("draggable", v)
}

// EncType creates an HTML attribute for 'encType'.
func EncType(v interface{}) gr.Modifier {
	return gr.Prop("encType", v)
}

// Form creates an HTML attribute for 'form'.
func Form(v interface{}) gr.Modifier {
	return gr.Prop("form", v)
}

// FormAction creates an HTML attribute for 'formAction'.
func FormAction(v interface{}) gr.Modifier {
	return gr.Prop("formAction", v)
}

// FormEncType creates an HTML attribute for 'formEncType'.
func FormEncType(v interface{}) gr.Modifier {
	return gr.Prop("formEncType", v)
}

// FormMethod creates an HTML attribute for 'formMethod'.
func FormMethod(v interface{}) gr.Modifier {
	return gr.Prop("formMethod", v)
}

// FormNoValidate creates an HTML attribute for 'formNoValidate'.
func FormNoValidate(v interface{}) gr.Modifier {
	return gr.Prop("formNoValidate", v)
}

// FormTarget creates an HTML attribute for 'formTarget'.
func FormTarget(v interface{}) gr.Modifier {
	return gr.Prop("formTarget", v)
}

// FrameBorder creates an HTML attribute for 'frameBorder'.
func FrameBorder(v interface{}) gr.Modifier {
	return gr.Prop("frameBorder", v)
}

// Headers creates an HTML attribute for 'headers'.
func Headers(v interface{}) gr.Modifier {
	return gr.Prop("headers", v)
}

// Height creates an HTML attribute for 'height'.
func Height(v interface{}) gr.Modifier {
	return gr.Prop("height", v)
}

// Hidden creates an HTML attribute for 'hidden'.
func Hidden(v interface{}) gr.Modifier {
	return gr.Prop("hidden", v)
}

// High creates an HTML attribute for 'high'.
func High(v interface{}) gr.Modifier {
	return gr.Prop("high", v)
}

// HRef creates an HTML attribute for 'href'.
func HRef(v interface{}) gr.Modifier {
	return gr.Prop("href", v)
}

// HRefLang creates an HTML attribute for 'hrefLang'.
func HRefLang(v interface{}) gr.Modifier {
	return gr.Prop("hrefLang", v)
}

// HTMLFor creates an HTML attribute for 'htmlFor'.
func HTMLFor(v interface{}) gr.Modifier {
	return gr.Prop("htmlFor", v)
}

// HTTPEquiv creates an HTML attribute for 'httpEquiv'.
func HTTPEquiv(v interface{}) gr.Modifier {
	return gr.Prop("httpEquiv", v)
}

// Icon creates an HTML attribute for 'icon'.
func Icon(v interface{}) gr.Modifier {
	return gr.Prop("icon", v)
}

// ID creates an HTML attribute for 'id'.
func ID(v interface{}) gr.Modifier {
	return gr.Prop("id", v)
}

// Inlist creates an HTML attribute for 'inlist'.
func Inlist(v interface{}) gr.Modifier {
	return gr.Prop("inlist", v)
}

// InputMode creates an HTML attribute for 'inputMode'.
func InputMode(v interface{}) gr.Modifier {
	return gr.Prop("inputMode", v)
}

// Integrity creates an HTML attribute for 'integrity'.
func Integrity(v interface{}) gr.Modifier {
	return gr.Prop("integrity", v)
}

// Is creates an HTML attribute for 'is'.
func Is(v interface{}) gr.Modifier {
	return gr.Prop("is", v)
}

// ItemProp creates an HTML attribute for 'itemProp'.
func ItemProp(v interface{}) gr.Modifier {
	return gr.Prop("itemProp", v)
}

// Key adds an optional, unique identifier. 
// When your component shuffles around during render passes, it might be destroyed 
// and recreated due to the diff algorithm. Assigning it a key that persists makes 
// sure the component stays.
func Key(v interface{}) gr.Modifier {
	return gr.Prop("key", v)
}

// KeyParams creates an HTML attribute for 'keyParams'.
func KeyParams(v interface{}) gr.Modifier {
	return gr.Prop("keyParams", v)
}

// KeyType creates an HTML attribute for 'keyType'.
func KeyType(v interface{}) gr.Modifier {
	return gr.Prop("keyType", v)
}

// Kind creates an HTML attribute for 'kind'.
func Kind(v interface{}) gr.Modifier {
	return gr.Prop("kind", v)
}

// Label creates an HTML attribute for 'label'.
func Label(v interface{}) gr.Modifier {
	return gr.Prop("label", v)
}

// Lang creates an HTML attribute for 'lang'.
func Lang(v interface{}) gr.Modifier {
	return gr.Prop("lang", v)
}

// List creates an HTML attribute for 'list'.
func List(v interface{}) gr.Modifier {
	return gr.Prop("list", v)
}

// Loop creates an HTML attribute for 'loop'.
func Loop(v interface{}) gr.Modifier {
	return gr.Prop("loop", v)
}

// Low creates an HTML attribute for 'low'.
func Low(v interface{}) gr.Modifier {
	return gr.Prop("low", v)
}

// Manifest creates an HTML attribute for 'manifest'.
func Manifest(v interface{}) gr.Modifier {
	return gr.Prop("manifest", v)
}

// MarginHeight creates an HTML attribute for 'marginHeight'.
func MarginHeight(v interface{}) gr.Modifier {
	return gr.Prop("marginHeight", v)
}

// MarginWidth creates an HTML attribute for 'marginWidth'.
func MarginWidth(v interface{}) gr.Modifier {
	return gr.Prop("marginWidth", v)
}

// Max creates an HTML attribute for 'max'.
func Max(v interface{}) gr.Modifier {
	return gr.Prop("max", v)
}

// MaxLength creates an HTML attribute for 'maxLength'.
func MaxLength(v interface{}) gr.Modifier {
	return gr.Prop("maxLength", v)
}

// Media creates an HTML attribute for 'media'.
func Media(v interface{}) gr.Modifier {
	return gr.Prop("media", v)
}

// MediaGroup creates an HTML attribute for 'mediaGroup'.
func MediaGroup(v interface{}) gr.Modifier {
	return gr.Prop("mediaGroup", v)
}

// Method creates an HTML attribute for 'method'.
func Method(v interface{}) gr.Modifier {
	return gr.Prop("method", v)
}

// Min creates an HTML attribute for 'min'.
func Min(v interface{}) gr.Modifier {
	return gr.Prop("min", v)
}

// MinLength creates an HTML attribute for 'minLength'.
func MinLength(v interface{}) gr.Modifier {
	return gr.Prop("minLength", v)
}

// Multiple creates an HTML attribute for 'multiple'.
func Multiple(v interface{}) gr.Modifier {
	return gr.Prop("multiple", v)
}

// Muted creates an HTML attribute for 'muted'.
func Muted(v interface{}) gr.Modifier {
	return gr.Prop("muted", v)
}

// Name creates an HTML attribute for 'name'.
func Name(v interface{}) gr.Modifier {
	return gr.Prop("name", v)
}

// NoValidate creates an HTML attribute for 'noValidate'.
func NoValidate(v interface{}) gr.Modifier {
	return gr.Prop("noValidate", v)
}

// Nonce creates an HTML attribute for 'nonce'.
func Nonce(v interface{}) gr.Modifier {
	return gr.Prop("nonce", v)
}

// Open creates an HTML attribute for 'open'.
func Open(v interface{}) gr.Modifier {
	return gr.Prop("open", v)
}

// Optimum creates an HTML attribute for 'optimum'.
func Optimum(v interface{}) gr.Modifier {
	return gr.Prop("optimum", v)
}

// Pattern creates an HTML attribute for 'pattern'.
func Pattern(v interface{}) gr.Modifier {
	return gr.Prop("pattern", v)
}

// Placeholder creates an HTML attribute for 'placeholder'.
func Placeholder(v interface{}) gr.Modifier {
	return gr.Prop("placeholder", v)
}

// Poster creates an HTML attribute for 'poster'.
func Poster(v interface{}) gr.Modifier {
	return gr.Prop("poster", v)
}

// Prefix creates an HTML attribute for 'prefix'.
func Prefix(v interface{}) gr.Modifier {
	return gr.Prop("prefix", v)
}

// Preload creates an HTML attribute for 'preload'.
func Preload(v interface{}) gr.Modifier {
	return gr.Prop("preload", v)
}

// Profile creates an HTML attribute for 'profile'.
func Profile(v interface{}) gr.Modifier {
	return gr.Prop("profile", v)
}

// Property creates an HTML attribute for 'property'.
func Property(v interface{}) gr.Modifier {
	return gr.Prop("property", v)
}

// RadioGroup creates an HTML attribute for 'radioGroup'.
func RadioGroup(v interface{}) gr.Modifier {
	return gr.Prop("radioGroup", v)
}

// ReadOnly creates an HTML attribute for 'readOnly'.
func ReadOnly(v interface{}) gr.Modifier {
	return gr.Prop("readOnly", v)
}

// Ref adds an ref to a component, see http://facebook.github.io/react/docs/more-about-refs.html
func Ref(v interface{}) gr.Modifier {
	return gr.Prop("ref", v)
}

// Rel creates an HTML attribute for 'rel'.
func Rel(v interface{}) gr.Modifier {
	return gr.Prop("rel", v)
}

// Required creates an HTML attribute for 'required'.
func Required(v interface{}) gr.Modifier {
	return gr.Prop("required", v)
}

// Resource creates an HTML attribute for 'resource'.
func Resource(v interface{}) gr.Modifier {
	return gr.Prop("resource", v)
}

// Results creates an HTML attribute for 'results'.
func Results(v interface{}) gr.Modifier {
	return gr.Prop("results", v)
}

// Reversed creates an HTML attribute for 'reversed'.
func Reversed(v interface{}) gr.Modifier {
	return gr.Prop("reversed", v)
}

// Role creates an HTML attribute for 'role'.
func Role(v interface{}) gr.Modifier {
	return gr.Prop("role", v)
}

// RowSpan creates an HTML attribute for 'rowSpan'.
func RowSpan(v interface{}) gr.Modifier {
	return gr.Prop("rowSpan", v)
}

// Rows creates an HTML attribute for 'rows'.
func Rows(v interface{}) gr.Modifier {
	return gr.Prop("rows", v)
}

// Sandbox creates an HTML attribute for 'sandbox'.
func Sandbox(v interface{}) gr.Modifier {
	return gr.Prop("sandbox", v)
}

// Scope creates an HTML attribute for 'scope'.
func Scope(v interface{}) gr.Modifier {
	return gr.Prop("scope", v)
}

// Scoped creates an HTML attribute for 'scoped'.
func Scoped(v interface{}) gr.Modifier {
	return gr.Prop("scoped", v)
}

// Scrolling creates an HTML attribute for 'scrolling'.
func Scrolling(v interface{}) gr.Modifier {
	return gr.Prop("scrolling", v)
}

// Seamless creates an HTML attribute for 'seamless'.
func Seamless(v interface{}) gr.Modifier {
	return gr.Prop("seamless", v)
}

// Security creates an HTML attribute for 'security'.
func Security(v interface{}) gr.Modifier {
	return gr.Prop("security", v)
}

// Selected creates an HTML attribute for 'selected'.
func Selected(v interface{}) gr.Modifier {
	return gr.Prop("selected", v)
}

// Shape creates an HTML attribute for 'shape'.
func Shape(v interface{}) gr.Modifier {
	return gr.Prop("shape", v)
}

// Size creates an HTML attribute for 'size'.
func Size(v interface{}) gr.Modifier {
	return gr.Prop("size", v)
}

// Sizes creates an HTML attribute for 'sizes'.
func Sizes(v interface{}) gr.Modifier {
	return gr.Prop("sizes", v)
}

// Span creates an HTML attribute for 'span'.
func Span(v interface{}) gr.Modifier {
	return gr.Prop("span", v)
}

// SpellCheck creates an HTML attribute for 'spellCheck'.
func SpellCheck(v interface{}) gr.Modifier {
	return gr.Prop("spellCheck", v)
}

// Src creates an HTML attribute for 'src'.
func Src(v interface{}) gr.Modifier {
	return gr.Prop("src", v)
}

// SrcDoc creates an HTML attribute for 'srcDoc'.
func SrcDoc(v interface{}) gr.Modifier {
	return gr.Prop("srcDoc", v)
}

// SrcLang creates an HTML attribute for 'srcLang'.
func SrcLang(v interface{}) gr.Modifier {
	return gr.Prop("srcLang", v)
}

// SrcSet creates an HTML attribute for 'srcSet'.
func SrcSet(v interface{}) gr.Modifier {
	return gr.Prop("srcSet", v)
}

// Start creates an HTML attribute for 'start'.
func Start(v interface{}) gr.Modifier {
	return gr.Prop("start", v)
}

// Step creates an HTML attribute for 'step'.
func Step(v interface{}) gr.Modifier {
	return gr.Prop("step", v)
}

// Style creates an HTML attribute for 'style'.
func Style(v interface{}) gr.Modifier {
	return gr.Prop("style", v)
}

// Summary creates an HTML attribute for 'summary'.
func Summary(v interface{}) gr.Modifier {
	return gr.Prop("summary", v)
}

// TabIndex creates an HTML attribute for 'tabIndex'.
func TabIndex(v interface{}) gr.Modifier {
	return gr.Prop("tabIndex", v)
}

// Target creates an HTML attribute for 'target'.
func Target(v interface{}) gr.Modifier {
	return gr.Prop("target", v)
}

// Title creates an HTML attribute for 'title'.
func Title(v interface{}) gr.Modifier {
	return gr.Prop("title", v)
}

// Type creates an HTML attribute for 'type'.
func Type(v interface{}) gr.Modifier {
	return gr.Prop("type", v)
}

// Typeof creates an HTML attribute for 'typeof'.
func Typeof(v interface{}) gr.Modifier {
	return gr.Prop("typeof", v)
}

// Unselectable creates an HTML attribute for 'unselectable'.
func Unselectable(v interface{}) gr.Modifier {
	return gr.Prop("unselectable", v)
}

// UseMap creates an HTML attribute for 'useMap'.
func UseMap(v interface{}) gr.Modifier {
	return gr.Prop("useMap", v)
}

// Value creates an HTML attribute for 'value'.
func Value(v interface{}) gr.Modifier {
	return gr.Prop("value", v)
}

// Vocab creates an HTML attribute for 'vocab'.
func Vocab(v interface{}) gr.Modifier {
	return gr.Prop("vocab", v)
}

// Width creates an HTML attribute for 'width'.
func Width(v interface{}) gr.Modifier {
	return gr.Prop("width", v)
}

// WMode creates an HTML attribute for 'wmode'.
func WMode(v interface{}) gr.Modifier {
	return gr.Prop("wmode", v)
}

// Wrap creates an HTML attribute for 'wrap'.
func Wrap(v interface{}) gr.Modifier {
	return gr.Prop("wrap", v)
}

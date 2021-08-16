#Token Constructor
### Documentation
######  A contructor has been written for compatibility with the [github.com/nguyenthenguyen/docx][1] module. This constructor adds more convenient control of document blocks, and allows you to create more new features.

------------
### Prepare
######  The [github.com/nguyenthenguyen/docx][1] module stores the content of the document (found inside the docx under the path 'word / document.xml') as a string. This is only useful for unsafe text substitutions in the document. 

```go
type Docx struct {
	files   []*zip.File
	content string
	links   string
	headers map[string]string
	footers map[string]string
}
```
######  To expand the capabilities and more convenient secure management, the Constructor submodule was written. It presents the document as a collection of Paragraph Tokens.
```go
type Document struct {
    WP     []WP //WP is Paragraph
    SectPr SectPr //section's properties.
}
```
######  Each paragraph is represented by a collection of tags(WPTokens) that are contained within it.
```go
type WP struct {
	Tag  string
	Body []WPTokens
}
```
######  WPTokens is element in the Paragraph. Example: r, pPr, bookmarkStart and another( read ooxmlDocs [http://officeopenxml.com/WPparagraph.php][2])
```go
type WPTokens struct {
	Tag    string
	Body   string
	Attr   string
	Status int
}
```
######  Tag  it  is name of tag. Example in tag  ```<w:pStyle w:val="Normal"/>``` , ``w:pStyle``  is name.
###### Body it is content on token.  Example in tag  ```<w:t xml:space="preserve">This is a word document.</w:t>``` , ``This is a word document.`` is a body.
###### Attr it is attributes on token.  Example in tag  ```<w:t xml:space="preserve">This is a word document.</w:t>``` , ``xml:space="preserve"`` is a attributes.
###### Status it is tag type. paired or not paired (self-closing) . Open or Self.
###### Example Open tag: `<w:t>This is a word document.</w:t>`
###### Example Self tag: `<w:pStyle/>`
###### In library, for set/check status using constants.
```go
const Open = 4
const Self = 5
```
######  For example:
```xml
<w:p>
            <w:pPr>
                <w:pStyle w:val="Normal"/>
                <w:rPr>
                    <w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs=""
						w:asciiTheme="minorHAnsi" w:cstheme="minorBidi"
						w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/>
                </w:rPr>
            </w:pPr>
            <w:bookmarkStart w:id="0" w:name="_GoBack"/>
            <w:bookmarkEnd w:id="0"/>
            <w:r>
                <w:rPr></w:rPr>
                <w:t xml:space="preserve">This is a word document.</w:t>
            </w:r>
        </w:p>
```
######  Parser convert this xml in struct:
```go
WP{
		Tag: "w:p",
		Body: []WPTokens{
			{
				Tag:    "w:pPr",
				Body:   `
					<w:pStyle w:val="Normal"/>
                	<w:rPr>
                    	<w:rFonts w:ascii="Calibri" w:hAnsi="Calibri" w:eastAsia="Calibri" w:cs=""
                        	w:asciiTheme="minorHAnsi" w:cstheme="minorBidi"
                        	w:eastAsiaTheme="minorHAnsi" w:hAnsiTheme="minorHAnsi"/>
               	 </w:rPr>`,
				Status: Open,
				Attr: ``,
			},
			{
				Tag:    "w:bookmarkStart",
				Body:   ``,
				Status: Self,
				Attr: ` w:id="0" w:name="_GoBack"`,
			},
			{
				Tag:    "w:r",
				Body:   `
					<w:rPr></w:rPr>
                	<w:t xml:space="preserve">
							This is a word document.
					</w:t>`,
				Status: Open,
				Attr: ``,
			},
		},
	}
```
### Beginning of work
######  Since the constructor uses its own struct type, you first need to parse the body of the document. And at the end, glue everything together again in one string.
```go
//read file
r, err := ReadDocxFile("TestDocument.docx")
	if err != nil {
		fmt.Println(err)
		return
	}
	docx1 := r.Editable()
//parse file
	doc, err := docx1.Parser()
	if err != nil {
		fmt.Println(err)
		return
	}
//work space
//.....
//.....
//.....
//compile
err = docx1.Compile("TestDocument2.docx", doc)
	if err != nil {
		fmt.Println(err)
		return
	}

```
### List of Methods
#### For (d *Docx)
- Compile(path string, doc Document) error
- Parser() (Document, error)

#### For (d *Document)
- AddNewBlock(s string)
- GetBlockIDByTag(s string) (int, error)
- EditBlockByID(id int) //not implemented
- AppendWPBlockInToEnd(block WP)
- EditBlockWithNewLine(oldTag, newString string) error
- ReplaceTagString(oldTag, newString string) error
- GetBlockByID(id int) WP
- BodyToString() string 
- CreateMarkedStringList(mp MarkerParams, letter ...string) []WP 

#### Service methods
- Screening(s string) string
- AtomicWPTokensToString(token WPTokens) string


[1]: http://github.com/nguyenthenguyen/docx "github.com/nguyenthenguyen/docx"
[2]: http://officeopenxml.com/WPparagraph.php "http://officeopenxml.com/WPparagraph.php"
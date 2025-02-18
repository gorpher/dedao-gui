package utils

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

func Svg2Raw(outputDir,title string, svgContents []*SvgContent, toc []EbookToc) (err error) {
	tocLevel = make(map[string]int, len(toc))
	for _, ebookToc := range toc {
		tocLevel[ebookToc.Text] = ebookToc.Level
	}
	fmt.Println(tocLevel)

	result, err := AllToRaw(svgContents, toc)
	if err != nil {
		return err
	}
	result2, err := AllToJSON(svgContents, toc)
	if err != nil {
		return err
	}
	p, err := Mkdir(outputDir, "Ebook")
	if err != nil {
		return err
	}
	fileName, err := FilePath(filepath.Join(p, FileName(title, "")), "html", false)
	if err != nil {
		return err
	}
	fileName2, err := FilePath(filepath.Join(p, FileName(title, "")), "json", false)
	if err != nil {
		return err
	}
	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", fileName)
	if err = WriteFileWithTrunc(fileName, result); err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return
	}
	fmt.Printf("正在生成文件：【\033[37;1m%s\033[0m】 ", fileName2)
	if err = WriteFileWithTrunc(fileName2, result2); err != nil {
		fmt.Printf("\033[31;1m%s\033[0m\n", "失败"+err.Error())
		return
	}
	fmt.Printf("\033[32;1m%s\033[0m\n", "完成")
	return
}

// AllToRaw generate ebook content all in one html file
func AllToRaw(svgContents []*SvgContent, toc []EbookToc) (result string, err error) {
	//result = GenHeadHtml()
	fnA, fnB = ParseBookFnDelimiter(svgContents)
	fmt.Println("fnA:", fnA, "fnB:", fnB)
	for _, ebookToc := range svgContents {

		for _, content := range ebookToc.Contents {
			result += content
		}
	}
	//	result += `
	//</body>
	//</html>`
	//	result = html.UnescapeString(result)
	return
}

func AllToJSON(svgContents []*SvgContent, toc []EbookToc) (result string, err error) {
	body, err := json.Marshal(map[string]interface{}{
		"contents": svgContents,
		"toc":      toc,
	})
	if err != nil {
		return
	}
	result = string(body)
	return
}

package excel

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/gogf/gf/v2/util/guid"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gogf/gf/v2/frame/g"
)

type Entity struct {
	Data       interface{} `json:"data"`     // [][]string || []map[string]string
	SortKey    []string    `json:"sort_key"` // 只有参数为[]map[string]string时，需要设置该值
	FileName   string      `json:"file_name"`
	SaveToPath string      `json:"save_to_path"`
	SheetName  string      `json:"sheet_name"`
}

func (e *Entity) initConf() {
	if e.FileName == "" {
		s := guid.S()
		e.FileName = fmt.Sprintf("%s_%s.xlsx", time.Now().Format("20060102150405"), s[len(s)-6:])
	}
	if e.SheetName == "" {
		e.SheetName = "Sheet1"
	}
}

func (e *Entity) formatData() [][]string {
	switch d := e.Data.(type) {
	case [][]string:
		return d
	case []map[string]string:
		if len(d) > 0 {
			sortKey := e.SortKey
			if len(sortKey) <= 0 {
				sortKey = make([]string, 0, len(d[0]))
				for k := range d[0] {
					sortKey = append(sortKey, k)
				}
				sort.Strings(sortKey)
			}

			data := make([][]string, 0, len(d))
			for _, v := range d {
				row := make([]string, 0, len(v))
				for _, vv := range sortKey {
					row = append(row, v[vv])
				}
				data = append(data, row)
			}
			return data
		}

	default:
		return [][]string{}
	}
	return nil
}

func (e *Entity) formatSavePath() string {
	if strings.TrimRight(e.SaveToPath, "/") != "" {
		return fmt.Sprintf("%s/%s", strings.TrimRight(e.SaveToPath, "/"), e.FileName)
	}
	return e.FileName
}

// Generate 生成文件byte
func (e *Entity) Generate() (*bytes.Buffer, error) {
	e.initConf()
	data := e.formatData()
	f := excelize.NewFile()
	for k, v := range data {
		for kk, vv := range v {
			cell := fmt.Sprintf("%s%d", string('A'+kk), k+1)
			f.SetCellValue(e.SheetName, cell, vv)
		}
	}
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buffer, err
}

// GenerateFile 生成文件
func (e *Entity) GenerateFile() error {
	e.initConf()
	data := e.formatData()
	f := excelize.NewFile()
	for k, v := range data {
		for kk, vv := range v {
			cell := fmt.Sprintf("%s%d", string('A'+kk), k+1)
			f.SetCellValue(e.SheetName, cell, vv)
		}
	}
	err := f.SaveAs(e.formatSavePath())
	if err != nil {
		return err
	}
	return nil
}

// GenerateSimple 直接下载
func (e *Entity) GenerateSimple(ctx context.Context) error {
	data := e.formatData()
	f := excelize.NewFile()
	for k, v := range data {
		for kk, vv := range v {
			cell := fmt.Sprintf("%s%d", string('A'+kk), k+1)
			f.SetCellValue(e.SheetName, cell, vv)
		}
	}
	r := g.RequestFromCtx(ctx)
	r.Response.Writer.Header().Set("Content-Type", "application/octet-stream")
	r.Response.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", e.FileName))
	r.Response.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	r.Response.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	return f.Write(r.Response.Writer)
}

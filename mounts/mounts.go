// Package mounts 处理坐骑收藏的计数与分档加成。
package mounts

import (
	_ "embed"
	"encoding/json"
	"sync"
)

//go:embed data/mounts.json
var rawTable []byte

var (
	tableOnce   sync.Once
	cachedTable *Table
)

// Mount 是坐骑表里的一条记录。
type Mount struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

// Tier 是一个加成档位。
type Tier struct {
	Min   int `json:"min"`
	Bonus int `json:"bonus"`
}

// Table 是整张坐骑表。
type Table struct {
	Mounts []Mount `json:"mounts"`
	Tiers  []Tier  `json:"tiers"`
}

// Counter 统计本轮重新解析的字节数。
type Counter struct {
	Scanned int
}

// Result 是一次结算的输出。
type Result struct {
	Mounts  int `json:"mounts"`
	Tier    int `json:"tier"`
	Bonus   int `json:"bonus"`
	Scanned int `json:"scanned"`
}

// Loader 读取内嵌的坐骑表。
type Loader struct{}

// Table 解析内嵌资源并返回坐骑表。
func (Loader) Table(c *Counter) *Table {
	tableOnce.Do(func() {
		var t Table
		if err := json.Unmarshal(rawTable, &t); err != nil {
			panic(err)
		}
		cachedTable = &t
		if c != nil {
			c.Scanned += len(rawTable)
		}
	})
	return cachedTable
}

// Compute 按拥有的坐骑算出收藏数量与档位加成。
func (Loader) Compute(owned []string, c *Counter) Result {
	t := Loader{}.Table(c)
	seen := make(map[string]bool, len(owned))
	count := 0
	for _, id := range owned {
		if seen[id] {
			continue
		}
		seen[id] = true
		for _, m := range t.Mounts {
			if m.ID == id {
				count++
				if m.Category == "limited" {
					count++
				}
				break
			}
		}
	}
	bonus := 0
	min := 0
	for _, tr := range t.Tiers {
		if count >= tr.Min && tr.Min > min {
			bonus = tr.Bonus
			min = tr.Min
		}
	}
	return Result{Mounts: count, Tier: min, Bonus: bonus, Scanned: c.Scanned}
}

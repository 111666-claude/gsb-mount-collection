// Command mount 是坐骑收藏结算的场景入口。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"example.com/mount-collection/mounts"
)

func emit(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	sample := flag.String("sample", "", "dedupe|tier|limited|work")
	flag.Parse()

	loader := mounts.Loader{}
	switch *sample {
	case "dedupe":
		emit(loader.Compute([]string{"m01", "m01", "m02"}, &mounts.Counter{}))
	case "tier":
		emit(loader.Compute([]string{"m01", "m02", "m03", "m04"}, &mounts.Counter{}))
	case "limited":
		emit(loader.Compute([]string{"L01"}, &mounts.Counter{}))
	case "work":
		c := &mounts.Counter{}
		for i := 0; i < 2000; i++ {
			loader.Compute([]string{"m01", "m02"}, c)
		}
		emit(map[string]int{"rounds": 2000, "scanned": c.Scanned})
	case "":
		flag.Usage()
		os.Exit(2)
	default:
		fmt.Fprintln(os.Stderr, "未知场景："+*sample)
		os.Exit(2)
	}
}

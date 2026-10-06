package mounts

import "testing"

func TestNoOwnedMounts(t *testing.T) {
	got := Loader{}.Compute(nil, &Counter{})
	if got.Mounts != 0 || got.Bonus != 0 {
		t.Fatalf("空收藏应得 0/0，得到 %+v", got)
	}
}

func TestUnknownMountIsIgnored(t *testing.T) {
	got := Loader{}.Compute([]string{"zz"}, &Counter{})
	if got.Mounts != 0 {
		t.Fatalf("未知坐骑不该计数，得到 %+v", got)
	}
}

func TestSingleMountBelowFirstTier(t *testing.T) {
	got := Loader{}.Compute([]string{"m01"}, &Counter{})
	if got.Mounts != 1 || got.Bonus != 0 {
		t.Fatalf("一只坐骑不该有加成，得到 %+v", got)
	}
}

func TestTwoDistinctMountsHitFirstTier(t *testing.T) {
	got := Loader{}.Compute([]string{"m01", "m02"}, &Counter{})
	if got.Mounts != 2 || got.Bonus != 5 {
		t.Fatalf("两只应命中第一档 5，得到 %+v", got)
	}
}

func TestTableHasThreeTiers(t *testing.T) {
	tbl := Loader{}.Table(&Counter{})
	if len(tbl.Tiers) != 3 {
		t.Fatalf("档位表应有 3 档，得到 %d", len(tbl.Tiers))
	}
}

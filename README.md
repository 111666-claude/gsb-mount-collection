# mount-collection

坐骑收藏的计数与分档加成。坐骑表随二进制走（`go:embed`），只用标准库。

```
go test ./...
go run ./cmd/mount --sample=dedupe
go run ./cmd/mount --sample=tier
go run ./cmd/mount --sample=limited
go run ./cmd/mount --sample=work
```

## 口径

- **去重**：同一只坐骑重复出现只算一次收藏。
- **只认表内坐骑**：不在坐骑表里的 id 不计入收藏数量。
- **覆盖层**：`category` 为 `limited` 的坐骑每只额外贡献 1 点收藏数量。
- **分档**：加成取"收藏数量 ≥ 档位下限"里下限最高的那一档；不到最低档给 0。

## 不变量

- 收藏数量不超过拥有的坐骑条数加上限定坐骑的额外点数。
- 加成只由收藏数量与档位表决定，与拥有列表的顺序无关。
- `scanned` 不随结算次数放大：两千次结算的 `scanned` 不超过 9000。

## 输出契约

每次结算打印一行 JSON：`mounts`、`tier`、`bonus`、`scanned`。
`--sample=work` 只打印 `{"rounds": N, "scanned": N}`。

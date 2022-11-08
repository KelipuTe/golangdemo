package insert

type Dialect interface {
	// quoter 返回一个引号，引用列名，表名的引号
	quoter() byte
	// buildUpsert 构造插入冲突部分
	buildUpsert(b *builder, odk *Upsert) error
}

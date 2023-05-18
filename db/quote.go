package db

import (
	"strings"
)

// QuoteForLikeEscape 转义字符串用于 LIKE 查询
// 例：
//
//	db.First(&row, "column LIKE ? ESCAPE ?", "%" + types.QuoteForLikeEscape(keyword, '\\') + "%", "\\")
//
// @see: https://www.sqlite.org/lang_expr.html#like
// @see: https://dev.mysql.com/doc/refman/5.7/en/string-comparison-functions.html#operator_like
// @see: https://www.postgresql.org/docs/12/functions-matching.html#FUNCTIONS-LIKE
func QuoteForLikeEscape(s string, escapeChar rune) string {
	escape := string(escapeChar)
	s = strings.ReplaceAll(s, escape, escape+escape)
	if escapeChar != '_' {
		s = strings.ReplaceAll(s, "_", escape+"_")
	}
	if escapeChar != '%' {
		s = strings.ReplaceAll(s, "%", escape+"%")
	}
	return s
}

// QuoteForLike 同 QuoteForLikeEscape，但固定使用反斜杠转义
// 这个行为于 MySQL 和 PostgreSQL 不带 ESCAPE 的 LIKE 一致，但 SQLite 必须手动指定 ESCAPE
// 例：
//    db.First(&row, "column LIKE ?", "%" + types.QuoteForLikeEscape(keyword) + "%")
//goland:noinspection GoUnusedExportedFunction
// func QuoteForLike(keyword string) string {
// 	return QuoteForLikeEscape(keyword, '\\')
// }

// QuoteForContain 同 QuoteForLike，但固定前后加百分号的模式
// 例：
//
//	db.First(&row, "column LIKE ?", types.QuoteForContain(keyword))
//
//goland:noinspection GoUnusedExportedFunction
func QuoteForContain(keyword string) string {
	if DB != nil && DB.Dialector.Name() == "sqlite" {
		// FIXME:
		//   由于 SQLite 需要显式增加 ESCAPE 才能支持 '_' 和 '%' 的转义
		//   在所有调用此函数的代码针对 SQLite 增加 ESCAPE 之前，关闭 SQlite 下的转义
		//   这将导致 keyword 中的 '_' 和 '%' 被识别为通配符，导致和 MySQL、PostgreSQL 不同的行为
		return "%" + keyword + "%"
	}
	return "%" + QuoteForLikeEscape(keyword, '\\') + "%"
}

// QuoteForStartWith 同 QuoteForLike，但固定后面加百分号的模式
//
//goland:noinspection GoUnusedExportedFunction
func QuoteForStartWith(keyword string) string {
	if DB != nil && DB.Dialector.Name() == "sqlite" {
		// FIXME:
		//   由于 SQLite 需要显式增加 ESCAPE 才能支持 '_' 和 '%' 的转义
		//   在所有调用此函数的代码针对 SQLite 增加 ESCAPE 之前，关闭 SQlite 下的转义
		//   这将导致 keyword 中的 '_' 和 '%' 被识别为通配符，导致和 MySQL、PostgreSQL 不同的行为
		return keyword + "%"
	}
	return QuoteForLikeEscape(keyword, '\\') + "%"
}

// QuoteForEndWith 同 QuoteForLike，但固定前面加百分号的模式
//
//goland:noinspection GoUnusedExportedFunction
func QuoteForEndWith(keyword string) string {
	if DB != nil && DB.Dialector.Name() == "sqlite" {
		// FIXME:
		//   由于 SQLite 需要显式增加 ESCAPE 才能支持 '_' 和 '%' 的转义
		//   在所有调用此函数的代码针对 SQLite 增加 ESCAPE 之前，关闭 SQlite 下的转义
		//   这将导致 keyword 中的 '_' 和 '%' 被识别为通配符，导致和 MySQL、PostgreSQL 不同的行为
		return "%" + keyword
	}
	return "%" + QuoteForLikeEscape(keyword, '\\')
}

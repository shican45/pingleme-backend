package serializer

import "PingLeMe-Backend/model"

// EvaluationTable 评审表序列化器
type EvaluationTable struct {
	TableName string `json:"table_name"`
	TableItems []EvaluationTableItem `json:"table_items"`
}

// EvaluationTableItem 评审表项序列化器
type EvaluationTableItem struct {
	Index           int                   `json:"index"`
	Content         string                `json:"content"`
	Score           int                   `json:"score"`
	Description     string                `json:"description"`
	ChildTableItems []EvaluationTableItem `json:"child_table_items"`
}

// BuildEvaluationTable 序列化评审表
func BuildEvaluationTable(tableModel model.EvaluationTable) EvaluationTable {
	var table EvaluationTable
	table.TableName = tableModel.TableName

	// TODO: Fix it later.
	//for index, item := range tableModel.TableItems {
	//
	//}

	return EvaluationTable{}
}
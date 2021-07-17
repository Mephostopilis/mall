package biz

import (
	"context"
	"fmt"
	"strings"

	ssopb "edu/api/sso/v1"
	pb "edu/api/sys/v1"
	stringsx "edu/pkg/strings"
	"edu/service/sys/internal/model"
	"edu/service/sys/internal/tools"
)

func (uc *AdminUsecase) convertSysColumns2Pb(columns []model.SysColumns) []*pb.SysColumn {
	list := make([]*pb.SysColumn, 0)
	for i := 0; i < len(columns); i++ {
		it := columns[i]
		d := &pb.SysColumn{
			ColumnId:     int32(it.ColumnId),
			TableId:      int32(it.TableId),
			ColumnName:   it.ColumnName,
			ColumnType:   it.ColumnType,
			GoType:       it.GoType,
			GoField:      it.GoField,
			JsonField:    it.JsonField,
			IsPk:         it.IsPk,
			IsIncrement:  it.IsIncrement,
			IsRequired:   it.IsRequired,
			IsInsert:     it.IsInsert,
			IsEdit:       it.IsEdit,
			IsList:       it.IsList,
			IsQuery:      it.IsQuery,
			QueryType:    it.QueryType,
			HtmlType:     it.HtmlType,
			DictType:     it.DictType,
			Sort:         int32(it.Sort),
			List:         it.List,
			Pk:           it.Pk,
			Required:     it.Required,
			SuperColumn:  it.SuperColumn,
			UsableColumn: it.UsableColumn,
			Increment:    it.Increment,
			Insert:       it.Insert,
			Edit:         it.Edit,
			Query:        it.Query,
			Remark:       it.Remark,
			CreatedAt:    it.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		list = append(list, d)
	}
	return list
}

func (uc *AdminUsecase) convertSysTables2Pb(tables []model.SysTables) []*pb.SysTable {
	list := make([]*pb.SysTable, 0)
	for i := 0; i < len(tables); i++ {
		it := tables[i]
		d := &pb.SysTable{
			TableId:             int32(it.TableId),
			TableName:           it.TBName,
			TableComment:        it.TableComment,
			ClassName:           it.ClassName,
			TplCategory:         it.TplCategory,
			PackageName:         it.PackageName,
			ModuleName:          it.ModuleName,
			BusinessName:        it.BusinessName,
			FunctionName:        it.FunctionName,
			FunctionAuthor:      it.FunctionAuthor,
			PkColumn:            it.PkColumn,
			PkGoField:           it.PkGoField,
			PkJsonField:         it.PkJsonField,
			Options:             it.Options,
			TreeCode:            it.TreeCode,
			TreeParentCode:      it.TreeParentCode,
			TreeName:            it.TreeName,
			Tree:                it.Tree,
			Crud:                it.Crud,
			Remark:              it.Remark,
			IsLogicalDelete:     it.IsLogicalDelete,
			LogicalDelete:       it.LogicalDelete,
			LogicalDeleteColumn: it.LogicalDeleteColumn,
			Columns:             uc.convertSysColumns2Pb(it.Columns),
			CreatedAt:           it.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		list = append(list, d)
	}
	return list
}

func (uc *AdminUsecase) genTableInit(ctx context.Context, dp *ssopb.DataPermission, toold tools.Dao, tbname string) (reply *model.SysTables, err error) {
	dbtable, err := toold.GetDBTables(uc.cfg.Gen.Dbname, &model.DBTables{TableName: tbname})
	if err != nil {
		return
	}
	var data model.SysTables
	data.TBName = tbname
	data.CreateBy = fmt.Sprint(dp.UserId)
	tablenamelist := strings.Split(tbname, "_")
	for i := 0; i < len(tablenamelist); i++ {
		strStart := string([]byte(tablenamelist[i])[:1])
		strend := string([]byte(tablenamelist[i])[1:])
		data.ClassName += strings.ToUpper(strStart) + strend
		data.PackageName += strings.ToLower(strStart) + strings.ToLower(strend)
		data.ModuleName += strings.ToLower(strStart) + strings.ToLower(strend)
	}
	data.TplCategory = "crud"
	data.Crud = true

	data.CreateBy = fmt.Sprint(dp.UserId)
	data.TableComment = dbtable.TableComment
	if dbtable.TableComment == "" {
		data.TableComment = data.ClassName
	}

	data.FunctionName = data.TableComment
	data.BusinessName = data.ModuleName
	data.IsLogicalDelete = "1"
	data.LogicalDelete = true
	data.LogicalDeleteColumn = "is_del"
	data.FunctionAuthor = "bing"

	// 获取列信息
	data.Columns = make([]model.SysColumns, 0)
	dbcolumns, err := toold.GetDBColumnsList(uc.cfg.Gen.Dbname, &model.DBColumns{TableName: tbname})
	for i := int32(0); i < int32(len(dbcolumns)); i++ {
		dbcolumn := dbcolumns[i]
		var column model.SysColumns
		column.ColumnComment = dbcolumn.ColumnComment
		column.ColumnName = dbcolumn.ColumnName
		column.ColumnType = dbcolumn.ColumnType
		column.Sort = i + 1
		column.Insert = true
		column.IsInsert = "1"
		column.QueryType = "EQ"
		column.IsPk = "0"

		namelist := strings.Split(dbcolumn.ColumnName, "_")
		for i := 0; i < len(namelist); i++ {
			strStart := string([]byte(namelist[i])[:1])
			strend := string([]byte(namelist[i])[1:])
			column.GoField += strings.ToUpper(strStart) + strend
			if i == 0 {
				column.JsonField = strings.ToLower(strStart) + strend
			} else {
				column.JsonField += strings.ToUpper(strStart) + strend
			}
		}
		if strings.Contains(dbcolumn.ColumnKey, "PR") {
			column.IsPk = "1"
			column.Pk = true
			data.PkColumn = dbcolumn.ColumnName
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
		}
		column.IsRequired = "0"
		if strings.Contains(dbcolumn.IsNullable, "NO") {
			column.IsRequired = "1"
			column.Required = true
		}

		if strings.Contains(dbcolumn.ColumnType, "int") {
			if strings.Contains(dbcolumn.ColumnKey, "PR") {
				column.GoType = "int"
			} else {
				column.GoType = "string"
			}
			column.HtmlType = "input"
		} else if strings.Contains(dbcolumn.ColumnType, "timestamp") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
		} else {
			column.GoType = "string"
			column.HtmlType = "input"
		}

		data.Columns = append(data.Columns, column)
	}
	reply = &data
	return
}

func (uc *AdminUsecase) GetSysTableList(ctx context.Context, req *pb.GetSysTableListRequest) (docs []*pb.SysTable, total int64, err error) {
	var pageIndex = int(req.PageIndex)
	var pageSize = int(req.PageSize)
	var data model.SysTables
	data.TBName = req.TableName
	data.TableComment = req.TableComment
	result, total, err := uc.d.GetSysTablesPage(&data, pageSize, pageIndex)
	if err != nil {
		return
	}
	docs = uc.convertSysTables2Pb(result)
	return
}

func (uc *AdminUsecase) GetSysTables(ctx context.Context, req *pb.GetSysTablesRequest) (doc *pb.GetSysTablesReply, err error) {
	var data model.SysTables
	data.TableId = (req.TableId)
	it, err := uc.d.GetSysTables(&data)
	if err != nil {
		return
	}
	// uc.log.Debugf("getsystables %v", it)
	doc = &pb.GetSysTablesReply{
		Info: &pb.SysTable{
			TableId:             int32(it.TableId),
			TableName:           it.TBName,
			TableComment:        it.TableComment,
			ClassName:           it.ClassName,
			TplCategory:         it.TplCategory,
			PackageName:         it.PackageName,
			ModuleName:          it.ModuleName,
			BusinessName:        it.BusinessName,
			FunctionName:        it.FunctionName,
			FunctionAuthor:      it.FunctionAuthor,
			PkColumn:            it.PkColumn,
			PkGoField:           it.PkGoField,
			PkJsonField:         it.PkJsonField,
			Options:             it.Options,
			TreeCode:            it.TreeCode,
			TreeParentCode:      it.TreeParentCode,
			TreeName:            it.TreeName,
			Tree:                it.Tree,
			Crud:                it.Crud,
			Remark:              it.Remark,
			IsLogicalDelete:     it.IsLogicalDelete,
			LogicalDelete:       it.LogicalDelete,
			LogicalDeleteColumn: it.LogicalDeleteColumn,
		},
		Rows: uc.convertSysColumns2Pb(it.Columns),
	}
	return
}

func (uc *AdminUsecase) InsertSysTable(ctx context.Context, tok string, req *pb.InsertSysTableRequest) (err error) {
	out, err := uc.mw.ValidationToken(tok)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	d, err := tools.New(uc.cfg, uc.logger)
	if err != nil {
		return
	}
	defer d.Close()
	tablesList := strings.Split(req.Tables, ",")
	for i := 0; i < len(tablesList); i++ {
		tbname := tablesList[i]
		data, err1 := uc.genTableInit(ctx, dp, d, tbname)
		if err1 != nil {
			err = err1
			return
		}
		_, err = uc.d.CreateSysTables(data)
		if err != nil {
			return
		}
	}
	return
}

func (uc *AdminUsecase) UpdateSysTable(ctx context.Context, tok string, req *pb.SysTable) (err error) {
	out, err := uc.mw.ValidationToken(tok)
	if err != nil {
		return
	}
	dp := out.(*ssopb.DataPermission)
	data := model.SysTables{
		TableId:             req.TableId,
		TBName:              req.TableName,
		TableComment:        req.TableComment,
		ClassName:           req.ClassName,
		TplCategory:         req.TplCategory,
		PackageName:         req.PackageName,
		ModuleName:          req.ModuleName,
		BusinessName:        req.BusinessName,
		FunctionName:        req.FunctionName,
		FunctionAuthor:      req.FunctionAuthor,
		PkColumn:            req.PkColumn,
		PkGoField:           req.PkGoField,
		PkJsonField:         req.PkJsonField,
		Options:             req.Options,
		TreeCode:            req.TreeCode,
		TreeParentCode:      req.TreeParentCode,
		TreeName:            req.TreeName,
		Tree:                req.Tree,
		Crud:                req.Crud,
		Remark:              req.Remark,
		IsLogicalDelete:     req.IsLogicalDelete,
		LogicalDelete:       req.LogicalDelete,
		LogicalDeleteColumn: req.LogicalDeleteColumn,
		UpdateBy:            fmt.Sprint(dp.UserId),
	}
	data.UpdateBy = fmt.Sprint(dp.UserId)
	_, err = uc.d.UpdateSysTables(&data)
	if err != nil {
		return
	}
	return
}

func (uc *AdminUsecase) DeleteSysTables(ctx context.Context, req *pb.DeleteSysTablesRequest) (err error) {
	ids, err := stringsx.SplitInts(req.Ids, ",")
	if err != nil {
		return
	}
	_, err = uc.d.BatchDeleteSysTables(&model.SysTables{}, ids)
	if err != nil {
		return
	}
	return
}

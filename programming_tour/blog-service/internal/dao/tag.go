package dao

import (
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
)

func (d *Dao) GetTag(id uint32, state uint8) (model.Tag, error) {
	tag := model.Tag{Model: &model.Model{ID: id}, State: state}
	return tag.Get(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetTagListByIDs(ids []uint32, state uint8) ([]*model.Tag, error) {
	tag := model.Tag{State: state}
	return tag.ListByIDs(d.engine, ids)
}

func (d *Dao) CountTag(name string, state uint8) (int, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) CreateTag(name string, state uint8, createdBy string) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: &model.Model{
			CreatedBy: createdBy,
		},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID: id,
		},
	}
	// 原代码:
	//tag := model.Tag{
	//	Model: &model.Model{
	//		ID: id,
	//		ModifiedBy: modifiedBy,
	//	},
	//  Name: name,
	//  State: state,
	//}
	/*
	发现问题
	在完成接口检验后，还需确定数据库内的数据变更是否正确。在经过一系列的对比后发现，在调用修改标签的接口时，
	通过接口入参，我们希望将id为1的标签状态修改为0，但是经对比后发现，在数据库内它的状态值仍然还是1，
	而且SQL语句内也没有出现state字段的设置.

	甚至在我们做其他类似的验证时，发现只要字段的值为零值，GORM就不会对该字段进行变更，这是为什么呢？
	实际上，我们先入为主地认为它一定会变更是不对的。在我们的程序中，是使用 struct 进行更新操作的。
	而在GORM中，当使用struct进行更新操作时，GORM是不会对值为零值的字段进行变更的。其根本原因在于，
	在识别这个结构体中的字段值时，很难判定是真的为零值，还是外部传入的该类型的值恰好为零值。对此，GORM并没有过多地去做特殊识别。
	*/
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeleteTag(id uint32) error {
	tag := model.Tag{Model: &model.Model{ID: id}}
	return tag.Delete(d.engine)
}

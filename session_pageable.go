package xorm

import (
	"errors"
	"github.com/go-various/xorm/pageable"
)
var ErrPageableCannotBeNil = errors.New("pageable cannot be nil")
func (session *Session) FindPagination(rowsSlicePtr interface{}, page pageable.Pageable, condiBean ...interface{}) (*pageable.Pagination, error) {
	if nil == page{
		return nil, ErrPageableCannotBeNil
	}

	total, err := session.Limit(page.Limit(), page.Skip()).FindAndCount(rowsSlicePtr, condiBean)
	if err != nil {
		return nil, err
	}

	return pageable.NewPagination(int(total), page), nil
}


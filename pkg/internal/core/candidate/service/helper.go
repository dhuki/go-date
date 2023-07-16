package service

func (c candidateServiceImpl) getTotalPage(totalData, limit int) (totalPage int) {
	if limit > 0 && totalData > 0 {
		additional := 0
		if totalData%limit != 0 {
			additional = 1
		}
		return (totalData / limit) + additional
	}
	return
}

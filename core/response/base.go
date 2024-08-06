package response

type Response struct {
	Code    int         `json:"status"` //返回 0，表示当前接口正确返回，否则按错误请求处理；
	Message string      `json:"msg"`    //返回接口处理信息，主要用于表单提交或请求失败时的 toast 显示；
	Data    interface{} `json:"data"`   //必须返回一个具有 key-value 结构的对象。
}

type ExResponse struct {
	Code    int            `json:"status"` //返回 0，表示当前接口正确返回，否则按错误请求处理；
	Message string         `json:"msg"`    //返回接口处理信息，主要用于表单提交或请求失败时的 toast 显示；
	Data    ResultResponse `json:"data"`   //必须返回一个具有 key-value 结构的对象。
}

type ResultResponse struct {
	Result interface{} `json:"result"`
}
type ListResponse struct {
	Total  int64       `json:"total" form:"total"`
	Items  interface{} `json:"items" form:"items"`
	AdList interface{} `json:"ad_list,omitempty" form:"ad_list"`
}

// amis Options 选择器
type ListOptionsResponse struct {
	Options interface{} `json:"options"`
}
type OptionsItemResponse struct {
	Value string `json:"value" form:"-" gorm:"-"`
	Label string `json:"label" form:"-" gorm:"-"`
}

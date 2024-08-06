package handler

// https://baidu.gitee.io/amis/zh-CN/components/form/input-file
type UploadResponse struct {
	Filename string `json:"filename"`
	Value    string `json:"value"`
	URL      string `json:"url"`
	Hash     string `json:"hash"`
}

// type UploadResponseEx struct {
// 	Filename string `json:"filename"`
// 	Value    string `json:"value"`
// 	URL      string `json:"url"`
// 	Hash     string `json:"hash"`
// }

// Request对象：接⼝封装了客户请求信息，如客户请求⽅式、参数、客户使⽤的协议、以 及发出请 求的远程主机信息等，
// 其主要⽅法：
// a) String getParamter(String paramName);//获取请求参数
// b) String[] getParamterValues(String paramName);
// c) void setCharacterEncoding(String encode);
// d) void setAttribute(String name,Object value);
// e) Object getAttribute(String name);
// f) RequestDispatcher getRequestDispatcher(String rediret)

// 这里是前端或者是前面传过来的数据通过request获取（片面的个人理解）

// Response对象：接⼝封装了服务器响应信息，主要⽤来输出信息到客户端浏览器，
// 其主要 ⽅法：response.setCharacterEncoding();
// Response.setContentType("text/html;charset=utf-8");设置响应的内容类型
// PrintWriter response.getWriter();获得响应的输出流
// response.sendRedirect(redirect)；重定向到指定的⽹址

package MercerServer

import (
	"MercerFrame/MercerLog"
	"strings"
)

func (c *Context) ParseParam(name string) string {
	//To do
	routerUri := strings.SplitN(c.RouterUri, "/", -1)[1:]
	uri := strings.SplitN(c.Uri, "/", -1)[1:]
	if len(routerUri) == len(uri) {
		for i := 0; i < len(routerUri); i++ {
			if routerUri[i] == name || routerUri[i] == ":"+name {
				c.Param = make(map[string]string, 10)
				c.Param[name] = uri[i]
				return c.Param[name]
			}
		}
	} else if len(routerUri) < len(uri) {
		uri1 := strings.SplitN(c.Uri, "/", len(routerUri))
		for i := 0; i < len(routerUri); i++ {
			if routerUri[i] == name || routerUri[i] == ":"+name {
				c.Param[name] = uri1[i]
				return c.Param[name]
			}
		}
	} else {
		//r := Router{}
		//rmap := r.RouterMap
		//if v,ok := rmap[routerUri[0]];ok {
		//}
		return "404"
	}
	return "404"
}

//Get方法取参数，/user/zjr类型的值，如果没有路由匹配/user/zjr/age,那么也会匹配成/user/zjr
func (c *Context) PraParam(name string) string {
	s := strings.SplitN(c.Uri,"/",-1)
	part1 := ""
	for i := 1;i <= c.RNum ;i++  {
		part1 += ("/" + s[i])
	}
	part2 := strings.Split(c.Uri,part1)[1]
	//temp := strings.Split(c.Uri,"/")
	head := part1
	//head := "/" + strings.SplitN(c.Uri, "/", -1)[1]
	//t1 := "/" + temp[1]
	RParamNum := strings.Count(part2,"/")
	//RParamNum := strings.Count(c.Uri, "/") - 2
	//uri := strings.SplitN(c.Uri,"/",-1)
	//uri := strings.SplitN(c.Uri, "/", -1)[2:]
	if strings.Count(c.Uri, "/") > 1 && head == c.RouterUri {
		if c.ParamNum == RParamNum {
			names := strings.SplitN(part2,"/",-1)[1:]
			for i := 0; i < len(c.ParamName); i++ {
				if c.ParamName[i] == name || c.ParamName[i] == ":"+name {
					c.Param = make(map[string]string, 10)
					c.Param[name] = names[i]
					return c.Param[name]
				}
			}
		} else if c.ParamNum < RParamNum {
			//uri1 := strings.SplitN(t1, "/", c.ParamNum)
			names := strings.SplitN(part2,"/",c.ParamNum + 1)[1:]
			for i := 0; i < len(c.ParamName); i++ {
				if c.ParamName[i] == name || c.ParamName[i] == ":"+name {
					c.Param = make(map[string]string, 10)
					c.Param[name] = names[i]
					return c.Param[name]
				}
			}
		} else {
			MercerLog.Error.Println("Param is nil")
			return "404"
		}
	} else {
		MercerLog.Error.Println("Method is wrong")
		return "404"
	}
	MercerLog.Error.Println("Method is wrong")
	return "404"
}

//Get方法拿参数，可设置默认返回值
func (c *Context) DefaultQuery(name string, defaultName string) string {
	var uri string
	if strings.Contains(c.Uri, "?") {
		uri = strings.Split(c.Uri, "?")[1]
		if strings.Contains(uri, "=") {
			c.Param = make(map[string]string, strings.Count(uri, "="))
			p := strings.Split(uri, "&")
			for _, str := range p {
				temp := strings.Split(str, "=")
				key := temp[0]
				value := temp[1]
				c.Param[key] = value
			}
			return c.Param[name]
		} else {
			return defaultName
		}
	}
	return defaultName
}

//Get方法拿参数，没有默认返回值
func (c *Context) Query(name string) string {
	var uri string
	if strings.Contains(c.Uri, "?") {
		uri = strings.Split(c.Uri, "?")[1]
		if strings.Contains(uri, "=") {
			c.Param = make(map[string]string, strings.Count(uri, "="))
			p := strings.Split(uri, "&")
			for _, str := range p {
				temp := strings.Split(str, "=")
				key := temp[0]
				value := temp[1]
				c.Param[key] = value
			}
			return c.Param[name]
		} else {
			return ""
		}
	}
	return ""
}

//post方法测试拿参数，有默认返回值
func (c *Context) DefaultPostForm(name string, defaultName string) string {
	str := c.Req.PostFormValue(name)
	if str == "" {
		return defaultName
	} else {
		return str
	}
}

//post方法测试拿参数，没有默认返回值
func (c *Context) PostForm(name string) string {
	str := c.Req.PostFormValue(name)
	return str
}
package model

import "github.com/hongshengjie/framework/app/api/blog"

var nav []*blog.Nav

func GetNav() []*blog.Nav {
	return nav
}

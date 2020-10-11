package model

import "framework/app/api/blog"

var nav []*blog.Nav

func GetNav() []*blog.Nav {
	return nav
}

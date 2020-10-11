package blog

import (
	"context"

	"github.com/hongshengjie/framework/app/api/blog"
	"github.com/hongshengjie/framework/app/service/blog/model"

	"github.com/golang/protobuf/ptypes/empty"
)

// Service Service
type Service struct {
	blog.UnimplementedBlogServer
}

// 首页获取的数据
func (s *Service) Index(ctx context.Context, req *empty.Empty) (*blog.IndexResp, error) {
	return &blog.IndexResp{Navs: model.GetNav()}, nil
}

// 列表页面
func (s *Service) List(ctx context.Context, req *blog.ListReq) (*blog.ListResp, error) {
	if list, err := model.PostList(req.GetPreId(), 20, req.GetAsc()); err != nil {
		return nil, err
	} else {
		return &blog.ListResp{Posts: list}, nil
	}

}

// 详情页
func (s *Service) Detail(ctx context.Context, req *blog.DetailReq) (*blog.Post, error) {
	p, err := model.PostByID(uint64(req.GetId()))
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 编辑页面
func (s *Service) Edit(ctx context.Context, req *blog.Post) (*empty.Empty, error) {
	if err := model.Upsert(req); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

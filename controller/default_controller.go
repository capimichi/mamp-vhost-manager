package controller

import "context"

type DefaultController struct {
	ctx context.Context
}

func (d *DefaultController) Startup(ctx context.Context) {
	d.ctx = ctx
}


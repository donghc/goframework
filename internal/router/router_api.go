package router

import "goframework/internal/api/helper"

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)
	helpers := r.mux.Group("/helper")
	{
		helpers.GET("/md5/:str",helperHandler.Md5())
		helpers.POST("/sign/:str",helperHandler.Sign())
	}
}

package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) {}

func add(c *gin.Context, db *sql.DB)        {}
func auth(c *gin.Context, db *sql.DB)       {}
func authCB(c *gin.Context, db *sql.DB)     {}
func privacy(c *gin.Context, db *sql.DB)    {}
func tos(c *gin.Context, db *sql.DB)        {}
func about(c *gin.Context, db *sql.DB)      {}
func guild(c *gin.Context, db *sql.DB)      {}
func guild_id(c *gin.Context, db *sql.DB)   {}
func members(c *gin.Context, db *sql.DB)    {}
func movies(c *gin.Context, db *sql.DB)     {}
func events(c *gin.Context, db *sql.DB)     {}
func events_new(c *gin.Context, db *sql.DB) {}
func events_id(c *gin.Context, db *sql.DB)  {}

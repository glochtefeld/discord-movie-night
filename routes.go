package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {}

func Add(c *gin.Context, db *sql.DB)        {}
func Auth(c *gin.Context, db *sql.DB)       {}
func AuthCB(c *gin.Context, db *sql.DB)     {}
func Privacy(c *gin.Context, db *sql.DB)    {}
func Tos(c *gin.Context, db *sql.DB)        {}
func About(c *gin.Context, db *sql.DB)      {}
func Guild(c *gin.Context, db *sql.DB)      {}
func Guild_id(c *gin.Context, db *sql.DB)   {}
func Members(c *gin.Context, db *sql.DB)    {}
func Movies(c *gin.Context, db *sql.DB)     {}
func Events(c *gin.Context, db *sql.DB)     {}
func Events_new(c *gin.Context, db *sql.DB) {}
func Events_id(c *gin.Context, db *sql.DB)  {}

package internal

import "embed"

//go:embed static/*
var StaticFiles embed.FS

package migrations

import "embed"

// Файлы с миграциями
//go:embed schemas/*
var Migrations embed.FS

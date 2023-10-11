package migrations

import "embed"

// Файлы с миграциями
//go:embed migrations/*
var Migrations embed.FS

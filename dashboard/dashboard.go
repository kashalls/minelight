package dashboard

import (
	"embed"
	"io/fs"
)

//go:embed all:.output/public
var distDir embed.FS
var DistDirFS, _ = fs.Sub(distDir, ".output/public")
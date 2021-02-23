package directories

import (
	"path"

	"github.com/spf13/viper"
	"github.com/wjdingman1/mediforui/pkg/config"
)

/* workingDir is the directory where the HTTP server will look for static assets
 * when running in production the workingDir and containerRoot will be the same */
var (
	conf          = initConfig()
	workingDir    = conf.Get("WORKINGDIR").(string)
	containerRoot = conf.Get("CONTAINERROOT").(string)
)

func initConfig() *viper.Viper {
	c, err := config.New()
	if err != nil {
		panic("-- CONFIG NOT FOUND: SHUTTING DOWN --")
	}
	return c
}

// Directories struct is a helper to organize all the backend directories
type Directories struct {
	ApplicationDirectory ApplicationDirectory
	ContainerOutputPath  string
	ContainerInputPath   string
}

// ApplicationDirectory struct hold all the directories the HTTP server will interact with
type ApplicationDirectory struct {
	InboxPath             string
	ApplicationOutputPath string
	ApplicationInputPath  string
	ThumbnailPath         string
	TranscodedVideoPath   string
}

// New returns a new Directories struct
func New() Directories {
	return Directories{
		ApplicationDirectory: ApplicationDirectory{
			InboxPath:             path.Join(workingDir, "inbox"),
			ApplicationOutputPath: path.Join(workingDir, "output"),
			ApplicationInputPath:  path.Join(workingDir, "input"),
			ThumbnailPath:         path.Join(workingDir, "input", "thumbnails"),
			TranscodedVideoPath:   path.Join(workingDir, "input", "transcoded"),
		},
		ContainerOutputPath: path.Join(containerRoot, "output"),
		ContainerInputPath:  path.Join(containerRoot, "input"),
	}
}

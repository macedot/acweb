package settings

import (
	"github.com/DeKugelschieber/go-util"
	log "github.com/sirupsen/logrus"
	"model"
)

func SaveSettings(folder, executable, args string) error {
	folder = util.Trim(folder)
	executable = util.Trim(executable)
	args = util.Trim(args)

	if folder == "" || executable == "" {
		return util.OpError{1, "Folder and executable must be set"}
	}

	settings, err := model.GetSettings()

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error reading settings")
		settings = &model.Settings{}
	}

	settings.Folder = folder
	settings.Executable = executable
	settings.Args = args

	if err := settings.Save(); err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error saving settings")
		return util.OpError{2, "Error saving settings"}
	}

	return nil
}

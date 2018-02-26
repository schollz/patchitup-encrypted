package patchitup

import (
	"path"

	"github.com/schollz/patchitup-encrypted/patchitup/keypair"
	"github.com/schollz/utils"
)

var DataFolder string
var sharedKey keypair.KeyPair

func init() {
	DataFolder = path.Join(utils.UserHomeDir(), ".patchitup")
	sharedKey = keypair.NewDeterministic("patchitup")
}

package airgap


import (
        "fmt"
        "bytes"
        "log"
        "os/exec"
	"github.com/openshift/installer/pkg/asset"
	//"github.com/openshift/installer/pkg/types"

)

const (
	releaseImage = "quay.io/openshift-release-dev/ocp-release:"
)

type mirrorReleaseMetaData struct {
}

var _ asset.Asset = (*mirrorReleaseMetaData)(nil)

// Dependencies returns no dependencies.
func (a *mirrorReleaseMetaData) Dependencies() []asset.Asset {
	return nil 
}

// Generate queries for the cluster name from the user.
func (a *mirrorReleaseMetaData) Generate(parents asset.Parents) error {
        return nil
}

// Name returns the human-friendly name of the asset.
func (a *mirrorReleaseMetaData) Name() string {
	return "Mirror Release Disconnected"
}


func (a *mirrorReleaseMetaData) pullClusterImages(airPackage *AirgapPackage) bool {
	fmt.Println("Pulling mirror images for OCP 4 mirror")

	command := exec.Command("oc", "adm", "release", "mirror", "-a",
		airPackage.pull_secret, releaseImage + airPackage.ocp_ver + "-x86_64",
		"--to=file://openshift/release", "--to-dir=" + airPackage.dest + "/registry")

	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		fmt.Println("Error occurred")
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Printf("Command Output: %s\n", out.String())

	return true
}

func (a *mirrorReleaseMetaData) pullRedHatOperators(airPackage *AirgapPackage) bool {
	return true
}

func (a *mirrorReleaseMetaData) pullCertifiedOperators(airPackage *AirgapPackage) bool {
        return true
}

func (a *mirrorReleaseMetaData) pullCommunityOperators(airPackage *AirgapPackage) bool {
        return true
}


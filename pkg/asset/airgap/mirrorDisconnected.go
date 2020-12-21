package airgap


import (
        "fmt"
        //"bytes"
        //"log"
        "os/exec"
	"github.com/openshift/installer/pkg/asset"
	//"bufio"
	//"github.com/openshift/installer/pkg/types"
	"io/ioutil"
	"strings"
)

const (
	releaseImage = "quay.io/openshift-release-dev/ocp-release:"
	rhOpIndex = "registry.redhat.io/redhat/redhat-operator-index:v"
	certOpIndex = "registry.redhat.io/redhat/certified-operator-index:v"
	commOpIndex ="registry.redhat.io/redhat/community-operator-index:v"
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
	fmt.Println("Pulling OpenShift 4 cluster images for package")

	command := exec.Command("oc", "adm", "release", "mirror", "-a",
		airPackage.pull_secret, releaseImage + airPackage.ocp_ver + "-x86_64",
		"--to=file://openshift/release", "--to-dir=" + airPackage.dest + "/registry")

	//stdout, _ := command.StderrPipe()
	command.Start()
	//scanner := bufio.NewScanner(stdout)
	//scanner.Split(bufio.ScanWords)
	//for scanner.Scan() {
	//	m := scanner.Text()
	//	fmt.Print(m)
	//}
	command.Wait()

/*
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		fmt.Println("Error occurred")
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Printf("Command Output: %s\n", out.String())
*/
	return true
}

func (a *mirrorReleaseMetaData) pullRedHatOperators(airPackage *AirgapPackage) bool {

        fmt.Println("Pulling Red Hat operators for package")

	// Create the manifests for the redhat operators catalog
	args := "adm catalog mirror --registry-config " + airPackage.pull_secret
	args += " " + rhOpIndex + "4.6 --manifests-only=true --to-manifests "
	args += airPackage.dest + "/redhat_operators_manifests replaceme"

	fmt.Println("Running: oc ", args)

	command := exec.Command("oc", strings.Split(args, " ")...)
	command.Start()
	command.Wait()

	// update the manifest on disk to point to our package registry location
	read, err := ioutil.ReadFile(airPackage.dest+"/redhat_operators_manifests/mapping.txt")
	if err != nil {
		panic(err)
	}
	newContents := strings.Replace(string(read), "replaceme", "file://"+ airPackage.dest + "/registry", -1)
	err = ioutil.WriteFile(airPackage.dest+"/redhat_operators_manifests/mapping.txt", []byte(newContents), 0)
	if err != nil {
		panic(err)
	}

	args = "image mirror --dir=" + airPackage.dest + "/registry "
	args += "--registry-config=" + airPackage.pull_secret
	args += " --filename=" + airPackage.dest+"/redhat_operators_manifests/mapping.txt"

        fmt.Println("Running: oc ", args)

	command = exec.Command("oc", strings.Split(args, " ")...)
	command.Start()
	command.Wait()

	//fmt.Println("RH Operators command: ", args)

	return true
}

func (a *mirrorReleaseMetaData) pullCertifiedOperators(airPackage *AirgapPackage) bool {
/*
        fmt.Println("Pulling certified operators for package")

        command := exec.Command("oc", "adm", "release", "mirror", "-a",
                airPackage.pull_secret, releaseImage + airPackage.ocp_ver + "-x86_64",
                "--to=file://openshift/release", "--to-dir=" + airPackage.dest + "/registry")

*/
        return true
}

func (a *mirrorReleaseMetaData) pullCommunityOperators(airPackage *AirgapPackage) bool {
/*
        fmt.Println("Pulling community operators for package")

        command := exec.Command("oc", "adm", "release", "mirror", "-a",
                airPackage.pull_secret, releaseImage + airPackage.ocp_ver + "-x86_64",
                "--to=file://openshift/release", "--to-dir=" + airPackage.dest + "/registry")

*/
        return true
}


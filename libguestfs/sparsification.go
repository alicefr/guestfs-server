package libguestfs

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func SparsifyImage(image string) error {
	args := []string{"--in-place", "-v", "-x", image}
	c := exec.Command("virt-sparsify", args...)
	os.Setenv("LIBGUESTFS_BACKEND", "direct")
	o, err := c.CombinedOutput()
	if err != nil {
		log.Errorf("Unable to run virt-sparsify: %v", string(o))
		return err
	}

	return nil
}

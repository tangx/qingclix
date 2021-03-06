package configure

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CustomizeCmd represents the customize command
var CustomizeCmd = &cobra.Command{
	Use:   "customize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("customize called")
	},
}

func init() {
}

type Qingtypes struct {
	InstanceTypes []InstanceType   `json:"instance_types"`
	VolumeTypes   []VolumeType     `json:"volume_types"`
	ImageTypes    []ImageType      `json:"image_types"`
	Zones         []CommonType     `json:"zones"`
	Vxnets        []CommonType     `json:"vxnets"`
	Keypairs      []CommonType     `json:"keypairs"`
	Relation      map[string][]int `json:"relation"`
}

type ImageType struct {
	Name  string `json:"name"`
	Image string `json:"image"`
	Desc  string `json:"desc"`
}

type InstanceType struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Class int    `json:"class"`
	Desc  string `json:"desc"`
}

type CommonType struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type VolumeType struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	Desc string `json:"desc"`
}

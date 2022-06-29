/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"go/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"opcap/internal/operator"

	"opcap/internal/capability"

	"github.com/spf13/cobra"
)

type CheckCommandFlags struct {
	CatalogSource          string `json:"catalogsource"`
	CatalogSourceNamespace string `json:"catalogsourcenamespace"`
}

var checkflags CheckCommandFlags

// TODO: provide godoc compatible comment for checkCmd
var checkCmd = &cobra.Command{
	Use: "check",
	// TODO: provide Short description for check command
	Short: "",
	// TODO: provide Long description for check command
	Long: ``,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Wait for CatalogSource provided to be ready
		operator.NewClient()

		// Confirm PackageManifests is not empty
		psc, err := operator.NewPackageServerClient()
		if err != nil {
			// TODO: handle error; should be fatal
		}

		pml, err := psc.OperatorsV1().PackageManifests("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			// TODO: handle error; should be fatal
		}

		if len(pml.Items) == 0 {
			// TODO: return non-nil because this means even though we are sure the catalogsource is ready we will
			// did not get back any packagemanifest resources
			return types.Error{Msg: "No packagemanifests"}
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		capability.OperatorInstallAllFromCatalog(checkflags.CatalogSource, checkflags.CatalogSourceNamespace)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	flags := checkCmd.Flags()

	flags.StringVar(&checkflags.CatalogSource, "catalogsource", "certified-operators",
		"")
	flags.StringVar(&checkflags.CatalogSourceNamespace, "catalogsourcenamespace", "openshift-marketplace",
		"")
}

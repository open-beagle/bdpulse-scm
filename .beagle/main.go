package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/open-beagle/go-scm/scm"
	"github.com/open-beagle/go-scm/scm/driver/beagle"
	"github.com/open-beagle/go-scm/scm/transport/oauth2"
)

func main() {
	client := beagle.NewDefault()
	client.Client = &http.Client{
		Transport: &oauth2.Transport{
			Source: oauth2.StaticTokenSource(
				&scm.Token{
					Token: "Zw5G2GM2AKl39t7VMD6vwmCYep2Kckyf",
				},
			),
		},
	}
	ctx := context.Background()

	// repo
	repo, _, _ := client.Repositories.Find(ctx, "test")
	fmt.Println(repo.Perm.Admin)
}

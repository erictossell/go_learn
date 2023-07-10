package main

import "fmt"

func main() {
	/* Define an array and let the compiler count its size */
	DeploymentOptions := [4]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
  for i := 0; i < len(DeploymentOptions); i++ {
    option := DeploymentOptions[i] 
		fmt.Println(i, option)
	}
}

package main

import "fmt"

func getAliasName(firstname, lastname string) string {
	return firstname[0:2] + lastname[0:2]
}

func getName() ([2]string, error) {
	var firstname, lastname string

	fmt.Print("Enter a firstname: ")
	_, err := fmt.Scan(&firstname)
	if err != nil {
		return [2]string{}, err
	}
	fmt.Print("Enter a lastname: ")
	_, err = fmt.Scan(&lastname)
	if err != nil {
		return [2]string{}, err
	}
	return [2]string{firstname, lastname}, nil
}

func main() {
	name, err := getName()
	if err != nil {
		fmt.Println(err)
		return
	}
	firstname, lastname := name[0], name[1]
	fmt.Println("Hello", firstname, lastname)
	fmt.Println(getAliasName(firstname, lastname))
}

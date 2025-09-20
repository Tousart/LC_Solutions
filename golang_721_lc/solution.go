package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	emailAccount := make(map[string][]int)
	visited := make([]bool, len(accounts))
	res := [][]string{}

	for i, account := range accounts {
		for _, email := range account[1:] {
			emailAccount[email] = append(emailAccount[email], i)
		}
	}

	for i, account := range accounts {
		if visited[i] {
			continue
		}

		account := []string{account[0]}

		emails := make(map[string]struct{})
		emailsList := []string{}
		dfs(i, emails, &emailsList, visited, accounts, emailAccount)
		sort.Strings(emailsList)

		account = append(account, emailsList...)
		res = append(res, account)
	}

	return res
}

func dfs(i int, emails map[string]struct{}, emailsList *[]string, visited []bool, accounts [][]string, emailAccount map[string][]int) {
	if visited[i] {
		return
	}
	visited[i] = true

	for j := 1; j < len(accounts[i]); j++ {
		email := accounts[i][j]
		if _, ok := emails[email]; !ok {
			*emailsList = append(*emailsList, email)
			emails[email] = struct{}{}
		}
		for _, neighbor := range emailAccount[email] {
			dfs(neighbor, emails, emailsList, visited, accounts, emailAccount)
		}
	}
}

func main() {
	accounts := [][]string{
		{"John", "huy@123", "privet@123"},
		{"John", "huy@123", "rotik@123"},
		{"Marry", "lala@123", "lala@123"},
		{"Pidor", "pizda@123"},
		{"Huesos", "zalupa@123"},
		{"Huesos", "zalupa@123"},
		{"John", "rotik@123"},
	}
	fmt.Println(accountsMerge(accounts))

	accounts1 := [][]string{
		{"John", "0", "1"},
		{"John", "3", "4"},
		{"John", "4", "5"},
		{"John", "2", "3"},
		{"John", "1", "2"},
	}
	fmt.Println(accountsMerge(accounts1))
}

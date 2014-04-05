package rockstar

import (
	"fmt"
	"github.com/k0kubun/gothub"
	"sort"
)

func ShowSummarization(username string) {
	var repositories Repositories = repositoriesOf(username)
	if len(repositories) == 0 {
		return
	}
	sort.Sort(repositories)

	user, _ := github().GetUser(username)
	fmt.Printf("★ %d %s (%s)\n\n", starCountOf(repositories), user.Name, username)

	starLabelWidth := len(fmt.Sprintf("%d", repositories[0].WatchersCount)) + 1
	format := fmt.Sprintf("★%%%dd %%s\n", starLabelWidth)
	for i := 0; i < 10 && i < len(repositories); i++ {
		repository := repositories[i]
		fmt.Printf(format, repository.WatchersCount, repository.Name)
	}
	fmt.Println()
}

func starCountOf(repositories Repositories) (count int) {
	for _, repository := range repositories {
		count += repository.WatchersCount
	}
	return
}

func repositoriesOf(username string) (repositories Repositories) {
	user, _ := github().GetUser(username)

	page := 1
	for {
		paginatedRepositories, _ := user.Repositories(page)
		if len(paginatedRepositories) == 0 {
			break
		}
		repositories = append(repositories, paginatedRepositories...)
		page++
	}
	return
}

func github() *gothub.GitHub {
	if !authenticated() {
		authenticate()
	}

	github, _ := gothub.BasicLogin(usernameAndPassword())
	return github
}

type Repositories []gothub.Repository

func (repositories Repositories) Len() int {
	return len(repositories)
}

func (repositories Repositories) Swap(i, j int) {
	repositories[i], repositories[j] = repositories[j], repositories[i]
}

func (repositories Repositories) Less(i, j int) bool {
	return repositories[i].WatchersCount > repositories[j].WatchersCount
}

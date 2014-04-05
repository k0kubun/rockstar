package rockstar

import (
	"fmt"
	"github.com/k0kubun/gothub"
	"os"
	"sort"
)

func ShowSummarization(username string) {
	var repositories Repositories = repositoriesOf(username)
	if len(repositories) == 0 {
		return
	}
	sort.Sort(repositories)

	user := getUser(username)
	fmt.Printf("★ %d %s (%s)\n", starCountOf(repositories), user.Name, username)
	fmt.Printf("%d repos, %d following, %d followers\n\n", len(repositories), user.Following, user.Followers)

	numberToShow := 10
	starLabelWidth := len(fmt.Sprintf("%d", repositories[0].WatchersCount))
	nameLabelWidth := repositoryNameMaxLength(repositories, numberToShow)
	format := fmt.Sprintf("★ %%%dd %%-%ds %%s\n", starLabelWidth, nameLabelWidth)

	for i := 0; i < numberToShow && i < len(repositories); i++ {
		repository := repositories[i]
		fmt.Printf(format, repository.WatchersCount, repository.Name, repository.Language)
	}
	fmt.Println()
}

func getUser(username string) *gothub.User {
	user, err := github().GetUser(username)
	if err != nil {
		fmt.Printf("%s does not exist\n", username)
		os.Exit(1)
	}
	return user
}

func starCountOf(repositories Repositories) (count int) {
	for _, repository := range repositories {
		count += repository.WatchersCount
	}
	return
}

func repositoriesOf(username string) (repositories Repositories) {
	user := getUser(username)

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

func repositoryNameMaxLength(repositories Repositories, number int) (length int) {
	for i := 0; i < number && i < len(repositories); i++ {
		currentLength := len(repositories[i].Name)
		if length < currentLength {
			length = currentLength
		}
	}
	return
}

func github() *gothub.GitHub {
	guestClient, _ := gothub.Guest()
	if guestClient != nil {
		return guestClient
	}

	if !authenticated() {
		authenticate()
	}
	//FIXME: Support Two Factor Authentication
	loginClient, _ := gothub.BasicLogin(usernameAndPassword())

	return loginClient
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

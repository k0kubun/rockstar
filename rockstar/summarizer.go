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
	starCount := starCountOf(repositories)
	name := fmt.Sprintf("★ %d %s (%s)\n", starCount, user.Name, username)
	fmt.Printf("%s", coloredUser(name, starCount))
	fmt.Printf("%d repos, %d following, %d followers\n", len(repositories), user.Following, user.Followers)
	summarizeLanguages(repositories)
	fmt.Printf("\n\n")

	numberToShow := 10
	starLabelWidth := len(fmt.Sprintf("%d", repositories[0].WatchersCount))
	nameLabelWidth := repositoryNameMaxLength(repositories, numberToShow)
	format := fmt.Sprintf("★ %%%dd %%-%ds", starLabelWidth, nameLabelWidth)

	for i := 0; i < numberToShow && i < len(repositories); i++ {
		repository := repositories[i]
		starCount = repository.WatchersCount
		repositoryLabel := fmt.Sprintf(format, starCount, repository.Name)
		fmt.Printf(coloredRepository(repositoryLabel, starCount))
		fmt.Printf(" %s\n", repository.Language)
	}
	fmt.Println()
}

func summarizeLanguages(repositories Repositories) {
	countByLanguage := map[string]int{}
	for _, repository := range repositories {
		if repository.Language != "" {
			countByLanguage[repository.Language]++
		}
	}

	languageCount := len(countByLanguage)
	for i := 0; i < 3 && i < languageCount; i++ {
		dumpMaxCoverage(countByLanguage, len(repositories))
		if i != 2 && i != len(countByLanguage)-1 {
			fmt.Printf(", ")
		}
	}
}

func dumpMaxCoverage(countByLanguage map[string]int, numOfRepositories int) {
	var maxCountLanguage string
	var maxCount int = -1

	for language, count := range countByLanguage {
		if maxCount < count {
			maxCount = count
			maxCountLanguage = language
		}
	}
	delete(countByLanguage, maxCountLanguage)

	coverage := 100.0 * float32(maxCount) / float32(numOfRepositories)
	fmt.Printf("%s: %.1f%%", maxCountLanguage, coverage)
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

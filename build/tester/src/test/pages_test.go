package test

import (
	"fmt"
	"os"
	"test/helper"
	"testing"
)

type PageTestDataProvider struct {
	UsersQuantity    int
	useIndexPageList []UseIndexPages
}

type UseIndexPages struct{
	UseIndex bool
	PageUrlList []string
}

func PagesDataProvider() []PageTestDataProvider {
	pageUrlList := []string{"user-profile-list", "user-profile-page?id=1"}
	useIndexPageList := []UseIndexPages{
		{
			UseIndex: false,
			PageUrlList: pageUrlList,
		},
		{
			UseIndex: true,
			PageUrlList: pageUrlList,
		},
	}
	return []PageTestDataProvider{
		{
			UsersQuantity:    10000,
			useIndexPageList: useIndexPageList,
		},
		{
			UsersQuantity:    50000,
			useIndexPageList: useIndexPageList,
		},
		{
			UsersQuantity:    100000,
			useIndexPageList: useIndexPageList,
		},
	}
}

/*func BenchmarkPagesSequential(b *testing.B){
	benchmarkPages(b, benchmarkPageSequential)
}*/

func BenchmarkPagesParallel(b *testing.B){
	benchmarkPages(b,
		func(url string) func (b *testing.B) { return benchmarkPageParallel(url, 20)},
	)
}

func benchmarkPages(b *testing.B, benchmarkPageFunction func(url string) func (b *testing.B) ) {
	pagesDataProvider := PagesDataProvider()
	seedDataParams := helper.GetSeedDataParams()

	masterHostname := os.ExpandEnv("$MYSQL_MASTER_HOSTNAME")
	slaveHostname := os.ExpandEnv("$MYSQL_SLAVE_HOSTNAME")

	for _, pageTestDataProvider := range pagesDataProvider {
		helper.RemoveIndexes(seedDataParams, slaveHostname)
		seedDataParams.Users = pageTestDataProvider.UsersQuantity
		helper.ApplyFixture(seedDataParams, masterHostname)

		for _, useIndexPage := range pageTestDataProvider.useIndexPageList {

			if useIndexPage.UseIndex {
				helper.AddIndexes(seedDataParams, slaveHostname)
			}

			for _, pageUrl := range useIndexPage.PageUrlList {
				b.Run(
					fmt.Sprintf(
						"usersQuantity = %d :: useIndex = %v :: pageUrl = %s",
						pageTestDataProvider.UsersQuantity,
						useIndexPage.UseIndex,
						pageUrl,
					),
					benchmarkPageFunction(pageUrl),
				)
			}
		}
	}
}

func benchmarkPageSequential(url string) func (b *testing.B) {
	return func(b *testing.B) {
		errorCounter := 0
		successCounter := 0
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StartTimer()
			err := helper.TestPageSingleThead(url)
			b.StopTimer()
			if err != nil {
				errorCounter++
			} else {
				successCounter++
			}
		}
		//defer b.Logf("\nsuccess: %d, errors: %d\n", successCounter, errorCounter)
	}
}

func benchmarkPageParallel(url string, threads int) func (b *testing.B) {
	return func(b *testing.B) {
		errorChain := make(chan int, threads)
		successChan := make(chan int, threads)
		errorCounter := 0
		successCounter := 0
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			b.StartTimer()
			go testPageSingleThead(url, errorChain, successChan, i)
			select {
				case <-successChan:
					successCounter += 1
				case <-errorChain:
					errorCounter += 1
			}
			b.StopTimer()
		}
		close(errorChain)
		close(successChan)
		fmt.Printf("\nsuccess: %d, errors: %d\n", successCounter, errorCounter)
	}
}

func testPageSingleThead(url string, errorChain chan int, successChan chan int, num int){
	
	err := helper.TestPageSingleThead(url)
	if err != nil {
		errorChain <- 1
	} else {
		successChan <- 1
	}
}
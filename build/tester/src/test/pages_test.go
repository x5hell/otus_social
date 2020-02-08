package test

import (
	"fmt"
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

func BenchmarkPages(b *testing.B) {
	pagesDataProvider := PagesDataProvider()
	seedDataParams := helper.GetSeedDataParams()

	for _, pageTestDataProvider := range pagesDataProvider {
		helper.RemoveIndexes(seedDataParams)
		seedDataParams.Users = pageTestDataProvider.UsersQuantity
		helper.ApplyFixture(seedDataParams)

		for _, useIndexPage := range pageTestDataProvider.useIndexPageList {

			if useIndexPage.UseIndex {
				helper.AddIndexes(seedDataParams)
			}

			for _, pageUrl := range useIndexPage.PageUrlList {
				b.Run(
					fmt.Sprintf(
						"usersQuantity = %d :: useIndex = %v :: pageUrl = %s",
						pageTestDataProvider.UsersQuantity,
						useIndexPage.UseIndex,
						pageUrl,
					),
					benchmarkPage(pageUrl),
				)
			}
		}
	}
}

func benchmarkPage(url string) func (b *testing.B) {
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
		//b.Logf("\nsuccess: %d, errors: %d\n", successCounter, errorCounter)
	}
}
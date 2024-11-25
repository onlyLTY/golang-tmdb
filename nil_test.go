package tmdb

func (suite *TMBDTestSuite) TestGetMovieDetailsCheckNil() {
	bumblebee, err := suite.client.GetMovieDetails(bumblebeeID, nil)
	suite.Nil(err)
	suite.NotNil(bumblebee)
	suite.NotNil(bumblebee.MovieAlternativeTitlesAppend)
	suite.NotNil(bumblebee.MovieAlternativeTitlesAppend.AlternativeTitles)
	suite.NotNil(bumblebee.MovieCreditsAppend.Credits.MovieCredits)
	suite.NotNil(bumblebee.MovieTranslationsAppend.Translations)
	suite.NotNil(bumblebee.MovieReleaseDatesAppend)
	suite.NotNil(bumblebee.MovieReleaseDatesAppend.ReleaseDates)
}

func (suite *TMBDTestSuite) TestGetSearchTVShowCheckNil() {
	search, err := suite.client.GetSearchTVShow(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.SearchTVShowsResults)
}

func (suite *TMBDTestSuite) TestGetSearchMoviesCheckNil() {
	search, err := suite.client.GetSearchMovies(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.SearchMoviesResults)
}

func (suite *TMBDTestSuite) TestGetTVDetailsCheckNil() {
	vikings, err := suite.client.GetTVDetails(vikingsID, nil)
	suite.Nil(err)
	suite.NotNil(vikings)
	suite.NotNil(vikings.TVCreditsAppend)
	suite.NotNil(vikings.TVCreditsAppend.Credits)
	suite.NotNil(vikings.TVCreditsAppend.Credits.TVCredits)
	suite.NotNil(vikings.TVAlternativeTitlesAppend)
	suite.NotNil(vikings.TVAlternativeTitlesAppend.AlternativeTitles)
	suite.NotNil(vikings.TVAlternativeTitlesAppend.AlternativeTitles.TVAlternativeTitlesResults)
	suite.NotNil(vikings.TVTranslationsAppend)
	suite.NotNil(vikings.TVTranslationsAppend.Translations)
	suite.NotNil(vikings.TVContentRatingsAppend)
	suite.NotNil(vikings.TVContentRatingsAppend.ContentRatings)
	suite.NotNil(vikings.TVContentRatingsAppend.ContentRatings.TVContentRatingsResults)
}

func (suite *TMBDTestSuite) TestGetTVSeasonDetailsCheckNil() {
	got, err := suite.client.GetTVSeasonDetails(gotID, 1, nil)
	suite.Nil(err)
	suite.NotNil(got.TVSeasonCreditsAppend)
	suite.NotNil(got.TVSeasonCreditsAppend.Credits)
	suite.NotNil(got.TVAggregateCreditsAppend)
	suite.NotNil(got.TVAggregateCreditsAppend.AggregateCredits)
}

func (suite *TMBDTestSuite) TestGetSearchMultiCheckNil() {
	search, err := suite.client.GetSearchMulti(defaultQuery, nil)
	suite.Nil(err)
	suite.NotNil(search.SearchMultiResults)
}

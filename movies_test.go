package tmdb

const (
	bumblebeeID   = 424783
	jackReacherID = 75780
)

func (suite *TMBDTestSuite) TestGetMovieDetails() {
	bumblebee, err := suite.GetMovieDetails(bumblebeeID, nil)
	suite.Nil(err)
	suite.Equal("Bumblebee", bumblebee.Title)
}

func (suite *TMBDTestSuite) TestGetMovieDetailsFail() {
	_, err := suite.GetMovieDetails(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieDetailsWithOptions() {
	options := make(map[string]string)
	options["language"] = "pt-BR"
	jackreacher, err := suite.GetMovieDetails(jackReacherID, options)
	suite.Nil(err)
	suite.Equal("Jack Reacher: O Último Tiro", jackreacher.Title)
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitles() {
	bumblebee, err := suite.GetMovieAlternativeTitles(bumblebeeID, nil)
	suite.Nil(err)
	for _, v := range bumblebee.Titles {
		if v.Iso3166_1 == "US" {
			suite.Equal("Transformers 6", v.Title)
		}
	}
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesFail() {
	_, err := suite.GetMovieAlternativeTitles(0, nil)
	suite.Equal("The resource you requested could not be found.", err.Error())
}

func (suite *TMBDTestSuite) TestGetMovieAlternativeTitlesWithOptions() {
	options := make(map[string]string)
	options["country"] = "RU"
	bumblebee, err := suite.GetMovieAlternativeTitles(bumblebeeID, options)
	suite.Nil(err)
	suite.Equal("RU", bumblebee.Titles[0].Iso3166_1)
}

func (suite *TMBDTestSuite) TestGetMovieChanges() {
	bumblebee, err := suite.GetMovieChanges(bumblebeeID, nil)
	suite.Nil(err)
	for _, v := range bumblebee.Changes {
		for _, v := range v.Items {
			suite.NotNil(v.ID)
		}
	}
}

// The API isn't handling zero values for this end-point.
// TODO: Fix this test later.
func (suite *TMBDTestSuite) TestGetMovieChangesFail() {
	_, err := suite.GetMovieChanges(0, nil)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestGetMovieChangesWithOptions() {
	options := make(map[string]string)
	options["start_date"] = "2019-01-01"
	options["end_date"] = "2019-01-12"
	options["page"] = "1"
	bumblebee, err := suite.GetMovieChanges(bumblebeeID, options)
	suite.Nil(err)
	for _, v := range bumblebee.Changes {
		for _, v := range v.Items {
			suite.NotNil(v.ID)
		}
	}
}

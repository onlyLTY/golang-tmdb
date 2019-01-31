package tmdb

func (suite *TMBDTestSuite) TestGetDiscoverMovie() {
	o := make(map[string]string)
	o["language"] = "en-US"
	o["include_adult"] = "false"
	o["include_video"] = "false"
	o["sort_by"] = "popularity.desc"
	discover, err := suite.GetDiscoverMovie(o)
	suite.Nil(err)
	suite.NotNil(discover)
}

func (suite *TMBDTestSuite) TestGetDiscoverMovieFail() {
	suite.APIKey = ""
	o := make(map[string]string)
	o["language"] = "en-US"
	o["include_adult"] = "false"
	o["include_video"] = "false"
	o["sort_by"] = "popularity.desc"
	_, err := suite.GetDiscoverMovie(o)
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestGetDiscoverTV() {
	o := make(map[string]string)
	o["page"] = "1"
	o["language"] = "en-US"
	o["timezone"] = "America/New_York"
	o["sort_by"] = "popularity.desc"
	o["include_null_first_air_dates"] = "false"
	discover, err := suite.GetDiscoverTV(o)
	suite.Nil(err)
	suite.NotNil(discover)
}

func (suite *TMBDTestSuite) TestGetDiscoverTVFail() {
	suite.APIKey = ""
	o := make(map[string]string)
	o["page"] = "1"
	o["language"] = "en-US"
	o["timezone"] = "America/New_York"
	o["sort_by"] = "popularity.desc"
	o["include_null_first_air_dates"] = "false"
	_, err := suite.GetDiscoverTV(o)
	suite.NotNil(err)
}

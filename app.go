package main

import (
    "time"
)

type App struct {
	Advisories []string `json:"advisories"`
	ArtistId int `json:"artistId"`
	ArtistName string `json:"artistName"`
	ArtistViewUrl string `json:"artistViewUrl"`
	ArtworkUrl100 string `json:"artworkUrl100"`
	ArtworkUrl512 string `json:"artworkUrl512"`
	ArtworkUrl60 string `json:"artworkUrl60"`
	AverageUserRating float32 `json:"averageUserRating"`
	AverageUserRatingForCurrentVersion float32 `json:"averageUserRatingForCurrentVersion"`
	BundleId string `json:"bundleId"`
	ContentAdvisoryRating string `json:"contentAdvisoryRating"`
	Currency string `json:"currency"`
	Description string `json:"description"`
	Features []string `json:"features"`
	FileSizeBytes string `json:"fileSizeBytes"`
	FormattedPrice string `json:"formattedPrice"`
	GenreIds []string `json:"genreIds"`
	Genres []string `json:"genres"`
	IpadScreenshotUrls []string `json:"ipadScreenshotUrls"`
	IsGameCenterEnabled bool `json:"isGameCenterEnabled"`
	//IsVppDeviceBasedLicensingEnabled bool `json:"isVppDeviceBasedLicensingEnabled"`
	Kind string `json:"kind"`
	LanguageCodesISO2A []string `json:"languageCodesISO2A"`
	MinimumOsVersion string `json:"minimumOsVersion"`
	Price float32 `json:"price"`
	PrimaryGenreId int `json:"primaryGenreId"`
	PrimaryGenreName string `json:"primaryGenreName"`
	ReleaseDate time.Time `json:"releaseDate"`
	ReleaseNotes string `json:"releaseNotes"`
	ScreenshotUrls []string `json:"screenshotUrls"`
	SellerName string `json:"sellerName"`
	SellerUrl string `json:"sellerUrl"`
	SupportedDevices []string `json:"supportedDevices"`
	TrackCensoredName string `json:"trackCensoredName"`
	TrackContentRating string `json:"trackContentRating"`
	TrackId int `json:"trackId"`
	TrackName string `json:"trackName"`
	TrackViewUrl string `json:"trackViewUrl"`
	UserRatingCount int `json:"userRatingCount"`
	UserRatingCountForCurrentVersion int `json:"userRatingCountForCurrentVersion"`
	Version string `json:"version"`
	WrapperType string `json:"wrapperType"`
}
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Applications struct {
	XMLName       xml.Name `xml:"applications"`
	Text          string   `xml:",chardata"`
	VersionsDelta string   `xml:"versions__delta"`
	AppsHashcode  string   `xml:"apps__hashcode"`
	Application   []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name"`
		Instance struct {
			Text             string `xml:",chardata"`
			InstanceId       string `xml:"instanceId"`
			HostName         string `xml:"hostName"`
			App              string `xml:"app"`
			IpAddr           string `xml:"ipAddr"`
			Status           string `xml:"status"`
			Overriddenstatus string `xml:"overriddenstatus"`
			Port             struct {
				Text    string `xml:",chardata"`
				Enabled string `xml:"enabled,attr"`
			} `xml:"port"`
			SecurePort struct {
				Text    string `xml:",chardata"`
				Enabled string `xml:"enabled,attr"`
			} `xml:"securePort"`
			CountryId      string `xml:"countryId"`
			DataCenterInfo struct {
				Text     string `xml:",chardata"`
				Class    string `xml:"class,attr"`
				Name     string `xml:"name"`
				Metadata struct {
					Text             string `xml:",chardata"`
					AccountId        string `xml:"accountId"`
					LocalHostname    string `xml:"local-hostname"`
					InstanceID       string `xml:"instance-id"`
					LocalIpv4        string `xml:"local-ipv4"`
					InstanceType     string `xml:"instance-type"`
					VpcID            string `xml:"vpc-id"`
					AmiID            string `xml:"ami-id"`
					Mac              string `xml:"mac"`
					AvailabilityZone string `xml:"availability-zone"`
				} `xml:"metadata"`
			} `xml:"dataCenterInfo"`
			LeaseInfo struct {
				Text                  string `xml:",chardata"`
				RenewalIntervalInSecs string `xml:"renewalIntervalInSecs"`
				DurationInSecs        string `xml:"durationInSecs"`
				RegistrationTimestamp string `xml:"registrationTimestamp"`
				LastRenewalTimestamp  string `xml:"lastRenewalTimestamp"`
				EvictionTimestamp     string `xml:"evictionTimestamp"`
				ServiceUpTimestamp    string `xml:"serviceUpTimestamp"`
			} `xml:"leaseInfo"`
			Metadata struct {
				Text          string `xml:",chardata"`
				Ec2InstanceId string `xml:"ec2InstanceId"`
				ImageId       string `xml:"imageId"`
			} `xml:"metadata"`
			AppGroupName                  string `xml:"appGroupName"`
			HomePageUrl                   string `xml:"homePageUrl"`
			StatusPageUrl                 string `xml:"statusPageUrl"`
			HealthCheckUrl                string `xml:"healthCheckUrl"`
			SecureHealthCheckUrl          string `xml:"secureHealthCheckUrl"`
			VipAddress                    string `xml:"vipAddress"`
			SecureVipAddress              string `xml:"secureVipAddress"`
			IsCoordinatingDiscoveryServer string `xml:"isCoordinatingDiscoveryServer"`
			LastUpdatedTimestamp          string `xml:"lastUpdatedTimestamp"`
			LastDirtyTimestamp            string `xml:"lastDirtyTimestamp"`
			ActionType                    string `xml:"actionType"`
			AsgName                       string `xml:"asgName"`
		} `xml:"instance"`
	} `xml:"application"`
}

func GetServices() (Applications, error) {
	response, err := http.Get("http://localhost:8761/eureka/apps")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return Applications{}, nil
	}

	// we initialize our Users array
	var aplications Applications
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err = xml.Unmarshal(responseData, &aplications)
	if err != nil {
		log.Fatal(err)
		return Applications{}, nil
	}
	return aplications, nil
}

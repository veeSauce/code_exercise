This is a Simple GoLang REST Application that gives back a random name and a joke with that name in it

API source for name: http://uinames.com/api/
API source for joke: http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]

The service once running has 2 endpoints, both can be reached with a GET request:

1- "/" This endpoint gives back the joke with a random name in it
2- "/health" This is for setting up healthcheck on the app (more discussed further)

The vendor directory has all the external dependency library and the folder is necessary strictly to deploy instances of Google Cloud Platform APP engine

If you are trying to run the app locally, follow these instructions:

1- Install GO 
2- Setup Goroot and Gopath correctly
3- clone the binary: <github link here>
4: cd into src directory and run go build main.go
5: run command ./main - app will run on the URL below

http://localhost:5000

You can curl it, postman it or just refresh your browser on http://localhost:5000 to see the application results



Deploying to GCP (Exposing your web server to the internet) involves the following steps:

1- Create a test [project](https://cloud.google.com/resource-manager/docs/creating-managing-projects)(call it test-project + make sure to enable billing !)



2- Install [gcloud] (https://cloud.google.com/sdk/docs/quickstart-macos) command line tool (alternative option is to do everything in the GUI)

3- For your new Test-Project setup a Service Account by running the following commands

    1) gcloud beta iam service-accounts create joke-web-app-account
    --description "this account helps create resources for deploying the joke web application"
    --display-name "joke-web-app"

    2) gcloud projects add-iam-policy-binding my-project-123 \
    --member serviceAccount:joke-web-app-account@test-project.iam.gserviceaccount.com \
    --role roles/owner
  
    3) gcloud iam service-accounts keys create ~/key.json
    --iam-account joke-web-app-account@test-project.iam.gserviceaccount.com

    4) gcloud auth activate-service-account --project=test-project --key-file=gcpcmdlineuser.json


For our purpose we can choose the App Engine Resource from Google Cloud Platform (Automatic scaling)

4- open main.go and change the listen&serve port from 5000 to 8080

The above step lets us skip the port forwarding and firewall creation steps. (If you want to deploy and run on a custom port other than 8080, in the app.yaml you need to create port forwarding and firewall rules)

5- run the commamd gcloud app deploy

Your terminal will show you the URL for your deployed application

***If you redeploy (gcloud app deploy), it will override the previous deployment with new version

To see the application running click [here] (https://kuber-test-239218.appspot.com) ! < My deployment :)


This is a Simple GoLang REST Application that gives back a random name and a joke with that name in it

API source for name: http://uinames.com/api/
API source for joke: http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]

If you are trying to run the app locally, follow these instructions:

1- Install GO 
2- Setup Goroot and Gopath correctly
3- clone the binary: <github link here>
4: cd into src directory and run go build main.go
5: run ./main - app will run on 

http://localhost:5000

You can curl it, postman it or just refresh your browser on http://localhost:5000 to see the application results



Deploying to GCP Involves steps from GUI and/or CLI:

Login to Google Cloud website and Create a test project  (call it test-project + make sure to enable billing !)

https://cloud.google.com/resource-manager/docs/creating-managing-projects

Install gcloud command line tool (alternative option is to do everything in the GUI)

For your new Test-Project we need a token to talk to it, so for that we create a Service Account with owner access

gcloud beta iam service-accounts create joke-web-app-account
    --description "this account helps create resources for deploying the joke web application"
    --display-name "joke-web-app"

gcloud projects add-iam-policy-binding my-project-123 \
  --member serviceAccount:joke-web-app-account@test-project.iam.gserviceaccount.com \
  --role roles/owner
  
gcloud iam service-accounts keys create ~/key.json
  --iam-account joke-web-app-account@test-project.iam.gserviceaccount.com

gcloud auth activate-service-account --project=someproject --key-file=gcpcmdlineuser.json in your local terminal

Now we can create a VM with a Linux image provided, run the fist command below then fill the appropriate values

gcloud compute images list

gcloud compute instances create [INSTANCE_NAME] \
--image-family [IMAGE_FAMILY] \
--image-project [IMAGE_PROJECT]

SSH into the instance either through the console GUI or gcloud CLI on your local (fill out zone and instance name)

gcloud compute ssh --project test-project --zone [ZONE] [INSTANCE_NAME]

Install GO just the way you have had to install it in your local machine (double check it is installed run go env)

Copy the source code (main.go) from your local to the cloud instance (fill out file path and instance name as corresponding to what you created)

gcloud compute scp [LOCAL_FILE_PATH] [INSTANCE_NAME]:~

run go install and run ./$GOPATH/bin/main where gopath is however you configured it

**For a full scaled app you need the following resources:

-Forwarding Rule, Target Proxy, URL Map, Backend Service, HealthCheck, Backend + Instance (all the above got the instance ready)


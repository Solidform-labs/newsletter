#!/bin/bash

# The ID of your GCP project
GCLOUD_PROJECT=weighty-flag-416118

# The name of your Cloud Run service
SERVICE=newsletter-test

# The GCP region your service is deployed in
REGION=europe-west1

# Path to your .env file
ENV_FILE=.env.yaml

gcloud builds submit --tag=gcr.io/$GCLOUD_PROJECT/$SERVICE .
gcloud run deploy $SERVICE --platform=managed --image=gcr.io/$GCLOUD_PROJECT/$SERVICE --region=$REGION --env-vars-file=$ENV_FILE
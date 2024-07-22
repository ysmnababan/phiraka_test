docker build -t phiraka_api .
docker tag phiraka_api gcr.io/test-gcp-427110/phiraka_api
docker push gcr.io/test-gcp-427110/phiraka_api
gcloud run deploy phiraka-api --image gcr.io/test-gcp-427110/phiraka_api --platform managed --region asia-southeast2 --allow-unauthenticated --project test-gcp-427110 --port 8080
pause
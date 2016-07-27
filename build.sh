docker build --rm --tag pixels .
docker run -it --rm -v $(pwd)/bin:/app/bin pixels
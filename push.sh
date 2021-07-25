IFS=

heroku container:login

DOCKER_BUILDKIT=1 docker build -t meli-operacion-fuego:latest --build-arg BUILDKIT_INLINE_CACHE=1 .

DOCKERFILE=$(cat Dockerfile)

echo ${DOCKERFILE} | DOCKER_BUILDKIT=1 docker build -t meli-operacion-fuego:latest --cache-from=meli-operacion-fuego:latest --build-arg BUILDKIT_INLINE_CACHE=1 -f - .
docker tag meli-operacion-fuego:latest registry.heroku.com/meli-operacion-fuego/web:latest
docker push registry.heroku.com/meli-operacion-fuego/web:latest
heroku container:release -a meli-operacion-fuego web

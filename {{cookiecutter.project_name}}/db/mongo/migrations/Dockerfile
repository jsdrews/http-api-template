FROM registry.gitlab.com/biglittletech/platform/packages/mongo-migrations:v2
COPY migrations /mongo_migrations/migrations
CMD [ "migrate-mongo", "up" ]